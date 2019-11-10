package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/go-sessions"
	// "os"
)

var db *sql.DB
var err error

type queryAll map[string]interface{}

type user struct {
	ID        int
	Username  string
	FirstName string
	LastName  string
	Password  string
	Role      string
}

func connect_db() {
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1)/go-login")

	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
}

func routes() {
	http.HandleFunc("/", home)
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/views", views)
	http.HandleFunc("/edit", edit)
	http.HandleFunc("/delete", delete)
}

func main() {
	connect_db()
	routes()

	defer db.Close()

	fmt.Println("Server running on port :8000")
	http.ListenAndServe(":8000", nil)
}

func checkErr(w http.ResponseWriter, r *http.Request, err error) bool {
	if err != nil {

		fmt.Println(r.Host + r.URL.Path)

		http.Redirect(w, r, r.Host+r.URL.Path, 301)
		return false
	}

	return true
}

func QueryUser(username string) user {
	var users = user{}
	err = db.QueryRow(`
		SELECT id, 
		username, 
		first_name, 
		last_name, 
		password,
		role
		FROM users 
		WHERE username=?
		`, username).
		Scan(
			&users.ID,
			&users.Username,
			&users.FirstName,
			&users.LastName,
			&users.Password,
			&users.Role,
		)
	return users
}
func QueryUserWithID(id string) user {
	var users = user{}
	err = db.QueryRow(`
		SELECT id, 
		username, 
		first_name, 
		last_name, 
		password,
		role
		FROM users 
		WHERE id=?
		`, id).
		Scan(
			&users.ID,
			&users.Username,
			&users.FirstName,
			&users.LastName,
			&users.Password,
			&users.Role,
		)
	return users
}

func home(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	if len(session.GetString("username")) == 0 {
		http.Redirect(w, r, "/login", 301)
	}

	var users []user
	var rows *sql.Rows

	if session.GetString("role") == "admin" {
		rows, _ = db.Query("select * from users")
	} else {
		rows, _ = db.Query("select * from users where username =?", session.GetString("username"))
	}

	defer rows.Close()

	for rows.Next() {
		var each = user{}

		err = rows.Scan(&each.ID, &each.Username, &each.FirstName, &each.LastName, &each.Password, &each.Role)
		checkErr(w, r, err)

		users = append(users, each)
	}

	data := queryAll{
		"username": session.GetString("username"),
		"message":  "Welcome to the Go !",
		"users":    users,
	}

	var t, err = template.ParseFiles("views/home.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	t.Execute(w, data)
	return
}

func views(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	if len(session.GetString("username")) == 0 {
		http.Redirect(w, r, "/login", 301)
	}

	id := r.URL.Query().Get("id")

	var users = QueryUserWithID(id)

	var t, err = template.ParseFiles("views/views.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	t.Execute(w, users)
	return
}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.ServeFile(w, r, "views/register.html")
		return
	}

	username := r.FormValue("email")
	first_name := r.FormValue("first_name")
	last_name := r.FormValue("last_name")
	password := r.FormValue("password")
	role := r.FormValue("role")

	users := QueryUser(username)

	if (user{}) == users {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

		if len(hashedPassword) != 0 && checkErr(w, r, err) {
			stmt, err := db.Prepare("INSERT INTO users SET username=?, password=?, first_name=?, last_name=?, role=?")
			if err == nil {
				_, err := stmt.Exec(&username, &hashedPassword, &first_name, &last_name, &role)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}
		}
	} else {
		http.Redirect(w, r, "/register", 302)
	}
}

func edit(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	if len(session.GetString("username")) == 0 {
		http.Redirect(w, r, "/login", 301)
	}

	if r.Method != "POST" {
		id := r.URL.Query().Get("id")
		users := QueryUserWithID(id)
		var t, err = template.ParseFiles("views/views.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		t.Execute(w, users)
		return
	}

	id := r.FormValue("uid")
	username := r.FormValue("email")
	first_name := r.FormValue("first_name")
	last_name := r.FormValue("last_name")
	password := r.FormValue("password")
	role := r.FormValue("role")

	// users := QueryUser(username)

	if len(password) > 0 {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

		if len(hashedPassword) != 0 && checkErr(w, r, err) {
			stmt, err := db.Prepare("UPDATE users SET username=?, password=?, first_name=?, last_name=?, role=? where id=?")
			if err == nil {
				_, err := stmt.Exec(&username, &hashedPassword, &first_name, &last_name, &role, &id)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				http.Redirect(w, r, "/home", http.StatusSeeOther)
				return
			}
		}
	} else {
		stmt, err := db.Prepare("UPDATE users SET username=?, first_name=?, last_name=?, role=? where id=?")
		if err == nil {
			_, err := stmt.Exec(&username, &first_name, &last_name, &role, &id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return
		}
	}
}

func delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	delForm, err := db.Prepare("DELETE FROM users WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(id)
	http.Redirect(w, r, "/home", 301)
}

func login(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	if len(session.GetString("username")) != 0 && checkErr(w, r, err) {
		http.Redirect(w, r, "/", 302)
	}
	if r.Method != "POST" {
		http.ServeFile(w, r, "views/login.html")
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")

	users := QueryUser(username)

	//deskripsi dan compare password
	var password_tes = bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(password))

	if password_tes == nil {
		//login success
		session := sessions.Start(w, r)
		session.Set("username", users.Username)
		session.Set("name", users.FirstName)
		session.Set("role", users.Role)
		http.Redirect(w, r, "/", 302)
	} else {
		//login failed
		http.Redirect(w, r, "/login", 302)
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	session.Clear()
	sessions.Destroy(w, r)
	http.Redirect(w, r, "/", 302)
}
