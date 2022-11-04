package main

import (
	"fmt"
	"time"
	"math"
	"strings"
	"strconv"
	"io/ioutil"
	"os"
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

		    e := os.Remove("./database/levels.csv")
		    if e != nil {
			    fmt.Println("[" + time.Now().Format("02/01/2006 15:04:05") + "]")
			    fmt.Println(e)
		    }


		    // create file
		    f, err := os.Create("./database/levels.csv")
		    if err != nil {
			    fmt.Println("[" + time.Now().Format("02/01/2006 15:04:05") + "]")
			    fmt.Println(err)
			}
		    // remember to close the file
		    defer f.Close()

		    for _, line := range Levels {
		        _, err := fmt.Fprintln(f, line)
		        if err != nil {
			    fmt.Println("[" + time.Now().Format("02/01/2006 15:04:05") + "]")
			    fmt.Println(err)
		        }
		    }
		}
	}
	if exists == false {
		f, err := os.OpenFile("./database/levels.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			    fmt.Println("[" + time.Now().Format("02/01/2006 15:04:05") + "]")
			    fmt.Println(err)
		}

		if _, err := f.Write([]byte(m.Author.ID + " # " + m.Author.Username + " # " + "0")); err != nil {
			f.Close() 
			    fmt.Println("[" + time.Now().Format("02/01/2006 15:04:05") + "]")
			    fmt.Println(err)
		}
		if err := f.Close(); err != nil {
			    fmt.Println("[" + time.Now().Format("02/01/2006 15:04:05") + "]")
			    fmt.Println(err)
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
	fileBytes, err := ioutil.ReadFile("./database/levels.csv")

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

