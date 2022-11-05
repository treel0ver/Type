package main

import (
	"sort"
	"strings"
	"strconv"
	"github.com/bwmarrin/discordgo"
)

type TwoSlices struct {
    main_slice  	 []float64
    other_slice  	 []string
}

type SortByOther TwoSlices

func (sbo SortByOther) Len() int {
    return len(sbo.main_slice)
}

func (sbo SortByOther) Swap(i, j int) {
    sbo.main_slice[i], sbo.main_slice[j] = sbo.main_slice[j], sbo.main_slice[i]
    sbo.other_slice[i], sbo.other_slice[j] = sbo.other_slice[j], sbo.other_slice[i]
}

func (sbo SortByOther) Less(i, j int) bool {
    return sbo.other_slice[i] < sbo.other_slice[j] 
}

func Tops(s *discordgo.Session, m *discordgo.MessageCreate) {

	var Leaderboard = make([]string, 1)
	var Leaderboard_WPM = make([]float64, 1)

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

	Load()
	for i := 0; i < len(DB); i++ {
		if strings.HasPrefix(DB[i], WHERE_str + " #") {
			if i == 0 {
				var CL = strings.Split(DB[i], " # ")
				WPM_f64, _ := strconv.ParseFloat(CL[3], 8)
				Leaderboard_WPM[0] = WPM_f64
				Leaderboard[0] = (CL[2] + " (" + CL[3] + " WPM) " + CL[4])
			} else {
				var CL = strings.Split(DB[i], " # ")
				WPM_f64, _ := strconv.ParseFloat(CL[3], 8)
				Leaderboard_WPM = append(Leaderboard_WPM, WPM_f64)
				Leaderboard = append(Leaderboard, (CL[2] + " (" + CL[3] + " WPM) " + CL[4]))
			}

			FOUND = true
			FOUND_how_many_times++
		}
	}

    var my_two_slices = TwoSlices{main_slice: Leaderboard_WPM, other_slice: Leaderboard}

    //fmt.Println("Not sorted : ", my_two_slices.main_slice)
    //fmt.Println("Not soerted : ", my_two_slices.other_slice)

    sort.Sort(sort.Reverse(sort.StringSlice(my_two_slices.other_slice)))
    //sort.Sort(sort.Reverse(sort.Float64Slice(my_two_slices.main_slice)))


    //fmt.Println("Sorted : ", my_two_slices.main_slice)
    //fmt.Println("Sorted : ", my_two_slices.other_slice)

	var DISPLAY string
	var C int = 1
	for i := 0; i < 5; i++ {
		if i != 0{
			if i>len(my_two_slices.other_slice)-1 {
				var C_str = strconv.Itoa(C)
				DISPLAY = DISPLAY + C_str + "." + "\n"
			} else {
				var C_str = strconv.Itoa(C)
				DISPLAY = DISPLAY + C_str + ". " + my_two_slices.other_slice[i] + "\n"
			}
		} else {
				var C_str = strconv.Itoa(C)
				DISPLAY = DISPLAY + C_str + ". " + my_two_slices.other_slice[0] + "\n"
		}
		C++
	}

	var FOUND_how_many_times_str = strconv.Itoa(FOUND_how_many_times)
	if FOUND {
		var Texts_arr = strings.Split(Texts[WHERE], " ")
		var λ string
		for i := 0; i < len(Texts_arr); i++ {
			if i != len(Texts_arr)-1 {
				λ = λ + Texts_arr[i] + "​ "
			} else {
				λ = λ + Texts_arr[i]
			}
		}
		s.ChannelMessageSend(m.ChannelID, "```" + λ + "```")
		s.ChannelMessageSend(m.ChannelID, "```diff\n+ Se encontraron " + FOUND_how_many_times_str + " marcas del texto ID:" + WHERE_str + "\n" + DISPLAY + "```")
	} else {
		s.ChannelMessageSend(m.ChannelID, "```diff\n- No se encontró```")
	}
	
}

func Top(s *discordgo.Session, m *discordgo.MessageCreate) {
	var Leaderboard = make([]string, 1)
	var Leaderboard_WPM = make([]float64, 1)

	var What = Current_text

	var WHERE int
	for i := 0; i < len(Texts); i++ {
		if strings.HasPrefix(Texts[i], What) {
			WHERE = i
		}
	}

	var WHERE_str = strconv.Itoa(WHERE)

	Load()
	for i := 0; i < len(DB); i++ {
		if strings.HasPrefix(DB[i], WHERE_str + " #") {
			if i == 0 {
				var CL = strings.Split(DB[i], " # ")
				WPM_f64, _ := strconv.ParseFloat(CL[3], 8)
				Leaderboard_WPM[0] = WPM_f64
				Leaderboard[0] = (CL[2] + " (" + CL[3] + " WPM) " + CL[4])
			} else {
				var CL = strings.Split(DB[i], " # ")
				WPM_f64, _ := strconv.ParseFloat(CL[3], 8)
				Leaderboard_WPM = append(Leaderboard_WPM, WPM_f64)
				Leaderboard = append(Leaderboard, (CL[2] + " (" + CL[3] + " WPM) " + CL[4]))
			}
		}
	}

    var my_two_slices = TwoSlices{main_slice: Leaderboard_WPM, other_slice: Leaderboard}

    //fmt.Println("Not sorted : ", my_two_slices.main_slice)
    //fmt.Println("Not soerted : ", my_two_slices.other_slice)

    sort.Sort(sort.Reverse(sort.StringSlice(my_two_slices.other_slice)))
    //sort.Sort(sort.Reverse(sort.Float64Slice(my_two_slices.main_slice)))


    //fmt.Println("Sorted : ", my_two_slices.main_slice)
    //fmt.Println("Sorted : ", my_two_slices.other_slice)

	var DISPLAY string
	var C int = 1
	for i := 0; i < 5; i++ {
		if i != 0{
			if i>len(my_two_slices.other_slice)-1 {
				var C_str = strconv.Itoa(C)
				DISPLAY = DISPLAY + C_str + "." + "\n"
			} else {
				var C_str = strconv.Itoa(C)
				DISPLAY = DISPLAY + C_str + ". " + my_two_slices.other_slice[i] + "\n"
			}
		} else {
				var C_str = strconv.Itoa(C)
				DISPLAY = DISPLAY + C_str + ". " + my_two_slices.other_slice[0] + "\n"
		}
		C++
	}

	s.ChannelMessageSend(m.ChannelID, "```🏆 LEADERBOARD 🏆\n" + DISPLAY + "```")
	
}

func Is_already_in_top(m *discordgo.MessageCreate) bool {
	Load()

	var Random_str = strconv.Itoa(Random)
	for i := 0; i < len(DB); i++ {
		if strings.HasPrefix(DB[i], Random_str) {
			var CL = strings.Split(DB[i], " # ")
			if CL[1] == m.Author.ID {
				CL_f64, _ := strconv.ParseFloat(CL[3], 8)
				if CL_f64 > WPM {
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
}