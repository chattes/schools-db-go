package database

import (
	"database/sql"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/chattes/schools-db-go/utils"

	_ "github.com/go-sql-driver/mysql"
)

type MySql struct{}

func getConnString() string {

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	conn_string := fmt.Sprintf("%s:%s@tcp(%s:%s)/", user, pass, host, port)
	fmt.Println(conn_string)
	return conn_string

}

func (db *MySql) CreateSchema(name string) (success bool, err error) {
	conn, err := sql.Open("mysql", getConnString())

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	_, err = conn.Exec("CREATE DATABASE " + name)

	if err != nil {
		panic(err)
	}

	return true, nil

}

func (db *MySql) CreateTable(name string, tableName string) {

	create_table_command := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s 
	( id INT NOT NULL AUTO_INCREMENT,
		name VARCHAR(200) NOT NULL,
		school_id INT,
		type VARCHAR(30),
		is_catholic BOOLEAN,
		language VARCHAR(20),
		level VARCHAR(20),
		city VARCHAR(60),
		city_slug VARCHAR(100),
		board VARCHAR(10),
		fraser_rating FLOAT,
		eqao_rating FLOAT,
		address varchar(200),
		grades varchar(20),
		website varchar(200),
		phone_number varchar(20),
		latitude DOUBLE,
		longitude DOUBLE,
		PRIMARY KEY (id),
		FULLTEXT KEY (name, address)
	)
	`, tableName)

	fmt.Println(create_table_command)
	conn, err := sql.Open("mysql", getConnString())

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	_, err = conn.Exec("USE " + name)

	if err != nil {
		panic(err)
	}

	_, err = conn.Exec(create_table_command)

	if err != nil {
		panic(err)
	}

}
func (db *MySql) DropDB(name string) {
	conn, err := sql.Open("mysql", getConnString())

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	_, err = conn.Exec("DROP DATABASE IF EXISTS " + name)

	if err != nil {
		panic(err)
	}

}
func (db *MySql) DropTable() (success bool, err error) {
	return true, nil

}

func (db *MySql) InsertValues(values utils.AllSchools) {
	start := time.Now()
	defer func() {
		duration := time.Since(start)
		fmt.Printf("Insert values %d milliseconds \n", duration.Milliseconds())
	}()
	batch := []utils.AllSchools{}
	chunk := []utils.School{}
	for i, v := range values {
		chunk = append(chunk, v)
		if i%500 == 0 {
			batch = append(batch, chunk)
			chunk = nil
		}
	}
	batch = append(batch, chunk)

	wg := &sync.WaitGroup{}
	wg.Add(len(batch))

	for _, batchRecs := range batch {
		go func(wg *sync.WaitGroup, chunk utils.AllSchools) {
			defer wg.Done()

			valStr := ""
			for _, data := range chunk {
				valStr += fmt.Sprintf(
					"(\"%s\", %d, \"%s\", %t, \"%s\", \"%s\", \"%s\", \"%s\", %.2f, %.2f, \"%s\", \"%s\", \"%s\", \"%s\", %f, %f),", data.Name, data.SchoolId, data.Type, data.IsCatholic, data.Language, data.City, data.CitySlug, data.Board, data.FraserRating, data.EQAORating, data.Address, data.Grades, data.Website, data.PhoneNumber, data.Latitude, data.Longitude)
			}

			preparedStatement := fmt.Sprintf("INSERT INTO school_info(name, school_id, type, is_catholic, language, city, city_slug, board, fraser_rating, eqao_rating, address, grades, website, phone_number, latitude, longitude) values %s", valStr)
			preparedStatement = preparedStatement[0 : len(preparedStatement)-1]
			conn, err := sql.Open("mysql", getConnString())
			conn.Exec("USE schools")
			if err != nil {
				panic(err)
			}

			defer conn.Close()

			_, err = conn.Exec(preparedStatement)

			if err != nil {
				panic(err)
			}

		}(wg, batchRecs)

	}

	wg.Wait()

	fmt.Println("All Records have been written succesfully")

}
