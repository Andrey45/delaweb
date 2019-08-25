package main

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	_ "github.com/heroku/x/hmetrics/onload"
	"log"
	"net/http"
	"os"

	_ "fmt"
	"github.com/gin-contrib/cors"
	_ "github.com/jmoiron/sqlx/types"
	_ "github.com/lib/pq"
	_ "html/template"
	_ "net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin/render"
)

type use struct {
	Email string `json:"email"`
	Password string `json:"password"`
}
type User struct {
	Email string `json:"email"`
	Username string `json:"username"`
	Country string `json:"country"`
	City string `json:"city"`
}
var db *sql.DB
func main() {
	port := os.Getenv("PORT")
	r := gin.Default()
	r.Use(cors.Default())
	dab, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	db = dab
	defer dab.Close()
	r.LoadHTMLGlob("templates/*.html")
	r.Static("/assets", "./assets")
	r.GET("/", html)
	r.POST("/login", autorization)
	r.POST("/check", check)
	r.PUT("/userUpt", update)
	r.Run(":"+ port)
}
func html(c*gin.Context)  { // Возвращвет index.html
	c.HTML(http.StatusOK, "index.html", nil)
}
func update(c*gin.Context)  { // Изменение данных пользователя
	var json User
	c.BindJSON(&json)
	email:= json.Email
	username := json.Username
	country := json.Country
	city := json.City
	rows, err := db.Query("UPDATE users SET username='"+username+"', city='"+city+"', country='"+country+"' WHERE email='"+email+"'")
	if err != nil {
		c.JSON(204, gin.H{
			"error": true,
		})
	}
	defer rows.Close()
	// проверки по паролю нет так как запрос отпраляет только авторизованный пользователь
	// но это можно обойти зная что нужно изменять и зная email пользователя, допустим можно провернуть через postman
	row := db.QueryRow("SELECT email, username, country, city FROM users WHERE email='"+email+"'") //
	user := []User{}

	p := User{}
	error := row.Scan(&p.Email, &p.Username, &p.Country, &p.City)
	if error != nil {
		fmt.Println(error)
	}
	user = append(user, p)

	c.JSON(200, user)

}
func check(c*gin.Context)  { // регистрация пользователя
	var json use
	c.BindJSON(&json)
	email := json.Email
	pass := json.Password
	row:=db.QueryRow("SELECT count(*) FROM users where email='"+email+"'")
	var num int
	err := row.Scan(&num)
	if num > 0 { // если пользовтель существует возвращаем ошибку
		c.JSON(302, gin.H{
			"error": true,
		})
	} else if err != nil{
		c.JSON(501, gin.H{
			"error": true,
		})
	} else {
		rows, err :=db.Query("insert into users(email, password) VALUES ('"+email+"','"+pass+"')")
		if err != nil {
			c.JSON(204, gin.H{
				"error": true,
			})
		}
		c.JSON(200, gin.H{
			"ok": true,
		})
		defer rows.Close()
	}
}
func autorization(c*gin.Context)  { // авторизация
	var json use
	c.BindJSON(&json)
	email := json.Email
	pass := json.Password
	row:=db.QueryRow("SELECT count(*) FROM users where email='"+email+"'")

	var num int
	err := row.Scan(&num)
	if num > 0 { // проверяем что пользоваетль зарегистрирован
		rows:= db.QueryRow("SELECT email, username, country, city FROM users WHERE email='"+email+"' and password='"+pass+"'")
		user := []User{}

		p := User{}
		err := rows.Scan(&p.Email, &p.Username, &p.Country, &p.City)
		if err != nil {
			fmt.Println(err)
		}
		user = append(user, p)

		c.JSON(200, user)

	} else if err != nil {
		c.JSON( 501, gin.H{
			"error": true,
		})
	} else {
		c.JSON( 404, gin.H{
			"error": true,
		})
	}
}
