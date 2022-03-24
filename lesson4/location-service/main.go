package main

import (
	postgres "PicusBootcamp/lesson4/location-service/common/db"
	"PicusBootcamp/lesson4/location-service/domain/city"
	"PicusBootcamp/lesson4/location-service/domain/country"
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := postgres.NewPsqlDB()
	if err != nil {
		log.Fatal("Postgres cannot init ", err)
	}
	log.Println("Postgres connected")

	//Repositories
	cityRepo := city.NewCityRepository(db)
	cityRepo.Migrations()
	cityRepo.InsertSampleData()

	fmt.Println(len(cityRepo.FindAll()))
	fmt.Println(cityRepo.FindByCountryCode("TR"))
	fmt.Println(cityRepo.FindByCountryCodeOrCityCode("03"))
	fmt.Println(cityRepo.FindByName("Ada"))

	countryRepo := country.NewCountryRepository(db)
	countryRepo.Migration()
	countryRepo.InsertSampleData()

	fmt.Println(countryRepo.GetAllCountriesWithCityInformation())
}
