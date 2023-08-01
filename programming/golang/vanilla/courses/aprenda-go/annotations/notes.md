# Aprenda Go

## Declarações

Tipo:

```go
type newtypename int
```

## Statements

### `switch`

```go
package main

var x interface{}

func main() {

	x = func() {}

	switch x.(type) {
	case string:
		println("is a string")
	case int:
		println("is a int")
	case float64:
		println("is a float")
	case bool:
		println("is a bool")
	default:
		println("is a composite type")
	}

}
```

## Builtins

### `fmt`

Pacote usado para trabalhar com impressão de *input* (qualquer tipo de dado de entrada). Básicamente há 3 tipos de *Prints* e cada um tem 3 variáções que são equivalentes.

#### Tipos

`Print`:

Imprime na *stdout* todos argumentos passados separados por *blank space*.

`Sprint`:

Ao invés de imprimir o *input* retorna uma `string` com a formatação que já esperada do `Print` clássico (não separa os argumentos com *blank space*).

`Fprint`:

Imprime em qualquer lugar que implemente um "escritor de arquivo" (*file*, *writer interface*).

#### Variações

*defaults*:

Executam **sem** *new line* por padrão.

`ln`:

Executam **com** *new line* por padrão.

`f`:

Executam com formatação da *string* (*like C printf*).

- `%v`: imprime qualquer tipo de variável.
- `%T`: imprime o tipo da variável.
- `%s`: imprime variáveis `string`.
- `%d`: imprime variáveis `int`.
- `%f`: imprime variáveis `float`.
- `%b`: imprime variáveis `bool`.
- `%9.2f`: comprimento 9, precisão 2.

## Observations

### slices

Quando você tem situações aonde precisar manipular duas *slices* que proveem do mesmo *array*, as duas *slices* manipularam esse *array*, *e.g.*:

```go
package main

import "fmt"

func main() {

	primeiroslice := []int{0, 1, 2, 3, 4}
	fmt.Println(primeiroslice)
	segundoslice := append(primeiroslice[:2], primeiroslice[4:]...)
	fmt.Println(segundoslice)
	fmt.Println(primeiroslice)

}
```

## Links

Gophers:

- https://github.com/ashleymcnamara/gophers
