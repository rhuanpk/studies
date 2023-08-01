Algoritmo "decimal_binario"

Var
   vet: vetor[1..10] de inteiro
   x: real
   y, z, c: inteiro

Inicio
   c <- 0
   enquanto c <> 1 faca
      escreval(" */* Converta Decimal em Binário */*")
      escreval( )
      escreva("Decimal: ")
      leia(x)
      y <- 1
      repita
         vet[y] <- x mod 2
         x <- int(x / 2)
         y <- y + 1
      ate (x = 0) ou (x = 1)
      escreva("Binário:", x)
      para z de (y -1) ate 1 passo -1 faca
         escreva(vet[z])
      fimpara
      escreval( )
      escreval( )
      escreval("sair <1>")
      escreva("novamente <0>")
      leia(c)
      limpatela( )
   fimenquanto

Fimalgoritmo