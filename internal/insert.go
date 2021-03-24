package internal

import (
	"log"

	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/internal/db"
	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/internal/reader"
	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/internal/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

// InsertData inserts all files data to database which are in filesNames.
func InsertData(client *mongo.Client, filename string) error {

	files, err := removeDuplicateFiles(client, filename)
	if err != nil {
		log.Println("Failed to remove duplicates files from input. Error:", err)
	}

	if len(files) > 0 {
		for _, v := range files {
			rawDevicesData, err := reader.ReadFileData(v.Name)
			if err != nil {
				log.Println("Failed to read file data. Error:", err)
				return err
			}

			file := db.FileInputToBSON(v)
			devicesData := db.DevicesInputToBSON(rawDevicesData)

			err = db.AddDeviceData(*client, devicesData)
			if err != nil {
				log.Println("Failed to add file data to database. Error:", err)
				return err
			}

			err = db.AddFile(*client, file)
			if err != nil {
				log.Println("Failed to add file to database. Error:", err)
				return err
			}
		}
		log.Println("Data succesfully added to database!")
	} else {
		log.Println("There were no new data to add to database!")
	}

	return nil
}

func removeDuplicateFiles(client *mongo.Client, filename string) ([]reader.File, error) {

	var (
		inputFiles []reader.File
		files      []db.BSONFile
		err        error
	)
	list := []reader.File{}

	if files, err = db.FilterAllFiles(client); err != nil {
		log.Println("Failed to extracts files from database. Error:", err)
		return list, err
	}

	if inputFiles, err = reader.InputFiles(filename); err != nil {
		log.Println("Failed to read input files names list. Error:", err)
		return list, err
	}

	for _, v := range inputFiles {
		if utils.IsFileAlreadyInDatabase(v, files) {
			continue
		}
		list = append(list, v)
	}

	return list, err
}
