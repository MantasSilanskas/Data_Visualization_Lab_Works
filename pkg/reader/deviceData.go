package reader

import "time"

type DeviceData struct {
	DeviceID    string    `csv:"device-id"`
	TimeStamp   time.Time `csv:"timestamp"`
	TimeStampMs int64     `csv:"timestamp-ms"`
	DataType    string    `csv:"type"`
	DataValue   float32   `csv:"value"`
}
