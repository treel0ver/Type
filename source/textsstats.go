package main

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func Text_stats(s *discordgo.Session, m *discordgo.MessageCreate) {
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
	temp2 = strconv.FormatInt(int64((temp1 / How_many_texts())), 10)

	var _0s = strconv.FormatInt(int64(_0), 10)
	var _100s = strconv.FormatInt(int64(_100), 10)
	var _200s = strconv.FormatInt(int64(_200), 10)
	var _300s = strconv.FormatInt(int64(_300), 10)
	var _400s = strconv.FormatInt(int64(_400), 10)
	var _500s = strconv.FormatInt(int64(_500), 10)
	var _600s = strconv.FormatInt(int64(_600), 10)

	s.ChannelMessageSend(m.ChannelID, "```diff\nLongitud promedio de textos: "+temp2+"\n\n"+"000-100: "+_0s+"\n100-200: "+_100s+"\n200-300: "+_200s+"\n300-400: "+_300s+"\n400-500: "+_400s+"\n500-600: "+_500s+"\n600-700: "+_600s+"```")
}
