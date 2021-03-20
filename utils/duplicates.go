package pkg

import (
	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/pkg/database"
	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/pkg/reader"
)

// IsFileAlreadyInDatabase checks if certain file already exist in database
func IsFileAlreadyInDatabase(input reader.File, files []database.BSONFile) bool {

	for _, v := range files {
		if input.Name == v.FileName {
			return true
		}
	}
	return false
}
