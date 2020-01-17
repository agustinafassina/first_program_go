package main

//libreria para poder imprimir en pantalla.
//para utilizar una variable go el compilador pide que la declares, para no consumir memoria
//1. no gastar recursos en algo que no se va a utilizar
//2. es parte del linter que declares variables que vas a usar, no variables que no usaras.

import "fmt"
//forma para formatear texto de manera correcta
//%s llama a las variables y le asigna el valor

func main() {
	var name string

	fmt.Print("Ingresa tu nombre: ")
	
	fmt.Scanf("%s", &name)
	
	//formatear la cadena de texto que quiero imprimir
	fmt.Printf("Helo world with Go %s, sii.\n", name)
	
	//Println salto de linea
	fmt.Println("Hello world")


}
