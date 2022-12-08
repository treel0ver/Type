package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	ADMIN_ID = []string{"910067180706627594"}
)

func Admin(s *discordgo.Session, m *discordgo.MessageCreate) {
	if Slice_contains(ADMIN_ID, m.Author.ID) {
		var args = strings.Split(m.Content, " ")
		if strings.HasPrefix(m.Content, "..say ") {
			var args_str = fmt.Sprint(args[1:])
			args_str = args_str[1:]
			args_str = args_str[:len(args_str)-1]

			s.ChannelMessageSend(m.ChannelID, args_str)
		}
	}
}
