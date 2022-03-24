package csv_utils

import (
	"PicusBootcamp/lesson4/file-examples/models"
	"encoding/csv"
	"fmt"
	"os"
	"sync"
)

func ReadLocationsWithWorkerPool(path string) error {
	const numJobs = 5
	jobs := make(chan []string, numJobs)
	results := make(chan models.Location, numJobs)
	wg := sync.WaitGroup{}

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go toStruct(jobs, results, &wg)
	}

	go func() {
		f, _ := os.Open(path)
		defer f.Close()

		lines, _ := csv.NewReader(f).ReadAll()
		for _, line := range lines[1:] {
			jobs <- line
		}
		close(jobs)
	}()
	go func() {
		wg.Wait()
		close(results)
	}()
	for v := range results {
		fmt.Println(v)
	}
	return nil
}

func toStruct(jobs <-chan []string, results chan<- models.Location, wg *sync.WaitGroup) {
	defer wg.Done()

	for j := range jobs {
		location := models.Location{
			CountryCode: j[0],
			CountryName: j[1],
			CityName:    j[2],
			IATACode:    j[3],
			Latitude:    j[4],
			Longitude:   j[5],
		}
		results <- location
	}
}
