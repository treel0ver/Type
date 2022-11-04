package main

import (
	"time"
	"strings"
	"strconv"
	"github.com/bwmarrin/discordgo"
) 

var Last_message string

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	Add_exp(s, m, 500)

	if m.ChannelID == "1015972766882738216" || m.ChannelID == "1031313220230709278"{
		Log(m)
	}

	if m.Author.ID == s.State.User.ID {
		Last_message = m.Message.ID
		return
	}

	if strings.ToLower(m.Content) == ".help" {
		s.ChannelMessageSend(m.ChannelID, "```css\n.t          empieza un test de velocidad\n.tops       muestra el leaderboard de un texto\n.stats      muestra tu número de tops\n.textstats  muestra información de los textos\n.info       infomación para desarrolladores```")
	}

	if strings.ToLower(m.Content) == ".help fun" {
		s.ChannelMessageSend(m.ChannelID, "```css\n.ch         pone un gif de Chaeyoung\n.mapache    pone gifs de mapaches\n.go         pone imágenes de Gopher\n.sha256     hashea una cadena de valores en SHA256\n.stb        pasa una cadena de valores al código binario\n.ntb        pasa un número al código binario```")
	}

	if m.ChannelID == "1015972766882738216" || m.ChannelID == "1031313220230709278" || m.ChannelID == "1035687048000053288" {
		if strings.ToLower(m.Content) == ".t" && !strings.HasPrefix(strings.ToLower(m.Content), ".tops") && !strings.HasPrefix(strings.ToLower(m.Content), ".textstats"){
			Typing_test(s, m)
		}

		if strings.ToLower(m.Content) == ".t short" && !strings.HasPrefix(strings.ToLower(m.Content), ".tops") && !strings.HasPrefix(strings.ToLower(m.Content), ".textstats"){
			TT_short(s, m)
		}
		
		if Is_started && m.Content == Current_text {
			Calculate(m)
			if !Is_already_in_top(m) {
				Show_result(s, m)
				Is_already_in_top_LOWER(m)
				Save_result(m)
				Add_exp(s, m, 10001)
			} else {
				Show_result_not_improved(s, m)
			}

			time.Sleep(500 * time.Millisecond); Top(s, m)

		} else if (Judge(m, m.Content) == 2) {
			if !Is_illegal(m.Content) {
				Calculate(m)
				Errors_calculate(m.Content, Current_text)
				Show_result_with_errors(s, m)

				time.Sleep(300 * time.Millisecond); Top(s, m)
			} else {
				Calculate(m)
				Errors_calculate(m.Content, Current_text)
				//s.ChannelMessageSend(m.ChannelID, "```🤔```")
				s.ChannelMessageSend("1034791654202294342", "[" + time.Now().Format("02/01/2006 15:04:05") + "] <#" + m.ChannelID + "> " + m.Author.Username + " ha hecho trampas en el texto: \"" + First_n(Current_text, 60) + "[…]\"\n" + Error_list + "`` <@910067180706627594>")
			}
		} else if (Judge(m, m.Content) == 4) {
			if !Is_illegal(m.Content) {
				s.ChannelMessageSend(m.ChannelID, "```diff\n- Te dejaste muchísimas palabras... Tu resultado (WPM) es una vaga representación de la realidad. 😬```")
				Calculate(m)
				Errors_calculate(m.Content, Current_text)
				Show_result_with_errors(s, m)
				time.Sleep(300 * time.Millisecond); Top(s, m)
			} else {
				Calculate(m)
				Errors_calculate(m.Content, Current_text)
				//s.ChannelMessageSend(m.ChannelID, "```🤔```")
				s.ChannelMessageSend("1034791654202294342", "[" + time.Now().Format("02/01/2006 15:04:05") + "] <#" + m.ChannelID + "> " + m.Author.Username + " ha hecho trampas en el texto: \"" + First_n(Current_text, 60) + "[…]\"\n" + Error_list + "`` <@910067180706627594>")
			}
		}
	}



	if strings.HasPrefix(m.Content, ".tops") {
		Tops(s, m)
	}

	if strings.HasPrefix(m.Content, ".stats") {
		s.ChannelMessageSend(m.ChannelID, "```css\nHas participado en " + Stat_list(s, m) + " textos```")
	}

	if strings.ToLower(m.Content) == ".textstats" {
		var temp1 int
		var temp2 string
		var _0 int
		var _100 int
		var _200 int
		var _300 int
		var _400 int 
		var _500 int
		var _600 int

		for i := 0; i < How_many_texts(); i++ {
			temp1 = temp1 + len(Texts[i])
			if len(Texts[i]) < 100 {
				_0++
			} 
			if len(Texts[i]) < 200 && len(Texts[i]) >= 100 {
				_100++
			}
			if len(Texts[i]) < 300 && len(Texts[i]) >= 200 {
				_200++
			}
			if len(Texts[i]) < 400 && len(Texts[i]) >= 300 {
				_300++
			}
			if len(Texts[i]) < 500 && len(Texts[i]) >= 400 {
				_400++
			}
			if len(Texts[i]) < 600 && len(Texts[i]) >= 500 {
				_500++
			}
			if len(Texts[i]) < 700 && len(Texts[i]) >= 600 {
				_600++
			}

		}
		temp2 = strconv.FormatInt(int64((temp1/How_many_texts())), 10) 

		var _0s = strconv.FormatInt(int64(_0), 10) 
		var _100s = strconv.FormatInt(int64(_100), 10) 
		var _200s = strconv.FormatInt(int64(_200), 10) 
		var _300s = strconv.FormatInt(int64(_300), 10) 
		var _400s = strconv.FormatInt(int64(_400), 10) 
		var _500s = strconv.FormatInt(int64(_500), 10) 
		var _600s = strconv.FormatInt(int64(_600), 10) 

		s.ChannelMessageSend(m.ChannelID, "```diff\nLongitud promedio de textos: " + temp2 + "\n\n" + "000-100: " + _0s + "\n100-200: " + _100s	+ "\n200-300: " + _200s + "\n300-400: " + _300s	+ "\n400-500: " + _400s	+ "\n500-600: " + _500s + "\n600-700: " + _600s + "```")	
	}

	if strings.HasPrefix(m.Content, ".level") {
		var args = strings.Split(m.Content, " ")
		Show_level(s, m, args)
	}

	if m.Content == ".info" {
		var Started_when_stringed = strconv.FormatInt(Started_when, 10)
		s.ChannelMessageSend(m.ChannelID, "Started_when: " +  Started_when_stringed)
		s.ChannelMessageSend(m.ChannelID, Date)
	}

	if m.Content == ".abc" {
	}

	Fun_commands(s, m)

}