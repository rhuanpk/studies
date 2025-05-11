Algoritmo "fibonacci"

Var
   ultimo, penultimo, valoratual, y: inteiro
   x: inteiro

Inicio
   escreva("Escolha o enésimo número: ")
   ultimo <- 1
   penultimo <- -1
   leia(x)
   escreval( )
   
   para y de 0 ate x faca
        valoratual <- ultimo + penultimo
        penultimo <- ultimo
        ultimo <- valoratual
        escreval(valoratual)
   fimpara

Fimalgoritmo