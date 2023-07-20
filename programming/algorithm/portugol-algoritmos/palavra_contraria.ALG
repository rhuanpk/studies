Algoritmo "palavra_contraria"

Var
   p: caractere
   x: inteiro

Inicio
   escreval("  */* Palavra ao contrario */*")
   escreval( )
   escreva("Palavra: ")
   leia(p)
   escreva("Contrario: ")
   para x de compr(p) ate 1 passo -1 faca
      escreva(copia(p, x, 1))
   fimpara
   escreval( )

Fimalgoritmo