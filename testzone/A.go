 package main

 import (
 	"fmt"
 	"io/ioutil"
 	"os"
 	"strings"
 )

 func main() {

 	var fileName = "fileName"
 	fileBytes, err := ioutil.ReadFile(fileName)

 	if err != nil {
 		fmt.Println(err)
 		os.Exit(1)
 	}

 	sliceData := strings.Split(string(fileBytes), "\n")

 	fmt.Println(sliceData[1])
 }