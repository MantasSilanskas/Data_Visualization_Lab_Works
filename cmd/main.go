package main

import (
	"log"

	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/pkg"
	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/pkg/db"
	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/pkg/filter"
	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/utils"
	"go.mongodb.org/mongo-driver/bson"
)

const filename = "filesNames.csv"

func main() {

	client, err := db.Connection()
	if err != nil {
		log.Println("Failed to connect to database. Error:", err)
		return
	}

	err = pkg.InsertData(client, filename)
	if err != nil {
		log.Println("Failed to insert data to database. Error:", err)
		return
	}

	names, err := utils.UniqueDevicesIDs(client)
	if err != nil {
		log.Println("Failed to extract file names from database. Error:", err)
	}

	for _, v := range names {
		data, err := db.FilterDevicesByID(client, bson.M{"deviceId": v})
		if err != nil {
			log.Println("Failed to extract", v, "data from database. Error:", err)
			return
		}
		filteredData, err := filter.FilterDeviceData(data, v)
		if err != nil {
			log.Println("Failed to filter out device data by type. Error:", err)
			return
		}

		results := utils.CalcDeviceData(filteredData)

		log.Println(results)
	}

}
