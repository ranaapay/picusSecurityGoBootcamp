package country

import (
	"PicusBootcamp/lesson4/location-service/domain/city"
	"gorm.io/gorm"
)

type Country struct {
	gorm.Model
	Name   string
	Code   string
	Cities []city.City `gorm:"foreignKey:country_code;references:Code"`
}
