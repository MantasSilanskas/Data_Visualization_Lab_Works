package pkg

import (
	"log"

	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/pkg/database"
	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/pkg/reader"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertData(client *mongo.Client, filename string) error {
	filesNames, err := reader.InputFiles(filename)
	if err != nil {
		log.Println("failed to read input files names list", err)
		return err
	}

	for _, v := range filesNames {
		data, err := reader.ReadFileData(v.Name)
		if err != nil {
			log.Println("failed to read file data", err)
			return err
		}
		err = database.AddFile(*client, v)
		if err != nil {
			log.Println("failed to add file to database. Error:", err)
			return err
		}
		err = database.AddDeviceData(*client, data)
		if err != nil {
			log.Println("failed to add file data to database. Error:", err)
			return err
		}
	}

	return nil
}
