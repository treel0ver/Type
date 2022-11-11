package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type Two_Slices struct {
	WPM      float64
	Username string
	Date     string
}

func Tops(s *discordgo.Session, m *discordgo.MessageCreate) {
	var What = strings.Replace(m.Content, ".tops ", "", -1)

	var WHERE int
	for i := 0; i < len(Texts); i++ {
		if strings.HasPrefix(Texts[i], What) {
			WHERE = i
		}
	}

	var WHERE_str = strconv.Itoa(WHERE)
	var FOUND bool = false
	var FOUND_how_many_times int = 0

	var ts []Two_Slices

	var slice Two_Slices

	Load()
	for i := 0; i < len(DB); i++ {
		if strings.HasPrefix(DB[i], WHERE_str+" #") {
			if i == 0 {

			} else {
				var CL = strings.Split(DB[i], " # ")
				WPM_f64, _ := strconv.ParseFloat(CL[3], 8)

				slice = Two_Slices{
					WPM_f64, CL[2], CL[4],
				}
				ts = append(ts, slice)
			}
			FOUND = true
			FOUND_how_many_times++
		}
	}

	sort.Slice(ts, func(i, j int) bool {
		return ts[i].WPM > ts[j].WPM
	})

	var DISPLAY string
	var C int = 1
	for i := 0; i < 5; i++ {
		if i != 0 {
			if i > len(ts)-1 {
				var C_str = strconv.Itoa(C)
				DISPLAY = DISPLAY + C_str + "." + "\n"
			} else {
				var C_str = strconv.Itoa(C)
				var WPM_str = fmt.Sprintf("%.1f", ts[i].WPM)
				DISPLAY = DISPLAY + C_str + ". " + ts[i].Username + " (" + WPM_str + ") " + ts[i].Date + "\n"
			}
		} else {
			if len(ts) > 0 {
				var C_str = strconv.Itoa(C)
				var WPM_str = fmt.Sprintf("%.1f", ts[i].WPM)
				DISPLAY = DISPLAY + C_str + ". " + ts[i].Username + " (" + WPM_str + ") " + ts[i].Date + "\n"
			}
		}
		C++
	}

	var FOUND_how_many_times_str = strconv.Itoa(FOUND_how_many_times)
	if FOUND {
		var Texts_arr = strings.Split(Texts[WHERE], " ")
		var λ string
		for i := 0; i < len(Texts_arr); i++ {
			if i != len(Texts_arr)-1 {
				λ = λ + Texts_arr[i] + "\u200b "
			} else {
				λ = λ + Texts_arr[i]
			}
		}
		s.ChannelMessageSend(m.ChannelID, "```"+λ+"```")
		s.ChannelMessageSend(m.ChannelID, "```diff\n+ Se encontraron "+FOUND_how_many_times_str+" marcas del texto ID:"+WHERE_str+"\n"+DISPLAY+"```")
	} else {
		s.ChannelMessageSend(m.ChannelID, "```diff\n- No se encontró```")
	}

}

func Top(s *discordgo.Session, m *discordgo.MessageCreate) {
	var What = Current_text

	var WHERE int
	for i := 0; i < len(Texts); i++ {
		if strings.HasPrefix(Texts[i], What) {
			WHERE = i
		}
	}

	var WHERE_str = strconv.Itoa(WHERE)

	var ts []Two_Slices

	var slice Two_Slices

	Load()
	for i := 0; i < len(DB); i++ {
		if strings.HasPrefix(DB[i], WHERE_str+" #") {
			if i == 0 {

			} else {
				var CL = strings.Split(DB[i], " # ")
				WPM_f64, _ := strconv.ParseFloat(CL[3], 8)

				slice = Two_Slices{
					WPM_f64, CL[2], CL[4],
				}
				ts = append(ts, slice)
			}
		}
	}

	sort.Slice(ts, func(i, j int) bool {
		return ts[i].WPM > ts[j].WPM
	})

	var DISPLAY string
	var C int = 1
	for i := 0; i < 5; i++ {
		if i != 0 {
			if i > len(ts)-1 {
				var C_str = strconv.Itoa(C)
				DISPLAY = DISPLAY + C_str + "." + "\n"
			} else {
				var C_str = strconv.Itoa(C)
				var WPM_str = fmt.Sprintf("%.1f", ts[i].WPM)
				DISPLAY = DISPLAY + C_str + ". " + ts[i].Username + " (" + WPM_str + ") " + ts[i].Date + "\n"
			}
		} else {
			if len(ts) > 0 {
				var C_str = strconv.Itoa(C)
				var WPM_str = fmt.Sprintf("%.1f", ts[i].WPM)
				DISPLAY = DISPLAY + C_str + ". " + ts[i].Username + " (" + WPM_str + ") " + ts[i].Date + "\n"
			}
		}
		C++
	}

	s.ChannelMessageSend(m.ChannelID, "```🏆 LEADERBOARD 🏆\n"+DISPLAY+"```")
}

