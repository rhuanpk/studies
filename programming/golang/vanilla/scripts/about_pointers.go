package main

import "fmt"

// Quando trabalhamos com ponteiros, estamos trablalhando com endereços de
// memória e não com o valor diretamente, ou seja, estamos manipulando o
// endereço de um espaço na memória que está reservado para guardar algum valor.

// go playground: https://go.dev/play/p/TD77XmpcxxA
func main() {
	println("-----> first example <-----")
	// 1. declaramos `foo` como um tipo ponteiro de string, ou seja, `foo`
	// poderá receber apenas endereços para espaços de memória destinados a
	// guardar valores do tipo string
	var foo *string
	fmt.Printf("\n1 step -> var foo *string\n\t- foo (ptr): %v; *foo (val): <panic> [segfault]\n", foo)

	// 2. reservamos (declaramos) um espaço na memória (variável) chamado `aux`
	// destinado a guardar valores do tipo string ao mesmo tempo que atribuímos
	// (inicializamos) o valor "hello" nesse espaço
	aux := "hello"
	fmt.Printf("\n2 step -> aux := \"hello\"\n\t- &aux (ptr): %v; aux (val): %#v\n", &aux, aux)

	// 3. atribuímos o ponteiro da variável `aux` para o ponteiro `foo`, ou seja
	// , agora o ponteiro `foo` aponta para o mesmo endereço de memória da
	// variável `aux` (&aux)
	foo = &aux
	fmt.Printf("\n3 step -> foo = &aux\n\t- foo (ptr): %v; *foo (val): %#v\n\t- &aux (ptr): %v; aux (val): %#v\n", foo, *foo, &aux, aux)

	// 4. agora, independendo se alterarmos o espeço na memória de `*foo` quanto
	// de `aux`, o valor se refletirá em ambos pois, ambos (`foo` e `&bar`)
	// apontam para o mesmo endereço na memória
	aux = "world"
	fmt.Printf("\n4 step -> aux = \"world\"\n\t- foo (ptr): %v; *foo (val): %#v\n\t- &aux (ptr): %v; aux (val): %#v\n", foo, *foo, &aux, aux)
	*foo = "xpto"
	fmt.Printf("\n4 step -> *foo = \"xpto\"\n\t- foo (ptr): %v; *foo (val): %#v\n\t- &aux (ptr): %v; aux (val): %#v\n", foo, *foo, &aux, aux)

	println("\n-----> second example <-----")
	// 1. `new()` irá alocar espaço na memória com o valor zero do tipo passado
	// e retornar o endereço desse espaço alocado
	var bar *string = new(string)
	fmt.Printf("\n1 step -> var bar *string = new(string)\n\t- bar (ptr): %v; *bar (val): %#v\n", bar, *bar)

	// 2. atribuímos o valor "xpto" no espaço da memória do ponteiro `bar`
	*bar = "xpto"
	fmt.Printf("\n2 step -> *bar = \"xpto\"\n\t- bar (ptr): %v; *bar (val): %#v\n", bar, *bar)
}
