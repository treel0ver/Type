package main

import (
	"time"
	"strings"
	"strconv"
	"github.com/bwmarrin/discordgo"
) 

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	if !strings.HasPrefix(m.Content, ".") && !(Judge(m, m.Content) == 1){
		if !(len(m.Content) > len(Current_text) - 3) {
			return
		}
	}

	if m.ChannelID == "1015972766882738216" || m.ChannelID == "1031313220230709278" {
		if strings.HasPrefix(strings.ToLower(m.Content), ".t") && !strings.HasPrefix(strings.ToLower(m.Content), ".tops"){
			Typing_test(s, m)
		}

		if Is_started && m.Content == Current_text {
			Calculate(m)
			if !Is_already_in_top(m) {
				Show_result(s, m)
				Is_already_in_top_LOWER(m)
				Save_result(m)
			} else {
				Show_result_not_improved(s, m)
			}

			time.Sleep(500 * time.Millisecond); Top(s, m)

		} else if (Judge(m, m.Content) == 2) {
			Calculate(m)
			Errors_calculate(m.Content, Current_text)
			Show_result_with_errors(s, m)

			time.Sleep(500 * time.Millisecond); Top(s, m)
		}
	}

	if strings.HasPrefix(m.Content, ".tops") {
		Tops(s, m)
	}

	if strings.HasPrefix(m.Content, ".stats") {
		s.ChannelMessageSend(m.ChannelID, Stats(s, m))
	}

	if m.Content == ".info" {
		var Started_when_stringed = strconv.FormatInt(Started_when, 10)
		s.ChannelMessageSend(m.ChannelID, "Started_when: " +  Started_when_stringed)
		s.ChannelMessageSend(m.ChannelID, Date)
	}

	if m.Content == ".test" {
		
	}
}