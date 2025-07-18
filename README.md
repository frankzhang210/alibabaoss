# AlibabaOSS Project

## Project Overview

This is a Go project based on the Beego framework, designed to interact with Alibaba Cloud Object Storage Service (OSS). The project provides an API interface that allows users to post, get and delete a json string stored in Alibaba Cloud OSS. Alibaba Cloud OSS stores the string as the file. 

## Technology Stack

- [Go](https://golang.org/) - Programming language
- [Beego v2](https://github.com/beego/beego) - Web framework
- [Alibaba Cloud OSS SDK](https://github.com/aliyun/aliyun-oss-go-sdk) - Alibaba Cloud Object Storage Service SDK

## Installation and Usage

### Prerequisites

- Go 1.16 or higher
- Alibaba Cloud OSS account and access credentials

### Installation Steps

1. Clone the repository
   ```bash
   git clone https://github.com/frankzhang210/alibabaoss.git
   cd alibabaoss
   ```

2. Install dependencies
   ```bash
   go mod tidy
   ```

3. Configure OSS credentials
   Set up your Alibaba Cloud OSS access credentials in `conf/app.conf`

4. Build and run
   ```bash
   go build
   ./alibabaoss
   ```
## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details