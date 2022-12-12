package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var Last_message string

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	Admin(s, m)

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

	if strings.ToLower(m.Content) == ".help" {
		s.ChannelMessageSend(m.ChannelID, "```css\n.t          empieza un test de velocidad\n.tops       muestra el leaderboard de un texto\n.stats      muestra tu número de tops\n.textstats  muestra información de los textos\n.info       infomación para desarrolladores```")
	}

	if strings.ToLower(m.Content) == ".help2" {
		s.ChannelMessageSend(m.ChannelID, "```css\nPerfil\n.perfil     muestra un perfil\n.frase      para cambiar la frase de tu perfil\n.mascota    para cambiar la mascota de tu perfil\n\nHerramientas\n.len        para ver la longitud de un texto\n.img        herramienta para pasar un texto en imagen\n.sha256     hashea una cadena de valores en SHA256\n.stb        pasa una cadena de valores al código binario\n.ntb        pasa un número al código binario\n.csv        herramienta random\n\nFun\n.level      para ver tu nivel\n.ch         pone un gif de chae-young\n.mapache    pone gifs de mapaches\n.go         pone imágenes de Gopher\n```")
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

	Fun_commands(s, m)

	//

	if strings.HasPrefix(m.Content, ".calc") {
		if len(m.Attachments) > 0 {
			var curl_output, _ = exec.Command("/usr/bin/curl", m.Attachments[0].URL).Output()
			var content = string(curl_output[:])

			var content_s = strings.Split(content, "\n")

			var average_wpm float64
			var average_puls float64

			for i := range content_s {
				var split = strings.Split(content_s[i], "\t")
				var wpm = split[1]
				var puls = split[3]
				var wpm_f, _ = strconv.ParseFloat(wpm, 64)
				var puls_f, _ = strconv.ParseFloat(puls, 64)
				average_wpm += wpm_f
				average_puls += puls_f
			}

			var result_wpm = average_wpm / float64(len(content_s))
			var result_puls = average_puls / float64(len(content_s))

			s.ChannelMessageSend(m.ChannelID, "```css\navg wpm: "+fmt.Sprint(result_wpm)+"\navg puls: "+fmt.Sprint(result_puls)+"```")

		}
	}
}
