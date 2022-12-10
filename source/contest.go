package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

type Current struct {
	Channel    string
	Started    float64
	Text_ID    int
	Requested  int64
	WPM        float64
	Errors     int
	Error_list string
}

var Currents []Current

var DB []string

var Texts []string

var Date string

var Text_message_ID string

var Random int

func Is_illegal(ζ string) bool {
	var π = strings.Contains(ζ, "\u200b")
	return π
}

func υ(Text string) string {

	var Texts_arr = strings.Split(Text, " ")
	var λ string
	for i := 0; i < len(Texts_arr); i++ {
		if i != len(Texts_arr)-1 {
			λ = λ + Texts_arr[i] + "\u200b "
		} else {
			λ = λ + Texts_arr[i]
		}
	}
	return λ
}

func Typing_test(s *discordgo.Session, m *discordgo.MessageCreate, Mode string) {
	var Now = time.Now().UnixMilli()
	var already_requested bool = true

	var exists = false
	for i, v := range Currents {
		if v.Channel == m.ChannelID {
			exists = true
			if Now-v.Requested < 3000 {
				already_requested = false
			}
			Currents[i].Requested = Now
		}
	}

	if !exists {
		Currents = append(Currents, Current{m.ChannelID, 1877777777, Random, Now, 0, 0, ""})
	}

	for i, v := range Currents {
		if v.Channel == m.ChannelID {
			if !already_requested {
				s.ChannelMessageSend(m.ChannelID, "```diff\n- Espera un poco.```")
			} else {
				Reset_typing_users()

				rand.Seed(time.Now().UnixNano())

				Currents[i].Channel = m.ChannelID
				Currents[i].Requested = Now

				switch Mode {
				case "Normal":
					Currents[i].Text_ID = rand.Intn(398-1) + 1
					var Test_message, _ = s.ChannelMessageSend(m.ChannelID, "```🔴 Preparados...```")
					Text_message_ID = Test_message.ID
					time.Sleep(1 * time.Second)
					s.ChannelMessageEdit(m.ChannelID, Test_message.ID, "```🟡 Listos...```")
					time.Sleep(1 * time.Second)

					var Text_shown, _ = s.ChannelMessageEdit(m.ChannelID, Test_message.ID, "**"+υ(Texts[Currents[i].Text_ID])+"**")
					_ = Text_shown
					var Started_when_time = Text_shown.EditedTimestamp
					var Started_when = Started_when_time.UnixMilli()

					Started_when = Started_when_time.UnixMilli()

					var Started_when_float float64 = float64(Started_when)

					Currents[i].Started = Started_when_float
				case "Short":
					Currents[i].Text_ID = rand.Intn(458-401) + 401

					var Test_message, _ = s.ChannelMessageSend(m.ChannelID, "```🔴🩳 (Textos cortos) Preparados...```")
					time.Sleep(1 * time.Second)
					s.ChannelMessageEdit(m.ChannelID, Test_message.ID, "```🟡🤏 (Textos cortos) Listos...```")
					time.Sleep(1 * time.Second)

					var Text_shown, _ = s.ChannelMessageEdit(m.ChannelID, Test_message.ID, "**"+υ(Texts[Currents[i].Text_ID])+"**")
					_ = Text_shown
					var Started_when_time = Text_shown.EditedTimestamp
					var Started_when = Started_when_time.UnixMilli()

					Started_when = Started_when_time.UnixMilli()

					var Started_when_float float64 = float64(Started_when)

					Currents[i].Started = Started_when_float
				case "Dev":
					Currents[i].Text_ID = rand.Intn(468-459) + 459

					var Test_message, _ = s.ChannelMessageSend(m.ChannelID, "```🔴💻 (Programación) Preparados...```")
					time.Sleep(1 * time.Second)
					s.ChannelMessageEdit(m.ChannelID, Test_message.ID, "```🟡👩🏾‍💻 (Programación) Listos...```")
					time.Sleep(1 * time.Second)

					var Text_shown, _ = s.ChannelMessageEdit(m.ChannelID, Test_message.ID, "**"+υ(Texts[Currents[i].Text_ID])+"**")
					_ = Text_shown
					var Started_when_time = Text_shown.EditedTimestamp
					var Started_when = Started_when_time.UnixMilli()

					Started_when = Started_when_time.UnixMilli()

					var Started_when_float float64 = float64(Started_when)

					Currents[i].Started = Started_when_float
				}
			}
		}
	}
}