var Date_temp string
var WPM_temp string

func Is_already_in_top(m *discordgo.MessageCreate) bool {
	Load()

	var Random_str = strconv.Itoa(Random)
	for i := 0; i < len(DB); i++ {
		if strings.HasPrefix(DB[i], Random_str) {
			var CL = strings.Split(DB[i], " # ")
			if CL[1] == m.Author.ID {
				CL_f64, _ := strconv.ParseFloat(CL[3], 8)
				if CL_f64 > WPM {
					Date_temp = CL[4]
					WPM_temp = CL[3]
					return true
				}
			}
		}
	}
	return false
}

func Is_already_in_top_LOWER(m *discordgo.MessageCreate) {
	Load()

	var Random_str = strconv.Itoa(Random)
	for i := 0; i < len(DB); i++ {
		if strings.HasPrefix(DB[i], Random_str) {
			var CL = strings.Split(DB[i], " # ")
			if CL[1] == m.Author.ID {
				CL_f64, _ := strconv.ParseFloat(CL[3], 8)
				if CL_f64 < WPM {
					DB[i] = ""
					Update()
				}
			}
		}
	}
}

func Stat_list(s *discordgo.Session, m *discordgo.MessageCreate) string {
	//var Tops_1 = 0

	var stats string
	var C int

	Load()
	for i := 0; i < len(DB); i++ {
		var CL = strings.Split(DB[i], " # ")
		/* Check if is not empty line */
		if len(CL) > 1 {
			if CL[1] == m.Author.ID {
				C++
				stats = stats + DB[i] + "\n"
			}
		}
	}

	C_str := strconv.Itoa(C)
	stats = C_str

	return stats
}

func Stats(s *discordgo.Session, m *discordgo.MessageCreate) {
	// 🚧👷
	var Texts_where_user_is_in []string

	var How_many_texts_in int
	var Not_tops_1 int

	Load()
	for i := 0; i < len(DB); i++ {
		var CL = strings.Split(DB[i], " # ")
		/* Check if is not empty line */
		if len(CL) > 1 {
			if CL[1] == m.Author.ID {
				if !(Slice_contains(Texts_where_user_is_in, CL[0])) {
					Texts_where_user_is_in = append(Texts_where_user_is_in, CL[0])
					How_many_texts_in++
					var what_text = CL[0]
					var what_WPM, _ = strconv.ParseFloat(CL[3], 64)
					for k := 0; k < len(DB); k++ {
						var CK = strings.Split(DB[k], " # ")
						if CK[0] == what_text && CK[1] != m.Author.ID {
							CK_float, _ := strconv.ParseFloat(CK[3], 64)
							if CK_float > what_WPM {
								Not_tops_1++
							}
						}
					}
				}
			}
		}
	}
	var TOPS = How_many_texts_in - Not_tops_1
	var TOPS_str = strconv.Itoa(TOPS)
	s.ChannelMessageSend(m.ChannelID, "```css\nHas participado en "+Stat_list(s, m)+" textos\n"+m.Author.Username+", tienes "+TOPS_str+" tops 1.```")
}
