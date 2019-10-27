package main

import (
	"fmt"
	"os"
	"github.com/pkg/errors"
	"bufio"
	"strings"
	"flag"
)

type Student struct {
	Nama string
    Alamat string
    Pekerjaan string
    Alasan string
}

func main(){
	getError();

	scanner2();
	
	GetID := os.Args[1]

	ReadArgs(GetID)
}
func scanner2(){
	nama := flag.String("nama","ibob","namanya")
	alamat := flag.String("alamat","depok","alamatnya")
	pekerjaan := flag.String("pekerjaan","developer","pekrjaannya")
	alasan := flag.String("alasan","belajar","alasannya")

	flag.Parse()

	fmt.Println("Nama : ",*nama)
	fmt.Println("Alamat : ",*alamat)
	fmt.Println("Pekerjaan : ",*pekerjaan)
	fmt.Println("Alasan : ",*alasan)
}
func scanner() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Input")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("hi", text) == 0 {
			fmt.Println("hello, Yourself")
		}
	}
}
func getError(){	
	err := errors.Wrap(func() error {
		return func() error {
			return errors.Errorf("hello %s", fmt.Sprintf("world"))
			}()
		}(), "failed")

		fmt.Printf("%v \n \n", err)
}
func ReadArgs (pram string){	
    student1 := Student{"Andi", "Jakarta", "Developer","Belajar"}
    student2 := Student{"Yoga", "Depok", "Software Engineer","Pen Belajar"}
    student3 := Student{"Patrick", "Bogor","Backend Developer","Aku mau belajar"}
    student4 := Student{"Arhian", "Bekasi", "Data Engineer","Sama kayak di atas gan"}
    student5 := Student{"Ibob","Kalimantan", "Fullstack Developer","ngikut aja"}
	student6 := Student{"Ojan","Tanggerang","Data Analys","serah lah"}

	if pram == "1"{
		fmt.Println("Namanya adalah ", student1.Nama)
		fmt.Println("Alamatnya adalah ", student1.Alamat)
		fmt.Println("Pakerjaannya adalah ", student1.Pekerjaan)
		fmt.Println("Alasannya adalah ", student1.Alasan)
	}else if pram == "2"{
		fmt.Println("Namanya adalah ", student2.Nama)
		fmt.Println("Alamatnya adalah ", student2.Alamat)
		fmt.Println("Pakerjaannya adalah ", student2.Pekerjaan)
		fmt.Println("Alasannya adalah ", student2.Alasan)
	}else if pram == "3"{
		fmt.Println("Namanya adalah ", student3.Nama)
		fmt.Println("Alamatnya adalah ", student3.Alamat)
		fmt.Println("Pakerjaannya adalah ", student3.Pekerjaan)
		fmt.Println("Alasannya adalah ", student3.Alasan)
	}else if pram == "4"{
		fmt.Println("Namanya adalah ", student4.Nama)
		fmt.Println("Alamatnya adalah ", student4.Alamat)
		fmt.Println("Pakerjaannya adalah ", student4.Pekerjaan)
		fmt.Println("Alasannya adalah ", student4.Alasan)
	}else if pram == "5"{
		fmt.Println("Namanya adalah ", student5.Nama)
		fmt.Println("Alamatnya adalah ", student5.Alamat)
		fmt.Println("Pakerjaannya adalah ", student5.Pekerjaan)
		fmt.Println("Alasannya adalah ", student5.Alasan)
	}else if pram == "6"{
		fmt.Println("Namanya adalah ", student6.Nama)
		fmt.Println("Alamatnya adalah ", student6.Alamat)
		fmt.Println("Pakerjaannya adalah ", student6.Pekerjaan)
		fmt.Println("Alasannya adalah ", student6.Alasan)
	}else{
		fmt.Println("kelebihan cuy")
	}
}