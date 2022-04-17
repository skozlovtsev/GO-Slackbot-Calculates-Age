package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent){
	for event := range analyticsChannel{
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main(){
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-3393407759590-3397219767205-C9yHg5QHfeIT3hlkWWuPJ8bK")  //переменной окружения SLACK_BOT_TOKEN присваиваем значеение полученого токена
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A03BBEB8T2B-3393442925174-389db31a7971ddd3c958de49d9e807d01ec394901daf75abb364b36b250d1c8f")  //переменной окружения SLACK_APP_TOKEN присваиваем значеение полученого токена

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))  //создаем экземпляр бота

	go printCommandEvents(bot.CommandEvents())  //запускаем функцию printCommandEvents в отдельном потоке (bot.CommandEvents() передаем в качестве канала)

	bot.Command("my yob is <year>", &slacker.CommandDefinition{  //добавляем обработчик для команды по патерну "my yob is <year>" 
		Description: "yob calculator",  //описание команды
		Example:     "my yob is 2020",  //пример команды
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {  //фуктция обработчик команды
			year := request.Param("year")  //в переменную year получаем параметр из патерна <year>
			yob, err := strconv.Atoi(year)
			if err != nil{
				println("error")
			}
			age := 2021 - yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)  //отправка r в качестве отввета
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}