func Judge(m *discordgo.MessageCreate, S string, Text_ID int) int8 {
	if S == Texts[Text_ID] {
		return 1
	}

	var Content_arr = strings.Split(m.Content, " ")
	var Text_arr = strings.Split(Texts[Text_ID], " ")

	if Content_arr[0] == Text_arr[0] {
		if len(m.Content) > len(Texts[Text_ID])-10 {
			return 2
		}

		if len(m.Content) > len(Texts[Text_ID])-30 {
			return 4
		}
	}
	return 3
}

func Calculate(m *discordgo.MessageCreate, Started_when_float float64, Text string) float64 {
	var sent_when, _ = SnowflakeTimestamp(m.Message.ID)
	Date = sent_when.Format("02/01/2006 15:04")
	var sent_when_unixmilli = sent_when.UnixMilli()

	var sent_when_unixmilli_float64 float64 = float64(sent_when_unixmilli)

	var length = len([]rune(Text))
	var length_as_a_float float64 = float64(length)

	var wpm = length_as_a_float / 5 / ((sent_when_unixmilli_float64 - Started_when_float) - 600) * 60000
	return wpm
}

/*
var Error_list string
var Errors int
var Errors_str string
*/
func Errors_calculate(sent string, current string) (int, string) {

	/* reseting */
	var Errors = 0
	var Error_list = ""

	A := sent
	sent_arrayed := strings.Split(A, " ")

	B := current
	text_arrayed := strings.Split(B, " ")

	if len(text_arrayed) == len(sent_arrayed) {
		for i := 0; i < len(text_arrayed); i++ {
			if text_arrayed[i] != sent_arrayed[i] {
				if Error_list != "" {
					Error_list = Error_list + ", " + sent_arrayed[i]
					Errors++
				} else {
					Error_list = sent_arrayed[i]
					Errors++
				}
			}
		}
	}

	/*
		if len(text_arrayed) > len(sent_arrayed) {

		}
	*/

	//Errors_str = strconv.FormatInt(int64(Errors), 10)

	return Errors, Error_list
}

var Delete_last_score_because_improved bool = false

func Show_result(s *discordgo.Session, m *discordgo.MessageCreate, wpm float64) {
	var wpm_rounded = (math.Round(wpm*10) / 10)
	var wpm_str = fmt.Sprint(wpm_rounded)

	if !Delete_last_score_because_improved {
		s.ChannelMessageSend(m.ChannelID, "```diff\n+ "+m.Author.Username+", has terminado.\nTu resultado es: "+wpm_str+" WPM```")
	} else {
		s.ChannelMessageSend(m.ChannelID, "```diff\n+ ¡"+m.Author.Username+", has superado tu anterior marca de "+WPM_temp+" WPM de la fecha "+Date_temp+"!\nTu resultado es: "+wpm_str+" WPM```")
		Delete_last_score_because_improved = false
	}
}

func Show_result_not_improved(s *discordgo.Session, m *discordgo.MessageCreate, wpm float64) {
	var WPM_rounded = (math.Round(wpm*10) / 10)
	var WPM_str = fmt.Sprint(WPM_rounded)
	s.ChannelMessageSend(m.ChannelID, "```diff\n- No has superado tu anterior marca de "+WPM_temp+" WPM de la fecha "+Date_temp+", "+m.Author.Username+".\nTu resultado es: "+WPM_str+" WPM```")
}

func Show_result_with_errors(s *discordgo.Session, m *discordgo.MessageCreate, wpm float64, errors int, error_list string) {
	var WPM_rounded = (math.Round(wpm*10) / 10)
	var WPM_str = fmt.Sprint(WPM_rounded)

	if errors < 1 {
		WPM_str = fmt.Sprint(WPM_rounded)
		s.ChannelMessageSend(m.ChannelID, "```diff\n- "+m.Author.Username+", no has terminado correctamente.\nPusiste una palabra de más o un doble espacio, así que no se pudieron calcular tus errores.\nTu resultado hubiera sido: "+WPM_str+" WPM```")

	} else {
		WPM_str = fmt.Sprint(WPM_rounded)
		s.ChannelMessageSend(m.ChannelID, "```diff\n- "+m.Author.Username+", no has terminado correctamente.\nHas cometido "+fmt.Sprint(errors)+" errores: "+error_list+"\nTu resultado hubiera sido: "+WPM_str+" WPM```")
	}
}

var check bool = false

