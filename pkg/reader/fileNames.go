package reader

import "time"

type File struct {
	Number    int    `csv:"file_number"`
	Name      string `csv:"file_name"`
	InputTime time.Time
}
