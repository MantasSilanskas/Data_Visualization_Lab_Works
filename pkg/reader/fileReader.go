package reader

import (
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

var dataFilesPath = "../cmd/data/"

func InputFiles(filename string) ([]File, error) {

	var (
		err  error
		list []File
	)

	file, err := os.Open(dataFilesPath + filename)
	if err != nil {
		log.Println(err)
		return list, err
	}
	defer file.Close()

	files := []*File{}

	if err := gocsv.UnmarshalFile(file, &files); err != nil {
		log.Println(err)
		return list, err
	}

	for _, v := range files {
		list = append(list, *v)
	}

	return list, nil
}

func ReadFileData(inputFile string) ([]DeviceData, error) {

	list := []DeviceData{}

	file, err := os.Open(dataFilesPath + inputFile)
	if err != nil {
		log.Println(err)
		return list, err
	}
	defer file.Close()

	data := []*DeviceData{}

	if err := gocsv.UnmarshalFile(file, &data); err != nil {
		log.Println(err)
		return list, err
	}

	for _, v := range data {
		list = append(list, *v)
	}

	return list, nil
}
