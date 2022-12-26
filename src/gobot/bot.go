package main

import (
    "os"
    "fmt"
    "bytes"
    "errors"
    "strings"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

// Handle bot requests
func BotHandler(res http.ResponseWriter, req *http.Request) {
    body := &WebhookReqBody{}
    if err := json.NewDecoder(req.Body).Decode(body); err != nil {
        fmt.Println("could not decode request body", err)
        return
    }

    // get user name if chat before
    user := GetUser(body.Message.Chat.ID)
    if user == "" {
        if strings.Contains(strings.ToLower(body.Message.Text), "@") {
            SetUserName(body)
        } else {
            GetUserName(body)
        }
    } else {
        if strings.Contains(strings.ToLower(body.Message.Text), "#") {
            GetWeatherAndNews(body)
        } else {
            AskLocation(body, user)
        }
    }
}

// Get weather and news of entered location
func GetWeatherAndNews(body *WebhookReqBody) {
    var msg string
    
    err, weather := getWeather(body)
    if err != nil{
        fmt.Println(err)
    } else{
        msg = "** Weather **\n\n"+weather
    }

    err, news := getNews(body)
    if err != nil{
        fmt.Println(err)
    } else {
        msg = msg+"** News **\n\n"+news
    }
    
    if err := reply(body.Message.Chat.ID, msg); err != nil {
        fmt.Println("error in sending reply:", err)
        return
    }
}

// Bot ask for user's location
func AskLocation(body *WebhookReqBody, name string) {
    msg := "Hi, "+name+"\n"+"I can tell you about weather and news of your location, please enter your location after #<city>,<country>"+"\n"+"for example #indore,india"
    if err := reply(body.Message.Chat.ID, msg); err != nil {
        fmt.Println("error in sending reply:", err)
        return
    }
}

// Bot ask for user's name
func GetUserName(body *WebhookReqBody) {
    msg := "Hi I am chat bot, please enter your name after @"+"\n"+"for example @YourName"
    if err := reply(body.Message.Chat.ID, msg); err != nil {
        fmt.Println("error in sending reply:", err)
        return
    }
}

// Set user's name in file for further conversation
func SetUserName(body *WebhookReqBody) {
    name := after(body.Message.Text, "@")
    SetUser(body.Message.Chat.ID, name)
    AskLocation(body, name)
}

// make reply to user
func reply(chatID int64, msg string) error {
    token     := os.Getenv("TOKEN")
    reqBody := &sendMessageReqBody{
        ChatID: chatID,
        Text:   msg,
    }
    reqBytes, err := json.Marshal(reqBody)
    if err != nil {
        return err
    }

    res, err := http.Post("https://api.telegram.org/bot"+token+"/sendMessage", "application/json", bytes.NewBuffer(reqBytes))
    if err != nil {
        return err
    }
    if res.StatusCode != http.StatusOK {
        return errors.New("unexpected status" + res.Status)
    }

    return nil
}

// for extract information feed by user
func after(value string, a string) string {
    pos := strings.LastIndex(value, a)
    if pos == -1 {
        return ""
    }
    adjustedPos := pos + len(a)
    if adjustedPos >= len(value) {
        return ""
    }
    return value[adjustedPos:len(value)]
}

// API call for get news
func getNews(reqBody *WebhookReqBody) (error, string) {
    newsToken := os.Getenv("NEWS_TOKEN")
    location := after(reqBody.Message.Text, "#")
    loc := strings.Split(location, ",")
    var newsString string

    if loc[0] != "" {
        url := "https://newsapi.org/v2/top-headlines?q="+loc[1]+"&apiKey="+newsToken
        res, err := http.Get(url)
        if err != nil {
            fmt.Println(err)
        }
        body, err := ioutil.ReadAll(res.Body)
        if err != nil {
            fmt.Println(err)
        }

        var data News
        err = json.Unmarshal(body, &data)
        if err != nil {
            fmt.Println(err)
        }
        
        if data.Status == "ok" && data.TotalResults > 0 {
            for i, d := range data.Articles {
                if i == 3 {
                    break
                }
                if i == 0{
                    newsString = fmt.Sprintf("%d. "+d.Description+"\n\n", (i+1))
                } else{
                    newsString = fmt.Sprintf(newsString+"%d. "+d.Description+"\n\n", (i+1))
                }
            }
        }
    }

    if newsString == "" {
        newsString = "Sorry could not get news for your location!\n"
    }

    return nil, newsString
}

// API call for get weather information
func getWeather(reqBody *WebhookReqBody) (error, string) {
    weatherToken := os.Getenv("WEATHER_TOKEN")
    location := after(reqBody.Message.Text, "#")
    var weatherStr string

    if location != "" {
        url := "https://api.weatherbit.io/v2.0/current?city="+location+"&key="+weatherToken
        res, err := http.Get(url)
        if err != nil {
            fmt.Println(err)
        }
        body, err := ioutil.ReadAll(res.Body)
        if err != nil {
            fmt.Println(err)
        }

        var data Weather
        err = json.Unmarshal(body, &data)
        if err != nil {
            fmt.Println(err)
        }
        
        if data.Count > 0 {
            for _, d := range data.Data {
                temp := fmt.Sprintf("%f", d.Temp)
                weatherStr = "Weather : "+d.Weather.Description+"\n"
                weatherStr = weatherStr+"Temparature : "+temp+"\n\n"
            }
        }
    }

    if weatherStr == "" {
        weatherStr = "Sorry could not get weather information for your location!\n\n"
    }

    return nil, weatherStr
}