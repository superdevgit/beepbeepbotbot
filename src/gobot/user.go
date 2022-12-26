package main

import (
    "os"
    "fmt"
    "io/ioutil"
    "encoding/json"
)

// Get user 
func GetUser(chatId int64) string {
    chatIdStr := fmt.Sprintf("%d", chatId)
    result := make(map[string]string)
    var _, err = os.Stat("users.json")
    
    // create file if not exists
    if os.IsNotExist(err) {
        var file, err = os.Create("users.json")
        if err != nil {
            fmt.Println(err)
        }
        defer file.Close()
    }
    
    // open the users file
    jsonFile, err := os.Open("users.json")
    if err != nil {
        fmt.Println(err)
    }
    defer jsonFile.Close()
    
    byteValue, _ := ioutil.ReadAll(jsonFile)
    json.Unmarshal([]byte(byteValue), &result)

    r := result[chatIdStr]
    if r != "" {
        if v := result[chatIdStr]; v != ""{
            return string(v)
        }
    }

    return ""
}

// Set user
func SetUser(chatId int64, name string) error {
    // open the users file
    jsonFile, err := os.Open("users.json")
    if err != nil {
        fmt.Println(err)
    }
    defer jsonFile.Close()

    chatIdStr := fmt.Sprintf("%d", chatId)
    result := make(map[string]string)

    byteValue, _ := ioutil.ReadAll(jsonFile)
    json.Unmarshal([]byte(byteValue), &result)
    result[chatIdStr] = name
    file, _ := json.Marshal(result)
    err = ioutil.WriteFile("users.json", file, 0777)
    if err != nil {
        return err
    }

    return nil
}