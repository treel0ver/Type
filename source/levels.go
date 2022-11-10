package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var Levels []string

var Doing_action bool

func Add_exp(s *discordgo.Session, m *discordgo.MessageCreate, n float64) {
	Load_levels()
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
				print("[")
				print(time.Now().Format("02/01/2006 15:04:05"))
				print("] " + m.Author.Username + " ha subido a nivel " + Current_level)
				s.ChannelMessageSend(m.ChannelID, "```diff\n+ Felicidades, "+m.Author.Username+", has subido al nivel "+Current_level+"```")
			}
			EXP_str := fmt.Sprintf("%f", EXP)

			Levels[i] = m.Author.ID + " # " + m.Author.Username + " # " + EXP_str

			if !Doing_action {
				Doing_action = true
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
				Doing_action = false
			}
		}
	}
	if !exists {
		f, err := os.OpenFile("./database/levels.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("[" + time.Now().Format("02/01/2006 15:04:05") + "]")
			fmt.Println(err)
		}

		if _, err := f.Write([]byte(m.Author.ID + " # " + m.Author.Username + " # " + "500\n")); err != nil {
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
	var level = 0.007 * math.Sqrt(exp)
	if level > 99 {
		var level_str = "+99"
		return level_str
	} else {
		var level_str = fmt.Sprintf("%.0f", level)
		return level_str
	}

}

func Load_levels() {
	fileBytes, err := ioutil.ReadFile("./database/levels.csv")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	Levels = strings.Split(string(fileBytes), "\n")

	// initialize the index of the element to delete
	i := len(Levels) - 1

	// check if the index is within array bounds
	if i < 0 || i >= len(Levels) {
		fmt.Println("The given index is out of bounds.")
	} else {
		// delete an element from the array
		newLength := 0
		for index := range Levels {
			if i != index {
				Levels[newLength] = Levels[index]
				newLength++
			}
		}

		// reslice the array to remove extra index
		newArray := Levels[:newLength]
		Levels = newArray
	}

}

func Show_level(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	Load_levels()
	if len(args) > 1 {
		for i := 0; i < len(Levels); i++ {
			var CL = strings.Split(Levels[i], " # ")
			if len(CL) > 1 {
				if CL[1] == args[1] || CL[0] == args[1] {
					EXP, _ := strconv.ParseFloat(CL[2], 8)
					s.ChannelMessageSend(m.ChannelID, "```css\n"+args[1]+" es nivel ["+Exp_to_level(EXP)+"]```")
				}
			}
		}
	} else {
		for i := 0; i < len(Levels); i++ {
			var CL = strings.Split(Levels[i], " # ")
			if CL[0] == m.Author.ID {
				EXP, _ := strconv.ParseFloat(CL[2], 8)
				s.ChannelMessageSend(m.ChannelID, "```css\n"+m.Author.Username+", eres nivel ["+Exp_to_level(EXP)+"]```")
			}
		}
	}
}

func Exp_until_next_level(n int) int {
	var adding = n

	for {
		var hipo_level = Exp_to_level(float64(adding))
		if hipo_level > Exp_to_level(float64(n)) {
			return adding
		}
		adding++
	}
}
