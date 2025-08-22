package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Product struct {
	Id         uint      `json:"id" orm:"auto;pk"`
	ProductId  string    `json:"product_id" orm:"size(20);unique"`
	Email      string    `json:"email" orm:"size(100)"`
	Subject    string    `json:"subject" orm:"size(100)"`
	Phone      string    `json:"phone" orm:"size(30);null"`
	CreateTime time.Time `json:"create_time" orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `json:"update_time" orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Product))
}

func InsertOrUpdateProduct(p *Product) (int64, error) {
	o := orm.NewOrm()
	// The second argument "ProductId" tells the ORM to check for a conflict
	// on the ProductId column. If a row with that ProductId exists, it updates it.
	// Otherwise, it inserts a new row.
	id, err := o.InsertOrUpdate(p, "ProductId")
	return id, err
}

func GetProduct(productId string) (*Product, error) {
	o := orm.NewOrm()
	p := &Product{ProductId: productId}
	err := o.Read(p, "ProductId")
	if err != nil {
		return nil, err
	}
	return p, nil
}
