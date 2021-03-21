package db

import (
	"time"

	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/pkg/reader"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Database name
var (
	database = "local"
)

type BSONFile struct {
	ID         primitive.ObjectID `bson:"_id"`
	FileNumber int                `bson:"number"`
	FileName   string             `bson:"name"`
	InputTime  time.Time          `bson:"inputTime"`
}

type BSONDeviceData struct {
	ID          primitive.ObjectID `bson:"_id"`
	DeviceID    string             `bson:"deviceId"`
	TimeStamp   time.Time          `bson:"timeStamp"`
	TimeStampMs int64              `bson:"timeStampMs"`
	DataType    string             `bson:"type"`
	DataValue   float32            `bson:"value"`
}

func FileInputToBSON(input reader.File) BSONFile {
	return BSONFile{
		ID:         primitive.NewObjectID(),
		FileNumber: input.Number,
		FileName:   input.Name,
		InputTime:  input.InputTime,
	}
}

func DevicesInputToBSON(input []reader.DeviceData) []BSONDeviceData {
	list := []BSONDeviceData{}
	for _, v := range input {
		list = append(list, DeviceInputToBSON(v))
	}
	return list
}

func DeviceInputToBSON(input reader.DeviceData) BSONDeviceData {
	return BSONDeviceData{
		ID:          primitive.NewObjectID(),
		DeviceID:    input.DeviceID,
		TimeStamp:   input.TimeStamp,
		TimeStampMs: input.TimeStampMs,
		DataType:    input.DataType,
		DataValue:   input.DataValue,
	}
}
