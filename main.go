package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

const url string = "http://localhost:3000"

type user struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"mail"`
	Cred  string `json:"cred"`
	Token string `json:"token"`
	Link  string `json:"link"`
}

func addRow(u user) {
	db, err := sql.Open("sqlite3", "main.db")
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare("INSERT INTO users(name, mail, cred, token, link) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()
	defer db.Close()

	if u.Link == "" {
		_, err = stmt.Exec(u.Name, u.Email, u.Cred, u.Token, nil)
	} else {
		_, err = stmt.Exec(u.Name, u.Email, u.Cred, u.Token, u.Link)
	}

	if err != nil {
		log.Fatal(err)
	}

}

func dropRow(name string) {
	db, err := sql.Open("sqlite3", "main.db")
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare("DELETE FROM users WHERE mail = ?")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()
	defer db.Close()

	_, err = stmt.Exec(name)

	if err != nil {
		log.Fatal(err)
	}
}

func getRows() []user {
	db, err := sql.Open("sqlite3", "main.db")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("SELECT * FROM users")
	var users []user

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	defer db.Close()

	for rows.Next() {

		var id int
		var name string
		var mail string
		var cred string
		var token string
		var link sql.NullString

		err = rows.Scan(&id, &name, &mail, &cred, &token, &link)

		if err != nil {
			log.Fatal(err)
		}

		if link.Valid {
			var user = user{ID: fmt.Sprint(id), Name: name, Email: mail, Cred: cred, Token: token, Link: link.String}
			users = append(users, user)
		} else {
			var user = user{ID: fmt.Sprint(id), Name: name, Email: mail, Cred: cred, Token: token}
			users = append(users, user)
		}
	}
	return users
}

func getRowByName(name string) user {
	for _, a := range getRows() {
		if a.Email == name {
			return a
		}
	}
	return user{}
}

func showRows() {
	db, err := sql.Open("sqlite3", "main.db")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("SELECT * FROM users")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	defer db.Close()

	for rows.Next() {

		var id int
		var name string
		var mail string
		var cred string
		var token string
		var link sql.NullString

		err = rows.Scan(&id, &name, &mail, &cred, &token, &link)

		if err != nil {
			log.Fatal(err)
		}
		if link.Valid {
			fmt.Printf("%d %s %s %s %s %s\n", id, name, mail, cred, token, link.String)
		} else {
			fmt.Printf("%d %s %s %s %s\n", id, name, mail, cred, token)
		}
	}
}

func getUsers(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", url)
	c.IndentedJSON(http.StatusOK, getRows())
}

func postUsers(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		c.Header("Access-Control-Allow-Origin", url)
		return
	}

	addRow(newUser)
	c.Header("Access-Control-Allow-Origin", url)
	c.IndentedJSON(http.StatusCreated, newUser)
	c.Header("Access-Control-Allow-Origin", url+"login")
}

func getUserBy(c *gin.Context) {
	name := c.Param("name")
	for _, a := range getRows() {
		if a.Email == name {
			c.Header("Access-Control-Allow-Origin", url)
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.Header("Access-Control-Allow-Origin", url)
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func updateToken(c *gin.Context) {
	name := c.Param("name")
	token := c.Param("token")
	row := getRowByName(name)
	if row == (user{}) {
		c.Header("Access-Control-Allow-Origin", url)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}
	row.Token = token
	dropRow(name)
	addRow(row)
	c.Header("Access-Control-Allow-Origin", url)
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user token replaced"})
}

func setupdb() {
	db, err := sql.Open("sqlite3", "main.db")
	fmt.Printf("%T\n", db)
	if err != nil {
		log.Fatal(err)
	}

	sts := `CREATE TABLE if not exists users(id INTEGER PRIMARY KEY, name TEXT, mail TEXT, cred TEXT, token TEXT, link TEXT);`
	defer db.Close()
	_, err = db.Exec(sts)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Users has been initialized!")
}

func main() {
	setupdb()
	showRows()
	router := gin.Default()

	router.GET("/users", getUsers)
	router.GET("/users/:name", getUserBy)
	router.GET("/users/:name/:token", updateToken)
	router.POST("/users", postUsers)

	router.Run("localhost:8080")
	fmt.Println("____")
	fmt.Println("REST API has ended its rest!")
}
