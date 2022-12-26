package main

type WebhookReqBody struct {
    Message struct {
        Text string `json:"text"`
        Chat struct {
            ID int64 `json:"id"`
        } `json:"chat"`
    } `json:"message"`
}

type sendMessageReqBody struct {
    ChatID int64  `json:"chat_id"`
    Text   string `json:"text"`
}

type News struct {
    Status       string `json:"status"`
    TotalResults int    `json:"totalResults"`
    Articles     []struct {
        Source struct {
            ID   interface{} `json:"id"`
            Name string      `json:"name"`
        } `json:"source"`
        Author      interface{} `json:"author"`
        Title       string      `json:"title"`
        Description string      `json:"description"`
        URL         string      `json:"url"`
        URLToImage  string      `json:"urlToImage"`
        Content     string      `json:"content"`
    } `json:"articles"`
}

type Weather struct {
    Data []struct {
        CountryCode  string  `json:"country_code"`
        Clouds       int     `json:"clouds"`
        WindCdirFull string  `json:"wind_cdir_full"`
        WindCdir     string  `json:"wind_cdir"`
        Snow         int     `json:"snow"`
        Precip       int     `json:"precip"`
        WindDir      int     `json:"wind_dir"`
        Weather      struct {
            Icon        string `json:"icon"`
            Code        string `json:"code"`
            Description string `json:"description"`
        } `json:"weather"`
        Datetime  string  `json:"datetime"`
        Temp      float64 `json:"temp"`
        Station   string  `json:"station"`
        AppTemp   float64 `json:"app_temp"`
    } `json:"data"`
    Count int `json:"count"`
}
