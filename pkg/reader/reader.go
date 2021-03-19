package reader

import (
	"fmt"
	"log"
	"os"
)

func ReadFileData() {

	var (
		fileName string
	)

	log.Println("Enter file name:")
	fmt.Scanln(&fileName)

	in, err := os.Open("/" + fileName + ".csv")
	if err != nil {
		log.Println(err)
		return
	}
	defer in.Close()

	data := []*DeviceData{}

	if err := gocsv.UnmarshalFile(in, &data); err != nil {
		log.Println(err)
		return
	}

	for _, v := range data {
		v.
	}

}
