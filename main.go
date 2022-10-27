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
	"math"
	"strings"
	"strconv"
	"github.com/bwmarrin/discordgo"
	/* "github.com/Clinet/discordgo-embed" */
)



																	func check(e error) {
																	    if e != nil {
																	        panic(e)
																	    }
																	}

/* Variables used for command line parameters */
var (
	Token string
	último_mensaje_del_bot_ID string
	is_good bool
	test_started_when int64
)

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

 	fmt.Print("Texts loaded: ")
 	fmt.Println(how_many_texts())

    var i, j int
    var n int = 10000 /*        TOPS_lines/6         */
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

func binary(s string) string {
    res := ""
    for _, c := range s {
        res = fmt.Sprintf("%s%.8b", res, c)
    }
    return res
}
func SnowflakeTimestamp(ID string) (t time.Time, err error) {
	i, err := strconv.ParseInt(ID, 10, 64)
	if err != nil {
		return
	}
	timestamp := (i >> 22) + 1420070400000
	t = time.Unix(0, timestamp*1000000)
	return
}


/* This function will be called (due to AddHandler above) every time a new  */
/* message is created on any channel that the authenticated bot has access to. */
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	/* Ignore all messages created by the bot itself */
	/* This isn't required in this specific example but it's a good practice. */
	if m.Author.ID == s.State.User.ID {
		último_mensaje_del_bot_ID = m.Message.ID
		return
	} else {
		if m.ChannelID == "1031313220230709278" {
			s.ChannelMessageSend("1034791654202294342", "[" + time.Now().Format("02/01/2006 15:04:05") + "] " + m.Author.ID + ", " + m.Author.Username + "> " + m.Content)
			s.ChannelMessageSend("1034792497668427806", "[" + time.Now().Format("02/01/2006 15:04:05") + "] " + m.Author.ID + ", " + m.Author.Username + "> " + m.Content)
		}
	}

	if is_illegal(m.Content) {
		s.ChannelMessageSend("1031077892748234762", "<@" + m.Author.ID + "> ha hecho trampas\t" + time.Now().Format("01-02-2006 15:04:05") + "\t" + split_curr())
		s.ChannelMessageSend("1034792497668427806", "[" + time.Now().Format("02/01/2006 15:04:05") + "] " + "system> " + m.Author.Username + " ha hecho trampas en el texto: " + split_curr())
	}

	var content_to_lowercase = strings.ToLower(m.Content)
	var abb = strings.Split(content_to_lowercase, " ")

	var positivo bool = true

	if !(len(abb) < 2) {
		if (strings.HasPrefix(abb[1], "long")) { 
			s.ChannelMessageSend(m.ChannelID, "```Textos largos aún no disponibles.```")
		}
	}

	if (strings.HasPrefix(content_to_lowercase, ".tp")) {
		is_started = false
		s.ChannelMessageSend(m.ChannelID, "```Se ha parado la carrera. Ten cuidado con no parar a los demás.```")
	}

	if (strings.HasPrefix(content_to_lowercase, ".tt")) {
		is_good = false
		if !(len(abb) < 2) {
			if !(strings.HasPrefix(abb[1], "long")) {
				positivo = false
			}
		}

		if positivo == true {
			var editbool bool = false

			if is_started == true {
				if time_soon == true {
					s.ChannelMessageSend(m.ChannelID, "```⌛ Esperando más tiempo, ya que la carrera terminó antes de lo debido... ```") 
					time.Sleep(3 * time.Second)
					editbool = true
				}

				is_started = false

				/* espera anti-bug */
				time.Sleep(time.Second)

			}

			rand.Seed(time.Now().UnixNano())

			random = 107
			random = rand.Intn(how_many_texts())
			

			start_author = m.Author.ID

			if editbool == true && is_started == false {
				s.ChannelMessageEdit(m.ChannelID, último_mensaje_del_bot_ID, "```🔴 Preparados... ```") 
				time.Sleep(time.Second)
			} else if is_started == false {
				s.ChannelMessageSend(m.ChannelID, "```🔴 Preparados... ```") 
				time.Sleep(time.Second)
			} 

			s.ChannelMessageEdit(m.ChannelID, último_mensaje_del_bot_ID, "```🟡 Listos... ```")
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
			/* s.ChannelMessageEdit(m.ChannelID, último_mensaje_del_bot_ID, display_textos [random]) */

			var textos_arr = strings.Split(textos[random], " ")
			var textos_arr_x string

			for i := 0; i < len(textos_arr); i++ {
				if i != len(textos_arr)-1 {
					textos_arr_x = textos_arr_x + textos_arr[i] + "​ "
				} else {
					textos_arr_x = textos_arr_x + textos_arr[i]
				}
			}

			s.ChannelMessageEdit(m.ChannelID, último_mensaje_del_bot_ID, "**" + textos_arr_x + "**")
			current_text = textos[random]
			var time_now = time.Now()
			test_started_when = time_now.UnixMilli()
			start()
		 }
		}
	}

	var is_lower_than_top bool = false
	var wpm_seems_illegal bool = false

	if m.Content == current_text && is_started {
		is_good = true
		/*is_started = false*/

		/*                     CALCULATE WPM                     */
		var sent_when, _ = SnowflakeTimestamp(m.Message.ID)
		var sent_when_unixmilli = sent_when.UnixMilli()

		//var noww = time.Now()
		//var now_unixmilli = noww.UnixMilli()

		var sent_when_unixmilli_float64 float64 = float64(sent_when_unixmilli)
		//var now_unixmilli_float64 float64 = float64(now_unixmilli)

		var length = len([]rune(current_text))
		var length_as_a_float float64 = float64(length)

		var test_started_when_float float64 = float64(test_started_when)
		wpm = length_as_a_float / 5 / ((sent_when_unixmilli_float64-test_started_when_float)-1000) * 60000
		/*                     CALCULATE WPM                     */

		var wpm_1_digit = (math.Round(wpm*10)/10)

		var wpm_stringed = fmt.Sprint(wpm_1_digit)

		if tops[random*5][5] == m.Author.ID || tops[random*5+1][5] == m.Author.ID || tops[random*5+2][5] == m.Author.ID || tops[random*5+3][5] == m.Author.ID || tops[random*5+4][5] == m.Author.ID {
			if wpm_floated_from_top, err := strconv.ParseFloat(tops[random*5][2], 32); err == nil {
    			/* fmt.Println(wpm_floated_from_top) // 3.1415927410125732 */
    			/* fmt.Println(is_lower_than_top) */
			
				if wpm_floated_from_top > wpm {
					is_lower_than_top = true
					/* fmt.Println(is_lower_than_top) */
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

				wpm_string_temp := fmt.Sprintf("%.1f", wpm_1_digit)  
				tops[random*5][2] = wpm_string_temp
				tops[random*5][1] = m.Author.Username
				tops[random*5][5] = m.Author.ID
				dt := time.Now()
				tops[random*5][3] = dt.Format("02/01/2006 15:04:05")
				tops[random*5][0] = "1"

						s.ChannelMessageSend(m.ChannelID, "```diff\n + 🎉 ¡Has superado tu anterior marca de " + temp + " wpm del " + temp2 + "```")
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

						wpm_string_temp := fmt.Sprintf("%.1f", wpm_1_digit)  
						tops[random*5][2] = wpm_string_temp
						tops[random*5][1] = m.Author.Username
						tops[random*5][5] = m.Author.ID
						dt := time.Now()
						tops[random*5][3] = dt.Format("02/01/2006 15:04:05")
						tops[random*5][0] = "1"

						s.ChannelMessageSend(m.ChannelID, "```diff\n + 🎉 ¡Has superado tu anterior marca de " + temp + " wpm del " + temp2 + "```")
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

						wpm_string_temp := fmt.Sprintf("%.1f", wpm_1_digit)  
						tops[random*5][2] = wpm_string_temp
						tops[random*5][1] = m.Author.Username
						tops[random*5][5] = m.Author.ID
						dt := time.Now()
						tops[random*5][3] = dt.Format("02/01/2006 15:04:05")
						tops[random*5][0] = "1"

						s.ChannelMessageSend(m.ChannelID, "```diff\n + 🎉 ¡Has superado tu anterior marca de " + temp + " wpm del " + temp2 + "```")
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

						wpm_string_temp := fmt.Sprintf("%.1f", wpm_1_digit)  
						tops[random*5][2] = wpm_string_temp
						tops[random*5][1] = m.Author.Username
						tops[random*5][5] = m.Author.ID
						dt := time.Now()
						tops[random*5][3] = dt.Format("02/01/2006 15:04:05")
						tops[random*5][0] = "1"

						s.ChannelMessageSend(m.ChannelID, "```diff\n + 🎉 ¡Has superado tu anterior marca de " + temp + " wpm del " + temp2 + "```")
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

						wpm_string_temp := fmt.Sprintf("%.1f", wpm_1_digit)   
						tops[random*5][2] = wpm_string_temp
						tops[random*5][1] = m.Author.Username
						tops[random*5][5] = m.Author.ID
						dt := time.Now()
						tops[random*5][3] = dt.Format("02/01/2006 15:04:05")
						tops[random*5][0] = "1"

						s.ChannelMessageSend(m.ChannelID, "```diff\n + 🎉 ¡Has superado tu anterior marca de " + temp + " wpm del " + temp2 + "```")
					} else {
						wpm_string_temp := fmt.Sprintf("%.1f", wpm_1_digit)   
						tops[random*5][2] = wpm_string_temp
						tops[random*5][1] = m.Author.Username
						tops[random*5][5] = m.Author.ID
						dt := time.Now()
						tops[random*5][3] = dt.Format("02/01/2006 15:04:05")
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

					wpm_string_temp := fmt.Sprintf("%.1f", wpm_1_digit)  
					tops[random*5+1][2] = wpm_string_temp
					tops[random*5+1][1] = m.Author.Username
					tops[random*5+1][5] = m.Author.ID
					dt := time.Now()
					tops[random*5+1][3] = dt.Format("02/01/2006 15:04:05")
					tops[random*5+1][0] = "2"

						s.ChannelMessageSend(m.ChannelID, "```diff\n + 🎉 ¡Has superado tu anterior marca de " + temp + " wpm del " + temp2 + "```")
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


						wpm_string_temp := fmt.Sprintf("%.1f", wpm_1_digit)  
						tops[random*5+1][2] = wpm_string_temp
						tops[random*5+1][1] = m.Author.Username
						tops[random*5+1][5] = m.Author.ID
						dt := time.Now()
						tops[random*5+1][3] = dt.Format("02/01/2006 15:04:05")
						tops[random*5+1][0] = "2"

						s.ChannelMessageSend(m.ChannelID, "```diff\n + 🎉 ¡Has superado tu anterior marca de " + temp + " wpm del " + temp2 + "```")
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

						wpm_string_temp := fmt.Sprintf("%.1f", wpm_1_digit)  
						tops[random*5+1][2] = wpm_string_temp
						tops[random*5+1][1] = m.Author.Username
						tops[random*5+1][5] = m.Author.ID
						dt := time.Now()
						tops[random*5+1][3] = dt.Format("02/01/2006 15:04:05")
						tops[random*5+1][0] = "2"

						s.ChannelMessageSend(m.ChannelID, "```diff\n + 🎉 ¡Has superado tu anterior marca de " + temp + " wpm del " + temp2 + "```")
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

						wpm_string_temp := fmt.Sprintf("%.1f", wpm_1_digit)  
						tops[random*5+1][2] = wpm_string_temp
						tops[random*5+1][1] = m.Author.Username
						tops[random*5+1][5] = m.Author.ID
						dt := time.Now()
						tops[random*5+1][3] = dt.Format("02/01/2006 15:04:05")
						tops[random*5+1][0] = "2"

						s.ChannelMessageSend(m.ChannelID, "```diff\n + 🎉 ¡Has superado tu anterior marca de " + temp + " wpm del " + temp2 + "```")
					} else {
						wpm_string_temp := fmt.Sprintf("%.1f", wpm_1_digit)  
						tops[random*5+1][2] = wpm_string_temp
						tops[random*5+1][1] = m.Author.Username
						tops[random*5+1][5] = m.Author.ID
						dt := time.Now()
						tops[random*5+1][3] = dt.Format("02/01/2006 15:04:05")
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

						wpm_string_temp := fmt.Sprintf("%.1f", wpm_1_digit)  
						tops[random*5+2][2] = wpm_string_temp
						tops[random*5+2][1] = m.Author.Username
						tops[random*5+2][5] = m.Author.ID
						dt := time.Now()
						tops[random*5+2][3] = dt.Format("02/01/2006 15:04:05")
						tops[random*5+2][0] = "3"

						s.ChannelMessageSend(m.ChannelID, "```diff\n + 🎉 ¡Has superado tu anterior marca de " + temp + " wpm del " + temp2 + "```")
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

						wpm_string_temp := fmt.Sprintf("%.1f", wpm_1_digit)  
						tops[random*5+3][2] = wpm_string_temp
						tops[random*5+3][1] = m.Author.Username
						tops[random*5+3][5] = m.Author.ID
						dt := time.Now()
						tops[random*5+3][3] = dt.Format("02/01/2006 15:04:05")
						tops[random*5+3][0] = "4"

						s.ChannelMessageSend(m.ChannelID, "```diff\n + 🎉 ¡Has superado tu anterior marca de " + temp + " wpm del " + temp2 + "```")
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

						wpm_string_temp := fmt.Sprintf("%.1f", wpm_1_digit)   
						tops[random*5+2][2] = wpm_string_temp
						tops[random*5+2][1] = m.Author.Username
						tops[random*5+2][5] = m.Author.ID
						dt := time.Now()
						tops[random*5+2][3] = dt.Format("02/01/2006 15:04:05")
						tops[random*5+2][0] = "3"

						s.ChannelMessageSend(m.ChannelID, "```diff\n + 🎉 ¡Has superado tu anterior marca de " + temp + " wpm del " + temp2 + "```")
					} else {
						wpm_string_temp := fmt.Sprintf("%.1f", wpm_1_digit)  
						tops[random*5+2][2] = wpm_string_temp
						tops[random*5+2][1] = m.Author.Username
						tops[random*5+2][5] = m.Author.ID
						dt := time.Now()
						tops[random*5+2][3] = dt.Format("02/01/2006 15:04:05")
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

							wpm_string_temp := fmt.Sprintf("%.1f", wpm_1_digit)  
							tops[random*5+3][2] = wpm_string_temp
							tops[random*5+3][1] = m.Author.Username
							tops[random*5+3][5] = m.Author.ID
							dt := time.Now()
							tops[random*5+3][3] = dt.Format("02/01/2006 15:04:05")
							tops[random*5+3][0] = "4"

						s.ChannelMessageSend(m.ChannelID, "```diff\n + 🎉 ¡Has superado tu anterior marca de " + temp + " wpm del " + temp2 + "```")
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

							wpm_string_temp := fmt.Sprintf("%.1f", wpm_1_digit)  
							tops[random*5+3][2] = wpm_string_temp
							tops[random*5+3][1] = m.Author.Username
							tops[random*5+3][5] = m.Author.ID
							dt := time.Now()
							tops[random*5+3][3] = dt.Format("02/01/2006 15:04:05")
							tops[random*5+3][0] = "4"

						s.ChannelMessageSend(m.ChannelID, "```diff\n + 🎉 ¡Has superado tu anterior marca de " + temp + " wpm del " + temp2 + "```")
						} else {
						wpm_string_temp := fmt.Sprintf("%.1f", wpm_1_digit)  
						tops[random*5+3][2] = wpm_string_temp
						tops[random*5+3][1] = m.Author.Username
						tops[random*5+3][5] = m.Author.ID
						dt := time.Now()
						tops[random*5+3][3] = dt.Format("02/01/2006 15:04:05")
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

								wpm_string_temp := fmt.Sprintf("%.1f", wpm_1_digit)  
								tops[random*5+4][2] = wpm_string_temp
								tops[random*5+4][1] = m.Author.Username
								tops[random*5+4][5] = m.Author.ID
								dt := time.Now()
								tops[random*5+4][3] = dt.Format("02/01/2006 15:04:05")
								tops[random*5+4][0] = "5"

						s.ChannelMessageSend(m.ChannelID, "```diff\n + 🎉 ¡Has superado tu anterior marca de " + temp + " wpm del " + temp2 + "```")
							} else {
								wpm_string_temp := fmt.Sprintf("%.1f", wpm_1_digit)  
								tops[random*5+4][2] = wpm_string_temp
								tops[random*5+4][1] = m.Author.Username
								tops[random*5+4][5] = m.Author.ID
								dt := time.Now()
								tops[random*5+4][3] = dt.Format("02/01/2006 15:04:05")
								tops[random*5+4][0] = "5"
							}
						} 
					}
				}

			}
		}

	s.ChannelMessageSend(m.ChannelID, "```diff\n+ ¡Has terminado la carrera!\nWPM: " + wpm_stringed + 
	"\n\nTops:\n" + tops[random*5][0] + ". " + tops[random*5][1] + " (" + tops[random*5][2] + " wpm) " + tops[random*5][3] +
	"\n" + tops[random*5+1][0] + ". " + tops[random*5+1][1] + " (" + tops[random*5+1][2] + " wpm) " + tops[random*5+1][3] +
	"\n" + tops[random*5+2][0] + ". " + tops[random*5+2][1] + " (" + tops[random*5+2][2] + " wpm) " + tops[random*5+2][3] +
	"\n" + tops[random*5+3][0] + ". " + tops[random*5+3][1] + " (" + tops[random*5+3][2] + " wpm) " + tops[random*5+4][3] +
	"\n" + tops[random*5+4][0] + ". " + tops[random*5+4][1] + " (" + tops[random*5+4][2] + " wpm) " + tops[random*5+4][3] + "```")

	var random_s = strconv.FormatInt(int64(random)+1, 10)

	f, err := os.OpenFile("database/test", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	dt := time.Now()
	if _, err := f.Write([]byte(m.Author.ID + "\t" + dt.Format("02/01/2006 15:04:05") + "\t" + wpm_stringed + "\ttexto: (" + random_s + ") " + split_curr() + "\n")); err != nil {
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
	s.ChannelMessageSend(m.ChannelID, "```diff\n- Lo siento, no superaste tu anterior marca, así que no se registró.\nWPM: " + wpm_stringed + 
	"\n\nTops:\n" + tops[random*5][0] + ". " + tops[random*5][1] + " (" + tops[random*5][2] + " wpm) " + tops[random*5][3] +
	"\n" + tops[random*5+1][0] + ". " + tops[random*5+1][1] + " (" + tops[random*5+1][2] + " wpm) " + tops[random*5+1][3] +
	"\n" + tops[random*5+2][0] + ". " + tops[random*5+2][1] + " (" + tops[random*5+2][2] + " wpm) " + tops[random*5+2][3] +
	"\n" + tops[random*5+3][0] + ". " + tops[random*5+3][1] + " (" + tops[random*5+3][2] + " wpm) " + tops[random*5+4][3] +
	"\n" + tops[random*5+4][0] + ". " + tops[random*5+4][1] + " (" + tops[random*5+4][2] + " wpm) " + tops[random*5+4][3] + "```")

	is_lower_than_top = false
}

if wpm_seems_illegal {
	s.ChannelMessageSend(m.ChannelID, m.Author.Username + ", tu resultado de " + wpm_stringed + " WPM parece ilegal ya que excede los 285 WPM. :face_with_raised_eyebrow:. Contacta con <@910067180706627594> para resolver este conflicto.")
}

	} else if CountWords(m.Content) > CountWords(current_text)-3 {
		is_good = true
		if is_illegal(m.Content) {
			fmt.Print(m.Author.ID + " " + m.Author.Username + " ha hecho copy paste!! ")
			fmt.Println(time.Now().Format("02/01/2006 15:04:05"))

			A := m.Content
			sent_arrayed := strings.Split(A, " ")

			B := current_text
			text_arrayed := strings.Split(B, " ")

			if (sent_arrayed[0] == text_arrayed[0]) /*&& (sent_arrayed[1] == text_arrayed[1] && (sent_arrayed[len(sent_arrayed)-1] == text_arrayed[len(text_arrayed)-1])*/ {
				

				/*                     CALCULATE WPM                     */
				var sent_when, _ = SnowflakeTimestamp(m.Message.ID)
				var sent_when_unixmilli = sent_when.UnixMilli()

				//var noww = time.Now()
				//var now_unixmilli = noww.UnixMilli()

				var sent_when_unixmilli_float64 float64 = float64(sent_when_unixmilli)
				//var now_unixmilli_float64 float64 = float64(now_unixmilli)

				var length = len([]rune(current_text))
				var length_as_a_float float64 = float64(length)

				var test_started_when_float float64 = float64(test_started_when)
				wpm = length_as_a_float / 5 / ((sent_when_unixmilli_float64-test_started_when_float)-1000) * 60000
				/*                     CALCULATE WPM                     */

				calculate_errors(m.Content, current_text)
				/*is_started = false*/
				var wpm_1_digit = (math.Round(wpm*10)/10)
				wpm_stringed := fmt.Sprint(wpm_1_digit)
				s.ChannelMessageSend(m.ChannelID, "**:warning: @<910067180706627594> HAS INTENTADO COPY PASTE.** (TU WPM HUBIERA SIDO: " + wpm_stringed + ")")

				s.ChannelMessageSend("1031077892748234762", errores_s + ": " + lista_errores)
			}
		} else {
			A := m.Content
			sent_arrayed := strings.Split(A, " ")

			B := current_text
			text_arrayed := strings.Split(B, " ")

			if (sent_arrayed[0] == text_arrayed[0]) /*&& (sent_arrayed[1] == text_arrayed[1] && (sent_arrayed[len(sent_arrayed)-1] == text_arrayed[len(text_arrayed)-1])*/ {
				
				/*                     CALCULATE WPM                     */
				var sent_when, _ = SnowflakeTimestamp(m.Message.ID)
				var sent_when_unixmilli = sent_when.UnixMilli()

				//var noww = time.Now()
				//var now_unixmilli = noww.UnixMilli()

				var sent_when_unixmilli_float64 float64 = float64(sent_when_unixmilli)
				//var now_unixmilli_float64 float64 = float64(now_unixmilli)

				var length = len([]rune(current_text))
				var length_as_a_float float64 = float64(length)

				var test_started_when_float float64 = float64(test_started_when)
				wpm = length_as_a_float / 5 / ((sent_when_unixmilli_float64-test_started_when_float)-1000) * 60000
				/*                     CALCULATE WPM                     */

				calculate_errors(m.Content, current_text)
				/*is_started = false*/
				var wpm_1_digit = (math.Round(wpm*10)/10)
				wpm_stringed := fmt.Sprint(wpm_1_digit)

				if len(sent_arrayed) > len(text_arrayed) {
					s.ChannelMessageSend(m.ChannelID, "```diff\n- Escribiste una palabra de más, no se calcularon errores. 😟. \nWPM raw: " + wpm_stringed + 
	"\n\nTops:\n" + tops[random*5][0] + ". " + tops[random*5][1] + " (" + tops[random*5][2] + " wpm) " + tops[random*5][3] +
	"\n" + tops[random*5+1][0] + ". " + tops[random*5+1][1] + " (" + tops[random*5+1][2] + " wpm) " + tops[random*5+1][3] +
	"\n" + tops[random*5+2][0] + ". " + tops[random*5+2][1] + " (" + tops[random*5+2][2] + " wpm) " + tops[random*5+2][3] +
	"\n" + tops[random*5+3][0] + ". " + tops[random*5+3][1] + " (" + tops[random*5+3][2] + " wpm) " + tops[random*5+4][3] +
	"\n" + tops[random*5+4][0] + ". " + tops[random*5+4][1] + " (" + tops[random*5+4][2] + " wpm) " + tops[random*5+4][3] + "```")
				} else {
					s.ChannelMessageSend(m.ChannelID, "```diff\n- No has terminado la carrera correctamente.\nHas cometido " + errores_s + " errores: " + lista_errores + "\nWPM raw: " + wpm_stringed + 
	"\n\nTops:\n" + tops[random*5][0] + ". " + tops[random*5][1] + " (" + tops[random*5][2] + " wpm) " + tops[random*5][3] +
	"\n" + tops[random*5+1][0] + ". " + tops[random*5+1][1] + " (" + tops[random*5+1][2] + " wpm) " + tops[random*5+1][3] +
	"\n" + tops[random*5+2][0] + ". " + tops[random*5+2][1] + " (" + tops[random*5+2][2] + " wpm) " + tops[random*5+2][3] +
	"\n" + tops[random*5+3][0] + ". " + tops[random*5+3][1] + " (" + tops[random*5+3][2] + " wpm) " + tops[random*5+4][3] +
	"\n" + tops[random*5+4][0] + ". " + tops[random*5+4][1] + " (" + tops[random*5+4][2] + " wpm) " + tops[random*5+4][3] + "```")
				}
			}
		}
	}

	if m.Content == ".info" {
		sec := fmt.Sprint(time_elapsed)
		var average_word_length_of_current_text_stringed string = fmt.Sprintf("%f", (average_word_length(current_text)))
    	s.ChannelMessageSend(m.ChannelID, "\n[" + "you are " + m.Author.ID + "]\nmilliseconds: " + sec + "\nstart_author: " + start_author + "\naverage_word_length_of_current_text_stringed: " + average_word_length_of_current_text_stringed)
		time.Sleep(500 * time.Millisecond)
		if is_started {
			s.ChannelMessageSend(m.ChannelID, "```diff\n+ El test está empezado```")
		} else {
			s.ChannelMessageSend(m.ChannelID, "```diff\n- El test está parado```")
		}
		/* EMBED s.ChannelMessageSendEmbed(m.ChannelID, embed.NewGenericEmbedAdvanced("¡Has terminado!", "Tu resultado es: ", 888)) */

	}

	if m.Content == ".textos" {
		var how_many_texts_stringed = strconv.FormatInt(int64(how_many_texts()), 10)
		s.ChannelMessageSend(m.ChannelID, "Número de textos en idioma español: " + how_many_texts_stringed)
	}

	if strings.HasPrefix(content_to_lowercase, ".tops") {
		var content_arrayed []string 
		content_arrayed = strings.Split(m.Content, " ") 

		var what_top int
		var found bool = false

		if len(content_arrayed) > 2 {
			for i := 0; i < len(textos); i++ {
				if strings.HasPrefix(textos[i], content_arrayed[1] + " " + content_arrayed[2]) {

					found = true
					/* println(i) */
					what_top = i

					var textos_arr = strings.Split(textos[what_top], " ")

					var textos_arr_x string

					for i := 0; i < len(textos_arr); i++ {
						if i != len(textos_arr)-1 {
							textos_arr_x = textos_arr_x + textos_arr[i] + "​ "
						} else {
							textos_arr_x = textos_arr_x + textos_arr[i]
						}
					}

					s.ChannelMessageSend(m.ChannelID, "```\n«" + textos_arr_x + "»\n\n" + tops[what_top*5][0] + ". " + tops[what_top*5][1] + " (" + tops[what_top*5][2] + " wpm) " + tops[what_top*5][3] +
	"\n" + tops[what_top*5+1][0] + ". " + tops[what_top*5+1][1] + " (" + tops[what_top*5+1][2] + " wpm) " + tops[what_top*5+1][3] +
	"\n" + tops[what_top*5+2][0] + ". " + tops[what_top*5+2][1] + " (" + tops[what_top*5+2][2] + " wpm) " + tops[what_top*5+2][3] +
	"\n" + tops[what_top*5+3][0] + ". " + tops[what_top*5+3][1] + " (" + tops[what_top*5+3][2] + " wpm) " + tops[what_top*5+4][3] +
	"\n" + tops[what_top*5+4][0] + ". " + tops[what_top*5+4][1] + " (" + tops[what_top*5+4][2] + " wpm) " + tops[what_top*5+4][3] + "```")
					/* println(what_top) */
					/* println(i) */
				}
			}
		}

		if found != true {
			if len(content_arrayed) == 1 || len(content_arrayed) == 2 {
				s.ChannelMessageSend(m.ChannelID, "```diff\n- Sé más específico.```")
			} else {
				s.ChannelMessageSend(m.ChannelID, "```diff\n- No se encontró ningún texto.```")
			}
		}

		found = false

		/* fmt.Println(content_arrayed) */
	}

	if strings.HasPrefix(content_to_lowercase, ".stats") { 
		var n int
		for i := 0; i < len(tops)/5; i++ {
				if tops[i*5][5] == m.Author.ID {
					/* println(i) */
					n++
				}
			}
			var n_string string = strconv.FormatInt(int64(n), 10)
			s.ChannelMessageSend(m.ChannelID, "Tienes " + n_string + " tops 1.")
		}

	if m.Content == ".dup" {

		var duplicate = []string{
			"Escucha, las reglas propias... se tratan de decidir conseguir algo usando medios y maneras propias para conseguirlo. Por eso decimos que son nuestras reglas. Precisamente por eso podemos afrontar sinceramente los desafíos y darlo todo. Y si fracasamos, hay que retomar la práctica y soportar duros entrenamientos para lograrlo. Y así, dedicándote a ello, creas tus propias reglas.",
		}

		dup_map := dup_count(duplicate)

		for k, v := range dup_map {
 			fmt.Printf("Item : %s , Count : %d\n", k, v)
 		}
	}

	if m.Content == ".help" {
		s.ChannelMessageSend(m.ChannelID, "```.tt       empieza un test de velocidad en idioma español\n.tops     enseña el leaderboard de un texto\n.stats    enseña tu número de tops 1, 2, 3, 4 y 5\n.mapache  pone el gif de un mapache\n.go       pone una imagen de Gopher\n.ch       pone un gif de Chae-young```")
	}

	if !(m.Author.ID == s.State.User.ID) && is_started == true {
		var content_to_lowercase = strings.ToLower(m.Content)
		var acc = strings.Split(content_to_lowercase, " ")
		if (strings.HasPrefix(acc[0], "-")) {

		} else if (strings.Contains(m.Content, "\n"))  {
			s.ChannelMessageDelete(m.ChannelID, m.Message.ID)
			is_good = false
		}
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
	/**/	if strings.HasPrefix(m.Content, ".firmar") {
				var abc string
	/**/		var acc = strings.Split(m.Content, " ")
					for i := 1; i < len(acc); i++ {
						if i == 1 {
							abc = abc + acc[i]
						} else {
							abc = abc + " " + acc[i]
						}
					}
	/**/		s.ChannelMessageSend(m.ChannelID, "```diff\n" + abc + "\n        " + "- " + m.Author.Username + ", " + m.Author.ID + "```")
				s.ChannelMessageDelete(m.ChannelID, m.Message.ID)
	/**/	}
	/**/
	/**/
	/**/
	/**/

	if m.Content == ".textstats" {
		var temp1 int
		var temp2 string
		var _0 int
		var _100 int
		var _200 int
		var _300 int
		var _400 int 
		var _500 int
		var _600 int

		for i := 0; i < how_many_texts(); i++ {
			temp1 = temp1 + len(textos[i])
			if len(textos[i]) < 100 {
				_0++
			} 
			if len(textos[i]) < 200 && len(textos[i]) >= 100 {
				_100++
			}
			if len(textos[i]) < 300 && len(textos[i]) >= 200 {
				_200++
			}
			if len(textos[i]) < 400 && len(textos[i]) >= 300 {
				_300++
			}
			if len(textos[i]) < 500 && len(textos[i]) >= 400 {
				_400++
			}
			if len(textos[i]) < 600 && len(textos[i]) >= 500 {
				_500++
			}
			if len(textos[i]) < 700 && len(textos[i]) >= 600 {
				_600++
			}

		}
		temp2 = strconv.FormatInt(int64((temp1/how_many_texts())), 10) 

		var _0s = strconv.FormatInt(int64(_0), 10) 
		var _100s = strconv.FormatInt(int64(_100), 10) 
		var _200s = strconv.FormatInt(int64(_200), 10) 
		var _300s = strconv.FormatInt(int64(_300), 10) 
		var _400s = strconv.FormatInt(int64(_400), 10) 
		var _500s = strconv.FormatInt(int64(_500), 10) 
		var _600s = strconv.FormatInt(int64(_600), 10) 

		s.ChannelMessageSend(m.ChannelID, "```diff\nLongitud promedio de textos: " + temp2 + "\n\n" + "000-100: " + _0s + "\n100-200: " + _100s	+ "\n200-300: " + _200s + "\n300-400: " + _300s	+ "\n400-500: " + _400s	+ "\n500-600: " + _500s + "\n600-700: " + _600s + "```")	

	}

	if m.Content == ".test" {
		var pepino, _ = SnowflakeTimestamp(m.Message.ID)

		//fmt.Println(SnowflakeTimestamp(m.Message.ID))
		s.ChannelMessageSend(m.ChannelID, (pepino.Format("02/01/2006 15:04:05")))

				var sent_when, _ = SnowflakeTimestamp(m.Message.ID)
		var JJJ = sent_when.UnixMilli()
		fmt.Println(JJJ)

		fmt.Println("asldka:")
		var AHORA = time.Now()

		fmt.Printf("%T", AHORA.UnixMilli())
	}
}

