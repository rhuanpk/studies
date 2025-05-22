package main

import (
	"fmt"
	"time"
)

// channelPopulate popula o canal recebido como parâmetro percorrendo um for estático 10x
func channelPopulate(channel chan int) {
	// percorre de 0 a 9 colocando o valor atual da iteração em `index`
	for index := 0; index < 10; index++ {
	
		// printa um log do que está acontecendo e atribui o valor da variável `index` na *stack* no canal
		fmt.Printf("write %d on channel!\n", index)
		channel <- index
	}

	// fecha o canal
	close(channel)
}

// go playground: https://go.dev/play/p/8Pioj5Gbeyw
func main() {
	// cria um novo canal que receberá tipo de dado inteiro e terá *buffer* de 3 dados por vez
	channel := make(chan int, 3)

	// chama o função `channelPopulate()` colocanod ela em paralelo com o restante do código (executa em *backround*)
	go channelPopulate(channel)

	// aguarda 3 segundo (para que encha a pilha do canal) e inicia um que itera sobre os valores do canal
	time.Sleep(time.Second * 3)
	for value := range channel {

		// printa um log do que está acontecendo e dorme 1 segundo.
		fmt.Printf("read %d from channel!\n", value)
		time.Sleep(time.Second * 1)
	
	}
}
