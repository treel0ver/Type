package main

import (
	"fmt"
	"time"
	//"bufio"
	"log"
	"io/ioutil"
	"os"
	"strings"
	"strconv"
	"github.com/bwmarrin/discordgo"
)

func Load() {
	fileBytes, err := ioutil.ReadFile("./database/texts.lbsv")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	DB = strings.Split(string(fileBytes), "\n")
}

func Save_result(m *discordgo.MessageCreate) {
	f, err := os.OpenFile("./database/texts.lbsv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	Random_str := strconv.Itoa(Random)
	if _, err := f.Write([]byte(Random_str + " # " + m.Author.ID + " # " + m.Author.Username + " # " + WPM_str + " # " + Date + " # " + WPM_str_save + " # " + m.Message.ID + " # " + Text_message_ID + "\n")); err != nil {
		f.Close() 
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func Update() {
    e := os.Remove("./database/texts.lbsv")
    if e != nil {
        log.Fatal(e)
    }



    // create file
    f, err := os.Create("./database/texts.lbsv")
    if err != nil {
        log.Fatal(err)
    }
    // remember to close the file
    defer f.Close()

    for _, line := range DB {
        _, err := fmt.Fprintln(f, line)
        if err != nil {
            log.Fatal(err)
        }
    }
}

func Log(m *discordgo.MessageCreate) {
	f, err := os.OpenFile("./database/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := f.Write([]byte("[" + time.Now().Format("02/01/2006 15:04:05") + "] <#" + m.ChannelID + "> " + m.Author.ID + ", " + m.Author.Username  + "> " + m.Content + "\n")); err != nil {
		f.Close() 
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}