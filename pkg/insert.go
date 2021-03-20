package pkg

import (
	"log"

	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/pkg/database"
	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/pkg/reader"
	"go.mongodb.org/mongo-driver/mongo"
)

// InsertData inserts all files data to database which are in filesNames.
func InsertData(client *mongo.Client, filename string) error {
	filesNames, err := reader.InputFiles(filename)
	if err != nil {
		log.Println("failed to read input files names list", err)
		return err
	}

	for _, v := range filesNames {
		rawDevicesData, err := reader.ReadFileData(v.Name)
		if err != nil {
			log.Println("failed to read file data", err)
			return err
		}
		file := database.FileInputToBSON(v)
		err = database.AddFile(*client, file)
		if err != nil {
			log.Println("failed to add file to database. Error:", err)
			return err
		}
		devicesData := database.DevicesInputToBSON(rawDevicesData)
		err = database.AddDeviceData(*client, devicesData)
		if err != nil {
			log.Println("failed to add file data to database. Error:", err)
			return err
		}
	}
	log.Println("Succesfully added data to database")
	return nil
}
