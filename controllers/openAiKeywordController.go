package controllers

import (
	"alibabaoss/conf"
	"context"
	"encoding/json"

	"net/http"

	"github.com/invopop/jsonschema"
	openai "github.com/openai/openai-go/v2"
	"github.com/openai/openai-go/v2/option"
)

type extractReq struct {
	Text        string `json:"text"`
	MaxKeywords int    `json:"max_keywords"`
	Language    string `json:"language"`
}

// Schema for keyword + score
type Keyword struct {
	Keyword string  `json:"keyword" jsonschema_description:"Single keyword or phrase"`
	Score   float64 `json:"score"   jsonschema:"minimum=0,maximum=1" jsonschema_description:"Relevance score [0,1]"`
}

type ExtractResult struct {
	Keywords []Keyword `json:"keywords" jsonschema_description:"Sorted by descending relevance"`
}

func jsonSchema[T any]() any {
	r := jsonschema.Reflector{
		AllowAdditionalProperties: false,
		DoNotReference:            true,
	}
	var zero T
	return r.Reflect(zero)
}

var extractSchema = jsonSchema[ExtractResult]()

type OpenAiKeywordController struct{ BaseController }

/*
	 Sample payload response
		{
		  "keywords": [
		    {"keyword":"Responses API","score":0.96},
		    {"keyword":"agent workflows","score":0.91},
		    {"keyword":"OpenAI","score":0.89}
		  ]
		}
*/
func (c *OpenAiKeywordController) Post() {
	// TODO: uncomment below code for production use
	// var req extractReq
	// if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil || len(req.Text) == 0 {
	// 	c.CustomAbort(http.StatusBadRequest, "invalid payload: need {text}")
	// 	return
	// }
	// if req.MaxKeywords <= 0 || req.MaxKeywords > 20 {
	// 	req.MaxKeywords = 8
	// }

	// Sample payload for testing
	req := extractReq{
		Text: `
SEATTLE — Seattle has temporarily closed three parks, citing ongoing safety concerns and misuse. But neighbors say the closures haven’t solved the problem — and a growing encampment across the street from an elementary school has them worried.
Seven Hills Park on Capitol Hill, Lake City Mini Park, and Blanche Lavizzo Park in the Central District were all closed Aug. 28. Seattle Parks and Recreation says the parks will remain closed for 60 days while the city considers changes such as new lighting, decorative fencing or removing certain amenities.

Neighbors who live near Seven Hills Park say conditions there had spiraled out of control.

“This problem has metastasized,” one Capitol Hill resident said. “The park has been marked as a safe space to come and do drugs and be an outlaw, and that cannot go forward.”

Before the fences went up, photos taken at Seven Hills Park showed tents, trash and used needles scattered across the grounds.

While residents welcome the temporary closure, they fear the activity is simply moving elsewhere.

“We’ve been bombarding the city with our Find It, Fix It reports, and it is only after months that something will miraculously happen,” one neighbor said.

Several blocks away, neighbors say an encampment across the street from Lowell Elementary School has been growing for weeks. Mary Lamery, who lives nearby, says she’s deeply concerned about students seeing or walking past the site.

`,
		MaxKeywords: 10,
		Language:    "en",
	}

	apiKey := conf.OPENAI_API_KEY //os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		c.CustomAbort(http.StatusInternalServerError, "OPENAI_API_KEY not set")
		return
	}

	client := openai.NewClient(option.WithAPIKey(apiKey))
	ctx := context.Background()

	// Structured output schema
	respFormat := openai.ChatCompletionNewParamsResponseFormatUnion{
		OfJSONSchema: &openai.ResponseFormatJSONSchemaParam{
			JSONSchema: openai.ResponseFormatJSONSchemaJSONSchemaParam{
				Name:        "extract_result",
				Description: openai.String("Extract keywords and their relevance scores"),
				Schema:      extractSchema,
				Strict:      openai.Bool(true),
			},
		},
	}

	systemPrompt := `You are a keyword extractor.
- Output strictly in JSON matching the provided schema.
- Extract up to N keywords or short phrases from the text.
- Score is relevance [0,1].
- Sort by score descending.
- No extra fields or explanations.`

	userPrompt := map[string]any{
		"text":         req.Text,
		"max_keywords": req.MaxKeywords,
		"language":     req.Language,
	}

	chat, err := client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Model: openai.ChatModelGPT4o2024_08_06,
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(systemPrompt),
			openai.UserMessage(mustJSON(userPrompt)),
		},
		ResponseFormat: respFormat,
		Seed:           openai.Int(42),
	})
	if err != nil {
		c.CustomAbort(http.StatusBadGateway, err.Error())
		return
	}

	var result ExtractResult
	if err := json.Unmarshal([]byte(chat.Choices[0].Message.Content), &result); err != nil {
		c.CustomAbort(http.StatusBadGateway, "failed to parse model JSON: "+err.Error())
		return
	}

	c.Data["json"] = result
	c.ServeJSON()
}

func mustJSON(v any) string {
	b, _ := json.Marshal(v)
	return string(b)
}
