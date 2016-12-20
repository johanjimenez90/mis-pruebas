package main

import "fmt"

func main() {
	aprender := false
	if aprender {
		fmt.Println(aprender, "que el conocimiento es una de las livertades mas lindas del mundo ")
		// para ver que tipo de dato es una variable utilizamos fmt.Printf ("aprender es de tipo %T \n",aprender)
		fmt.Printf("aprender es de tipo %T \n", aprender)
		//ojo el else  debe ir siempre adelante del la llave donde termina el if.  } else
	} else {
		fmt.Println(aprender, "cuando no aprendemos somos unos hermitaños")
	}
	fmt.Println("fin del programa")
}
