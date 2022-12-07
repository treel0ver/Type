package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	Load_texts()

	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		println("error creating Discord session,", err)
		return
	}
	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentGuildMessageTyping
	dg.AddHandler(messageCreate)
	dg.AddHandler(Typing_start_handler)

	err = dg.Open()
	if err != nil {
		println("error opening connection,", err)
		return
	}

	err = dg.UpdateGameStatus(0, ".help, .help2")
	if err != nil {
		println("Unable to set activity: ", err)
	} else {

	}

	println("[" + time.Now().Format("02/01/2006 15:04:05") + "] ✔️  Type is running")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}
