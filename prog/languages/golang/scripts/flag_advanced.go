package main

import (
	"flag"
	"fmt"
)

// Para criarmos _flags_ de CLI e obter seus arugmentos, podemos fazer de duas formas:
//  1. Passando o ponteiro da nossa variável pré-definida para as funções `flag.<type>Var()`
//  1. Inicilizar uma nova variável diretamente pelas funções `flag.<type>()`
//
// Sobre _flags_ do tipo [bool]:
//   - Se NÃO passar a _flag_ na CLI: O valor da variável será o _default_ passado na função (seja [true] ou [false])
//   - Se passar a _flag_ na CLI: O valor da variável será sempre [true]
//   - Se passar a _flag_ na CLI com argumento: O valor da variável será o argumento passado para _flag_ com `=`
//
// Sobre _flags_ de outros tipos:
//   - Se NÃO passar a _flag_ na CLI: O valor da variável será o _default_ passado na função (seja o valor zero do tipo o qualquer outro)
//   - Se passar a _flag_ na CLI com argumento: O valor da variável será o argumento passado para _flag_ com `=` ou espaço em branco
//
// Funções do pacte [flag]:
//   - [flag.Set]: Define/Troca o valor de uma _flag_ (e consequentemente também da variável) de forma manual
//   - [flag.Lookup]: Pega uma _flag_ ([*flag.Flag], que contem seu nome, valor, _default_ e descrição)
//   - [flag.Visit]: Itera sobre todas as _flags_ ([*flag.Flag]) que foram passadas na CLI
//   - [flag.VisitAll]: Itera sobre todas as _flags_ ([*flag.Flag]) definidas (incluindo as não utilizadas na CLI)
//
// Exemplos:
//
//	// build
//	go build -o app ./
//
//	// tipo bool
//	app                // sem flags
//	app -b             // com flag
//	app -b=true        // com flag e argumento
//
//	// outros tipos
//	app                // sem flags
//	app -a=42          // com flag e argumento
//	app -a 42          // com flag e argumento
//
// Go playground: https://go.dev/play/p/QEYUB0rRTYB
func main() {
	// declare the vars and bind the flags after
	var (
		boolFlag bool
		intFlag  int
	)
	flag.BoolVar(&boolFlag, "b", false, "usage")
	flag.IntVar(&intFlag, "i", 42, "usage")

	// create the vars already binding the flags
	stringFlag := flag.String("s", "xpto", "usage")
	floatFlag := flag.Float64("f", 0.0, "usage")

	// parse the cli flags
	flag.Parse()

	// get the values from vars and directly from flags too
	println("from cli\n")
	printVars(boolFlag, intFlag, *stringFlag, *floatFlag)
	printLookups("b", "i", "s", "f")

	// manually change the value of flags (consequently of vars too)
	flag.Set("i", "-1")
	flag.Set("s", "hello")

	// get the values from vars and from visit function
	println("\nafter set\n")
	printVars(boolFlag, intFlag, *stringFlag, *floatFlag)
	flag.VisitAll(func(f *flag.Flag) {
		println("visit", f.Name)
		fmt.Printf("- f.Name: %#v\n", f.Name)
		fmt.Printf("- f.Value.String(): %#v\n", f.Value.String())
		// fmt.Printf("- f.DefValue: %#v\n", f.DefValue)
		// fmt.Printf("- f.Usage: %#v\n", f.Usage)
	})
}

func printVars(boolFlag bool, intFlag int, stringFlag string, floatFlag float64) {
	println("variables")
	fmt.Printf("- boolFlag: %#v\n", boolFlag)
	fmt.Printf("- intFlag: %#v\n", intFlag)
	fmt.Printf("- stringFlag: %#v\n", stringFlag)
	fmt.Printf("- floatFlag: %#v\n", floatFlag)
}

func printLookups(flags ...string) {
	for _, f := range flags {
		println("lookup", f)
		lookup := flag.Lookup(f)
		fmt.Printf("- lookup.Name: %#v\n", lookup.Name)
		fmt.Printf("- lookup.Value.String(): %#v\n", lookup.Value.String())
		// fmt.Printf("- lookup.DefValue: %#v\n", lookup.DefValue)
		// fmt.Printf("- lookup.Usage: %#v\n", lookup.Usage)
	}
}
