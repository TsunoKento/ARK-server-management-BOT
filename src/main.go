package main

import (
	"TsunoKento/AWS-server-management-BOT/discord"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var session *discordgo.Session

func init() {
	var err error

	session, err = discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	session.AddHandler(discord.OnMessageCreate)

	if err := session.Open(); err != nil {
		log.Fatal(err)
	}

	log.Println("BOT Running...")

	stop_bot := make(chan os.Signal, 1)

	signal.Notify(stop_bot, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	<-stop_bot

	if err := session.Close(); err != nil {
		log.Fatal(err)
	}

}
