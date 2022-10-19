package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

func main() {

    // Open our jsonFile
    jsonFile, err := os.Open("NAME.json")
    // if we os.Open returns an error then handle it
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("Successfully Opened .json")
    // defer the closing of our jsonFile so that we can parse it later on
    defer jsonFile.Close()

    byteValue, _ := ioutil.ReadAll(jsonFile)

    var result struct{}
    json.Unmarshal([]byte(byteValue), &result)

    fmt.Println(result["users"])

}