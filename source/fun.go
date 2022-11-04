package main

import (
	"fmt"
	"math/rand"
	"time"
	"strings"
	"strconv"
    "crypto/sha256"
	"github.com/bwmarrin/discordgo"
)
func Fun_commands(s *discordgo.Session, m *discordgo.MessageCreate) {
		if strings.ToLower(m.Content) == ".mapache" {
			rand.Seed(time.Now().UnixNano()); var Fun_random = rand.Intn(3)
			switch Fun_random {
			case 0: s.ChannelMessageSend(m.ChannelID, "https://static.boredpanda.com/blog/wp-content/uploads/2017/06/adorable-cute-raccoons-52-59561ef26a4a2__700.gif")
			case 1: s.ChannelMessageSend(m.ChannelID, "https://tenor.com/view/froze-stop-moving-not-moving-still-standing-gif-16669739")
	 		case 2: s.ChannelMessageSend(m.ChannelID, "https://www.eulixe.com/media/eulixe/images/2019/02/08/adorable3.gif")
			case 3: s.ChannelMessageSend(m.ChannelID, "https://64.media.tumblr.com/83c825536d1bcf79969eafa9d029c6f3/tumblr_oey49g9AcZ1rmin5no1_500.gifv")
			}
		}

		if strings.ToLower(m.Content) == ".go" {
			rand.Seed(time.Now().UnixNano()); var Fun_random = rand.Intn(1)
			switch Fun_random {
			case 0: s.ChannelMessageSend(m.ChannelID, "https://cdn.dribbble.com/userupload/2624050/file/original-59266f4dea1c2aa43f2064cc0f3b165a.png?resize=400x0")
			case 1: s.ChannelMessageSend(m.ChannelID, "https://camo.githubusercontent.com/833cfd306ac2bef74ddf0560ee3b4112321c5b6939e52a1629f0aed8aec46922/687474703a2f2f692e696d6775722e636f6d2f485379686177742e6a7067")
			}
		}
	
		if strings.ToLower(m.Content) == ".ch" {
			s.ChannelMessageSend(m.ChannelID, "https://thumbs.gfycat.com/AdmirableLikableFlea-mobile.mp4")
		}

		if strings.ToLower(m.Content) == ".chaeyoung" {
			rand.Seed(time.Now().UnixNano()); var Fun_random = rand.Intn(1)
			switch Fun_random {
			case 0: s.ChannelMessageSend(m.ChannelID, "https://cdn.dribbble.com/userupload/2624050/file/original-59266f4dea1c2aa43f2064cc0f3b165a.png?resize=400x0")
			case 1: s.ChannelMessageSend(m.ChannelID, "https://camo.githubusercontent.com/833cfd306ac2bef74ddf0560ee3b4112321c5b6939e52a1629f0aed8aec46922/687474703a2f2f692e696d6775722e636f6d2f485379686177742e6a7067")
			}
		}

		if strings.HasPrefix(strings.ToLower(m.Content), ".sha256") {
			var content = strings.Replace(m.Content, ".sha256 ", "", -1)

			h := sha256.New()

			h.Write([]byte(content))

			bs := h.Sum(nil)

			var result_str = fmt.Sprintf("%x", bs)

			s.ChannelMessageSend(m.ChannelID, "```" + result_str + "```")
		}

		if strings.HasPrefix(strings.ToLower(m.Content), ".stb") {
			var content = strings.Replace(m.Content, ".stb ", "", -1)

			s.ChannelMessageSend(m.ChannelID, "```" + String_to_binary(content) + "```")
		}

		if strings.HasPrefix(strings.ToLower(m.Content), ".ntb") {
			var content = strings.Replace(m.Content, ".ntb ", "", -1)

			content2, _ := strconv.Atoi(content)
			var content3 = int64(content2)

			s.ChannelMessageSend(m.ChannelID, "```" + strconv.FormatInt(content3, 2) + "```")
		}

		if strings.HasPrefix(strings.ToLower(m.Content), ".len") {
			var content = strings.Replace(m.Content, ".len ", "", -1)
			var length = len(content)
			var length_str = strconv.Itoa(length)

			s.ChannelMessageSend(m.ChannelID, "```css\n[" + length_str + "]```")
		}

}