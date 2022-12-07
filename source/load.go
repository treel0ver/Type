package main

import (
	"fmt"
	"time"

	//"bufio"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Load() {
	fileBytes, err := ioutil.ReadFile("./database/saved_results.csv")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	DB = strings.Split(string(fileBytes), "\n")
}

var changing_saved_results bool = false

func Save_result(m *discordgo.MessageCreate, Text_ID int) {
	for true {
		if !changing_saved_results {
			changing_saved_results = true
			f, err := os.OpenFile("./database/saved_results.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatal(err)
			}
			Random_str := strconv.Itoa(Text_ID)
			if _, err := f.Write([]byte(Random_str + " # " + m.Author.ID + " # " + m.Author.Username + " # " + WPM_str + " # " + Date + " # " + WPM_str_save + " # " + m.Message.ID + " # " + Text_message_ID + "\n")); err != nil {
				f.Close()
				log.Fatal(err)
			}
			if err := f.Close(); err != nil {
				log.Fatal(err)
			}
			changing_saved_results = false
			break
		}
	}
}

func Update() {
	for true {
		if !changing_saved_results {
			changing_saved_results = true
			e := os.Remove("./database/saved_results.csv")
			if e != nil {
				log.Fatal(e)
			}

			f, err := os.Create("./database/saved_results.csv")
			if err != nil {
				log.Fatal(err)
			}

			defer f.Close()

			for _, line := range DB {
				_, err := fmt.Fprintln(f, line)
				if err != nil {
					log.Fatal(err)
				}
			}
			changing_saved_results = false
			break
		}
	}
}

func Log(m *discordgo.MessageCreate) {
	f, err := os.OpenFile("./database/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := f.Write([]byte("[" + time.Now().Format("02/01/2006 15:04:05") + "] <#" + m.ChannelID + "> " + m.Author.ID + ", " + m.Author.Username + "> " + m.Content + "\n")); err != nil {
		f.Close()
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func Load_texts() {
	fileBytes, err := ioutil.ReadFile("./database/texts.csv")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	Texts = strings.Split(string(fileBytes), "\n")
}

func Free_texts() {
	Texts = nil
}
