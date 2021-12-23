package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

/*
	Before first usage of the program make sure to install all dependencies with the command "go get" (dependencies will be resolved using the versions declared in the go.mod file).
	Also make sure to run the createscript.sql file and edit the credentials bellow to match your configuration.

	For security reasons it is HIGHLY recommended to create a new sql user and only give it the required permissions, especially when running this API on public webservers.
*/

const (
	serverAddress = "127.0.0.1:8080" // example value
	//sQLServerAddress  = "127.0.0.1:10000"							// example value, if you want support for a different sql server address add it to the dbConnectionString
	sQLUsername     = "giftcode_api"                                // example value
	sQLPassword     = "6IXTcEH9diGlDcnq89lSRm7ZIJUsH1c+AfHtWO6tn+0" // example value
	sQLDatabaseName = "giftcode_api_db"                             // example value
	giftcodeID      = 1
)

type Giftcode struct {
	CODE string `json:"giftcode"`
}

func main() {

	router := gin.Default()

	router.GET("/giftcode", getGiftcode)

	router.Run(serverAddress) // Listen and serve on server address
}

func getGiftcode(c *gin.Context) {
	CODE := queryDB()
	c.JSON(200, CODE)
}

/* Returns the result of a static query */
func queryDB() string {

	// set return and configuration values
	var result string
	dBConnectionString := fmt.Sprintf("%s:%s@/%s", sQLUsername, sQLPassword, sQLDatabaseName) // Set the database credentials

	db, err := sql.Open("mysql", dBConnectionString)
	if err != nil {
		panic(err.Error()) // Proper error handling instead of panic
	}

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	fmt.Printf("[DEBUG]: db ping result (should be <nil>): %v\n", err)
	if err != nil {
		panic(err.Error()) // Proper error handling instead of panic
	}

	fmt.Println("[DEBUG]: Connected to database.")

	/* Database query */

	/*
		The database table named "giftcode" looks like this:

		+------------+--------------+------+-----+---------------------+-------------------------------+
		| Field      | Type         | Null | Key | Default             | Extra                         |
		+------------+--------------+------+-----+---------------------+-------------------------------+
		| id         | int(11)      | NO   | PRI | NULL                | auto_increment                |
		| giftcode   | varchar(255) | NO   |     | NULL                |                               |
		| origin_url | varchar(255) | YES  |     | NULL                |                               |
		| reg_date   | timestamp    | NO   |     | current_timestamp() | on update current_timestamp() |
		+------------+--------------+------+-----+---------------------+-------------------------------+
	*/

	/* Variable block to store some query values */
	var (
		id       int
		giftcode string
	)

	rows, err := db.Query("select id, giftcode from giftcode where id = ?", 1) // This value is static. If you need a dynamic query implement it above and pass the resulting id variable here
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &giftcode)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, giftcode)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	/* Update return value with query result */
	result = giftcode

	/* Close database connection */
	defer db.Close() // This is optional as the connection is also closed on program exit
	fmt.Println("[DEBUG]: Database connection closed.")

	/* Return query result */
	fmt.Printf("[DEBUG]: Sending giftcode \"%v\"...\n", result)
	return result
}
