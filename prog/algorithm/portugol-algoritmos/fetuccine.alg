Algoritmo "fetuccine"

Var
   a, b, soma: inteiro
   x: inteiro
   
Inicio
   escreva("Pirmeiro número: ")
   leia(b)
   escreva("Segundo número: ")
   leia(a)
   limpatela( )
   escreval("       Série de Fetuccine!")
   escreval( )
   escreval(" 1)", b)
   escreval(" 2)", a)
   para x de 3 ate 10 faca
      escreva(x, ")")
      se x mod 2 = 1 entao
         soma <- a + b
         escreval(soma)
         b <- a
         a <- soma
      senao
         se x mod 2 = 0 entao
            soma <- a - b
            escreval(soma)
            b <- a
            a <- soma
         fimse
      fimse
   fimpara

Fimalgoritmo