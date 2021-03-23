package results

import (
	"log"

	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/internal/db"
	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/internal/filter"
	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CalcResults(input []string, client *mongo.Client) (results []utils.CalculatedData, err error) {
	list := []utils.CalculatedData{}

	for _, v := range input {
		data, err := db.FilterDevicesByID(client, bson.M{"deviceId": v})
		if err != nil {
			log.Println("Failed to extract", v, "data from database. Error:", err)
			return list, err
		}
		filteredData, err := filter.FilterDeviceData(data, v)
		if err != nil {
			log.Println("Failed to filter out device data by type. Error:", err)
			return list, err
		}

		results := utils.CalcDeviceData(filteredData)

		list = append(list, results)
	}

	return list, nil
}
