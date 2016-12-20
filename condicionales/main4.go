package main

import "fmt"

//el marcador de posicion para mostrar un tipo de dato booleano es %t
func main() {
	if noembarre, disculpas := true, 1; noembarre && disculpas == 4 {
		fmt.Println("no rogar mas dejar que ella te busque")
	} else if noembarre && disculpas == 3 {
		fmt.Println("se siente que se acaba la paciencia pero la quieres mucho sigues intentando")
	} else if noembarre && disculpas == 2 {
		fmt.Printf("llevas %d  disculpas, aun debes intentarlo mas veces y estas desesperado", disculpas)
	} else if noembarre && disculpas == 1 {
		fmt.Printf("apenas llavas %d disculpa, y es %t que no lo embarraste debes hacer hasta la imposible para que ella te crea por que la amas", disculpas, noembarre)
	} else if noembarre && disculpas == 0 {
		fmt.Println("te amo mi jessi y no me imagino estar sin ti, quiero que estemos siempre juntos")
	} else {
		fmt.Println("la embarraste y los mas logico es que terminen por que si te pasar a ty no lo perdonarias")
	}

}
