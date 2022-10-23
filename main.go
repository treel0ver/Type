package main

import (
	"flag"
 	"fmt"
 	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
	"log"
	"time"
	"math/rand"
	"strings"
	"strconv"
	"reflect"
	"github.com/bwmarrin/discordgo"
	"github.com/Clinet/discordgo-embed"
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

var último_mensaje_del_bot_ID string

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	var fileName = "database/TOPS"
 	fileBytes, err := ioutil.ReadFile(fileName)

 	if err != nil {
 		fmt.Println(err)
 		os.Exit(1)
 	}

 	sliceData := strings.Split(string(fileBytes), "\n")

 	fmt.Println("number of texts: ")
 	fmt.Println(how_many_texts())

   var i, j int
   var n int = 300
   var cuenta int = 0
   for  i = 0; i < n; i++ {
      for j = 0; j < 6; j++ {
         tops[i][j] = sliceData[cuenta]
         cuenta++
      }
   }	

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
		último_mensaje_del_bot_ID = m.Message.ID
		return
	}

	if is_illegal(m.Content) {
		s.ChannelMessageSend("1031077892748234762", "<@" + m.Author.ID + "> ha hecho trampas\t" + time.Now().Format("01-02-2006 15:04:05") + "\t" + split_curr())
	}


	var content_to_lowercase = strings.ToLower(m.Content)
	var abb = strings.Split(content_to_lowercase, " ")

	var positivo bool = true

	if !(len(abb) < 2) {
		if (strings.HasPrefix(abb[1], "long")) { 
			s.ChannelMessageSend(m.ChannelID, "Textos largos aún no disponibles.")
		}
	}

	if strings.HasPrefix(content_to_lowercase, ".tt") {

		if !(len(abb) < 2) {
			if !(strings.HasPrefix(abb[1], "long")) {
				positivo = false
			}
		}

		if positivo == true {
			var editbool bool = false

			if is_started == true {
				if time_soon == true {
					s.ChannelMessageSend(m.ChannelID, "⌛ Esperando más tiempo, ya que la carrera terminó antes de lo debido... ") 
					time.Sleep(3 * time.Second)
					editbool = true
				}

				is_started = false

				/* espera anti-bug */
				time.Sleep(time.Second)

			}

			rand.Seed(time.Now().UnixNano())

			random = 0
			random = rand.Intn(how_many_texts())
			
			start_author = m.Author.ID

			if editbool == true && is_started == false {
				s.ChannelMessageEdit(m.ChannelID, último_mensaje_del_bot_ID, "🔴 Preparados... ") 
				time.Sleep(time.Second)
			} else if is_started == false {
				s.ChannelMessageSend(m.ChannelID, "🔴 Preparados... ") 
				time.Sleep(time.Second)
			} 

			s.ChannelMessageEdit(m.ChannelID, último_mensaje_del_bot_ID, "🟡 Listos... ")
			time.Sleep(time.Second)

	/*
			chanl, err := s.Channel(m.ChannelID)
			if err != nil {
				return
			}

			s.ChannelMessageEdit(m.ChannelID, m.Message.ID, display_textos [random])
			fmt.Println(m.Message.ID)
	*/	
		if is_started == false {
			s.ChannelMessageEdit(m.ChannelID, último_mensaje_del_bot_ID, display_textos [random])
			current_text = textos[random]
			start()
		 }
		}
	}

	var is_lower_than_top bool = false
	var wpm_seems_illegal bool = false

	if m.Content == current_text && is_started {
		calculate_wpm()
		/*is_started = false*/

		wpm_stringed := fmt.Sprint(wpm)



		if tops[random*5][5] == m.Author.ID || tops[random*5+1][5] == m.Author.ID || tops[random*5+2][5] == m.Author.ID || tops[random*5+3][5] == m.Author.ID || tops[random*5+4][5] == m.Author.ID {
			if wpm_floated_from_top, err := strconv.ParseFloat(tops[random*5][2], 32); err == nil {
    			fmt.Println(wpm_floated_from_top) // 3.1415927410125732
    			fmt.Println(is_lower_than_top)
			
				if wpm_floated_from_top > wpm {
					is_lower_than_top = true
					fmt.Println(is_lower_than_top)
				}
			}
		}

		if wpm > 285 {
			wpm_seems_illegal = true
		}

if !is_lower_than_top && !wpm_seems_illegal {
		// CALCULATE IF GREATER OR SMALER
		
		var AAA float64

		/* IF TOP 1 */
		/*Set to float*/ if AA, err := strconv.ParseFloat(tops[random*5][2], 64); err == nil {AAA = AA}
		if wpm > AAA {
			if tops[random*5][5] == m.Author.ID {
				var temp string
				var temp2 string

				temp = tops[random*5][2]
				temp2 = tops[random*5][3]

				wpm_string_temp := fmt.Sprintf("%f", wpm)  
				tops[random*5][2] = wpm_string_temp
				tops[random*5][1] = m.Author.Username
				tops[random*5][5] = m.Author.ID
				dt := time.Now()
				tops[random*5][3] = dt.Format("01-02-2006 15:04:05")
				tops[random*5][0] = "1"

				s.ChannelMessageSend(m.ChannelID, "--------------------------------------------------------------------\n¡Has superado tu anterior marca de " + temp + " wpm del " + temp2)
			} else if tops[random*5+4][5] == m.Author.ID {
						var temp string
						var temp2 string

						temp = tops[random*5+4][2]
						temp2 = tops[random*5+4][3]

						tops[random*5+4][2] = tops[random*5+3][2]
						tops[random*5+4][1] = tops[random*5+3][1]
						tops[random*5+4][5] = tops[random*5+3][5]
						tops[random*5+4][3] = tops[random*5+3][3]
						tops[random*5+4][0] = "5"

						tops[random*5+3][2] = tops[random*5+2][2]
						tops[random*5+3][1] = tops[random*5+2][1]
						tops[random*5+3][5] = tops[random*5+2][5]
						tops[random*5+3][3] = tops[random*5+2][3]
						tops[random*5+3][0] = "4"

						tops[random*5+2][2] = tops[random*5+1][2]
						tops[random*5+2][1] = tops[random*5+1][1]
						tops[random*5+2][5] = tops[random*5+1][5]
						tops[random*5+2][3] = tops[random*5+1][3]
						tops[random*5+2][0] = "3"

						tops[random*5+1][2] = tops[random*5][2]
						tops[random*5+1][1] = tops[random*5][1]
						tops[random*5+1][5] = tops[random*5][5]
						tops[random*5+1][3] = tops[random*5][3]
						tops[random*5+1][0] = "2"

						wpm_string_temp := fmt.Sprintf("%f", wpm)  
						tops[random*5][2] = wpm_string_temp
						tops[random*5][1] = m.Author.Username
						tops[random*5][5] = m.Author.ID
						dt := time.Now()
						tops[random*5][3] = dt.Format("01-02-2006 15:04:05")
						tops[random*5][0] = "1"

						s.ChannelMessageSend(m.ChannelID, "--------------------------------------------------------------------\n¡Has superado tu anterior marca de " + temp + " wpm del " + temp2)
					} else if tops[random*5+3][5] == m.Author.ID {
						var temp string
						var temp2 string

						temp = tops[random*5+3][2]
						temp2 = tops[random*5+3][3]

						tops[random*5+3][2] = tops[random*5+2][2]
						tops[random*5+3][1] = tops[random*5+2][1]
						tops[random*5+3][5] = tops[random*5+2][5]
						tops[random*5+3][3] = tops[random*5+2][3]
						tops[random*5+3][0] = "4"

						tops[random*5+2][2] = tops[random*5+1][2]
						tops[random*5+2][1] = tops[random*5+1][1]
						tops[random*5+2][5] = tops[random*5+1][5]
						tops[random*5+2][3] = tops[random*5+1][3]
						tops[random*5+2][0] = "3"

						tops[random*5+1][2] = tops[random*5][2]
						tops[random*5+1][1] = tops[random*5][1]
						tops[random*5+1][5] = tops[random*5][5]
						tops[random*5+1][3] = tops[random*5][3]
						tops[random*5+1][0] = "2"

						wpm_string_temp := fmt.Sprintf("%f", wpm)  
						tops[random*5][2] = wpm_string_temp
						tops[random*5][1] = m.Author.Username
						tops[random*5][5] = m.Author.ID
						dt := time.Now()
						tops[random*5][3] = dt.Format("01-02-2006 15:04:05")
						tops[random*5][0] = "1"

						s.ChannelMessageSend(m.ChannelID, "--------------------------------------------------------------------\n¡Has superado tu anterior marca de " + temp + " wpm del " + temp2)
					} else if tops[random*5+2][5] == m.Author.ID {
						var temp string
						var temp2 string

						temp = tops[random*5+2][2]
						temp2 = tops[random*5+2][3]

						tops[random*5+2][2] = tops[random*5+1][2]
						tops[random*5+2][1] = tops[random*5+1][1]
						tops[random*5+2][5] = tops[random*5+1][5]
						tops[random*5+2][3] = tops[random*5+1][3]
						tops[random*5+2][0] = "3"

						tops[random*5+1][2] = tops[random*5][2]
						tops[random*5+1][1] = tops[random*5][1]
						tops[random*5+1][5] = tops[random*5][5]
						tops[random*5+1][3] = tops[random*5][3]
						tops[random*5+1][0] = "2"

						wpm_string_temp := fmt.Sprintf("%f", wpm)  
						tops[random*5][2] = wpm_string_temp
						tops[random*5][1] = m.Author.Username
						tops[random*5][5] = m.Author.ID
						dt := time.Now()
						tops[random*5][3] = dt.Format("01-02-2006 15:04:05")
						tops[random*5][0] = "1"

						s.ChannelMessageSend(m.ChannelID, "--------------------------------------------------------------------\n¡Has superado tu anterior marca de " + temp + " wpm del " + temp2)
					} else if tops[random*5+1][5] == m.Author.ID {
						var temp string
						var temp2 string

						temp = tops[random*5+1][2]
						temp2 = tops[random*5+1][3]

						tops[random*5+1][2] = tops[random*5][2]
						tops[random*5+1][1] = tops[random*5][1]
						tops[random*5+1][5] = tops[random*5][5]
						tops[random*5+1][3] = tops[random*5][3]
						tops[random*5+1][0] = "2"

						wpm_string_temp := fmt.Sprintf("%f", wpm)  
						tops[random*5][2] = wpm_string_temp
						tops[random*5][1] = m.Author.Username
						tops[random*5][5] = m.Author.ID
						dt := time.Now()
						tops[random*5][3] = dt.Format("01-02-2006 15:04:05")
						tops[random*5][0] = "1"

						s.ChannelMessageSend(m.ChannelID, "--------------------------------------------------------------------\n¡Has superado tu anterior marca de " + temp + " wpm del " + temp2)
					} else {
						wpm_string_temp := fmt.Sprintf("%f", wpm)  
						tops[random*5][2] = wpm_string_temp
						tops[random*5][1] = m.Author.Username
						tops[random*5][5] = m.Author.ID
						dt := time.Now()
						tops[random*5][3] = dt.Format("01-02-2006 15:04:05")
						tops[random*5][0] = "1"
						
					}
		} else {
			/* IF TOP 2 */
			/*Set to float*/ if AA, err := strconv.ParseFloat(tops[random*5+1][2], 64); err == nil {AAA = AA}
			if wpm > AAA {
				if tops[random*5+1][5] == m.Author.ID {
					var temp string
					var temp2 string

					temp = tops[random*5+1][2]
					temp2 = tops[random*5+1][3]

					wpm_string_temp := fmt.Sprintf("%f", wpm)  
					tops[random*5+1][2] = wpm_string_temp
					tops[random*5+1][1] = m.Author.Username
					tops[random*5+1][5] = m.Author.ID
					dt := time.Now()
					tops[random*5+1][3] = dt.Format("01-02-2006 15:04:05")
					tops[random*5+1][0] = "2"

					s.ChannelMessageSend(m.ChannelID, "--------------------------------------------------------------------\n¡Has superado tu anterior marca de " + temp + " wpm del " + temp2)
				} else if tops[random*5+4][5] == m.Author.ID {
						var temp string
						var temp2 string

						temp = tops[random*5+4][2]
						temp2 = tops[random*5+4][3]

						tops[random*5+4][2] = tops[random*5+3][2]
						tops[random*5+4][1] = tops[random*5+3][1]
						tops[random*5+4][5] = tops[random*5+3][5]
						tops[random*5+4][3] = tops[random*5+3][3]
						tops[random*5+4][0] = "5"

						tops[random*5+3][2] = tops[random*5+2][2]
						tops[random*5+3][1] = tops[random*5+2][1]
						tops[random*5+3][5] = tops[random*5+2][5]
						tops[random*5+3][3] = tops[random*5+2][3]
						tops[random*5+3][0] = "4"

						tops[random*5+2][2] = tops[random*5+1][2]
						tops[random*5+2][1] = tops[random*5+1][1]
						tops[random*5+2][5] = tops[random*5+1][5]
						tops[random*5+2][3] = tops[random*5+1][3]
						tops[random*5+2][0] = "3"


						wpm_string_temp := fmt.Sprintf("%f", wpm)  
						tops[random*5+1][2] = wpm_string_temp
						tops[random*5+1][1] = m.Author.Username
						tops[random*5+1][5] = m.Author.ID
						dt := time.Now()
						tops[random*5+1][3] = dt.Format("01-02-2006 15:04:05")
						tops[random*5+1][0] = "2"

						s.ChannelMessageSend(m.ChannelID, "--------------------------------------------------------------------\n¡Has superado tu anterior marca de " + temp + " wpm del " + temp2)
					} else if tops[random*5+3][5] == m.Author.ID {
						var temp string
						var temp2 string

						temp = tops[random*5+3][2]
						temp2 = tops[random*5+3][3]

						tops[random*5+3][2] = tops[random*5+2][2]
						tops[random*5+3][1] = tops[random*5+2][1]
						tops[random*5+3][5] = tops[random*5+2][5]
						tops[random*5+3][3] = tops[random*5+2][3]
						tops[random*5+3][0] = "4"

						tops[random*5+2][2] = tops[random*5+1][2]
						tops[random*5+2][1] = tops[random*5+1][1]
						tops[random*5+2][5] = tops[random*5+1][5]
						tops[random*5+2][3] = tops[random*5+1][3]
						tops[random*5+2][0] = "3"

						wpm_string_temp := fmt.Sprintf("%f", wpm)  
						tops[random*5+1][2] = wpm_string_temp
						tops[random*5+1][1] = m.Author.Username
						tops[random*5+1][5] = m.Author.ID
						dt := time.Now()
						tops[random*5+1][3] = dt.Format("01-02-2006 15:04:05")
						tops[random*5+1][0] = "2"

						s.ChannelMessageSend(m.ChannelID, "--------------------------------------------------------------------\n¡Has superado tu anterior marca de " + temp + " wpm del " + temp2)
					} else if tops[random*5+2][5] == m.Author.ID {
						var temp string
						var temp2 string

						temp = tops[random*5+2][2]
						temp2 = tops[random*5+2][3]

						tops[random*5+2][2] = tops[random*5+1][2]
						tops[random*5+2][1] = tops[random*5+1][1]
						tops[random*5+2][5] = tops[random*5+1][5]
						tops[random*5+2][3] = tops[random*5+1][3]
						tops[random*5+2][0] = "3"

						wpm_string_temp := fmt.Sprintf("%f", wpm)  
						tops[random*5+1][2] = wpm_string_temp
						tops[random*5+1][1] = m.Author.Username
						tops[random*5+1][5] = m.Author.ID
						dt := time.Now()
						tops[random*5+1][3] = dt.Format("01-02-2006 15:04:05")
						tops[random*5+1][0] = "2"

						s.ChannelMessageSend(m.ChannelID, "--------------------------------------------------------------------\n¡Has superado tu anterior marca de " + temp + " wpm del " + temp2)
					} else {
						wpm_string_temp := fmt.Sprintf("%f", wpm)  
						tops[random*5+1][2] = wpm_string_temp
						tops[random*5+1][1] = m.Author.Username
						tops[random*5+1][5] = m.Author.ID
						dt := time.Now()
						tops[random*5+1][3] = dt.Format("01-02-2006 15:04:05")
						tops[random*5+1][0] = "2"
						
					}
			} else {
				/* IF TOP 3 */
				/*Set to float*/ if AA, err := strconv.ParseFloat(tops[random*5+2][2], 64); err == nil {AAA = AA}
				if wpm > AAA {
					if tops[random*5+2][5] == m.Author.ID {
						var temp string
						var temp2 string

						temp = tops[random*5+2][2]
						temp2 = tops[random*5+2][3]

						wpm_string_temp := fmt.Sprintf("%f", wpm)  
						tops[random*5+2][2] = wpm_string_temp
						tops[random*5+2][1] = m.Author.Username
						tops[random*5+2][5] = m.Author.ID
						dt := time.Now()
						tops[random*5+2][3] = dt.Format("01-02-2006 15:04:05")
						tops[random*5+2][0] = "3"

						s.ChannelMessageSend(m.ChannelID, "--------------------------------------------------------------------\n¡Has superado tu anterior marca de " + temp + " wpm del " + temp2)
					} else if tops[random*5+4][5] == m.Author.ID {
						var temp string
						var temp2 string

						temp = tops[random*5+4][2]
						temp2 = tops[random*5+4][3]

						tops[random*5+4][2] = tops[random*5+3][2]
						tops[random*5+4][1] = tops[random*5+3][1]
						tops[random*5+4][5] = tops[random*5+3][5]
						tops[random*5+4][3] = tops[random*5+3][3]
						tops[random*5+4][0] = "5"

						tops[random*5+3][2] = tops[random*5+2][2]
						tops[random*5+3][1] = tops[random*5+2][1]
						tops[random*5+3][5] = tops[random*5+2][5]
						tops[random*5+3][3] = tops[random*5+2][3]
						tops[random*5+3][0] = "4"

						wpm_string_temp := fmt.Sprintf("%f", wpm)  
						tops[random*5+3][2] = wpm_string_temp
						tops[random*5+3][1] = m.Author.Username
						tops[random*5+3][5] = m.Author.ID
						dt := time.Now()
						tops[random*5+3][3] = dt.Format("01-02-2006 15:04:05")
						tops[random*5+3][0] = "4"

						s.ChannelMessageSend(m.ChannelID, "--------------------------------------------------------------------\n¡Has superado tu anterior marca de " + temp + " wpm del " + temp2)
					} else if tops[random*5+3][5] == m.Author.ID {
						var temp string
						var temp2 string

						temp = tops[random*5+3][2]
						temp2 = tops[random*5+3][3]

						tops[random*5+3][2] = tops[random*5+2][2]
						tops[random*5+3][1] = tops[random*5+2][1]
						tops[random*5+3][5] = tops[random*5+2][5]
						tops[random*5+3][3] = tops[random*5+2][3]
						tops[random*5+3][0] = "4"

						wpm_string_temp := fmt.Sprintf("%f", wpm)  
						tops[random*5+2][2] = wpm_string_temp
						tops[random*5+2][1] = m.Author.Username
						tops[random*5+2][5] = m.Author.ID
						dt := time.Now()
						tops[random*5+2][3] = dt.Format("01-02-2006 15:04:05")
						tops[random*5+2][0] = "3"

						s.ChannelMessageSend(m.ChannelID, "--------------------------------------------------------------------\n¡Has superado tu anterior marca de " + temp + " wpm del " + temp2)
					} else {
						wpm_string_temp := fmt.Sprintf("%f", wpm)  
						tops[random*5+2][2] = wpm_string_temp
						tops[random*5+2][1] = m.Author.Username
						tops[random*5+2][5] = m.Author.ID
						dt := time.Now()
						tops[random*5+2][3] = dt.Format("01-02-2006 15:04:05")
						tops[random*5+2][0] = "3"
						
					}
				} else {
					/* IF TOP 4 */
					/*Set to float*/ if AA, err := strconv.ParseFloat(tops[random*5+3][2], 64); err == nil {AAA = AA}
					if wpm > AAA {
						if tops[random*5+3][5] == m.Author.ID {
							var temp string
							var temp2 string

							temp = tops[random*5+3][2]
							temp2 = tops[random*5+3][3]

							wpm_string_temp := fmt.Sprintf("%f", wpm)  
							tops[random*5+3][2] = wpm_string_temp
							tops[random*5+3][1] = m.Author.Username
							tops[random*5+3][5] = m.Author.ID
							dt := time.Now()
							tops[random*5+3][3] = dt.Format("01-02-2006 15:04:05")
							tops[random*5+3][0] = "4"

							s.ChannelMessageSend(m.ChannelID, "--------------------------------------------------------------------\n¡Has superado tu anterior marca de " + temp + " wpm del " + temp2)
						} else if tops[random*5+4][5] == m.Author.ID {
							var temp string
							var temp2 string

							temp = tops[random*5+4][2]
							temp2 = tops[random*5+4][3]

							tops[random*5+4][2] = tops[random*5+3][2]
							tops[random*5+4][1] = tops[random*5+3][1]
							tops[random*5+4][5] = tops[random*5+3][5]
							tops[random*5+4][3] = tops[random*5+3][3]
							tops[random*5+4][0] = "5"

							wpm_string_temp := fmt.Sprintf("%f", wpm)  
							tops[random*5+3][2] = wpm_string_temp
							tops[random*5+3][1] = m.Author.Username
							tops[random*5+3][5] = m.Author.ID
							dt := time.Now()
							tops[random*5+3][3] = dt.Format("01-02-2006 15:04:05")
							tops[random*5+3][0] = "4"

							s.ChannelMessageSend(m.ChannelID, "--------------------------------------------------------------------\n¡Has superado tu anterior marca de " + temp + " wpm del " + temp2)
						} else {
						wpm_string_temp := fmt.Sprintf("%f", wpm)  
						tops[random*5+3][2] = wpm_string_temp
						tops[random*5+3][1] = m.Author.Username
						tops[random*5+3][5] = m.Author.ID
						dt := time.Now()
						tops[random*5+3][3] = dt.Format("01-02-2006 15:04:05")
						tops[random*5+3][0] = "4"
						
					}
					} else {
						/* IF TOP 5 */
						/*Set to float*/ if AA, err := strconv.ParseFloat(tops[random*5+4][2], 64); err == nil {AAA = AA}
						if wpm > AAA {
							if tops[random*5+3][5] == m.Author.ID {
								var temp string
								var temp2 string

								temp = tops[random*5+4][2]
								temp2 = tops[random*5+4][3]

								wpm_string_temp := fmt.Sprintf("%f", wpm)  
								tops[random*5+4][2] = wpm_string_temp
								tops[random*5+4][1] = m.Author.Username
								tops[random*5+4][5] = m.Author.ID
								dt := time.Now()
								tops[random*5+4][3] = dt.Format("01-02-2006 15:04:05")
								tops[random*5+4][0] = "5"

								s.ChannelMessageSend(m.ChannelID, "--------------------------------------------------------------------\n¡Has superado tu anterior marca de " + temp + " wpm del " + temp2)
							} else {
								wpm_string_temp := fmt.Sprintf("%f", wpm)  
								tops[random*5+4][2] = wpm_string_temp
								tops[random*5+4][1] = m.Author.Username
								tops[random*5+4][5] = m.Author.ID
								dt := time.Now()
								tops[random*5+4][3] = dt.Format("01-02-2006 15:04:05")
								tops[random*5+4][0] = "5"
							}
						} 
					}
				}

			}
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

	/* WRITE TO FILE */
    if err != nil {
        log.Fatalf("readLines: %s", err)
    }

    escribir_tops_a_base_de_datos()


	if err := writeLines(); err != nil {
   		 log.Fatalf("writeLines: %s", err)
  	}
	/* WRITE TO FILE */
}

if is_lower_than_top {
	s.ChannelMessageSend(m.ChannelID, "--------------------------------------------------------------------\nLo siento, no has podido superar tu anterior marca.\nWPM: " + wpm_stringed + 
	"\nTops:\n" + tops[random*5][0] + ". " + tops[random*5][1] + " (" + tops[random*5][2] + " wpm) " + tops[random*5][3] +
	"\n" + tops[random*5+1][0] + ". " + tops[random*5+1][1] + " (" + tops[random*5+1][2] + " wpm) " + tops[random*5+1][3] +
	"\n" + tops[random*5+2][0] + ". " + tops[random*5+2][1] + " (" + tops[random*5+2][2] + " wpm) " + tops[random*5+2][3] +
	"\n" + tops[random*5+3][0] + ". " + tops[random*5+3][1] + " (" + tops[random*5+3][2] + " wpm) " + tops[random*5+4][3] +
	"\n" + tops[random*5+4][0] + ". " + tops[random*5+4][1] + " (" + tops[random*5+4][2] + " wpm) " + tops[random*5+4][3] +

	"\n--------------------------------------------------------------------")

	is_lower_than_top = false
}

if wpm_seems_illegal {
	s.ChannelMessageSend(m.ChannelID, m.Author.Username + ", tu resultado de " + wpm_stringed + " WPM parece ilegal ya que excede los 285 WPM. Contacta con <@910067180706627594> para resolver este conflicto.")
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

	if m.Content == ".textos" {
		var how_many_texts_stringed = strconv.FormatInt(int64(how_many_texts()), 10)
		s.ChannelMessageSend(m.ChannelID, "Número de textos en idioma español: " + how_many_texts_stringed)
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
	/**/	if m.Content == ".go" {
	/**/		s.ChannelMessageSend(m.ChannelID, "https://camo.githubusercontent.com/833cfd306ac2bef74ddf0560ee3b4112321c5b6939e52a1629f0aed8aec46922/687474703a2f2f692e696d6775722e636f6d2f485379686177742e6a7067")
	/**/	}
	/**/
	/**/	if m.Content == ".ch" {
	/**/		s.ChannelMessageSend(m.ChannelID, "https://thumbs.gfycat.com/AdmirableLikableFlea-mobile.mp4")
	/**/	}
	/**/
	/**/
	/**/
	/**/
	/**/
	/**/
	/**/
	/**/
	if m.Content == ".m" {
			s.ChannelMessageSend(m.ChannelID, "pues...")
			time.Sleep(100 * time.Millisecond)
			if is_started {
				s.ChannelMessageEdit(m.ChannelID, último_mensaje_del_bot_ID, "está empezado")
			} else {
				s.ChannelMessageEdit(m.ChannelID, último_mensaje_del_bot_ID, "no está empezado")
			}
			println(m.Message.ID)
			println("ajsd: " + último_mensaje_del_bot_ID)
			s.ChannelMessageSendEmbed(m.ChannelID, embed.NewGenericEmbedAdvanced("¡Has terminado!", "Tu resultado es: ", 888))

	}	



}

