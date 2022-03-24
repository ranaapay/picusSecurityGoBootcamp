package main

import (
	csv_utils "PicusBootcamp/lesson4/file-examples/csv"
	"bufio"
	"fmt"
	"log"
	"os"
)

var filename = "patikadev.txt"

func main() {
	//Create empty file
	//_, err := CreateEmptyFile()
	//if err != nil {
	//	log.Fatal(err)
	//}

	//open file
	//file, err := os.OpenFile(filename, os.O_RDWR, 0755)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//WriteFile(file)

	//Get information about the file
	//err := GetFileInfo()
	//if err != nil {
	//	log.Fatal(err)
	//}

	//Read file line by line
	//err := ReadFileLines()
	//if err != nil {
	//	log.Fatal(err)
	//}

	//Read file word by word
	//err := ReadFileWords()
	//if err != nil {
	//	log.Fatal(err)
	//}

	//Read csv
	//locations, err := csv_utils.ReadCSV("locations.csv")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("Count: ", len(locations))

	//JSON to CSV
	//if err := csv_utils.JSONToCSV("locations.json","custom_locations.csv"); err != nil {
	//	fmt.Println(err)
	//}

	err := csv_utils.ReadLocationsWithWorkerPool("locations.csv")
	if err != nil {
		fmt.Println(err)
	}
}

func CreateEmptyFile() (*os.File, error) {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) == false {
		return nil, fmt.Errorf("file already exists")
	}
	myFile, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	log.Println("File created", myFile.Name())

	return myFile, nil
}

func WriteFile(file *os.File) {
	defer file.Close()
	w := bufio.NewWriter(file)
	w.WriteString("patika dev\n")
	w.WriteString("bootcamp\n")
	w.WriteString("golang is the best\n")
	w.Flush()
}

func GetFileInfo() error {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return err
	}

	fmt.Println("File Name: ", fileInfo.Name())
	fmt.Println("Size: ", fileInfo.Size())
	fmt.Println("Permission: ", fileInfo.Mode())
	fmt.Println("Last Modified: ", fileInfo.ModTime())
	fmt.Println("Is Directory: ", fileInfo.IsDir())

	return nil
}

func ReadFileLines() error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Printf("line: %s\n", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func ReadFileWords() error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Printf("line: %s\n", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
