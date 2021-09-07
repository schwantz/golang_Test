package MysqlTest

import (
	"fmt"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host     = "127.0.0.1:3306"
	database = "bank"
	user     = "root"
	password = "123hshh"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func MysqlTest() {

	fmt.Printf("----------------------------------------------------------\n")

	// Initialize connection string.
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s)/%s?allowNativePasswords=true", user, password, host, database)

	// Initialize connection object.
	db, err := sql.Open("mysql", connectionString)
	checkError(err)
	defer db.Close()
	err = db.Ping()
	checkError(err)
	fmt.Println("Successfully created connection to database.")

	// Variables for printing column data when scanned.
	var (
		UserID  int32
		RegTime string
		balance float64
	)

	// Read some data from the table.
	rows, err := db.Query("SELECT UserID, RegTime, balance from tb_userinfo limit 10;")
	checkError(err)
	defer rows.Close()
	fmt.Println("Reading data:")
	for rows.Next() {
		err := rows.Scan(&UserID, &RegTime, &balance)
		checkError(err)
		fmt.Printf("Data row = (%d, %s, %f)\n", UserID, RegTime, balance)
	}
	err = rows.Err()
	checkError(err)
	fmt.Println("Done.")

}
