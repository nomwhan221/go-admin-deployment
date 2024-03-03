package models

import "gorm.io/gorm"

type Order struct {
	Id         uint        `json:"id"`
	FirstName  string      `json:"first_name"`
	LastName   string      `json:"last_name"`
	Name 	   string		`json:"name" gorm:"-"`
	Email      string      `json:"email"`
	Total      float32     		`json:"total" gorm:"-"`
	UpdatedAt  string     `json:"updated_at"`
	CreatedAt  string     `json:"created_at"`
	OrderItems []OrderItem `json:"order_items" gorm:"foreignKey:OrderId"`
}

type OrderItem struct {
	Id           uint    `json:"id"`
	OrderId      uint    `json:"order_id"`
	ProductTitle string  `json:"product_title"`
	Price        float32 `json:"price"`
	Quantity     uint    `json:"quantity"`
}

func (order *Order) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Order{}).Count(&total)
	return total
}

func (order *Order) Take(db *gorm.DB, limit int, offset int) interface{} {

	var orders []Order
	db.Preload("OrderItems").Offset(offset).Limit(limit).Find(&orders)

	for i, _ := range orders{
		orders[i].Name = orders[i].FirstName + " " + orders[i].LastName
		var total float32 = 0

		for j, _ := range orders[i].OrderItems {
			var itemorder =  orders[i].OrderItems[j]
			total += (itemorder.Price) * float32(itemorder.Quantity)
		}
		orders[i].Total = total
	}

	return orders
}