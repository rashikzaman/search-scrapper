## Description

This application lets user to search `google search results` page with uploaded keywords. User must be authenticated to use this application.


It is build with `golang` and `postgres` is used for database. 

## Installation

Please make sure you have `golang` and `postgres` installed on your system. 

After cloning the repository, install project dependencies by running:

```sh
go mod download
```

Create a .env file and enter proper postgres db config:

```sh
cp .env.example .env
```

Run the appliation with:
```
go run main.go
```

To Test:
```
go test ./...
```

Application will run on port `8080` by default. To change it, update `PORT` variable in `.env` file and restart the server.


## How it works

When the authenticated user uploads a csv file containing keywords, those keyowrds are store in the database and marked as `pending`. A scheduler will be running periodically which will pick up a `pending` keyword and make request to `https://www.google.ru/search` page. Then the application process the response html and parses `Total Search Result`, `Total Adwords` and `Total link`. Then it'll save these information into the database and save the html file in `public\results` folder. After that, keyword status will be updated to `complete`. 

Different user-agents are used to carry out the search. Those user-agents are stored in `user_agents.txt` file. 

Scheduler is run on repeat after every 5 seconds by default. To speed things up, this value can be changed by updating `SCHEDULER_INTERVAL` variable in .env file. This value must be whole number. However, make sure that you don't face ip ban by requesting too many searches in short time. 

Note that, for same keyword, the application doesn't look for cache but instead make the search again, cause search result can vary even if the search key is the same.

There is a sample csv file in project root named `sample_keywords.csv`. Please make sure the uploaded csv file adheres to the same format.

To parse the total search result, `#result-stats` element is looked for in the html

For, total ads, total divs in `#tads` elements is counted

To count total link, all `a` elements is counted