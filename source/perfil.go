package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var Profiles []string

func Load_profiles() {
	fileBytes, err := ioutil.ReadFile("./database/profiles.csv")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	Profiles = strings.Split(string(fileBytes), "\n")

	// initialize the index of the element to delete
	i := len(Profiles) - 1

	// check if the index is within array bounds
	if i < 0 || i >= len(Profiles) {
		fmt.Println("The given index is out of bounds.")
	} else {
		// delete an element from the array
		newLength := 0
		for index := range Profiles {
			if i != index {
				Profiles[newLength] = Profiles[index]
				newLength++
			}
		}

		// reslice the array to remove extra index
		newArray := Profiles[:newLength]
		Profiles = newArray
	}

}

func Profile(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {

	Load_profiles()
	var exists bool = false
	for i := range Profiles {
		var CL = strings.Split(Profiles[i], " # ")
		if CL[0] == m.Author.ID {
			exists = true
		}
	}
	if !exists {
		var new_quote string = "¡Hola a todos, soy " + m.Author.Username + "!"
		f, err := os.OpenFile("./database/profiles.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("[" + time.Now().Format("02/01/2006 15:04:05") + "]")
			fmt.Println(err)
		}

		if _, err := f.Write([]byte(m.Author.ID + " # " + m.Author.Username + " # " + new_quote + "\n")); err != nil {
			f.Close()
			fmt.Println("[" + time.Now().Format("02/01/2006 15:04:05") + "]")
			fmt.Println(err)
		}
		if err := f.Close(); err != nil {
			fmt.Println("[" + time.Now().Format("02/01/2006 15:04:05") + "]")
			fmt.Println(err)
		}
	}

	var username string
	var quote string
	var mascot string = "🐳"
	Load_profiles()
	for i := range Profiles {
		var CL = strings.Split(Profiles[i], " # ")
		if CL[0] == m.Author.ID {
			quote = CL[2]
			username = CL[1]
		}
	}
	s.ChannelMessageSend(m.ChannelID, "```¡Perfil de "+username+"! "+mascot+"\n\n"+"«"+quote+"»```" /*+m.Author.AvatarURL("")*/)
}

var Doing_action2 bool

func Quote(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	Load_profiles()
	var exists bool = false
	var where int
	for i := range Profiles {
		var CL = strings.Split(Profiles[i], " # ")
		if CL[0] == m.Author.ID {
			exists = true
			where = i
		}
	}
	if !exists {
		var new_quote string

		for i := range args {
			if !(i == 0) {
				new_quote += args[i] + "\u200b "
			}
		}
		new_quote = new_quote[:len(new_quote)-1]
		f, err := os.OpenFile("./database/profiles.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("[" + time.Now().Format("02/01/2006 15:04:05") + "]")
			fmt.Println(err)
		}

		if _, err := f.Write([]byte(m.Author.ID + " # " + m.Author.Username + " # " + new_quote + "\n")); err != nil {
			f.Close()
			fmt.Println("[" + time.Now().Format("02/01/2006 15:04:05") + "]")
			fmt.Println(err)
		}
		if err := f.Close(); err != nil {
			fmt.Println("[" + time.Now().Format("02/01/2006 15:04:05") + "]")
			fmt.Println(err)
		}

		s.ChannelMessageSend(m.ChannelID, "```diff\n+ Tu frase se ha añadido correctamente.```")
	} else {

		var new_quote string

		for i := range args {
			if !(i == 0) {
				new_quote += args[i] + "\u200b "
			}
		}
		new_quote = new_quote[:len(new_quote)-1]

		if len(args) > 1 {
			Profiles[where] = m.Author.ID + " # " + m.Author.Username + " # " + new_quote
			if !Doing_action2 {
				Doing_action2 = true
				e := os.Remove("./database/profiles.csv")
				if e != nil {
					fmt.Println("[" + time.Now().Format("02/01/2006 15:04:05") + "]")
					fmt.Println(e)
				}

				// create file
				f, err := os.Create("./database/profiles.csv")
				if err != nil {
					fmt.Println("[" + time.Now().Format("02/01/2006 15:04:05") + "]")
					fmt.Println(err)
				}
				// remember to close the file
				defer f.Close()

				for _, line := range Profiles {
					_, err := fmt.Fprintln(f, line)
					if err != nil {
						fmt.Println("[" + time.Now().Format("02/01/2006 15:04:05") + "]")
						fmt.Println(err)
					}
				}

				s.ChannelMessageSend(m.ChannelID, "```diff\n+ Tu frase se ha añadido correctamente.```")

				Doing_action2 = false
			}
		} else {
			s.ChannelMessageSend(m.ChannelID, "```diff\n- Necesitas más argumentos.```")
		}
	}

}
