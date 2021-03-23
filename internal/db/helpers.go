package db

import (
	"log"
	"path/filepath"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
)

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
