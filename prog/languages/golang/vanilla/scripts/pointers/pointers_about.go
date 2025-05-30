package main

import "fmt"

// Quando trabalhamos com ponteiros, estamos trablalhando com endereços de
// memória e não com o valor diretamente, ou seja, estamos manipulando o
// endereço de um espaço na memória que está reservado para guardar algum valor.

// Devemos sempre lembrar que o que gera _segfault_ é a desreferência de um
// ponteiro não inicializado (que é `nil`) e não algum outro tipo de dado `nil`,
// por exemplo, num _for-range_, Go consegue lidar tranquilamente com uma
// _slice_ `nil` (basicamente não haverá nenhuma iteração).

// go playground: https://go.dev/play/p/8eLOPap_D-S
func main() {
	println("-----> first example <-----")
	// 1. declaramos `foo` como um tipo ponteiro de string, ou seja, `foo`
	// poderá receber apenas endereços para espaços de memória destinados a
	// guardar valores do tipo string
	var foo *string
	fmt.Printf("\n1 step -> var foo *string\n\t- foo (val): %v; *foo (drf): <segfault>\n", foo)

	// 2. reservamos (declaramos) um espaço na memória (variável) chamado `aux`
	// destinado a guardar valores do tipo string ao mesmo tempo que atribuímos
	// (inicializamos) o valor "hello" nesse espaço
	aux := "hello"
	fmt.Printf("\n2 step -> aux := \"hello\"\n\t- &aux (ptr): %v; aux (val): %#v\n", &aux, aux)

	// 3. atribuímos o endereço da variável `aux` para o ponteiro `foo`, ou seja
	// , agora o ponteiro `foo` aponta para o mesmo endereço de memória da
	// variável `aux` (&aux)
	foo = &aux
	fmt.Printf("\n3 step -> foo = &aux\n\t- foo (val): %v; *foo (drf): %#v\n\t- &aux (ptr): %v; aux (val): %#v\n", foo, *foo, &aux, aux)

	// 4. agora, independente se alterarmos o espeço na memória de `*foo` quanto
	// de `aux`, o valor se refletirá em ambos pois, ambos apontam para o mesmo
	// endereço na memória
	aux = "world"
	fmt.Printf("\n4 step -> aux = \"world\"\n\t- foo (val): %v; *foo (drf): %#v\n\t- &aux (ptr): %v; aux (val): %#v\n", foo, *foo, &aux, aux)
	*foo = "xpto"
	fmt.Printf("\n4 step -> *foo = \"xpto\"\n\t- foo (val): %v; *foo (drf): %#v\n\t- &aux (ptr): %v; aux (val): %#v\n", foo, *foo, &aux, aux)

	println("\n-----> second example <-----")
	// 1. `new()` irá alocar espaço na memória com o valor zero do tipo passado
	// e retornar o endereço desse espaço alocado
	var bar *string = new(string)
	fmt.Printf("\n1 step -> var bar *string = new(string)\n\t- bar (val): %v; *bar (drf): %#v\n", bar, *bar)

	// 2. atribuímos o valor "xpto" no espaço da memória do ponteiro `bar`
	*bar = "xpto"
	fmt.Printf("\n2 step -> *bar = \"xpto\"\n\t- bar (val): %v; *bar (drf): %#v\n", bar, *bar)

	println("\n-----> third example <-----")
	// 1. não devemos esquecer também que variáveis do tipo ponteiro são
	// variáveis como quaisquer outras, ou seja, quando criamos uma variável do
	// tipo ponteiro, simplesmente estamos reservando um espaço na memória para
	// guardar valores de endereço de memória (de outros espaços na memória, que
	// por sua vez, estão reservados para alocar valor de algum tipo de dado),
	// ou seja, como variáveis do tipo ponteiros também são variáveis "comuns",
	// elas também tem seu próprio endereço na memória
	var boo *any
	fmt.Printf("\n1 step -> var boo *any\n\t- &boo (ptr): %v; boo (val): %#v; *boo (drf): <segfault>\n", &boo, boo)

	// 2. se tentamos atribuir valores diretamente a desreferência de ponteiros
	// sem antes o inicializarmos, pegamos um erro de "invalid memory address or
	// nil pointer dereference" (segfault) pois, se ainda não o inicializamos,
	// o valor do ponteiro (que é um endereço de memória) está nil, ou seja,
	// como não há valor algum, não está apontando para lugar nenhum na memória,
	// então quando desreferênciamos o ponteiro (ou seja, acessamos o espaço na memória
	// no qual o ponteiro aponta), ele não existe, então como podemos atribuir
	// valor num espaço da memória que não existe?
	var baz *any = new(any)
	fmt.Printf("\n2 step -> var baz *any = new(any)\n\t- &baz (ptr): %v; baz (val): %#v; *baz (drf): %#v\n", &baz, baz, *baz)
}
