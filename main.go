package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
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

	// Get the base file from S3

	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2")},
	)

	svc := s3.New(sess)
	bucketName := os.Getenv("BUCKET_NAME")
	file := os.Getenv("FILE_NAME")
	response, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: &bucketName,
		Key:    &file,
	})

	if err != nil {
		panic("Unable to read file from S3...")
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic("Error occured reading file contents")
	}

	var result utils.AllSchools

	err = json.Unmarshal([]byte(body), &result)

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
