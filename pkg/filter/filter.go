package filter

import "github.com/MantasSilanskas/Data_Visualization_Lab_Works/pkg/db"

const (
	Humidity    = "humidity"
	Temperature = "temperature"
	Co2         = "CO2"
)

// FilterDeviceData filters out devide data by type.
func FilterDeviceData(data []db.BSONDeviceData, id string) (FilteredData, error) {

	list := FilteredData{}

	list.DeviceID = id

	for _, v := range data {
		switch v.DataType {
		case Humidity:
			list.Humidity = append(list.Humidity, HumidityData{
				v.DataType,
				v.DataValue,
				v.TimeStamp,
			})
		case Temperature:
			list.Temperature = append(list.Temperature, TemperatureData{
				v.DataType,
				v.DataValue,
				v.TimeStamp,
			})
		case Co2:
			list.Co2 = append(list.Co2, Co2Data{
				v.DataType,
				v.DataValue,
				v.TimeStamp,
			})
		}
	}

	return list, nil
}
