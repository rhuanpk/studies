Algoritmo "diagonal_principal_matriz"

Var
   m1: vetor[0..3,0..3] de inteiro
   m2: vetor[0..3] de inteiro
   flag: logico
   x, y: inteiro
   c1, c2: inteiro

Inicio
   c1 <- 0
   c2 <- 3

   escreval("Matriz")
   para x de 0 ate 4 faca
      para y de 0 ate 3 faca
         se x < 4 entao
            flag <- falso
            m1[x,y] <- randi(10)
            escreva(m1[x,y])
            se (x = c1) e (y = c2) entao
               m2[c1] <- m1[x,y]
               flag <- verdadeiro
            fimse
            se flag entao
               c1 <- c1 + 1
               c2 <- c2 - 1
            fimse
         senao
            se y = 0 entao
               escreval( )
               escreval("Diagonal principal")
            fimse
            escreva(m2[y])
         fimse
      fimpara
      escreval( )
   fimpara

Fimalgoritmo