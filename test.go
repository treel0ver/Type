package main

import (
	"time"
    "strings"
    "strconv"
    "bufio"
    "os"
    "fmt"

	//"strconv"
	//"reflect"
)

var almacenar_tops_en_archivo [3000]string

func escribir_tops_a_base_de_datos() {
   var i, j int
   var n int = 200
   var cuenta int = 0
   for  i = 0; i < n; i++ {
      for j = 0; j < 6; j++ {
         almacenar_tops_en_archivo[cuenta] = tops[i][j]
         cuenta++
      }
   }
}

func writeLines() error {
    file, err := os.Create("database/TOPS")
    if err != nil {
        return err
    }
    defer file.Close()

    w := bufio.NewWriter(file)
    for _, line := range almacenar_tops_en_archivo {
        fmt.Fprintln(w, line)
    }
    return w.Flush()
}

var textos [1000]string = [1000]string{
	"Escucha, las reglas propias... se tratan de decidir conseguir algo usando medios y maneras propias para conseguirlo. Por eso decimos que son nuestras reglas. Precisamente por eso podemos afrontar sinceramente los desafíos y darlo todo. Y si fracasamos, hay que retomar la práctica y soportar duros entrenamientos para lograrlo. Y así, dedicándote a ello, creas tus propias reglas.",
	"Pocos ven lo que somos, pero todos ven lo que aparentamos. Nunca intentes ganar por la fuerza lo que puede ser ganado por la mentira. Nunca se logró nada grandioso sin peligro.",
	"Tal vez solo nos espere un camino oscuro por delante. Pero aún así necesitas creer y seguir adelante. Cree en que las estrellas iluminarán tu camino, incluso un poco. Vamos, ¡emprendamos una aventura!",
	"Sólo soy verdaderamente libre cuando todos los seres humanos que me rodean, hombres y mujeres, son igualmente libres, de manera que cuanto más numerosos son los hombres libres que me rodean y más profunda y más amplia es su libertad, más extensa, más profunda y más amplia viene a ser mi libertad.",
	"Cada tipo que crece podría ser un repetidor, un pequeño maestro, el desencadenante de una reacción en cadena que en sí misma es capaz de cambiar el mundo.",
	"Es natural que estas cosas se produzcan necesariamente así a partir de tales hombres. Y el que así no lo acepta, pretende que la higuera no produzca su zumo.",
	"Su cabello empezó a caer, sus ojos a derretirse, su boca se abría y veía solamente mierda. El castaño desaparece, el blanco se arruga; una dulce niña sin más. Y la veía, y no miraba; y en ella pensaba, se escondía. En ese instante se fue la furia, el dolor cambió a escalofríos, el desaliento en calma; y ahora moriré solo como siempre lo soñé y como siempre lo esperé.",
	"Cada tipo que crece podría ser un repetidor, un pequeño maestro, el desencadenante de una reacción en cadena que en sí misma es capaz de cambiar el mundo.",
	"También sabía que era uno de los músicos del coro, y aunque nunca se había atrevido a levantar la vista para comprobarlo durante la misa, un domingo tuvo la revelación de que mientras los otros instrumentos tocaban para todos, el violín tocaba solo para ella.",
	"No me mires que nos miran, nos miran que nos miramos. Miremos que no nos miren, y cuando no nos miren, nos miraremos. Porque si nos miran que nos miramos, miran que nos amamos.",
}
var display_textos [1000]string = [1000]string{
	"**Escucha,​ las reglas​ propias... se tratan de decidir conseguir algo​ usando medios y maneras propias​ para conseguirlo. Por eso decimos que son nuestras​ reglas. Precisamente por eso podemos afrontar sinceramente los desafíos y darlo todo. Y si​ fracasamos, hay que retomar la práctica​ y soportar duros entrenamientos para lograrlo. Y así, dedicándote a ello​, creas tus propias reglas.​**",
	"**Pocos ven lo que somos, pero todos ven lo que aparentamos. Nunca intentes ganar por la fuerza lo que puede ser ganado por la mentira. Nunca se logró nada grandioso sin peligro.**",
	"**Tal​ vez solo nos espere un camino oscuro por​ delante. Pero aún así necesitas​ creer y seguir adelante. Cree en que las estrellas​ iluminarán tu camino, incluso un poco. Vamos,​ ¡emprendamos una aventura!**",
	"**Sólo soy​ verdaderamente libre cuando todos los seres humanos​ que me rodean, hombres y mujeres,​ son igualmente libres, de manera que cuanto​ más numerosos son los hombres libres​ que me rodean y más profunda y más​ amplia es su libertad, más extensa, más profunda y más amplia viene a ser mi libertad.​**",
	"**Cada​ tipo que crece podría​ ser un repetidor, un pequeño maestro,​ el desencadenante de una reacción en cadena que en sí misma es capaz​ de cambiar el mundo.**",
	"**Es natural​ que estas​ cosas​ se​ produzcan​ necesariamente​ así a partir​ de tales​ hombres. Y el​ que así no​ lo acepta​, pretende que la higuera​ no produzca​ su zumo.**",
	"**Su cabello​ empezó a caer,​ sus ojos a derretirse,​ su boca se​ abría y veía solamente​ mierda. El castaño desaparece,​ el blanco se arruga; una dulce niña​ sin más. Y la veía, y no miraba; y en ella pensaba, se escondía. En ese instante se fue la furia, el dolor cambió a escalofríos, el desaliento en calma; y ahora moriré solo como siempre lo soñé y como siempre lo esperé.**",
	"**Cada​ tipo que crece podría ser un​ repetidor,​ un pequeño maestro,​ el desencadenante de una reacción​ en cadena que en sí misma​ es capaz de cambiar​ el mundo.**",
	"**También​ sabía​ que era uno​ de los músicos del coro, y aunque nunca se​ había atrevido​ a levantar la vista para comprobarlo durante​ la misa, un domingo tuvo​​ la revelación​ de que mientras​ los otros instrumentos​ tocaban para todos,​ el violín​ tocaba solo para​ ella.**",
	"**No​ me mires​ que​ nos miran,​ nos miran​ que nos​ miramos.​ Miremos​ que no​ nos miren,​ y cuando​ no nos​ miren, nos​ miraremos.​ Porque​ si nos miran​ que nos miramos,​ miran​ que nos amamos.​**",
}