func Contest(s *discordgo.Session, m *discordgo.MessageCreate) {
	if true {
		if strings.ToLower(m.Content) == ".t" && !strings.HasPrefix(strings.ToLower(m.Content), ".tops") && !strings.HasPrefix(strings.ToLower(m.Content), ".textstats") {
			Typing_test(s, m, "Normal")
		}

		if strings.ToLower(m.Content) == ".t short" && !strings.HasPrefix(strings.ToLower(m.Content), ".tops") && !strings.HasPrefix(strings.ToLower(m.Content), ".textstats") {
			Typing_test(s, m, "Short")
		}

		if strings.ToLower(m.Content) == ".t dev" && !strings.HasPrefix(strings.ToLower(m.Content), ".tops") && !strings.HasPrefix(strings.ToLower(m.Content), ".textstats") {
			Typing_test(s, m, "Dev")
		}

		for true {
			if !check {
				check = true
				for i, v := range Currents {
					if m.ChannelID == v.Channel {
						if m.Content == Texts[v.Text_ID] {
							Currents[i].WPM = Calculate(m, Currents[i].Started, Texts[v.Text_ID])
							/*var date_temp string
							var wpm_temp string*/
							if !Is_already_in_top(m, v.Text_ID, Currents[i].WPM) {
								Is_already_in_top_LOWER(s, m, v.Text_ID, Currents[i].WPM)
								Save_result(m, v.Text_ID, Currents[i].WPM)
								Add_exp(s, m, 10001)
								Show_result(s, m, Currents[i].WPM)
							} else {
								Show_result_not_improved(s, m, Currents[i].WPM)
							}

							time.Sleep(100 * time.Millisecond)
							Top(s, m, v.Text_ID)

						} else if Judge(m, m.Content, v.Text_ID) == 2 {
							if !Is_illegal(m.Content) {
								Currents[i].WPM = Calculate(m, v.Started, Texts[v.Text_ID])
								Currents[i].Errors, Currents[i].Error_list = Errors_calculate(m.Content, Texts[v.Text_ID])
								Show_result_with_errors(s, m, Currents[i].WPM, Currents[i].Errors, Currents[i].Error_list)

								time.Sleep(100 * time.Millisecond)
								Top(s, m, v.Text_ID)
							} else {
								Currents[i].WPM = Calculate(m, v.Started, Texts[v.Text_ID])
								Currents[i].Errors, Currents[i].Error_list = Errors_calculate(m.Content, Texts[v.Text_ID])
								//s.ChannelMessageSend(m.ChannelID, "```🤔```")
								s.ChannelMessageSend("1034791654202294342", "["+time.Now().Format("02/01/2006 15:04:05")+"] <#"+m.ChannelID+"> "+m.Author.Username+" ha hecho trampas en el texto: \""+First_n(Texts[v.Text_ID], 60)+"[…]\"\n"+Currents[i].Error_list+"`` <@910067180706627594>")
							}
						} else if Judge(m, m.Content, v.Text_ID) == 4 {
							if !Is_illegal(m.Content) {
								/**** s.ChannelMessageSend(m.ChannelID, "```diff\n- Te dejaste muchísimas palabras... 😬```") */
								Currents[i].WPM = Calculate(m, v.Started, Texts[v.Text_ID])
								Currents[i].Errors, Currents[i].Error_list = Errors_calculate(m.Content, Texts[v.Text_ID])
								//Show_result_with_errors(s, m)
								time.Sleep(500 * time.Millisecond)
								/**** Top(s, m) */
							} else {
								Calculate(m, v.Started, Texts[v.Text_ID])
								Currents[i].Errors, Currents[i].Error_list = Errors_calculate(m.Content, Texts[v.Text_ID])
								//s.ChannelMessageSend(m.ChannelID, "```🤔```")
								s.ChannelMessageSend("1034791654202294342", "["+time.Now().Format("02/01/2006 15:04:05")+"] <#"+m.ChannelID+"> "+m.Author.Username+" ha hecho trampas en el texto: \""+First_n(Texts[v.Text_ID], 60)+"[…]\"\n"+Currents[i].Error_list+"`` <@910067180706627594>")
							}
						}
					}
				}
				check = false
				break
			}
		}
	}

}

var Users_typing []string
var How_many_times_added []int
var Timestamp_temp_typing int

func Reset_typing_users() {
	Users_typing = nil
	How_many_times_added = nil
}

func Typing_start_handler(s *discordgo.Session, t *discordgo.TypingStart) {
	Timestamp_temp_typing = (t.Timestamp)
	if !Slice_contains(Users_typing, t.UserID) {
		Users_typing = append(Users_typing, t.UserID)
	} else {
		// add +1 to How_many_times_added in position of Users_typing
	}
}
