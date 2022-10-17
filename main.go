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

	 var fileName = "database/TOPS"
 	fileBytes, err := ioutil.ReadFile(fileName)

 	if err != nil {
 		fmt.Println(err)
 		os.Exit(1)
 	}

 	sliceData := strings.Split(string(fileBytes), "\n")

 	fmt.Println(sliceData[21])
 	fmt.Println(how_many_texts())

	tops[0][0] = sliceData[0]	; tops[0][1] = sliceData[1]	; tops[0][2] = sliceData[2]	; tops[0][3] = sliceData[3]	; tops[0][4] = sliceData[4]	; tops[0][5] = sliceData[5]	; tops[1][0] = sliceData[6]	; tops[1][1] = sliceData[7]	; tops[1][2] = sliceData[8]	; tops[1][3] = sliceData[9]	; tops[1][4] = sliceData[10]	; tops[1][5] = sliceData[11]	; tops[2][0] = sliceData[12]	; tops[2][1] = sliceData[13]	; tops[2][2] = sliceData[14]	; tops[2][3] = sliceData[15]	; tops[2][4] = sliceData[16]	; tops[2][5] = sliceData[17]	; tops[3][0] = sliceData[18]	; tops[3][1] = sliceData[19]	; tops[3][2] = sliceData[20]	; tops[3][3] = sliceData[21]	; tops[3][4] = sliceData[22]	; tops[3][5] = sliceData[23]	; tops[4][0] = sliceData[24]	; tops[4][1] = sliceData[25]	; tops[4][2] = sliceData[26]	; tops[4][3] = sliceData[27]	; tops[4][4] = sliceData[28]	; tops[4][5] = sliceData[29]	; tops[5][0] = sliceData[30]	; tops[5][1] = sliceData[31]	; tops[5][2] = sliceData[32]	; tops[5][3] = sliceData[33]	; tops[5][4] = sliceData[34]	; tops[5][5] = sliceData[35]	; tops[6][0] = sliceData[36]	; tops[6][1] = sliceData[37]	; tops[6][2] = sliceData[38]	; tops[6][3] = sliceData[39]	; tops[6][4] = sliceData[40]	; tops[6][5] = sliceData[41]	; tops[7][0] = sliceData[42]	; tops[7][1] = sliceData[43]	; tops[7][2] = sliceData[44]	; tops[7][3] = sliceData[45]	; tops[7][4] = sliceData[46]	; tops[7][5] = sliceData[47]	; tops[8][0] = sliceData[48]	; tops[8][1] = sliceData[49]	; tops[8][2] = sliceData[50]	; tops[8][3] = sliceData[51]	; tops[8][4] = sliceData[52]	; tops[8][5] = sliceData[53]	; tops[9][0] = sliceData[54]	; tops[9][1] = sliceData[55]	; tops[9][2] = sliceData[56]	; tops[9][3] = sliceData[57]	; tops[9][4] = sliceData[58]	; tops[9][5] = sliceData[59]	; tops[10][0] = sliceData[60]	; tops[10][1] = sliceData[61]	; tops[10][2] = sliceData[62]	; tops[10][3] = sliceData[63]	; tops[10][4] = sliceData[64]	; tops[10][5] = sliceData[65]	; tops[11][0] = sliceData[66]	; tops[11][1] = sliceData[67]	; tops[11][2] = sliceData[68]	; tops[11][3] = sliceData[69]	; tops[11][4] = sliceData[70]	; tops[11][5] = sliceData[71]	; tops[12][0] = sliceData[72]	; tops[12][1] = sliceData[73]	; tops[12][2] = sliceData[74]	; tops[12][3] = sliceData[75]	; tops[12][4] = sliceData[76]	; tops[12][5] = sliceData[77]	; tops[13][0] = sliceData[78]	; tops[13][1] = sliceData[79]	; tops[13][2] = sliceData[80]	; tops[13][3] = sliceData[81]	; tops[13][4] = sliceData[82]	; tops[13][5] = sliceData[83]	; tops[14][0] = sliceData[84]	; tops[14][1] = sliceData[85]	; tops[14][2] = sliceData[86]	; tops[14][3] = sliceData[87]	; tops[14][4] = sliceData[88]	; tops[14][5] = sliceData[89]	; tops[15][0] = sliceData[90]	; tops[15][1] = sliceData[91]	; tops[15][2] = sliceData[92]	; tops[15][3] = sliceData[93]	; tops[15][4] = sliceData[94]	; tops[15][5] = sliceData[95]	; tops[16][0] = sliceData[96]	; tops[16][1] = sliceData[97]	; tops[16][2] = sliceData[98]	; tops[16][3] = sliceData[99]	; tops[16][4] = sliceData[100]	; tops[16][5] = sliceData[101]	; tops[17][0] = sliceData[102]	; tops[17][1] = sliceData[103]	; tops[17][2] = sliceData[104]	; tops[17][3] = sliceData[105]	; tops[17][4] = sliceData[106]	; tops[17][5] = sliceData[107]	; tops[18][0] = sliceData[108]	; tops[18][1] = sliceData[109]	; tops[18][2] = sliceData[110]	; tops[18][3] = sliceData[111]	; tops[18][4] = sliceData[112]	; tops[18][5] = sliceData[113]	; tops[19][0] = sliceData[114]	; tops[19][1] = sliceData[115]	; tops[19][2] = sliceData[116]	; tops[19][3] = sliceData[117]	; tops[19][4] = sliceData[118]	; tops[19][5] = sliceData[119]	; tops[20][0] = sliceData[120]	; tops[20][1] = sliceData[121]	; tops[20][2] = sliceData[122]	; tops[20][3] = sliceData[123]	; tops[20][4] = sliceData[124]	; tops[20][5] = sliceData[125]	; tops[21][0] = sliceData[126]	; tops[21][1] = sliceData[127]	; tops[21][2] = sliceData[128]	; tops[21][3] = sliceData[129]	; tops[21][4] = sliceData[130]	; tops[21][5] = sliceData[131]	; tops[22][0] = sliceData[132]	; tops[22][1] = sliceData[133]	; tops[22][2] = sliceData[134]	; tops[22][3] = sliceData[135]	; tops[22][4] = sliceData[136]	; tops[22][5] = sliceData[137]	; tops[23][0] = sliceData[138]	; tops[23][1] = sliceData[139]	; tops[23][2] = sliceData[140]	; tops[23][3] = sliceData[141]	; tops[23][4] = sliceData[142]	; tops[23][5] = sliceData[143]	; tops[24][0] = sliceData[144]	; tops[24][1] = sliceData[145]	; tops[24][2] = sliceData[146]	; tops[24][3] = sliceData[147]	; tops[24][4] = sliceData[148]	; tops[24][5] = sliceData[149]	; tops[25][0] = sliceData[150]	; tops[25][1] = sliceData[151]	; tops[25][2] = sliceData[152]	; tops[25][3] = sliceData[153]	; tops[25][4] = sliceData[154]	; tops[25][5] = sliceData[155]	; tops[26][0] = sliceData[156]	; tops[26][1] = sliceData[157]	; tops[26][2] = sliceData[158]	; tops[26][3] = sliceData[159]	; tops[26][4] = sliceData[160]	; tops[26][5] = sliceData[161]	; tops[27][0] = sliceData[162]	; tops[27][1] = sliceData[163]	; tops[27][2] = sliceData[164]	; tops[27][3] = sliceData[165]	; tops[27][4] = sliceData[166]	; tops[27][5] = sliceData[167]	; tops[28][0] = sliceData[168]	; tops[28][1] = sliceData[169]	; tops[28][2] = sliceData[170]	; tops[28][3] = sliceData[171]	; tops[28][4] = sliceData[172]	; tops[28][5] = sliceData[173]	; tops[29][0] = sliceData[174]	; tops[29][1] = sliceData[175]	; tops[29][2] = sliceData[176]	; tops[29][3] = sliceData[177]	; tops[29][4] = sliceData[178]	; tops[29][5] = sliceData[179]	; tops[30][0] = sliceData[180]	; tops[30][1] = sliceData[181]	; tops[30][2] = sliceData[182]	; tops[30][3] = sliceData[183]	; tops[30][4] = sliceData[184]	; tops[30][5] = sliceData[185]	; tops[31][0] = sliceData[186]	; tops[31][1] = sliceData[187]	; tops[31][2] = sliceData[188]	; tops[31][3] = sliceData[189]	; tops[31][4] = sliceData[190]	; tops[31][5] = sliceData[191]	; tops[32][0] = sliceData[192]	; tops[32][1] = sliceData[193]	; tops[32][2] = sliceData[194]	; tops[32][3] = sliceData[195]	; tops[32][4] = sliceData[196]	; tops[32][5] = sliceData[197]	; tops[33][0] = sliceData[198]	; tops[33][1] = sliceData[199]	; tops[33][2] = sliceData[200]	; tops[33][3] = sliceData[201]	; tops[33][4] = sliceData[202]	; tops[33][5] = sliceData[203]	;

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
		random = rand.Intn(how_many_texts())
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

		/*Set to float*/ if AA, err := strconv.ParseFloat(tops[random*5][2], 64); err == nil {AAA = AA}
		if wpm > AAA {
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

			s := fmt.Sprintf("%f", wpm)  
			tops[random*5][2] = s
			tops[random*5][1] = m.Author.Username
			tops[random*5][5] = m.Author.ID
			dt := time.Now()
			tops[random*5][3] = dt.Format("01-02-2006 15:04:05")
			tops[random*5][0] = "1"

		} else {
			/*Set to float*/ if AA, err := strconv.ParseFloat(tops[random*5+1][2], 64); err == nil {AAA = AA}
			if wpm > AAA {
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

				s := fmt.Sprintf("%f", wpm)  
				tops[random*5+1][2] = s
				tops[random*5+1][1] = m.Author.Username
				tops[random*5+1][5] = m.Author.ID
				dt := time.Now()
				tops[random*5+1][3] = dt.Format("01-02-2006 15:04:05")
				tops[random*5+1][0] = "2"
			} else {
				/*Set to float*/ if AA, err := strconv.ParseFloat(tops[random*5+2][2], 64); err == nil {AAA = AA}
				if wpm > AAA {
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

					s := fmt.Sprintf("%f", wpm)  
					tops[random*5+2][2] = s
					tops[random*5+2][1] = m.Author.Username
					tops[random*5+2][5] = m.Author.ID
					dt := time.Now()
					tops[random*5+2][3] = dt.Format("01-02-2006 15:04:05")
					tops[random*5+2][0] = "3"
				} else {
				/*Set to float*/ if AA, err := strconv.ParseFloat(tops[random*5+3][2], 64); err == nil {AAA = AA}
				if wpm > AAA {
					tops[random*5+4][2] = tops[random*5+3][2]
					tops[random*5+4][1] = tops[random*5+3][1]
					tops[random*5+4][5] = tops[random*5+3][5]
					tops[random*5+4][3] = tops[random*5+3][3]
					tops[random*5+4][0] = "5"

					s := fmt.Sprintf("%f", wpm)  
					tops[random*5+3][2] = s
					tops[random*5+3][1] = m.Author.Username
					tops[random*5+3][5] = m.Author.ID
					dt := time.Now()
					tops[random*5+3][3] = dt.Format("01-02-2006 15:04:05")
					tops[random*5+3][0] = "4"
				} else {
				/*Set to float*/ if AA, err := strconv.ParseFloat(tops[random*5+4][2], 64); err == nil {AAA = AA}
					if wpm > AAA {

						s := fmt.Sprintf("%f", wpm)  
						tops[random*5+4][2] = s
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
	/**/	if m.Content == ".go" {
	/**/		s.ChannelMessageSend(m.ChannelID, "https://camo.githubusercontent.com/833cfd306ac2bef74ddf0560ee3b4112321c5b6939e52a1629f0aed8aec46922/687474703a2f2f692e696d6775722e636f6d2f485379686177742e6a7067")
	/**/	}
	/**/

}

