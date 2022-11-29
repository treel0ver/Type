package main

import (
	"strconv"
	"strings"
	"github.com/bwmarrin/discordgo"
)

var Last_message string

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	Add_exp(s, m, 500)

	if m.ChannelID == "1015972766882738216" || m.ChannelID == "1031313220230709278" {
		//Log(m)
	}

	Contest(s, m)

	/* ignore messages from bot */
	if m.Author.ID == s.State.User.ID {
		Last_message = m.Message.ID
		return
	}

	/* ignore messages that don't contain prefix "." */
	if !(strings.HasPrefix(m.Content, ".")) {
		return
	}
	
	if m.Content == "jajajear" {
		s.ChannelMessageSend(m.ChannelID, "noooriii")
	}
	if strings.ToLower(m.Content) == ".help" {
		s.ChannelMessageSend(m.ChannelID, "```css\n.t          empieza un test de velocidad\n.tops       muestra el leaderboard de un texto\n.stats      muestra tu número de tops\n.textstats  muestra información de los textos\n.info       infomación para desarrolladores```")
	}

	if strings.ToLower(m.Content) == ".help fun" {
		s.ChannelMessageSend(m.ChannelID, "```css\n.ch         pone un gif de Chaeyoung\n.mapache    pone gifs de mapaches\n.go         pone imágenes de Gopher\n.sha256     hashea una cadena de valores en SHA256\n.stb        pasa una cadena de valores al código binario\n.ntb        pasa un número al código binario```")
	}

	if strings.HasPrefix(m.Content, ".tops") && !(strings.HasPrefix(m.Content, ".topsID")) {
		Tops(s, m)
	}

	if strings.ToLower(m.Content) == ".lb" || strings.ToLower(m.Content) == ".leaderboard" || strings.ToLower(m.Content) == ".leaderboards" {
		Leaderboards(s, m)
	}

	if strings.HasPrefix(m.Content, ".topsID") {
		TopsID(s, m)
	}

	if strings.HasPrefix(m.Content, ".stats") {
		Stats(s, m)
	}

	if strings.ToLower(m.Content) == ".textstats" {
		Text_stats(s, m)
	}

	if m.Content == ".test" {
		var Started_when_stringed = strconv.FormatInt(Started_when, 10)
		s.ChannelMessageSend(m.ChannelID, "```cs\nStarted_when: "+Started_when_stringed+"```"+Date)

	}
	/*
		if m.Content == ".load" {
			Load_texts()
			var len_str = strconv.Itoa(len(Texts))
			s.ChannelMessageSend(m.ChannelID, "```css\nCargados ["+len_str+"] elementos```")
		}

		if m.Content == ".free" {
			var len_str = strconv.Itoa(len(Texts))
			s.ChannelMessageSend(m.ChannelID, "```css\nLiberados ["+len_str+"] elementos```")
			Free_texts()
		}
	*/
	Fun_commands(s, m)
}
