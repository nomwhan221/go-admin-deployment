package models

import (
	"math"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Paginate(db *gorm.DB,entity Entity,page int) fiber.Map {
	limit := 5
	offset := (page-1) * limit
	//var total int64
	//var products []Product

	data := entity.Take(db,limit,offset)

	total := entity.Count(db)

	// db.Offset(offset).Limit(limit).Find(&products)

	// db.Model(&Product{}).Count(&total)

	return fiber.Map{
		"data" :data,
		"meta":fiber.Map{
			"total":total,
			"page":page,
			"last_page":math.Ceil((float64(total) / float64(limit))),
		},
	}
}