package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var DB []string

var Texts []string

var Date string

var Current_text string = "Escucha, las reglas propias... se tratan de decidir conseguir algo usando medios y maneras propias para conseguirlo. Por eso decimos que son nuestras reglas. Precisamente por eso podemos afrontar sinceramente los desafíos y darlo todo. Y si fracasamos, hay que retomar la práctica y soportar duros entrenamientos para lograrlo. Y así, dedicándote a ello, creas tus propias reglas."
var Is_started bool

var Started_when = time.Now().UnixMilli()
var Requested_when = time.Now().UnixMilli()

var Text_message_ID string

var WPM float64
var WPM_str string
var WPM_str_save string

var Random int

func Is_illegal(ζ string) bool {
	var π = strings.Contains(ζ, "\u200b")
	return π
}

func υ() string {

	var Texts_arr = strings.Split(Current_text, " ")
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

func Typing_test(s *discordgo.Session, m *discordgo.MessageCreate) {
	Reset_typing_users()
	var Now = time.Now().UnixMilli()

	/*
		println("-----")
		println("Now: ")
		println(Now)
		println("Started_when: ")
		println(Started_when)
		println("diff: ")
		println(Now - Started_when)
	*/

	if Now-Requested_when < 3000 {
		s.ChannelMessageSend(m.ChannelID, "```diff\n- Espera un poco.```")
	} else {

		rand.Seed(time.Now().UnixNano())
		Random = rand.Intn(How_many_texts())
		Random = rand.Intn(399) + 1
		//Random = 1

		Current_text = Texts[Random]

		var Requested_when_time = time.Now()
		Requested_when = Requested_when_time.UnixMilli()

		var Test_message, _ = s.ChannelMessageSend(m.ChannelID, "```🔴 Preparados...```")
		Text_message_ID = Test_message.ID
		time.Sleep(1 * time.Second)
		s.ChannelMessageEdit(m.ChannelID, Test_message.ID, "```🟡 Listos...```")
		time.Sleep(1 * time.Second)

		var Text_shown, _ = s.ChannelMessageEdit(m.ChannelID, Test_message.ID, "**"+υ()+"**")
		_ = Text_shown
		var Started_when_time = Text_shown.EditedTimestamp
		Started_when = Started_when_time.UnixMilli()

	}
}

func TT_short(s *discordgo.Session, m *discordgo.MessageCreate) {
	Reset_typing_users()
	var Now = time.Now().UnixMilli()

	if Now-Requested_when < 3000 {
		s.ChannelMessageSend(m.ChannelID, "```diff\n- Espera un poco.```")
	} else {
		rand.Seed(time.Now().UnixNano())
		Random = rand.Intn(458-401) + 401
		//Random = 76
		Current_text = Texts[Random]

		var Requested_when_time = time.Now()
		Requested_when = Requested_when_time.UnixMilli()

		var Test_message, _ = s.ChannelMessageSend(m.ChannelID, "```🔴🩳 (Textos cortos) Preparados...```")
		time.Sleep(1 * time.Second)
		s.ChannelMessageEdit(m.ChannelID, Test_message.ID, "```🟡🤏 (Textos cortos) Listos...```")
		time.Sleep(1 * time.Second)
		var Text_shown, _ = s.ChannelMessageEdit(m.ChannelID, Test_message.ID, "**"+υ()+"**")
		_ = Text_shown
		var Started_when_time = Text_shown.EditedTimestamp
		Started_when = Started_when_time.UnixMilli()
	}
}

func TT_dev(s *discordgo.Session, m *discordgo.MessageCreate) {
	Reset_typing_users()
	var Now = time.Now().UnixMilli()

	if Now-Requested_when < 3000 {
		s.ChannelMessageSend(m.ChannelID, "```diff\n- Espera un poco.```")
	} else {
		rand.Seed(time.Now().UnixNano())
		Random = rand.Intn(468-459) + 459
		//Random = 459
		Current_text = Texts[Random]

		var Requested_when_time = time.Now()
		Requested_when = Requested_when_time.UnixMilli()

		var Test_message, _ = s.ChannelMessageSend(m.ChannelID, "```🔴💻 (Programación) Preparados...```")
		time.Sleep(1 * time.Second)
		s.ChannelMessageEdit(m.ChannelID, Test_message.ID, "```🟡👩🏾‍💻 (Programación) Listos...```")
		time.Sleep(1 * time.Second)
		var Text_shown, _ = s.ChannelMessageEdit(m.ChannelID, Test_message.ID, "**"+υ()+"**")
		_ = Text_shown
		var Started_when_time = Text_shown.EditedTimestamp
		Started_when = Started_when_time.UnixMilli()
	}
}

func Judge(m *discordgo.MessageCreate, S string) int8 {
	if S == Current_text {
		return 1
	}

	var Content_arr = strings.Split(m.Content, " ")
	var Current_text_arr = strings.Split(Current_text, " ")

	if Content_arr[0] == Current_text_arr[0] {
		if len(m.Content) > len(Current_text)-10 {
			return 2
		}

		if len(m.Content) > len(Current_text)-30 {
			return 4
		}
	}
	return 3
}

func Calculate(m *discordgo.MessageCreate) {
	var sent_when, _ = SnowflakeTimestamp(m.Message.ID)
	Date = sent_when.Format("02/01/2006 15:04")
	var sent_when_unixmilli = sent_when.UnixMilli()

	var sent_when_unixmilli_float64 float64 = float64(sent_when_unixmilli)

	var length = len([]rune(Current_text))
	var length_as_a_float float64 = float64(length)

	var Started_when_float float64 = float64(Started_when)
	WPM = length_as_a_float / 5 / ((sent_when_unixmilli_float64 - Started_when_float) - 600) * 60000
}

var Error_list string
var Errors int
var Errors_str string

func Errors_calculate(sent string, current string) {

	/* reseting */
	Errors = 0
	Error_list = ""

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

	Errors_str = strconv.FormatInt(int64(Errors), 10)
}

var Delete_last_score_because_improved bool = false

func Show_result(s *discordgo.Session, m *discordgo.MessageCreate) {
	if !Delete_last_score_because_improved {
		s.ChannelMessageSend(m.ChannelID, "```diff\n+ "+m.Author.Username+", has terminado.\nTu resultado es: "+WPM_str+" WPM```")
	} else {
		s.ChannelMessageSend(m.ChannelID, "```diff\n+ ¡"+m.Author.Username+", has superado tu anterior marca de "+WPM_temp+" WPM de la fecha "+Date_temp+"!\nTu resultado es: "+WPM_str+" WPM```")
		Delete_last_score_because_improved = false
	}
}

func Calc_WPM(s *discordgo.Session, m *discordgo.MessageCreate) {
	var WPM_rounded = (math.Round(WPM*10) / 10)
	WPM_str = fmt.Sprint(WPM_rounded)

	WPM_str_save = fmt.Sprint(WPM)
}

func Show_result_not_improved(s *discordgo.Session, m *discordgo.MessageCreate) {
	var WPM_rounded = (math.Round(WPM*10) / 10)
	WPM_str = fmt.Sprint(WPM_rounded)

	s.ChannelMessageSend(m.ChannelID, "```diff\n- No has superado tu anterior marca de "+WPM_temp+" WPM de la fecha "+Date_temp+", "+m.Author.Username+".\nTu resultado es: "+WPM_str+" WPM```")
}

func Show_result_with_errors(s *discordgo.Session, m *discordgo.MessageCreate) {
	var WPM_rounded = (math.Round(WPM*10) / 10)
	if Errors < 1 {
		WPM_str = fmt.Sprint(WPM_rounded)
		s.ChannelMessageSend(m.ChannelID, "```diff\n- "+m.Author.Username+", no has terminado correctamente.\nPusiste una palabra de más o un doble espacio, así que no se pudieron calcular tus errores.\nTu resultado hubiera sido: "+WPM_str+" WPM```")

	} else {
		WPM_str = fmt.Sprint(WPM_rounded)
		s.ChannelMessageSend(m.ChannelID, "```diff\n- "+m.Author.Username+", no has terminado correctamente.\nHas cometido "+Errors_str+" errores: "+Error_list+"\nTu resultado hubiera sido: "+WPM_str+" WPM```")
	}
}

func Contest(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.ChannelID == "1015972766882738216" || m.ChannelID == "1031313220230709278" || m.ChannelID == "1035687048000053288" {
		if strings.ToLower(m.Content) == ".t" && !strings.HasPrefix(strings.ToLower(m.Content), ".tops") && !strings.HasPrefix(strings.ToLower(m.Content), ".textstats") {
			Typing_test(s, m)
		}

		if strings.ToLower(m.Content) == ".t short" && !strings.HasPrefix(strings.ToLower(m.Content), ".tops") && !strings.HasPrefix(strings.ToLower(m.Content), ".textstats") {
			TT_short(s, m)
		}

		if strings.ToLower(m.Content) == ".t dev" && !strings.HasPrefix(strings.ToLower(m.Content), ".tops") && !strings.HasPrefix(strings.ToLower(m.Content), ".textstats") {
			TT_dev(s, m)
		}

		if m.Content == Current_text {
			Calculate(m)
			if !Is_already_in_top(m) {
				Calc_WPM(s, m)
				Is_already_in_top_LOWER(s, m)
				Save_result(m)
				Add_exp(s, m, 10001)
				Show_result(s, m)
			} else {
				Show_result_not_improved(s, m)
			}

			time.Sleep(100 * time.Millisecond)
			Top(s, m)

		} else if Judge(m, m.Content) == 2 {
			if !Is_illegal(m.Content) {
				Calculate(m)
				Errors_calculate(m.Content, Current_text)
				Show_result_with_errors(s, m)

				time.Sleep(100 * time.Millisecond)
				Top(s, m)
			} else {
				Calculate(m)
				Errors_calculate(m.Content, Current_text)
				//s.ChannelMessageSend(m.ChannelID, "```🤔```")
				s.ChannelMessageSend("1034791654202294342", "["+time.Now().Format("02/01/2006 15:04:05")+"] <#"+m.ChannelID+"> "+m.Author.Username+" ha hecho trampas en el texto: \""+First_n(Current_text, 60)+"[…]\"\n"+Error_list+"`` <@910067180706627594>")
			}
		} else if Judge(m, m.Content) == 4 {
			if !Is_illegal(m.Content) {
				/**** s.ChannelMessageSend(m.ChannelID, "```diff\n- Te dejaste muchísimas palabras... 😬```") */
				Calculate(m)
				Errors_calculate(m.Content, Current_text)
				//Show_result_with_errors(s, m)
				time.Sleep(500 * time.Millisecond)
				/**** Top(s, m) */
			} else {
				Calculate(m)
				Errors_calculate(m.Content, Current_text)
				//s.ChannelMessageSend(m.ChannelID, "```🤔```")
				s.ChannelMessageSend("1034791654202294342", "["+time.Now().Format("02/01/2006 15:04:05")+"] <#"+m.ChannelID+"> "+m.Author.Username+" ha hecho trampas en el texto: \""+First_n(Current_text, 60)+"[…]\"\n"+Error_list+"`` <@910067180706627594>")
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
