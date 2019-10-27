package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Users struct {
	Users []User `json:"users"`
}

// User struct which contains a name
// a type and a list of social links
type User struct {
	Name   string `json:"name"`
	Job    string `json:"job"`
	Age    int    `json:"age"`
	Social Social `json:"social"`
}
type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

type ManyStatus struct {
	ManyStatus []Status `json:"status"`
	// WindStatus  string
	// WaterStatus string
}

// Social struct which contains a
// list of links
type Social struct {
	Instagram string `json:"instagram"`
	Twitter   string `json:"twitter"`
}

func index(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "It's Works!!") // send data to client side
	ibob := struct {
		Name string
		Umur int
	}{"Ibob", 22}
	tmplt, err := template.New("index").Parse("Nama saya {{.Name}} dan Umur nya {{.Umur}} tahun")

	if err != nil {
		log.Fatal(err)
	}

	err = tmplt.Execute(w, ibob)

	if err != nil {
		log.Fatal(err)
	}
}

func html(w http.ResponseWriter, r *http.Request) {
	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	// t, err := template.New("index").Parse(tpl)
	var t = template.Must(template.ParseFiles(
		"base.gtpl",
		"Header.gtpl",
		"Body.gtpl",
		"Footer.gtpl",
	))

	// check(err)

	data := struct {
		Title string
		Items []string
		Posts string
	}{
		Title: "My Page",
		Items: []string{
			"My photos",
			"My blog",
			"My Personal",
			"My Testo",
		},
		Posts: "test posts",
	}

	var err = t.Execute(w, data)
	check(err)
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "testo!!") // send data to client side
}

func student(w http.ResponseWriter, r *http.Request) {
	jsonFile, err := os.Open("student.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("sukses buka file student.json")

	byteValue, _ := ioutil.ReadAll(jsonFile)

	defer jsonFile.Close()

	var users Users

	json.Unmarshal([]byte(byteValue), &users)

	t, err := template.ParseFiles(
		"student.html",
	)
	t.ExecuteTemplate(w, "student.html", users)
	// t.Execute(w, users)

}

func exercise(w http.ResponseWriter, r *http.Request) {
	jsonFile, err := os.Open("exercise.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("sukses buka file exercise.json")

	byteValue, _ := ioutil.ReadAll(jsonFile)

	defer jsonFile.Close()

	var status []ManyStatus

	json.Unmarshal([]byte(byteValue), &status)

	t, err := template.ParseFiles(
		"exercise.html",
	)
	if err != nil {
		log.Fatal(err)
	}
	t.ExecuteTemplate(w, "exercise.html", status)
}

func main() {
	http.Handle("/assets/", http.FileServer(http.Dir("."))) //serve other files in assets dir
	http.HandleFunc("/", index)                             // set router
	http.HandleFunc("/html", html)                          // set router
	http.HandleFunc("/test", test)                          // set router
	http.HandleFunc("/student", student)                    // set router
	http.HandleFunc("/Exercise", exercise)                  // set router

	fmt.Println("Lagi jalan coy !")
	err := http.ListenAndServe(":1808", nil) // set listen port to 9090

	if err != nil {
		log.Fatal("Error running service: ", err)
	}
}