var tops [5000][6]string = [5000][6]string{
	{"1", "Username", "125.00 (wpm)", "2022 (fecha)", "0", "123456 (User.ID)"},
	{"2", "0", "0", "0", "0", "0"},
	{"3", "0", "0", "0", "0", "0"},
	{"4", "0", "0", "0", "0", "0"},
	{"5", "0", "0", "0", "0", "0"},

	{"1", "0", "0", "0", "0", "0"},
	{"2", "0", "0", "0", "0", "0"},
	{"3", "0", "0", "0", "0", "0"},
	{"4", "0", "0", "0", "0", "0"},
	{"5", "0", "0", "0", "0", "0"},
}

var levels [500][2]string = [500][2]string{
	{"12345678910...", "0"},
	{"87652345654...", "0"},

}
const INVISIBLE string = "​"

func how_many_texts() int {
	var cuenta int
	for i := 1; i <= len(textos); i++ {
	    if textos[i] != "" {
	    	cuenta++
	    } else { return cuenta+1 }
	}
	return 0
}

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

var time_soon bool = true

var wpm float64 = 0

var start_author string

var result_splited[]string

var conseguiste_alguna_marca bool = false
var conseguiste_superar_alguna_marca bool = true

func start() {
	time_elapsed = 0
	time_soon = true
	if is_started == false {
		is_started = true
		time.Sleep(500 * time.Millisecond)
		for is_started {
			time.Sleep(1 * time.Millisecond)
			time_elapsed = time_elapsed + 1
			if time_elapsed == 6000 {
				time_soon = false
			}
			if time_elapsed == 180000 {
				is_started = false
			}
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
