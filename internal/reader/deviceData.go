package reader

import "time"

// DeviceData defines what kind of data we receive from device
type DeviceData struct {
	DeviceID    string    `csv:"device-id"`
	TimeStamp   time.Time `csv:"timestamp"`
	TimeStampMs int64     `csv:"timestamp-ms"`
	DataType    string    `csv:"type"`
	DataValue   float32   `csv:"value"`
}
