package main

import (
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Text_stats(s *discordgo.Session, m *discordgo.MessageCreate) {
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
