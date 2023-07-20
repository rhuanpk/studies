package main

import (
	typeexample "predefinedtype/type_example"
)

func main() {
	// como a função typeexample.AnyFunction() só aceita dados do tipo typeexample.newType e esse tipo é privado, os únicos aceitos nessa função são as constantes que foram criadas dentro do seu pacote e que são exportáveis.
	typeexample.AnyFunction(typeexample.FirstConst)
}
