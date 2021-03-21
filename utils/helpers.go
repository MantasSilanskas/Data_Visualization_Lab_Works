package utils

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/pkg/db"
	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/pkg/reader"
	"go.mongodb.org/mongo-driver/mongo"
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

// UniqueDevicesIDs returns all unique devices ID list
func UniqueDevicesIDs(client *mongo.Client) ([]string, error) {

	var (
		files []db.BSONFile
		err   error
	)
	list := []string{}

	if files, err = db.FilterAllFiles(client); err != nil {
		log.Println("Failed to extracts files from database. Error:", err)
		return list, err
	}

	for _, v := range files {
		ID := strings.TrimSuffix(v.FileName, filepath.Ext(v.FileName))
		list = append(list, ID)
	}

	return list, nil
}
