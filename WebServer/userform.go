package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method: ", r.Method) // get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("Username: ", r.Form["username"])
		fmt.Println("Password: ", r.Form["password"])

		if len(r.Form["username"][0]) == 0 {
			// code if field is empty
			fmt.Println("kosongs")
		}

		getint, err := strconv.Atoi(r.Form.Get("age"))
		if err != nil {
			// error when convert to number, it may not a number
			fmt.Println("gak ada angka")
		}
		fmt.Println(getint)

		if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get("Age")); !m {
			fmt.Println("gak ada angka")
		}
		if m, _ := regexp.MatchString("^[a-zA-Z]+$", r.Form.Get("name")); !m {
			fmt.Println("gak huruf")
		}

		if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, r.Form.Get("email")); !m {
			fmt.Println("salah email")
		}

		fruits := []string{"apple", "banana", "pear"}

		for _, v := range fruits {
			if v == r.Form.Get("fruit") {
				fmt.Println("ada buah")
				// return true
			}
		}
		// gender := []int{1, 2}

		// for _, v := range gender {
		// 	if v,_ == strconv.Atoi(r.Form.Get("gender")); v {
		// 		fmt.Println("ada gender")
		// 		// return true
		// 	}
		// }

		v := url.Values{}
		v.Set("name", "Adam")
		v.Add("hobby", "reading")
		v.Add("hobby", "cooking")
		v.Add("hobby", "coding")

		fmt.Println(v.Get("name"))
		fmt.Println(v.Get("hobby"))
		fmt.Println(v["hobby"])

		fmt.Println(v.Encode())
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(
		"userform.gtpl",
	)
	t.Execute(w, "")
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/login", login) // set router

	fmt.Println("Lagi jalan coy !")
	err := http.ListenAndServe(":1808", nil) // set listen port to 9090

	if err != nil {
		log.Fatal("Error running service: ", err)
	}
}
