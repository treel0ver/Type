package main

import (
	"time"
    "strings"
    "strconv"

    //"fmt"

	//"strconv"
	//"reflect"
)
/*
type TOP struct {
	Pos 		int
	User_ID 	string
	WPM 		float64
	Made_at		string
}

type TEXTO struct {
	Texto_ID 		int
	Texto_original	string
	Texto_display 	string
	Top  			[]TOP

}

type TEXTOS struct {
	Texto 		[]TEXTO
}

var new_texto = TEXTO{
	Texto_ID: 1,
	Texto_original: "tas papeao mano bueno bueno el sapo sapote no come unas cosas como esas hola buenas",
}

var Textos = TEXTOS{

}
*/

var textos [1000]string = [1000]string{
	"Escucha, las reglas propias... se tratan de decidir conseguir algo usando medios y maneras propias para conseguirlo. Por eso decimos que son nuestras reglas. Precisamente por eso podemos afrontar sinceramente los desafíos y darlo todo. Y si fracasamos, hay que retomar la práctica y soportar duros entrenamientos para lograrlo. Y así, dedicándote a ello, creas tus propias reglas.",
	"Pocos ven lo que somos, pero todos ven lo que aparentamos. Nunca intentes ganar por la fuerza lo que puede ser ganado por la mentira. Nunca se logró nada grandioso sin peligro.",
	"Tal vez solo nos espere un camino oscuro por delante. Pero aún así necesitas creer y seguir adelante. Cree en que las estrellas iluminarán tu camino, incluso un poco. Vamos, ¡emprendamos una aventura!",
	"Sólo soy verdaderamente libre cuando todos los seres humanos que me rodean, hombres y mujeres, son igualmente libres, de manera que cuanto más numerosos son los hombres libres que me rodean y más profunda y más amplia es su libertad, más extensa, más profunda y más amplia viene a ser mi libertad.",
	"Cada tipo que crece podría ser un repetidor, un pequeño maestro, el desencadenante de una reacción en cadena que en sí misma es capaz de cambiar el mundo.",
}
var display_textos [1000]string = [1000]string{
	"**Escucha,​ las reglas​ propias... se tratan de decidir conseguir algo​ usando medios y maneras propias​ para conseguirlo. Por eso decimos que son nuestras​ reglas. Precisamente por eso podemos afrontar sinceramente los desafíos y darlo todo. Y si​ fracasamos, hay que retomar la práctica​ y soportar duros entrenamientos para lograrlo. Y así, dedicándote a ello​, creas tus propias reglas.​**",
	"**Pocos ven lo que somos, pero todos ven lo que aparentamos. Nunca intentes ganar por la fuerza lo que puede ser ganado por la mentira. Nunca se logró nada grandioso sin peligro.**",
	"**Tal​ vez solo nos espere un camino oscuro por​ delante. Pero aún así necesitas​ creer y seguir adelante. Cree en que las estrellas​ iluminarán tu camino, incluso un poco. Vamos,​ ¡emprendamos una aventura!**",
	"**Sólo soy​ verdaderamente libre cuando todos los seres humanos​ que me rodean, hombres y mujeres,​ son igualmente libres, de manera que cuanto​ más numerosos son los hombres libres​ que me rodean y más profunda y más​ amplia es su libertad, más extensa, más profunda y más amplia viene a ser mi libertad.​**",
	"**Cada​ tipo que crece podría​ ser un repetidor, un pequeño maestro,​ el desencadenante de una reacción en cadena que en sí misma es capaz​ de cambiar el mundo.**",
}

var tops [5000][6]string = [5000][6]string{
	{"1", "José", "197.36333542095375", "14-10-2012", "123456789"},
	{"2", "", "", ""},
	{"3", "", "", ""},
	{"4", "", "", ""},
	{"5", "", "", ""},

	{"1", "", "", ""},
	{"2", "", "", ""},
	{"3", "", "", ""},
	{"4", "", "", ""},
	{"5", "", "", ""},

	{"1", "", "", ""},
	{"2", "", "", ""},
	{"3", "", "", ""},
	{"4", "", "", ""},
	{"5", "", "", ""},

	{"1", "", "", ""},
	{"2", "", "", ""},
	{"3", "", "", ""},
	{"4", "", "", ""},
	{"5", "", "", ""},

	{"1", "", "", ""},
	{"2", "", "", ""},
	{"3", "", "", ""},
	{"4", "", "", ""},
	{"5", "", "", ""},


}

const INVISIBLE string = "​"

func split_curr() string {
	arrayed := strings.Split(current_text, " ")
	return arrayed[0] + " " + arrayed[1] + " " + arrayed[2] + " " + arrayed[3] + " " + arrayed[4] + " " + arrayed[5] + "..."
}

func is_illegal(s string) bool{
	x := strings.Contains(s, INVISIBLE)
	return x
}

func split(tosplit string, sep rune) []string {
	var fields []string

	last := 0
	for i, c := range tosplit {
		if c == sep {
			fields = append(fields, string(tosplit[last:i]))
			last = i + 1
		}
	}
	fields = append(fields, string(tosplit[last:]))

	return fields
}

func CountWords(s string) int {
	return len(strings.Fields(s))
}

func average_word_length(s string) float64 {
	arrayed := strings.Split(s, " ")
	var total = 0
	for i := 0; i < len(arrayed); i++ {
		total += len([]rune(arrayed[i]))
	}
	var x float64 = float64(total)/float64(len(arrayed))
	return x
}

var random int
var current_text = textos[0]
var is_started bool = false
var time_elapsed float64 = 0

var wpm float64 = 0

var start_author string

var result_splited[]string

func start() {
	time_elapsed = 0
	if is_started == false {
		is_started = true
		for is_started {
			time.Sleep(1 * time.Millisecond)
			time_elapsed = time_elapsed + 1
		}
	}

}

func calculate_wpm() {
	var length = len([]rune(current_text))/* - (CountWords(current_text)) */
	var length_as_a_float float64 = float64(length)
	wpm = length_as_a_float / 5 / time_elapsed * 60000
}

var errores int = 0
var errores_s string

var lista_errores string

func calculate_errors(sent string, current string) {
	/* reseting */
	errores = 0
	lista_errores = ""

	A := sent
	sent_arrayed := strings.Split(A, " ")

	B := current
	text_arrayed := strings.Split(B, " ")

	if len(text_arrayed) == len(sent_arrayed) {
		for i := 0; i < len(text_arrayed); i++ {
				if text_arrayed[i] != sent_arrayed[i] {
					if lista_errores != "" {
						lista_errores = lista_errores + ", " + sent_arrayed[i]
						errores++
					} else {lista_errores = sent_arrayed[i]}
				}
			}
	}

	if  len(text_arrayed) > len(sent_arrayed) {

	}


	errores_s = strconv.FormatInt(int64(errores), 10)
}
/*
func calculate_top(pos int) {
	_1 := tops[pos][0]
	_1_arrayed := strings.Split(_1, " ")

	var a = string(_1[1])

	fmt.Println(a)
	fmt.Println(_1_arrayed)
}
*/
