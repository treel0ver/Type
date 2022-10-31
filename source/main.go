package main

import (	
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"
	"github.com/bwmarrin/discordgo"
)

/* Variables used for command line parameters */
var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	/* Create a new Discord session using the provided bot token. */
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		println("error creating Discord session,", err)
		return
	}

	/* Register the messageCreate func as a callback for MessageCreate events. */
	dg.AddHandler(messageCreate)

	/* In this example, we only care about receiving message events. */
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	/* Open a websocket connection to Discord and begin listening. */
	err = dg.Open()
	if err != nil {
		println("error opening connection,", err)
		return
	}

	err = dg.UpdateGameStatus(0, ".help")
	if err != nil {
		println("Unable to set activity: ", err)
	} else {
		
	}

	/* Wait here until CTRL-C or other term signal is received. */
	println("[" + time.Now().Format("02/01/2006 15:04:05") + "] ✅ RUNNING")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	/* Cleanly close down the Discord session. */
	dg.Close()
}


