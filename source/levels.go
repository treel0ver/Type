package main

import (
	"fmt"
	"math"
	"strings"
	"strconv"
	"io/ioutil"
	"os"
	"log"
	"github.com/bwmarrin/discordgo"
)

var Levels []string

func Add_exp(s *discordgo.Session, m *discordgo.MessageCreate, n float64) {
	Load_levels()
	/* println(m.Author.ID) */
	var exists bool = false
	for i := 0; i < len(Levels); i++ {
		var CL = strings.Split(Levels[i], " # ")
		if CL[0] == m.Author.ID {
			exists = true
			EXP, _ := strconv.ParseFloat(CL[2], 8)
			var Memorized_level = Exp_to_level(EXP)
			EXP = EXP + n
			var Current_level = Exp_to_level(EXP)
			if Current_level != Memorized_level {
				s.ChannelMessageSend(m.ChannelID, "```diff\n+ Felicidades, " + m.Author.Username + ", has subido al nivel " + Current_level + "```")
			}
			EXP_str := fmt.Sprintf("%f", EXP)

			Levels[i] = m.Author.ID + " # " + m.Author.Username + " # " + EXP_str

		    e := os.Remove("./database/levels.lbsv")
		    if e != nil {
		        log.Fatal(e)
		    }


		    // create file
		    f, err := os.Create("./database/levels.lbsv")
		    if err != nil {
		        log.Fatal(err)
		    }
		    // remember to close the file
		    defer f.Close()

		    for _, line := range Levels {
		        _, err := fmt.Fprintln(f, line)
		        if err != nil {
		            log.Fatal(err)
		        }
		    }
		}
	}
	if exists == false {
		f, err := os.OpenFile("./database/levels.lbsv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}

		if _, err := f.Write([]byte(m.Author.ID + " # " + m.Author.Username + " # " + "0")); err != nil {
			f.Close() 
			log.Fatal(err)
		}
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}
}

func Exp_to_level(exp float64) string {
	var level = 0.02 * math.Sqrt(exp)
	if level > 99 {
		var level_str = "99+"
		return level_str
	} else {
		var level_str = fmt.Sprintf("%.0f", level)
		return level_str
	}
	return "error"
}

func Load_levels() {
	fileBytes, err := ioutil.ReadFile("./database/levels.lbsv")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	Levels = strings.Split(string(fileBytes), "\n")
}

func Show_level(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	Load_levels()
	if len(args) > 1{
		for i := 0; i < len(Levels); i++ {
			var CL = strings.Split(Levels[i], " # ")
			if len(CL) > 1 {
				if CL[1] == args[1] || CL[0] == args[1] {
					EXP, _ := strconv.ParseFloat(CL[2], 8)
					s.ChannelMessageSend(m.ChannelID, "```css\n" + args[1] + " es nivel [" + Exp_to_level(EXP) + "]```")
				}
			}
		}
	} else {
		for i := 0; i < len(Levels); i++ {
			var CL = strings.Split(Levels[i], " # ")
			if CL[0] == m.Author.ID {
				EXP, _ := strconv.ParseFloat(CL[2], 8)
				s.ChannelMessageSend(m.ChannelID, "```css\n" + m.Author.Username + ", eres nivel [" + Exp_to_level(EXP) + "]```")
			}
		}
	}
}

