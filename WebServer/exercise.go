package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

type ManyStatus struct {
	ManyStatus Status `json:"status"`
}

func index(w http.ResponseWriter, r *http.Request) {
	jsonFile, err := os.Open("exercise.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("sukses buka file exercise.json")

	byteValue, _ := ioutil.ReadAll(jsonFile)

	defer jsonFile.Close()

	var mStatus ManyStatus

	json.Unmarshal([]byte(byteValue), &mStatus)

	t, err := template.ParseFiles(
		"exercise.html",
	)
	// t, err := template.New("exercise").ParseFiles("./exercise.html")
	if err != nil {
		log.Fatal(err)
	}
	// err = t.Execute(w, mStatus)
	err = t.ExecuteTemplate(w, "exercise.html", mStatus)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("sukses semua")
}

func random() {
	ms := &ManyStatus{
		Status{
			Water: rand.Intn(100),
			Wind:  rand.Intn(100),
		},
	}

	byteVal, err := json.Marshal(&ms)
	if err != nil {
		log.Fatalln(err)
	}

	err = ioutil.WriteFile("exercise.json", byteVal, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("sukses update")

}

func main() {

	go func() {
		for {
			select {
			case <-time.After(3 * time.Second):
				random()
			}
		}
	}()

	http.HandleFunc("/", index) // set router

	fmt.Println("Lagi jalan coy !")
	err := http.ListenAndServe(":1808", nil) // set listen port to 9090

	if err != nil {
		log.Fatal("Error running service: ", err)
	}
}
