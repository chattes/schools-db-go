package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/chattes/schools-db-go/utils"

	"github.com/chattes/schools-db-go/database"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Starting DB Load...")

	err := godotenv.Load()

	if err != nil {
		panic("Cannot read env file")
	}

	filePath := os.Getenv("FILE_PATH")
	asbPath, err := filepath.Abs(filePath)
	if err != nil {
		fmt.Println("Error finding path")
	}

	jsonFile, err := os.Open(asbPath)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer jsonFile.Close()

	schoolsBytes, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	var result utils.AllSchools

	err = json.Unmarshal([]byte(schoolsBytes), &result)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Saveing data")

	_, err = saveData(result)

	if err != nil {
		fmt.Println("An error occured")
	}
	fmt.Println("All good...")

}

func saveData(data utils.AllSchools) (success bool, err error) {

	db := new(database.MySql)

	db.DropDB("schools")
	_, err = db.CreateSchema("schools")

	if err != nil {
		panic(err)
	}

	db.CreateTable("schools", "school_info")

	db.InsertValues(data)

	return true, nil

}
