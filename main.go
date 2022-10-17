package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
    "log"
    "time"
    "math/rand"
	"github.com/bwmarrin/discordgo"
	"strings"
	//"strconv"
	"strconv"
	"reflect"
/*
	"encoding/json"
	"io/ioutil"
	"time"
	"strconv"
    "bufio"
    "io"
*/
)

																	func check(e error) {
																	    if e != nil {
																	        panic(e)
																	    }
																	}

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
		fmt.Println("error creating Discord session,", err)
		return
	}

	/* Register the messageCreate func as a callback for MessageCreate events. */
	dg.AddHandler(messageCreate)

	/* In this example, we only care about receiving message events. */
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	/* Open a websocket connection to Discord and begin listening. */
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	/* Wait here until CTRL-C or other term signal is received. */
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	/* Cleanly close down the Discord session. */
	dg.Close()
}

/* This function will be called (due to AddHandler above) every time a new  */
/* message is created on any channel that the authenticated bot has access to. */
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	/* Ignore all messages created by the bot itself */
	/* This isn't required in this specific example but it's a good practice. */
	if m.Author.ID == s.State.User.ID {
		return
	}

	if is_illegal(m.Content) {
		s.ChannelMessageSend("1031077892748234762", "<@" + m.Author.ID + "> ha hecho trampas\t" + time.Now().Format("01-02-2006 15:04:05") + "\t" + split_curr())
	}
	if m.Content == ".tt" {
		random = rand.Intn(5)
		start_author = m.Author.ID
		s.ChannelMessageSend(m.ChannelID, "Preparados...") 
		time.Sleep(2 * time.Second)
/*
		chanl, err := s.Channel(m.ChannelID)
		if err != nil {
			return
		}

		s.ChannelMessageEdit(m.ChannelID, m.Message.ID, display_textos [random])
		fmt.Println(m.Message.ID)
*/
		s.ChannelMessageSend(m.ChannelID, display_textos [random])
		current_text = textos[random]
		start()
		
	}

	if m.Content == current_text && is_started {
		calculate_wpm()
		/*is_started = false*/


		wpm_stringed := fmt.Sprint(wpm)

		// CALCULATE IF GREATER OR SMALER
		
		var AAA float64

		if AA, err := strconv.ParseFloat(tops[random*5][2], 64); err == nil {
		    AAA = AA
		}
		
		if AAA < wpm {
			s := fmt.Sprintf("%f", wpm)  
			tops[random*5][2] = s
			tops[random*5][1] = m.Author.Username
			tops[random*5][5] = m.Author.ID
			dt := time.Now()
			tops[random*5][3] = dt.Format("01-02-2006 15:04:05")
			tops[random*5][0] = "1"

		/* CHECKING IF EMPTY */
		} else if tops[random*5][2] == "" { 
			s := fmt.Sprintf("%f", wpm)  
			tops[random*5][2] = s
			tops[random*5][1] = m.Author.Username
			tops[random*5][5] = m.Author.ID
			dt := time.Now()
			tops[random*5][3] = dt.Format("01-02-2006 15:04:05")
			tops[random*5][0] = "1"
		} else if tops[random*5+1][2] == "" { 
			s := fmt.Sprintf("%f", wpm)  
			tops[random*5+1][2] = s
			tops[random*5+1][1] = m.Author.Username
			tops[random*5+1][5] = m.Author.ID
			dt := time.Now()
			tops[random*5+1][3] = dt.Format("01-02-2006 15:04:05")
			tops[random*5+1][0] = "2"
		} else if tops[random*5+2][2] == "" { 
			s := fmt.Sprintf("%f", wpm)  
			tops[random*5+2][2] = s
			tops[random*5+2][1] = m.Author.Username
			tops[random*5+2][5] = m.Author.ID
			dt := time.Now()
			tops[random*5+2][3] = dt.Format("01-02-2006 15:04:05")
			tops[random*5+2][0] = "3"
		} else if tops[random*5+3][2] == "" { 
			s := fmt.Sprintf("%f", wpm)  
			tops[random*5+3][2] = s
			tops[random*5+3][1] = m.Author.Username
			tops[random*5+3][5] = m.Author.ID
			dt := time.Now()
			tops[random*5+3][3] = dt.Format("01-02-2006 15:04:05")
			tops[random*5+3][0] = "4"
		} else if tops[random*5+4][2] == "" { 
			s := fmt.Sprintf("%f", wpm)  
			tops[random*5+4][2] = s
			tops[random*5+4][1] = m.Author.Username
			tops[random*5+3][5] = m.Author.ID
			dt := time.Now()
			tops[random*5+4][3] = dt.Format("01-02-2006 15:04:05")
			tops[random*5+4][0] = "5"
		} else { 

		}

		s.ChannelMessageSend(m.ChannelID, "--------------------------------------------------------------------\n¡Has terminado la carrera!\nWPM: " + wpm_stringed + 
		"\nTops:\n" + tops[random*5][0] + ". " + tops[random*5][1] + " (" + tops[random*5][2] + " wpm) " + tops[random*5][3] +
		"\n" + tops[random*5+1][0] + ". " + tops[random*5+1][1] + " (" + tops[random*5+1][2] + " wpm) " + tops[random*5+1][3] +
		"\n" + tops[random*5+2][0] + ". " + tops[random*5+2][1] + " (" + tops[random*5+2][2] + " wpm) " + tops[random*5+2][3] +
		"\n" + tops[random*5+3][0] + ". " + tops[random*5+3][1] + " (" + tops[random*5+3][2] + " wpm) " + tops[random*5+4][3] +
		"\n" + tops[random*5+4][0] + ". " + tops[random*5+4][1] + " (" + tops[random*5+4][2] + " wpm) " + tops[random*5+4][3] +

		"\n--------------------------------------------------------------------")

		var random_s = strconv.FormatInt(int64(random)+1, 10)

		f, err := os.OpenFile("database/test", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		dt := time.Now()
		if _, err := f.Write([]byte(m.Author.ID + "\t" + dt.Format("01-02-2006 15:04:05") + "\t" + wpm_stringed + "\ttexto: (" + random_s + ") " + split_curr() + "\n")); err != nil {
			f.Close() 
			log.Fatal(err)
		}
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}


	} else if CountWords(m.Content) > CountWords(current_text)-3 {
		if is_illegal(m.Content) {
			for i := 0; i < 4; i++ {
				fmt.Println("¡¡<@" + m.Author.ID + "> HA COPY PASTE!!")
			}
			

			A := m.Content
			sent_arrayed := strings.Split(A, " ")

			B := current_text
			text_arrayed := strings.Split(B, " ")

			if (sent_arrayed[0] == text_arrayed[0]) /*&& (sent_arrayed[1] == text_arrayed[1] && (sent_arrayed[len(sent_arrayed)-1] == text_arrayed[len(text_arrayed)-1])*/ {
				calculate_wpm()
				calculate_errors(m.Content, current_text)
				/*is_started = false*/
				wpm_stringed := fmt.Sprint(wpm)
				s.ChannelMessageSend(m.ChannelID, "--------------------------------------------------------------------\nHAS INTENTADO COPY PASTE. (TU WPM HUBIERA SIDO: " + wpm_stringed + ")\n--------------------------------------------------------------------")

				s.ChannelMessageSend("1031077892748234762", errores_s + ": " + lista_errores)
			}
		} else {

			A := m.Content
			sent_arrayed := strings.Split(A, " ")

			B := current_text
			text_arrayed := strings.Split(B, " ")

			if (sent_arrayed[0] == text_arrayed[0]) /*&& (sent_arrayed[1] == text_arrayed[1] && (sent_arrayed[len(sent_arrayed)-1] == text_arrayed[len(text_arrayed)-1])*/ {
				calculate_wpm()
				calculate_errors(m.Content, current_text)
				/*is_started = false*/
				wpm_stringed := fmt.Sprint(wpm)
				s.ChannelMessageSend(m.ChannelID, "--------------------------------------------------------------------\nNos has terminado la carrera correctamente.\nHas cometido " + errores_s + " errores: " + lista_errores + "\nWPM raw: " + wpm_stringed + "\n--------------------------------------------------------------------")
			}
		}
	}

	if m.Content == ".info" {
		sec := fmt.Sprint(time_elapsed)
		sec2 := fmt.Sprint(time_elapsed/1000)

		var average_word_length_of_current_text_stringed string = fmt.Sprintf("%f", (average_word_length(current_text)))
	
	    dat, err := os.ReadFile("database/test")
    	check(err)
    	s.ChannelMessageSend(m.ChannelID, string(dat) + "\n[" + m.Author.ID + "]\nmilliseconds: " + sec + "\nseconds: " + sec2 + "\nstart_author: " + start_author + "\naverage_word_length_of_current_text_stringed: " + average_word_length_of_current_text_stringed)
	}

	if m.Content == ".test" {
		f, err := os.OpenFile("database/test", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		dt := time.Now()
		if _, err := f.Write([]byte(m.Author.ID + "\t" + dt.Format("02-21-2006 15:04:05") + "\tappended some data\n")); err != nil {
			f.Close() 
			log.Fatal(err)
		}
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}

	if m.Content == ".A" {
		ff := string(tops[random*5][2])
		fmt.Println(ff)
		fmt.Println(tops[random*5][2])
		f, err := strconv.ParseFloat(ff, 8)

		
		fmt.Println(f, err, reflect.TypeOf(f))


	}

	if m.Content == ".help" {
		s.ChannelMessageSend(m.ChannelID, "```.tt       empieza un test de velocidad en idioma español\n.info     enseña información aleatoria para desarrolladores\n.mapache  pone el gif de un mapache```")
	}

	/* FUN COMMANDS */
	/**/	
	/**/	if m.Content == ".mapache" {
	/**/		s.ChannelMessageSend(m.ChannelID, "https://tenor.com/view/froze-stop-moving-not-moving-still-standing-gif-16669739")
	/**/	}
	/**/
	/**/	if m.Content == ".logo" {
	/**/		s.ChannelMessageSend(m.ChannelID, "𝗧𝗬𝗣𝗘 𝗕𝗢𝗧")
	/**/	}	
	/**/
	/**/
	/**/
	/**/
	/**/

}

