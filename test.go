package main

import (
		"time"
	    "strings"
	    "strconv"
	    "bufio"
	    "os"
	    "fmt"
)

var almacenar_tops_en_archivo [60000]string /*        TOPS_lines          */

func escribir_tops_a_base_de_datos() {
   var i, j int
   var n int = 10000 /*        TOPS_lines/6         */
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

var tops [10000][6]string = [10000][6]string{
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
					} else {
						lista_errores = sent_arrayed[i]
						errores++
					}
				}
			}
	}

	if  len(text_arrayed) > len(sent_arrayed) {

	}


	errores_s = strconv.FormatInt(int64(errores), 10)
}











var textos [1000]string = [1000]string{
	"Escucha, las reglas propias... se tratan de decidir conseguir algo usando medios y maneras propias para conseguirlo. Por eso decimos que son nuestras reglas. Precisamente por eso podemos afrontar sinceramente los desafíos y darlo todo. Y si fracasamos, hay que retomar la práctica y soportar duros entrenamientos para lograrlo. Y así, dedicándote a ello, creas tus propias reglas.",
	"Pocos ven lo que somos, pero todos ven lo que aparentamos. Nunca intentes ganar por la fuerza lo que puede ser ganado por la mentira. Nunca se logró nada grandioso sin peligro.",
	"Tal vez solo nos espere un camino oscuro por delante. Pero aún así necesitas creer y seguir adelante. Cree en que las estrellas iluminarán tu camino, incluso un poco. Vamos, ¡emprendamos una aventura!",
	"Sólo soy verdaderamente libre cuando todos los seres humanos que me rodean, hombres y mujeres, son igualmente libres, de manera que cuanto más numerosos son los hombres libres que me rodean y más profunda y más amplia es su libertad, más extensa, más profunda y más amplia viene a ser mi libertad.",
	"Cada tipo que crece podría ser un repetidor, un pequeño maestro, el desencadenante de una reacción en cadena que en sí misma es capaz de cambiar el mundo.",
	"Es natural que estas cosas se produzcan necesariamente así a partir de tales hombres. Y el que así no lo acepta, pretende que la higuera no produzca su zumo.",
	"Su cabello empezó a caer, sus ojos a derretirse, su boca se abría y veía solamente mierda. El castaño desaparece, el blanco se arruga; una dulce niña sin más. Y la veía, y no miraba; y en ella pensaba, se escondía. En ese instante se fue la furia, el dolor cambió a escalofríos, el desaliento en calma; y ahora moriré solo como siempre lo soñé y como siempre lo esperé.",
	"También sabía que era uno de los músicos del coro, y aunque nunca se había atrevido a levantar la vista para comprobarlo durante la misa, un domingo tuvo la revelación de que mientras los otros instrumentos tocaban para todos, el violín tocaba solo para ella.",
	"No me mires que nos miran, nos miran que nos miramos. Miremos que no nos miren, y cuando no nos miren, nos miraremos. Porque si nos miran que nos miramos, miran que nos amamos.",
	"En el símbolo yin-yang hay un punto blanco sobre la parte negra y un punto negro sobre la parte blanca. Esto sirve para ilustrar el equilibrio en la vida, porque nada puede sobrevivir mucho tiempo yendo a cualquiera de los extremos, tanto si es el yin puro (negativismo) como el yang puro (positivismo). El calor extremo mata, como también lo hace el frío extremo. Ningún extremo violento perdura. Nada permanece, excepto la sobria moderación.",
	"Hace mucho tiempo que me hubiera suicidado de no haber leído en alguna parte que es un pecado quitarse voluntariamente la vida mientras pueda hacerse todavía una buena acción. La vida es hermosa, pero la mía está envenenada para siempre.",
	"En un lugar de la Mancha, de cuyo nombre no quiero acordarme, no hace mucho tiempo que vivía un hidalgo de los de lanza en astillero, adarga antigua, rocín flaco y galgo corredor. Una olla de algo más vaca que carnero, salpicón las más noches, duelos y quebrantos los sábados, lentejas los viernes, algún palomino de añadidura los domingos, consumían las tres partes de su hacienda.",
	"Tú dices que amas la lluvia, sin embargo usas un paraguas cuando llueve. Tú dices que amas el sol, pero siempre buscas una sombra cuando el sol brilla. Tú dices que amas al viento, pero cierras las ventanas cuando el viento sopla. Por eso es que tengo miedo cuando dices que me amas.",
	"El miedo es como un fuego. Te va quemando por dentro. Si lo controlas, entras en calor. Pero si llega a dominarte, te quemará a ti y a todo lo que te rodea.",
	"El sapo sapote no sabe decir \"power\", el sapo sapote creo que sabe decir \"power\", el sapo sapote sí sabe decir \"power\".",
	"Raras veces nos topamos con un individuo capaz de revisar la idea que tiene de su propia inteligencia y sus propios límites bajo la presión de los acontecimientos, hasta el punto de abrirse a nuevas perspectivas sobre lo que aún puede aprender.",
	"Yo aconsejo a mis pacientes que evalúen cuidadosamente el momento adecuado para la confrontación. No se trata de disparar sin haber apuntado, pero tampoco de postergar indefinidamente la confrontación.",
	"El día de pascua nevó mucho, pero al siguiente sopló de improviso un viento cálido, amontonándose las nubes, y por espacio de tres días con sus noches no dejó de caer una lluvia tibia. El viento se calmó el jueves y entonces se extendió sobre la tierra una espesa bruma de color gris, como para ocultar los misterios que se producían en la naturaleza.",
	"Las personas con alta tolerancia a la frustración pueden hacer frente a los contratiempos con éxito. Las personas con baja tolerancia a la frustración pueden sentirse frustradas por inconvenientes cotidianos aparentemente menores, como atascos de tráfico o esperar en la cola del supermercado. Las personas con baja tolerancia a la frustración pueden renunciar a tareas difíciles de inmediato. La idea de tener que esperar en la cola o trabajar en una tarea que no entienden puede ser intolerable.",
	"Te sientas frente a un tablero y repentinamente tu corazón brinca. Tu mano tiembla al tomar una pieza y moverla. Pero lo que el ajedrez te enseña es que tú debes permanecer ahí con calma y pensar si realmente es una buena idea o si hay otras ideas mejores.",
	"Sabes, Cristo me contó una vez el secreto de la resurrección. Pero estábamos en una boda en Canaán y me emborraché y se me olvidó. ¿Que si lo conocía? El tipo me debe doce dólares.",
	"Debía de tener unos diez u once años cuando desapareció su madre. Era una mujer alta, elegante, más bien callada, de movimientos lentos y precioso cabello rubio. A su padre lo recordaba de manera más vaga como un hombre moreno y delgado, vestido siempre pulcramente de negro.",
	"Leyendo novelas de detectives suelo sentir una sensación de alivio ante el asesinato de un personaje que al pasar la página no me molestará con su existencia. De todos modos, las historias de detectives siempre tienen demasiados personajes. Muchos mencionados al principio pero nunca vistos, que permanecen fuera de la historia, adquiriendo una calidad portentosa horrible. Es mejor que se hayan ido.",
	"La forma en que actuó muestra que le falta la madurez que es necesaria para este trabajo. Necesito a alguien que tenga, pues primero, un mínimo de cuidado para este trabajo, pero también, alguien que quiera aprender y ampliar sus conocimientos. Y por lo que puedo decir, él no tiene ninguna de esas cualidades. Entiendo que él necesita dinero y trabajo, pero yo no soy caridad para los que viven en los sótanos de sus padres. Dile que necesita salir de su trasero y trabajar.",
	"Me giré sobre mis pasos y a medida que me acercaba a la habitación de mi compañero de piso, el olor se hacía más espeso. Vacilé antes de abrir la puerta. Las cortinas flotaban sin rumbo, despavoridas y sobre su lecho, mi amigo yacía con los ojos vidriosos y el cuerpo inerte, sin vida, acompañado por ella.",
	"No te obstines, pues, en mantener como única opinión la tuya creyéndola la única razonable. Todos los que creen que ellos solos poseen una inteligencia, una elocuencia o un genio superior a los de los demás, cuando se penetra dentro muestran solo la desnudez de su alma.",
	"¿Por qué salimos corriendo cuando notamos la menor conexión? ¿Por qué luchamos contra lo que más nos atrae? Quizá es porque cuando encontramos algo o alguien a quién aferrarnos lo ansiamos como al aire. Y nos aterra perderlo... Y creedme, uno aprende a estar solo. Pero la mayoría de las cosas son mejores si las compartes.",
	"Pues si vemos lo presente cómo en un punto se es ido y acabado, si juzgamos sabiamente, daremos lo no venido por pasado. No se engañe nadie, no, pensando que ha de durar lo que espera más que duró lo que vio, pues que todo ha de pasar por tal manera.",
	"Había una casa abajo, junto al estruendo de las olas desbaratándose contra los cantiles, donde el amor era más intenso porque tenía algo de naufragio.",
	"El hombre intentó mover la cabeza en vano. Echó una mirada de reojo a la empuñadura del machete, húmeda aún del sudor de su mano. Apreció mentalmente la extensión y la trayectoria del machete dentro de su vientre, y adquirió fría, matemática e inexorable, la seguridad de que acababa de llegar al término de su existencia. La muerte.",
	"Todos mis instintos son una forma, y todos los hechos son otras, y me temo mucho que los jurados británicos todavía no han alcanzado ese tono de inteligencia cuando van a dar la preferencia a mis teorías.",
	"No somos más que unos hombres elaborados para actuar en nombre de la venganza que consideramos justicia. Pero cuando llamamos a nuestra venganza justicia, solo engendramos más venganza, forjando el primer eslabón de las cadenas de odio.",
	"La oscuridad, la algarabía y sobre todo el acto ilegal que estaba a punto de cometer no formaban parte de mi vida cotidiana, sin embargo, proseguí y llegué a la habitación. Respiré hondamente. Abrí la puerta con el temblor de la mano como testigo y reuniendo el poco valor que me quedaba crucé el umbral prohibido.",
	"Existe una conexión fundamental entre lo que uno parece y lo que uno es. Todos nos contamos una historia sobre nosotros mismos. Siempre. Continuamente. Esa historia es la que nos convierte en lo que somos. Nos construimos a nosotros mismos a partir de esa historia.",
	"No hay razón para sufrir. La única razón por la que sufres es porque así tú lo exiges. Si observas tu vida encontrarás muchas excusas para sufrir, pero ninguna razón válida. Lo mismo es aplicable a la felicidad. La única razón por la que eres feliz es porque tú decides ser feliz. La felicidad es una elección, como también lo es el sufrimiento.",
	"Si aún respiras, eres de los que más suerte tienen, porque la mayoría de nosotros jadeamos con pulmones corruptos y nos desatendemos a nosotros mismos en nuestro afán de recolectar los nombres de los amantes con quienes no fue bien.",
	"Si hay algo más importante que mi ego en los alrededores, lo quiero capturado y matado cuanto antes.- Por un momento, no pasó nada. Luego, después de un segundo más o menos, nada continuó sucediendo.",
	"Mi madre me dice que no mastico mi comida lo suficiente. Dice que estoy haciendo que sea más difícil para mi cuerpo obtener los nutrientes esenciales que necesita. Si ella estuviera aquí, le recordaría que estoy comiendo una tarta de arándanos.",
	"El libro es una criatura frágil. Sufre el paso del tiempo, el acoso de los roedores y las manos torpes, así que el bibliotecario protege los libros no sólo contra el género humano sino también contra la naturaleza, dedicando su vida a esta guerra contra las fuerzas del olvido.",
	"No es de extrañar que nos pasemos la vida entera soñando. Estar despierto y verlo todo tal y como es en realidad... nadie podría soportarlo mucho tiempo.",
	"Despertó empapado en sudor, con el corazón desbocado y casi sin aire. Al abrir los ojos, todo era oscuridad. Intentó incorporarse en la cama, estirar los brazos, pero el techo y las paredes de madera se lo impedían.",
	"Sabía que esto enviudaría a tu madre y te dejaría huérfano. Pero había encontrado mi destino. Así que abandoné a mi hijo. Ahora tengo trabajo que hacer, una cantidad infinita. Debo encontrar vida inteligente.",
	"La gente que más me gusta es la que ha fallado, ha sido lastimada, ha llorado, ha visto cosas terribles, y sin embargo, no ha perdido su capacidad para seguir amando.",
	"Generalmente, esa sensación de estar solo en el mundo aparece mezclada con un orgulloso sentimiento de superioridad: desprecio a los hombres, los veo sucios, feos, incapaces, ávidos, groseros, mezquinos; mi soledad no me asusta, es casi olímpica.",
	"Un ejemplo de conjunto de palabras aleatorias es: como pasa mismo nos oh creo mejor nada tenemos papas qué sea acuerdo ellos un vas tiene fue nadie estamos va parece mi tus sin ahí mira alguien",
	"A primera vista, la llave y la cerradura en la que encaja pueden parecer muy distintas. Diferentes en su forma, diferentes en su función, diferentes en su diseño. El hombre que las mira sin conocer su verdadera naturaleza puede pensar que son opuestas, pues una sirve para abrir y la otra para mantener cerrado. Sin embargo, examinándolas con atención, se ve que sin una la otra no sirve para nada. El hombre sabio ve que la cerradura y la llave fueron creadas para el mismo propósito.",
	"No conocerás el miedo. El miedo mata la mente. El miedo es la pequeña muerte que conduce a la destrucción total. Afrontaré mi miedo. Permitiré que pase sobre mí y a través de mí. Y cuando haya pasado giraré mi ojo interior para escrutar su camino. Allá donde haya pasado el miedo ya no habrá nada. Solo estaré yo.",
	"No tengo hambre, tengo ansiedad. Ver tanta gente acá reunida me dan ganas de fumar. Lo que pasa es que en casa muy solo se está. Ahí va otro inadaptado intentando encajar. No sé muy bien en dónde... Encajar, no sé ni para qué... Si no hay botón de pausa, no hay rebobinar. Por eso es que no me pierdo ningún evento social. Hace tiempo estoy pensando tengo que parar, mientras tanto por las dudas sigo a toda velocidad.",
	"Amar es prolongar el breve instante de angustia de ansiedad y de tormento en que, mientras espero, te presiento en la sombra suspenso y delirante. Yo quisiera anular de tu cambiante y fugitivo ser el movimiento, y cautivarte con el pensamiento y por él sólo ser tu solo amante. Pues si no quiero ver, mientras avanza el tiempo indiferente, a quien más quiero, para soñar despierto en tu tardanza. La sola posesión de lo que espero, es porque cuando llega mi esperanza es cuando ya sin esperanza muero.",
	"En las noches de luna llena, la lívida luz se refleja en los adoquines del pavimento, que refulgen con un aura fantasmal, alimentando la imaginación de los cuentos de abuelas generación tras generación. Ese brillo mortecino abraza la soledad de sus calles en las noches eternas y se une al tenue ulular del viento evocando un lamento por la grandeza perdida de los tiempos pasados.",
	"El muro de tu olvido absorbe la caída de mis pupilas en la quieta mirada del retiro... La noche exhala sus tules. Estás y no... Con el ojo estrangulado en el vano intento de acariciar tu imagen verifico tu ausencia. Estás y no... Te figuro en un rayo de la luna, parece que me observas, levanto la cabeza, para asirme en tus cabellos. Estás y no... En el muro, en la morada, en mis ojos y en la luna en el hoy y en el ayer. Estás y no...",
	"Estás preocupado por lo que fue y por lo que va a ser... Hay un dicho: \"El ayer es historia, el mañana es un misterio, pero el hoy es un obsequio. Por eso se llama presente.\"",
	"Mi vida comenzó un otoño del 96, y acabó mucho antes de los 16, mi habla fue precoz, pero qué poquito le duró la voz, cuando el tiempo la cercenó con su hoz.",
	"Cerca de quinientas almas viven en esta pequeña villa amurallada de origen medieval que se eleva orgullosa sobre la cima de una colina, rodeada de verdes pastos y tierras de labranza. A lo largo de los siglos sus vecinos se han distinguido como gente sencilla y trabajadora.",
	"De Tarde, arrimaba a la puerta una de las sillas y mateaba con seriedad, puestos los ojos en la enredadera del muro de la inmediata casa de altos. Años de soledad le habían enseñado que los días, en la memoria, tienden a ser iguales, pero que no hay un día, ni siquiera de cárcel o de hospital, que no traiga sorpresas, que no sea al trasluz una red de mínimas sorpresas.",
	"No sabía en qué estaba pensando, no comprendía el significado de su aroma. Estos aromas eran brillantes, cinéticos, me quemaban la nariz. Pero por debajo de ellos, estaba Joe. Era como humo y tierra y lluvia. Eran los aromas que siempre había asociado a él, pero intensificados por miles de veces más. Quería enterrarme en ellos, enroscándome hasta que su esencia me cubriera por completo.",
	"Me hubiera gustado que viese usted mi nombre en un libro, aunque no pudiese leerlo. Me hubiera gustado que estuviese aquí, conmigo, para ver que su hijo conseguía abrirse camino y llegaba a hacer algunas de las cosas que a usted nunca le dejaron. Me hubiera gustado conocerle, padre, y que usted me hubiera conocido a mí. Le convertí a usted en un extraño para olvidarle y ahora el extraño soy yo.",
	"La vida de un crítico es sencilla en muchos aspectos. Arriesgamos poco y tenemos poder sobre aquellos que ofrecen su trabajo y su servicio a nuestro juicio. Preferimos las críticas negativas, que es divertida de escribir y de leer. Pero la triste verdad que debemos afrontar es que en el gran orden de las cosas, cualquier basura tiene más significado que lo que deja ver nuestra crítica.",
	"Cualquiera puede enfadarse, eso es algo muy sencillo. Pero enfadarse con la persona adecuada, en el grado exacto, en el momento oportuno, con el propósito justo y del modo correcto, eso, ciertamente, no resulta tan sencillo.",
	"Nuestro conocimiento nos ha hecho cínicos. Nuestra inteligencia, duros y secos. Pensamos demasiado y sentimos muy poco. Más que máquinas, necesitamos humanidad. Más que inteligencia, tener bondad y dulzura. Sin estas cualidades, la vida será violenta. Se perderá todo... A los que puedan oírme, les digo; no desesperéis. La desdicha que padecemos no es más que la pasajera codicia y la amargura de hombres que temen seguir el camino del progreso humano. El odio de los hombres pasará.",
	"Si escuchas mi voz y al girar tu cabeza no me encuentro ahí, el que te hable será mi espíritu que viaja a través del tiempo y del espacio en tus recuerdos.",
	"La guerra es lo más importante para el estado, el terreno de la vida y de la muerte, el camino a la supervivencia o la desaparición. No puede ser ignorada.",
	"No me gustan los términos de buena o mala persona porque es imposible ser totalmente bueno con todo el mundo, o totalmente malo con todo el mundo. Para algunos, eres una buena persona, mientras que para otros eres una mala persona.",
	"Las palabras son pálidas sombras de nombres olvidados. Los nombres tienen poder, y las palabras también. Las palabras pueden hacer prender el fuego en la mente de los hombres. Las palabras pueden arrancarles lágrimas a los corazones más duros. Existen siete palabras que harán que una persona te ame. Existen diez palabras que minarán la más poderosa voluntad de un hombre. Pero una palabra no es más que la representación de un fuego. Un nombre es el fuego en sí.",
	"En el camino hacia aquí se sientan cómodos y disfrutan el trayecto. A veces nos paramos y vemos los pájaros donde no los hay, y miramos atardeceres cuando está lloviendo. Nos lo pasamos fantásticamente y recibo muy buena propina.",
	"Si la felicidad tuviera una forma, tendría forma de cristal, porque puede estar a tu alrededor sin que la notes. Pero si cambias de perspectiva, puede reflejar una luz capaz de iluminarlo todo.",
	"Había una mujer, fue la primera vez que encontré a alguien que estuviera verdaderamente vivo. Al menos, eso fue lo que pensé. Ella era... la parte de mí que perdí en algún lugar del camino, la parte que faltaba, la que deseaba.",
	"Las pesadillas son muy raras. Tu cerebro es el autor, espectador y escenario de una película de terror cuyo guión está siendo escrito mientras lo ves.",
	"Ya no jugamos tres en raya porque es un juego aburrido y siempre acaba en empate; no hay forma de ganar, el juego en sí no tiene sentido. Pero en la sala de guerra creen que puedes ganar una guerra nuclear, que puede haber pérdidas aceptables.",
	"Súbitamente desapareció en el fondo de la bahía, sin duda para embestirme desde abajo. Primero fue un terrible costalazo. Luego se asomó para mostrarme sus nueve hileras de colmillos blancos, sus ojos apagados, sin odio ni crueldad, como los de quien ejecuta una rutina.",
	"Qué extraño. Esto es irracional. Sin embargo, fueron los seres humanos quienes prefirieron las emociones en lugar de las ganancias o las pérdidas. Yo solo quería que mi hermano mayor me reconociera, pero me faltaba algo. Lo siento, no pude cambiar hasta el final.",
	"Eso es porque todos queremos no tener problemas, arreglarnos a nosotros mismos. Buscamos una solución mágica para mejorarnos, pero ninguno de nosotros sabe realmente lo que está haciendo. Eso es todo lo que los humanos podemos hacer: adivinar, intentar, esperar.",
	"Ahora que hemos establecido el dilema de nuestro protagonista, pasemos a nuestro antagonista. A muchos kilómetros de distancia, a través de las llanuras abiertas, otra bella bestia salvaje se abre paso hasta un abrevadero.",
	"Sabes es increíble tener la misma rutina desde hace tres años, levantarte, escuchar unas rolitas, imaginar cosas para ser feliz, solo te haces daño a ti mismo, pero- por lo menos eres feliz en ese mundo, en tu imaginación, eso es lo único que te mantiene feliz, ver esa sonrisa en su rostro, imaginarte un beso, dormir con esa persona, una salida al parque y que por primera vez te dé una señal de que le gustas...",
	"Crearía un perfume que no sólo fuera humano, sino sobrehumano. Un aroma de ángel, tan indescriptiblemente bueno y pletórico de vigor que quien lo oliera quedaría hechizado y no tendría más remedio que amar a la persona que lo llevara, o sea, amarle a él, Grenouille, con todo su corazón.",
	"No podía moverme, no podía hablar. Solo podía escuchar y ver. Mi grado de desesperación se acercaba a dimensiones infinitas mientras la silueta seguía ahí, muy bajito, desde su sucio rincón, alimentando mi nuevo estado de locura.",
	"Tenía una mezcla de miedo y resaca porque había escuchado desde chaval miles de historias sobre heroína, putas y problemas. Comencé a andar más rápido para salir de allí cuanto antes y a unos metros de mí escuché una gran carcajada seguido de una voz que se acercaba, pero no acerté a entender lo que decía.",
	"¿Realmente cree que el enemigo atacaría sin provocación, usando tantos misiles, bombarderos y submarinos, para que no tengamos más remedio que aniquilarlos por completo? General, está escuchando a una máquina, hágale un favor al mundo y no actúe como una.",
	"Pero luego el tráfico, la estampida, los autos, los camiones, las bocinas... Esta pobre mujer gritaba y me di cuenta de que no era un renacimiento, era una especie de ilusión vertiginosa de renovación que ocurre en el momento final antes de la muerte.",
	"Deseaba ver el infinito, pero en realidad nunca salía de los límites de su propia cabeza. Decirle la verdad sería como dar una patada a un perro de aguas.",
	"Abrí la puerta del coche para, como un día cualquiera, recorrer el tramo que separa mi casa del trabajo. El tráfico era el habitual, como habitual es también el cabreo que te produce.",
	"La gente que más me gusta es la que ha fallado, ha sido lastimada, ha llorado, ha visto cosas terribles, y sin embargo, no ha perdido su capacidad para seguir amando.",
	"En segundo lugar, opino que no es tan importante cuantas medidas tomen las autoridades: no basta que estén asociados a campañas nacionales de concientización y enseñanza, si no pueden transmitir eficazmente esta conciencia. Estoy seguro de que un gran porcentaje de la población desconoce las consecuencias del uso del plástico o que este se obtiene del petróleo, un recurso natural no renovable. Cuando se desconoce, no se entiende. Y, si no se entiende, no se hace.",
	"La vida es un viaje y no nos damos cuenta, nos anclamos a lo que sea que permita que nuestra existencia tenga algún sentido, somos adictos a creer en algo pero muchos son incapaces de creer en ellos mismos.",
	"Hay mucho más. Hay todo lo que va más allá, todo lo que está afuera. Y todo lo de atrás, desde hace muchísimo, muchísimo tiempo. Yo recibí todas esas cosas cuando me seleccionaron. Y aquí en esta habitación, yo solo, las vuelvo a experimentar una y otra vez. Es así como viene la sabiduría. Y como configuramos nuestro futuro.",
	"Ahora se sorprendía a menudo enfadado: irracionalmente enfadado con sus compañeros de grupo, porque estaban satisfechos con unas vidas que no tenían nada de la vibración que la suya estaba adquiriendo. Y enfadado consigo mismo, por no poder cambiar esa situación.",
	"Un amor por ocultar. Aunque en cueros no hay donde esconderlo, lo disfrazan de amistad cuando sale a pasear por la ciudad. Una opina que aquello no está bien. La otra opina que qué se le va a hacer. Y lo que opinen los demás está de más. Quien detiene palomas al vuelo volando a ras de suelo. Mujer contra mujer.",
	"Entrará el mar lentamente en tus venas, oh nadador que esperaste la noche y la soledad para medir tus fuerzas con la tormenta, digo con tu propio destino, desde el principio con algo que sabías superior a ti mismo.",
	"Fue cuestión de segundos, bebió más de la cuenta, estaba alegre, aún oigo su carcajada. Para él no tiene remedio y para mí ya nada tiene sentido. Bajo del metro y recorro lentamente los trescientos sesenta y cinco escalones que me conducen a él.",
	"Al recuperarme, abrí los ojos y vi a mi mujer a mi lado. Estaba tumbado en la cama de ese hospital que había visto desde la ventana. Un sentimiento de alivio recorrió todo mi cuerpo. Había vuelto de un viaje que nunca recomendaré a nadie.",
	"El dolor desaparece con el tiempo. Pero no deseo ser curado por el tiempo, porque cuando huyes del dolor, con el anhelo de olvidar, lo único que logras es quedarte atascado. Te vuelves incapaz de seguir adelante.",
	"Otro ejemplo de conjunto de palabras aleatorias es: en sabe estoy papas tipo espero fuera mejor ahora quiere favor necesito ser podría nada ella sé papas desde estado era mejor noche amigo creo puedes",
	"Otro ejemplo de conjunto de palabras aleatorias es: tipo sabe nadie no nadie ya usted has sea dos espero de sobre día desde antes sin algo papas todos así porque con te nuevo ni por otro cosa hablar",
	"El problema es, Leslie, que el mundo no se va a acabar mañana. El sol se va a levantar por ese lado, será un viernes común y corriente, y tú seguirás en la misma posición en la que estás ahora.",
	"Nos sentimos más apasionados cuando besamos, más humildes cuando nos arrodillamos y más enojados cuando agitamos los puños. Es decir, que el beso no solamente expresa pasión sino que la origina. Los papeles llevan consigo al mismo tiempo ciertas acciones y las emociones y actitudes que corresponden a esas acciones. El profesor que pone en escena un acto que finge sabiduría, llega a sentirse sabio. El predicador llega a creer en lo que predica.",
	"En realidad, a uno no le gobierna en modo alguno la mente. Lo que la mente revela es una corriente interminable de opciones, todas ellas disfrazadas en la forma de recuerdos, fantasías, temores, conceptos, etc. Para liberarse del dominio de la mente, solo hay que darse cuenta de que su desfile de temas no más que un autoservicio arbitrario de opciones que pasan a través de la pantalla de la mente.",
	"Mientras sacaba las llaves del bolso, oí de repente unos pasos que se acercaban. Al volverme, vi la figura de un hombre joven que se dirigía hacia mí con una amplia sonrisa en la cara. Según avanzaba comenzó a hablar...",
	"Quisiera que me vieras y descubrieras que solo un murmullo me haría feliz, quisiera que me dieras una sonrisa y poder sentir, que estoy cerca de ti. Quisiera, que quisieras quererme, olvidarme para siempre y de tu mano andar; quisiera que quisieras quereme, y el primer beso me dieras. Eso quiera.",
	"Cuando algo sobre una persona resuena profundamente dentro de ti, es el sentimiento más maravilloso. Tu corazón se libera, tu mente se abre y te das cuenta de que hay más en el mundo de lo que nunca imaginaste.",
	"Se movía como un depredador, sus anchos hombros se balanceaban siguiendo los vaivenes de su andar y la cabeza giraba de un lado a otro, explorando. Ella tuvo la incómoda sensación de que si él quisiera, podía matar a todos los presentes sin usar más arma que sus manos.",
	"Solo podía apreciar el típico ruido que hacen esos scanners. De repente, bajo mis pies se abrió una puerta negra. Caí en picado por un túnel. Las imágenes pasaban rápidamente a mi alrededor. Eran todos los acontecimientos de mi vida. Sentí un fuerte golpe y perdí la consciencia.",
	"Te escribí cuando mi tren salía de su andén. Fue una madrugada bendecida por el beso de Diciembre. Mi mano temblaba de confesiones de cobardía; pero no pude dar marcha atrás. Así que comencé a escribirte un poema mientras mi tren salía. Y al hacerlo, el peso dentro de mí disminuyó, y solo lo hacía en viajes como estos.",
	"La fantasía es una bicicleta estática para la mente. Es posible que no te lleve a ninguna parte, pero tonifica los músculos que pueden. Por supuesto, podría estar equivocado.",
	"Rendirse es lo que destruye a la gente, cuando te niegas con todo tu corazón a rendirte entonces trasciendes tu humanidad, incluso ante la muerte nunca te rindas.",
	"Yo tengo miedo de todo. Y estoy loca. Como si tal vez piensas que estoy un poco loca, pero la gente sólo ve la punta del iceberg de la locura. Por debajo de esta apariencia de ligeramente loca y socialmente inepta, soy un completo desastre.",
	"En la vida se tienen que tomar muchas decisiones; si esas decisiones son correctas o no, nadie lo sabe. Por eso las personas suelen elegir lo que ellos creen que es correcto.",
	"Ahora se sorprendía a menudo enfadado: irracionalmente enfadado con sus compañeros de grupo, porque estaban satisfechos con unas vidas que no tenían nada de la vibración que la suya estaba adquiriendo. Y enfadado consigo mismo, por no poder cambiar esa situación.",
	"Caminó rápidamente a través de la oscuridad con el paso franco de alguien que al menos estaba seguro de que el bosque, en esta noche húmeda y ventosa, contenía cosas extrañas y terribles y ella las era.",
	"Pues, no, hermano, no la quiero, que es historia muy cansada ver que al pasar del arroyo te llegue a la boca el agua. La mujer que ha de ser propia ha de estar en una caja como el gusano de seda, hasta ser paloma blanca. Si fuiste abeja en su rosa, que buen provecho te haga; que lo que no fue posible olvidar con la mudanza de su traje, ni acabaron sus desdenes y desgracias, con lo que me has dicho sólo, hoy para siempre se acaba.",
	"Indicaciones terapéuticas del ibuprofeno: artritis reumatoide (incluyendo artritis reumatoide juvenil), espondilitis anquilopoyética, artrosis y otros procesos reumáticos agudos o crónicos. Alteraciones musculoesqueléticas y traumáticas con dolor e inflamación. Tto. sintomático del dolor leve o moderado (dolor de origen dental, dolor posquirúrgico, dolor de cabeza, migraña). Dismenorrea primaria. Cuadros febriles.",
	"Es una verdad mundialmente reconocida que un hombre soltero, poseedor de una gran fortuna, necesita una esposa. Sin embargo, poco se sabe de los sentimientos u opiniones de un hombre de tales condiciones cuando entra a formar parte de un vecindario. Esta verdad está tan arraigada en las mentes de algunas de las familias que lo rodean, que algunas le consideran de su legítima propiedad y otras de la de sus hijas.",
	"Los hombres se equivocan al creerse libres, opinión que obedece al solo hecho de que son conscientes de sus acciones e ignorantes de las causas que las determinan. Y, por tanto, su idea de \"libertad\" se reduce al desconocimiento de las causas de sus acciones, pues todo eso que dicen de que las acciones humanas dependen de la voluntad son palabras, sin idea alguna que les corresponda.",
	"Hay una fuerza motriz más poderosa que el vapor, la electricidad y la energía atómica: la voluntad.",
	"Hora crepuscular. Un guardillón con ventano angosto, lleno de sol. Retratos, grabados, autógrafos repartidos por las paredes, sujetos con chinches de dibujante. Conversación lánguida de un hombre ciego y una mujer pelirrubia, triste y fatigada. El hombre ciego es un hiperbólico andaluz, poeta de odas y madrigales, Máximo Estrella. A la pelirrubia, por ser francesa, le dicen en la vecindad Madama Collet.",
	"Es un gran placer para mí darles la bienvenida a la 19 Conferencia de Plenipotenciarios (PP-14), el evento de mayor trascendencia en el calendario de la UIT, que congrega a casi tres mil delegados de los Estados Miembros junto con otros Miembros de la UIT, entre ellos empresas del sector privado, organizaciones internacionales, la sociedad civil e Instituciones Académicas. En la Conferencia se abordarán cuestiones importantes relativas al futuro de la Unión, desde la integración digital a la implantación de la banda ancha.",
	"La manera en la que este bot calcula el WPM, es mediante la siguiente función: func calculate_wpm() {var length = len([]rune(current_text)); var length_as_a_float float64 = float64(length); wpm = length_as_a_float / 5 / time_elapsed * 60000}. Como puedes ver, solo consta de aplicar una pequeña fórmula a la longitud del texto y los milisegundos transcurridos.",
	"Estoy formado por los recuerdos de mis padres y mis abuelos y todos mis antepasados. Están en mi aspecto, en el color de mi cabello. Estoy hecho de todos los que he conocido y han cambiado mi forma de pensar.",
	"Aprecio mi pasado en detrimento de mi futuro; aunque encuentro excelentes ambos, no puedo otorgar primacía a ninguno, y solo debo censurar la injusticia de la providencia que tanto me favorece.",
	"No más que un soplo de viento es el rumor de la aprobación humana, que tan pronto viene de un extremo como del opuesto y cambia de nombre al cambiar de dirección.",
	"Deberías ser más desafiante, porque solo vives una vez, deberías enfrentar, lo que sea que quieres hacer cuando eres joven, cuando seas viejo y pienses en el pasado, no será algo malo de recordar incluso si fallaste.",
	"Así es como me gusta tenerlo. Bajo mi poder. Entre mis manos. No por ahí conspirando y tramando planes, hablando con vampiros. Ahora te tengo, pienso, finalmente te tengo donde quería tenerte.",
	"Muchos de los que viven merecen morir y algunos de los que mueren merecen la vida. ¿Puedes devolver la vida? Entonces no te apresures a dispensar la muerte, pues ni el más sabio conoce el fin de todos los caminos.",
	"Eres grande Ox. Más grande que cualquier lobo que haya visto antes. Más grande que yo, que mi padre, pero tiene sentido, ¿sabes? Porque siempre has sido así para mí. Más grande que cualquier cosa. El día que te vi, supe que las cosas jamás volverían a ser como antes. Lo abarcas todo, empequeñeces todo lo demás y cuando te veo, Ox, eres todo lo que puedo ver.",
	"La llegada de su amigo y compañero que por un contratiempo permanente la puerta nunca abrirá, aunque me digan que nunca llegará, lo esperaré y cuando cruce la puerta lo recibiré. O tal vez, solo tal vez me reciba él.",
	"La guerra deberá existir mientras defendamos nuestras vidas contra un destructor que lo devoraría todo. Pero no amo la espada por su filo, ni la flecha por su rapidez, ni al guerrero por su gloria. Solo amo lo que defienden.",
	"Significa poder enfrentar los problemas, aguantar el dolor y los golpes que te da la vida, no significa que no puedas sentir dolor ni ponerte triste, pero alguien fuerte no deja que su vida sea afectada por cualquier desdicha.",
	"La gente tiene diferentes formas de pensar, incluso cuando se comete un error... Si la persona se da cuenta de su error es posible que lo enmiende, si mantienes tu visión clara verás el futuro, es de lo que se trata esta vida...",
	"El problema es, Leslie, que el mundo no se va a acabar mañana. El sol se va a levantar por ese lado, será un viernes común y corriente, y tú seguirás en la misma posición en la que estás ahora.",
	"Al lado, se encontraba la enorme escultura de un payaso realizada en un descolorido material que Cristina tampoco pudo reconocer. Sorprendida, reparó en la falta de nariz tan significativa en todos los payasos.",
	"La gente que más me gusta es la que ha fallado, ha sido lastimada, ha llorado, ha visto cosas terribles, y sin embargo, no ha perdido su capacidad para seguir amando.",
	"Lo que realmente significa el poder del hombre sobre la naturaleza es el poder ejercido por algunas personas sobre otras, con el conocimiento de la naturaleza como instrumento.",
}
