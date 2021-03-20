package pkg

import (
	"log"

	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/pkg/database"
	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/pkg/reader"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// InsertData inserts all files data to database which are in filesNames.
func InsertData(client *mongo.Client, filename string) error {

	files, err := RemoveDuplicateFiles(filename)
	if err != nil {
		log.Println("Failed to remove duplicates files from input", err)
	}

	for _, v := range files {
		rawDevicesData, err := reader.ReadFileData(v.Name)
		if err != nil {
			log.Println("failed to read file data", err)
			return err
		}

		file := database.FileInputToBSON(v)
		devicesData := database.DevicesInputToBSON(rawDevicesData)

		err = database.AddDeviceData(*client, devicesData)
		if err != nil {
			log.Println("failed to add file data to database. Error:", err)
			return err
		}

		err = database.AddFile(*client, file)
		if err != nil {
			log.Println("failed to add file to database. Error:", err)
			return err
		}
	}

	log.Println("Succesfully added data to database")

	return nil
}

func RemoveDuplicateFiles(client *mongo.Client, filename string) ([]reader.File, error) {

	var (
		inputFiles []reader.File
		files      []database.BSONFile
		err        error
	)
	list := []reader.File{}

	if files, err = database.FilterAll(client, bson.M{}); err != nil {
		log.Println("failed to extracts files from database", err)
		return list, err
	}

	if inputFiles, err = reader.InputFiles(filename); err != nil {
		log.Println("failed to read input files names list", err)
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
