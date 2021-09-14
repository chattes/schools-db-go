package database

import "fmt"

type MySql struct{}

func (db *MySql) CreateSchema() (success bool, err error) {

	return true, nil

}

func (db *MySql) CreateTable() (success bool, err error) {

	return true, nil

}
func (db *MySql) DropTable() (success bool, err error) {

	return true, nil

}

func (db *MySql) InsertValues(values int) (success bool, err error) {

	return true, nil

}

func TestMyFunc() {
	fmt.Println("Test")
}
