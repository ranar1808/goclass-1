package main

import "fmt"

// func sum(arrayInt []int, channelInt chan int, rank string) {
func sum(arrayInt []int, rank string)(total2 int) {
	fmt.Println(rank)
    total := 0
    for _, value := range arrayInt {
        total += value
    }
	fmt.Println("total :", total)
	total2 = total
    // channelInt <- total // send total to channelInt
}

func main() {
    arrayInt := []int{7, 2, 8, -9, 4, 0}

    channelInt := make(chan int, 1)
    // channelInt2 := make(chan int, 2)
    go channelInt <- sum(arrayInt[:len(arrayInt)/2],  "1")
    go sum(arrayInt[len(arrayInt)/2:],  "2")
    // go sum(arrayInt, channelInt2,"3")
    // result1, result2, result3 := <-channelInt, <-channelInt, <-channelInt2 // receive from channelInt
    result1, result2 := <-channelInt, <-channelInt  // receive from channelInt

	fmt.Println(result1, result2,result1+ result2 )
	// fmt.Println(result1, result2,  result3)
}
