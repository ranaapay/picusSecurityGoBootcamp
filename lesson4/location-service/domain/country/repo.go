package country

import "gorm.io/gorm"

type CountryRepository struct {
	db *gorm.DB
}

func NewCountryRepository(db *gorm.DB) *CountryRepository {
	return &CountryRepository{db: db}
}

func (c *CountryRepository) GetAllCountriesWithCityInformation() ([]Country, error) {
	var countries []Country
	result := c.db.Preload("Cities").Find(&countries)
	if result.Error != nil {
		return nil, result.Error
	}
	return countries, nil
}

func (c *CountryRepository) Migration() {
	c.db.AutoMigrate(&Country{})
}

func (c *CountryRepository) InsertSampleData() {
	countries := []Country{
		{Name: "Türkiye", Code: "TR"},
		{Name: "Amerika", Code: "US"},
	}
	for _, country := range countries {
		c.db.Create(&country)
	}
}
