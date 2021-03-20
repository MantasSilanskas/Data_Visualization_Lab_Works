package main

import (
	"log"

	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/pkg/database"
	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/pkg/reader"
)

func main() {

	_, err := database.Connection()
	if err != nil {
		log.Println("failed to connect to database", err)
		return
	}

	filename := "filesNames.csv"
	filesNames, err := reader.InputFiles(filename)
	if err != nil {
		log.Println("failed to read input files names list", err)
		return
	}

	for _, v := range filesNames {
		_, err = reader.ReadFileData(v.Name)
		if err != nil {
			log.Println("failed to read file data", err)
			return
		}
	}

}
