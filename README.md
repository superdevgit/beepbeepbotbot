# Beep Beep Bot (Golang)
Create a Telegram chatbot which talks with a user and provides him/her relevant information like the weather, news and other interesting tidbits on any particular day, at a geographic location. 

Basic feature of the application should include:
The bot should first ask the name of the user, and start addressing him/her by that name
Then the bot should ask about the location of the person
The bot will tell the weather of the day in that location
The bot will also tell the top 3 news in that country

## Prerequisites
Need to have 'go' installed on your machine. 
To check its working bot has to be deployed on Heroku server. So you should have an account on Heroku and Telegram.
* Create a bot in [Telegram](https://web.telegram.org/#/login). Search BotFather and follow the instructions. It will give you a bot token.
* Clone this repository.
* **This app is using some API's to get weather, and news.**
  * For Weather: [https://www.weatherbit.io/](https://www.weatherbit.io/)
  * For News: [https://newsapi.org](https://newsapi.org/) (currently it supported are 54 countries)
* Create accounts in these 2 sites and get the api key.
* Now Heroku login to Heroku account use Heroku CLI. Run Heroku login.
    ```
    heroku login
    heroku create APPNAME
    ```
* Once app is created it will give you an app url.
* and push this code to Heroku repository.
    ```
    git push heroku master
    ```
* Once the build is successful. 
* Now we have to give all API keys and bot token to the app.
* Set webhook of telegram by public URL of heroku app.
    ```
    https://api.telegram.org/bot{my_bot_token}/setWebhook?url={url_to_send_updates_to}
    ```
* 3 environment variables need to be set on heroku.
    ```
    NEWS_TOKEN
    PUBLIC_URL
    WEATHER_TOKEN
    ```
* Search your bot in telegram and start chat.

**Search beepgolang_bot in Telegram for live example.**
