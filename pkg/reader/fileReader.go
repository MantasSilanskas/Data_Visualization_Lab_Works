package reader

import (
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

var dataFilesPath = "../cmd/data/"

// InputFiles reads csv file where we keep all our data files names and returns list of them
func InputFiles(filename string) ([]File, error) {

	var (
		err  error
		list []File
	)

	file, err := os.Open(dataFilesPath + filename)
	if err != nil {
		log.Println("Failed to open file input file. Error:", err)
		return list, err
	}
	defer file.Close()

	files := []*File{}

	if err := gocsv.UnmarshalFile(file, &files); err != nil {
		log.Println("Failed to unmarshal data. Error:", err)
		return list, err
	}

	for _, v := range files {
		list = append(list, *v)
	}

	return list, nil
}

// ReadFileData opens file by given name and returns list of all data of that file
func ReadFileData(filename string) ([]DeviceData, error) {

	list := []DeviceData{}

	file, err := os.Open(dataFilesPath + filename)
	if err != nil {
		log.Println("Failed to open data file.  Error:", err)
		return list, err
	}
	defer file.Close()

	data := []*DeviceData{}

	if err := gocsv.UnmarshalFile(file, &data); err != nil {
		log.Println("Failed to unmarshal data. Error:", err)
		return list, err
	}

	for _, v := range data {
		list = append(list, *v)
	}

	return list, nil
}
