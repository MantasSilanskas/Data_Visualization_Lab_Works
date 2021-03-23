package utils

import (
	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/internal/db"
	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/internal/reader"
)

// IsFileAlreadyInDatabase checks if certain file already exist in database
func IsFileAlreadyInDatabase(input reader.File, files []db.BSONFile) bool {

	for _, v := range files {
		if input.Name == v.FileName {
			return true
		}
	}
	return false
}
