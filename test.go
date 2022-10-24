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
}

var display_textos [1000]string = [1000]string{
	"**Escucha,​ las reglas​ propias... se tratan de decidir conseguir algo​ usando medios y maneras propias​ para conseguirlo. Por eso decimos que son nuestras​ reglas. Precisamente por eso podemos afrontar sinceramente los desafíos y darlo todo. Y si​ fracasamos, hay que retomar la práctica​ y soportar duros entrenamientos para lograrlo. Y así, dedicándote a ello​, creas tus propias reglas.​**",
	"**Pocos ven lo que somos, pero todos ven lo que aparentamos. Nunca intentes ganar por la fuerza lo que puede ser ganado por la mentira. Nunca se logró nada grandioso sin peligro.**",
	"**Tal​ vez solo nos espere un camino oscuro por​ delante. Pero aún así necesitas​ creer y seguir adelante. Cree en que las estrellas​ iluminarán tu camino, incluso un poco. Vamos,​ ¡emprendamos una aventura!**",
	"**Sólo soy​ verdaderamente libre cuando todos los seres humanos​ que me rodean, hombres y mujeres,​ son igualmente libres, de manera que cuanto​ más numerosos son los hombres libres​ que me rodean y más profunda y más​ amplia es su libertad, más extensa, más profunda y más amplia viene a ser mi libertad.​**",
	"**Cada​ tipo que crece podría​ ser un repetidor, un pequeño maestro,​ el desencadenante de una reacción en cadena que en sí misma es capaz​ de cambiar el mundo.**",
	"**Es natural​ que estas​ cosas​ se​ produzcan​ necesariamente​ así a partir​ de tales​ hombres. Y el​ que así no​ lo acepta​, pretende que la higuera​ no produzca​ su zumo.**",
	"**Su cabello​ empezó a caer,​ sus ojos a derretirse,​ su boca se​ abría y veía solamente​ mierda. El castaño desaparece,​ el blanco se arruga; una dulce niña​ sin más. Y la veía, y no miraba; y en ella pensaba, se escondía. En ese instante se fue la furia, el dolor cambió a escalofríos, el desaliento en calma; y ahora moriré solo como siempre lo soñé y como siempre lo esperé.**",
	"**También​ sabía​ que era uno​ de los músicos del coro, y aunque nunca se​ había atrevido​ a levantar la vista para comprobarlo durante​ la misa, un domingo tuvo​​ la revelación​ de que mientras​ los otros instrumentos​ tocaban para todos,​ el violín​ tocaba solo para​ ella.**",
	"**No me mires​ que nos miran,​ nos miran​ que nos miramos.​ Miremos que no nos​ miren, y cuando no​ nos miren, nos​ miraremos. Porque si nos miran que nos miramos,​ miran​ que nos amamos.​**",
	"**En el​ símbolo yin-yang​ hay un punto​ blanco sobre​ la parte negra​ y un punto negro​ sobre la parte blanca.​ Esto​ sirve para ilustrar​ el equilibrio​ en la vida,​ porque nada puede sobrevivir mucho tiempo​ yendo a cualquiera de​ los extremos, tanto​ si es el yin​ puro (negativismo)​ como el yang​ puro (positivismo).​ El calor extremo mata,​ como también lo​ hace el frío extremo. Ningún​ extremo violento perdura. Nada permanece,​ excepto la sobria moderación.​**",
	"**Hace​ mucho​ tiempo​ que me hubiera​ suicidado de no haber​ leído en alguna parte que​ es un pecado quitarse​ voluntariamente la​ vida mientras pueda​ hacerse todavía una buena acción. La​ vida es hermosa,​ pero la mía está envenenada​ para siempre.​**",
	"**En un lugar de la Mancha, de cuyo nombre no quiero acordarme, no hace mucho tiempo que vivía​ un hidalgo de los de lanza​ en astillero, adarga antigua,​ rocín flaco y galgo corredor. Una olla​ de algo más vaca que carnero,​ salpicón las más​ noches, duelos y quebrantos los sábados, lentejas​ los viernes, algún palomino de añadidura​ los domingos, consumían​ las tres partes de su hacienda.​**",
	"**Tú ​dices ​que ​amas ​la ​lluvia, ​sin ​embargo ​usas ​un ​paraguas ​cuando ​llueve. ​Tú ​dices ​que ​amas ​el ​sol,​ ​pero ​siempre ​buscas​ ​una ​sombra ​cuando ​el ​sol ​brilla. ​Tú ​dices ​que​ ​amas ​al ​viento, ​pero​ ​cierras ​las ​ventanas ​cuando ​el ​viento ​sopla. ​Por ​eso ​es ​que​ ​tengo ​miedo ​cuando ​dices ​que ​me ​amas.​**",
	"**El ​miedo ​es ​como ​un ​fuego. ​Te ​va ​quemando ​por ​dentro. ​Si ​lo ​controlas, ​entras ​en ​calor. ​Pero ​si ​llega ​a ​dominarte, ​te ​quemará ​a ​ti ​y ​a ​todo ​lo ​que ​te ​rodea.**",
	"**El ​sapo ​sapote ​no ​sabe ​decir ​\"power\", ​el ​sapo ​sapote ​creo ​que ​sabe ​decir ​\"power\", ​el ​sapo ​sapote ​sí ​sabe ​decir ​\"power\".**",
	"**Raras ​veces ​nos ​topamos ​con ​un ​individuo ​capaz ​de ​revisar ​la ​idea ​que ​tiene ​de ​su ​propia ​inteligencia ​y ​sus ​propios ​límites ​bajo ​la ​presión ​de ​los ​acontecimientos, ​hasta ​el ​punto ​de ​abrirse ​a ​nuevas ​perspectivas ​sobre ​lo ​que ​aún ​puede ​aprender.**",
	"**Yo ​aconsejo ​a ​mis ​pacientes ​que ​evalúen ​cuidadosamente ​el ​momento ​adecuado ​para ​la ​confrontación. ​No ​se ​trata ​de ​disparar ​sin ​haber ​apuntado, ​pero ​tampoco ​de ​postergar ​indefinidamente ​la ​confrontación.**",
	"**El ​día ​de ​pascua ​nevó ​mucho, ​pero ​al ​siguiente ​sopló ​de ​improviso ​un ​viento ​cálido, ​amontonándose ​las ​nubes, ​y ​por ​espacio ​de ​tres ​días ​con ​sus ​noches ​no ​dejó ​de ​caer ​una ​lluvia ​tibia. ​El ​viento ​se ​calmó ​el ​jueves ​y ​entonces ​se ​extendió ​sobre ​la ​tierra ​una ​espesa ​bruma ​de ​color ​gris, ​como ​para ​ocultar ​los ​misterios ​que ​se ​producían ​en ​la ​naturaleza.**",
	"**Las ​personas ​con ​alta ​tolerancia ​a ​la ​frustración ​pueden ​hacer ​frente ​a ​los ​contratiempos ​con ​éxito. ​Las ​personas ​con ​baja ​tolerancia ​a ​la ​frustración ​pueden ​sentirse ​frustradas ​por ​inconvenientes ​cotidianos ​aparentemente ​menores, ​como ​atascos ​de ​tráfico ​o ​esperar ​en ​la ​cola ​del ​supermercado. ​Las ​personas ​con ​baja ​tolerancia ​a ​la ​frustración ​pueden ​renunciar ​a ​tareas ​difíciles ​de ​inmediato. ​La ​idea ​de ​tener ​que ​esperar ​en ​la ​cola ​o ​trabajar ​en ​una ​tarea ​que ​no ​entienden ​puede ​ser ​intolerable.**",
	"**Te ​sientas ​frente ​a ​un ​tablero ​y ​repentinamente ​tu ​corazón ​brinca. ​Tu ​mano ​tiembla ​al ​tomar ​una ​pieza ​y ​moverla. ​Pero ​lo ​que ​el ​ajedrez ​te ​enseña ​es ​que ​tú ​debes ​permanecer ​ahí ​con ​calma ​y ​pensar ​si ​realmente ​es ​una ​buena ​idea ​o ​si ​hay ​otras ​ideas ​mejores.**",
	"**Sabes, ​Cristo ​me ​contó ​una ​vez ​el ​secreto ​de ​la ​resurrección. ​Pero ​estábamos ​en ​una ​boda ​en ​Canaán ​y ​me ​emborraché ​y ​se ​me ​olvidó. ​¿Que ​si ​lo ​conocía? ​El ​tipo ​me ​debe ​doce ​dólares.**",
	"**Debía ​de ​tener ​unos ​diez ​u ​once ​años ​cuando ​desapareció ​su ​madre. ​Era ​una ​mujer ​alta, ​elegante, ​más ​bien ​callada, ​de ​movimientos ​lentos ​y ​precioso ​cabello ​rubio. ​A ​su ​padre ​lo ​recordaba ​de ​manera ​más ​vaga ​como ​un ​hombre ​moreno ​y ​delgado, ​vestido ​siempre ​pulcramente ​de ​negro.**",
	"**Leyendo ​novelas ​de ​detectives ​suelo ​sentir ​una ​sensación ​de ​alivio ​ante ​el ​asesinato ​de ​un ​personaje ​que ​al ​pasar ​la ​página ​no ​me ​molestará ​con ​su ​existencia. ​De ​todos ​modos, ​las ​historias ​de ​detectives ​siempre ​tienen ​demasiados ​personajes. ​Muchos ​mencionados ​al ​principio ​pero ​nunca ​vistos, ​que ​permanecen ​fuera ​de ​la ​historia, ​adquiriendo ​una ​calidad ​portentosa ​horrible. ​Es ​mejor ​que ​se ​hayan ​ido.**",
	"**La ​forma ​en ​que ​actuó ​muestra ​que ​le ​falta ​la ​madurez ​que ​es ​necesaria ​para ​este ​trabajo. ​Necesito ​a ​alguien ​que ​tenga, ​pues ​primero, ​un ​mínimo ​de ​cuidado ​para ​este ​trabajo, ​pero ​también, ​alguien ​que ​quiera ​aprender ​y ​ampliar ​sus ​conocimientos. ​Y ​por ​lo ​que ​puedo ​decir, ​él ​no ​tiene ​ninguna ​de ​esas ​cualidades. ​Entiendo ​que ​él ​necesita ​dinero ​y ​trabajo, ​pero ​yo ​no ​soy ​caridad ​para ​los ​que ​viven ​en ​los ​sótanos ​de ​sus ​padres. ​Dile ​que ​necesita ​salir ​de ​su ​trasero ​y ​trabajar.**",
	"**Me ​giré ​sobre ​mis ​pasos ​y ​a ​medida ​que ​me ​acercaba ​a ​la ​habitación ​de ​mi ​compañero ​de ​piso, ​el ​olor ​se ​hacía ​más ​espeso. ​Vacilé ​antes ​de ​abrir ​la ​puerta. ​Las ​cortinas ​flotaban ​sin ​rumbo, ​despavoridas ​y ​sobre ​su ​lecho, ​mi ​amigo ​yacía ​con ​los ​ojos ​vidriosos ​y ​el ​cuerpo ​inerte, ​sin ​vida, ​acompañado ​por ​ella.**",
	"**No ​te ​obstines, ​pues, ​en ​mantener ​como ​única ​opinión ​la ​tuya ​creyéndola ​la ​única ​razonable. ​Todos ​los ​que ​creen ​que ​ellos ​solos ​poseen ​una ​inteligencia, ​una ​elocuencia ​o ​un ​genio ​superior ​a ​los ​de ​los ​demás, ​cuando ​se ​penetra ​dentro ​muestran ​solo ​la ​desnudez ​de ​su ​alma.**",
	"**¿Por ​qué ​salimos ​corriendo ​cuando ​notamos ​la ​menor ​conexión? ​¿Por ​qué ​luchamos ​contra ​lo ​que ​más ​nos ​atrae? ​Quizá ​es ​porque ​cuando ​encontramos ​algo ​o ​alguien ​a ​quién ​aferrarnos ​lo ​ansiamos ​como ​al ​aire. ​Y ​nos ​aterra ​perderlo... ​Y ​creedme, ​uno ​aprende ​a ​estar ​solo. ​Pero ​la ​mayoría ​de ​las ​cosas ​son ​mejores ​si ​las ​compartes.**",
	"**Pues ​si ​vemos ​lo ​presente ​cómo ​en ​un ​punto ​se ​es ​ido ​y ​acabado, ​si ​juzgamos ​sabiamente, ​daremos ​lo ​no ​venido ​por ​pasado. ​No ​se ​engañe ​nadie, ​no, ​pensando ​que ​ha ​de ​durar ​lo ​que ​espera ​más ​que ​duró ​lo ​que ​vio, ​pues ​que ​todo ​ha ​de ​pasar ​por ​tal ​manera.**",
	"**Había ​una ​casa ​abajo, ​junto ​al ​estruendo ​de ​las ​olas ​desbaratándose ​contra ​los ​cantiles, ​donde ​el ​amor ​era ​más ​intenso ​porque ​tenía ​algo ​de ​naufragio.**",
	"**El ​hombre ​intentó ​mover ​la ​cabeza ​en ​vano. ​Echó ​una ​mirada ​de ​reojo ​a ​la ​empuñadura ​del ​machete, ​húmeda ​aún ​del ​sudor ​de ​su ​mano. ​Apreció ​mentalmente ​la ​extensión ​y ​la ​trayectoria ​del ​machete ​dentro ​de ​su ​vientre, ​y ​adquirió ​fría, ​matemática ​e ​inexorable, ​la ​seguridad ​de ​que ​acababa ​de ​llegar ​al ​término ​de ​su ​existencia. ​La ​muerte.**",
	"**Todos ​mis ​instintos ​son ​una ​forma, ​y ​todos ​los ​hechos ​son ​otras, ​y ​me ​temo ​mucho ​que ​los ​jurados ​británicos ​todavía ​no ​han ​alcanzado ​ese ​tono ​de ​inteligencia ​cuando ​van ​a ​dar ​la ​preferencia ​a ​mis ​teorías.**",
	"**No ​somos ​más ​que ​unos ​hombres ​elaborados ​para ​actuar ​en ​nombre ​de ​la ​venganza ​que ​consideramos ​justicia. ​Pero ​cuando ​llamamos ​a ​nuestra ​venganza ​justicia, ​solo ​engendramos ​más ​venganza, ​forjando ​el ​primer ​eslabón ​de ​las ​cadenas ​de ​odio.**",
	"**La ​oscuridad, ​la ​algarabía ​y ​sobre ​todo ​el ​acto ​ilegal ​que ​estaba ​a ​punto ​de ​cometer ​no ​formaban ​parte ​de ​mi ​vida ​cotidiana, ​sin ​embargo, ​proseguí ​y ​llegué ​a ​la ​habitación. ​Respiré ​hondamente. ​Abrí ​la ​puerta ​con ​el ​temblor ​de ​la ​mano ​como ​testigo ​y ​reuniendo ​el ​poco ​valor ​que ​me ​quedaba ​crucé ​el ​umbral ​prohibido.**",
	"**Existe ​una ​conexión ​fundamental ​entre ​lo ​que ​uno ​parece ​y ​lo ​que ​uno ​es. ​Todos ​nos ​contamos ​una ​historia ​sobre ​nosotros ​mismos. ​Siempre. ​Continuamente. ​Esa ​historia ​es ​la ​que ​nos ​convierte ​en ​lo ​que ​somos. ​Nos ​construimos ​a ​nosotros ​mismos ​a ​partir ​de ​esa ​historia.**",
	"**No ​hay ​razón ​para ​sufrir. ​La ​única ​razón ​por ​la ​que ​sufres ​es ​porque ​así ​tú ​lo ​exiges. ​Si ​observas ​tu ​vida ​encontrarás ​muchas ​excusas ​para ​sufrir, ​pero ​ninguna ​razón ​válida. ​Lo ​mismo ​es ​aplicable ​a ​la ​felicidad. ​La ​única ​razón ​por ​la ​que ​eres ​feliz ​es ​porque ​tú ​decides ​ser ​feliz. ​La ​felicidad ​es ​una ​elección, ​como ​también ​lo ​es ​el ​sufrimiento.**",
	"**Si ​aún ​respiras, ​eres ​de ​los ​que ​más ​suerte ​tienen, ​porque ​la ​mayoría ​de ​nosotros ​jadeamos ​con ​pulmones ​corruptos ​y ​nos ​desatendemos ​a ​nosotros ​mismos ​en ​nuestro ​afán ​de ​recolectar ​los ​nombres ​de ​los ​amantes ​con ​quienes ​no ​fue ​bien.**",
	"**Si ​hay ​algo ​más ​importante ​que ​mi ​ego ​en ​los ​alrededores, ​lo ​quiero ​capturado ​y ​matado ​cuanto ​antes.- ​Por ​un ​momento, ​no ​pasó ​nada. ​Luego, ​después ​de ​un ​segundo ​más ​o ​menos, ​nada ​continuó ​sucediendo.**",
	"**Mi ​madre ​me ​dice ​que ​no ​mastico ​mi ​comida ​lo ​suficiente. ​Dice ​que ​estoy ​haciendo ​que ​sea ​más ​difícil ​para ​mi ​cuerpo ​obtener ​los ​nutrientes ​esenciales ​que ​necesita. ​Si ​ella ​estuviera ​aquí, ​le ​recordaría ​que ​estoy ​comiendo ​una ​tarta ​de ​arándanos.**",
	"**El ​libro ​es ​una ​criatura ​frágil. ​Sufre ​el ​paso ​del ​tiempo, ​el ​acoso ​de ​los ​roedores ​y ​las ​manos ​torpes, ​así ​que ​el ​bibliotecario ​protege ​los ​libros ​no ​sólo ​contra ​el ​género ​humano ​sino ​también ​contra ​la ​naturaleza, ​dedicando ​su ​vida ​a ​esta ​guerra ​contra ​las ​fuerzas ​del ​olvido.**",
	"**No ​es ​de ​extrañar ​que ​nos ​pasemos ​la ​vida ​entera ​soñando. ​Estar ​despierto ​y ​verlo ​todo ​tal ​y ​como ​es ​en ​realidad... ​nadie ​podría ​soportarlo ​mucho ​tiempo.**",
	"**Despertó ​empapado ​en ​sudor, ​con ​el ​corazón ​desbocado ​y ​casi ​sin ​aire. ​Al ​abrir ​los ​ojos, ​todo ​era ​oscuridad. ​Intentó ​incorporarse ​en ​la ​cama, ​estirar ​los ​brazos, ​pero ​el ​techo ​y ​las ​paredes ​de ​madera ​se ​lo ​impedían.**",
	"**Sabía ​que ​esto ​enviudaría ​a ​tu ​madre ​y ​te ​dejaría ​huérfano. ​Pero ​había ​encontrado ​mi ​destino. ​Así ​que ​abandoné ​a ​mi ​hijo. ​Ahora ​tengo ​trabajo ​que ​hacer, ​una ​cantidad ​infinita. ​Debo ​encontrar ​vida ​inteligente.**",
	"**La ​gente ​que ​más ​me ​gusta ​es ​la ​que ​ha ​fallado, ​ha ​sido ​lastimada, ​ha ​llorado, ​ha ​visto ​cosas ​terribles, ​y ​sin ​embargo, ​no ​ha ​perdido ​su ​capacidad ​para ​seguir ​amando.**",
	"**Generalmente, ​esa ​sensación ​de ​estar ​solo ​en ​el ​mundo ​aparece ​mezclada ​con ​un ​orgulloso ​sentimiento ​de ​superioridad: ​desprecio ​a ​los ​hombres, ​los ​veo ​sucios, ​feos, ​incapaces, ​ávidos, ​groseros, ​mezquinos; ​mi ​soledad ​no ​me ​asusta, ​es ​casi ​olímpica.**",
	"**Un ​ejemplo ​de ​conjunto ​de ​palabras ​aleatorias ​es: ​como ​pasa ​mismo ​nos ​oh ​creo ​mejor ​nada ​tenemos ​papas ​qué ​sea ​acuerdo ​ellos ​un ​vas ​tiene ​fue ​nadie ​estamos ​va ​parece ​mi ​tus ​sin ​ahí ​mira ​alguien**",

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
