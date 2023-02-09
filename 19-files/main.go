package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main(){
	fmt.Println("Handling files in golang..")
	content := "This needs to go in a file - Learncodeonline"

	file, err := os.Create("./mylcogofile.txt")

	checkNilErr(err)
	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("Length is: ", length)
	defer file.Close()
	readFile("./mylcogofile.txt")
}

func readFile(filename string){
	contentByte, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("Text data in file:\n", string(contentByte))
}

func checkNilErr(err error){
	if err != nil{
		panic(err)
	}
}