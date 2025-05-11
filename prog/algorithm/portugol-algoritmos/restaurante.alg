Algoritmo "restaurante"

Var
   prato, bebida, sobre: inteiro
   c1, c2, c3: caractere
   ponto: caractere
   x: inteiro

Inicio
   escreval("  Bem-vindo ao Restaurante!")
   escreval( )
   escreval("- Escolha as opções, a seguir;")
   escreval("- Retornaremos as calorias totais")
   escreval( )
   escreva("'Enter' Prosseguir!")
   leia(ponto)
   
   limpatela( )
   escreval("/*/ Bem-vindo ao Restaurante /*/")
   escreval( )
   escreval("   ************************")
   escreval("   *        Pratos        *")
   escreval("   ************************")
   escreval( )
   timer 500
   escreval("-> 1 Vegetariano...180cal.")
   escreval("-> 2 Peixe.........230cal.")
   escreval("-> 3 Frango........250cal.")
   escreval("-> 4 Carne.........350cal.")
   timer 0
   escreval( )
   escreva("Escolha: ")
   leia(prato)
   escolha prato
      caso 1
         x <- x + 180
         c1 <- "Vegetariano 180cal."
      caso 2
         x <- x + 230
         c1 <- "Peixe 230cal."
      caso 3
         x <- x + 250
         c1 <- "Frango 250cal."
      caso 4
         x <- x + 350
         c1 <- "Carne 350cal."
   fimescolha
   
   limpatela( )
   escreval("/*/ Bem-vindo ao Restaurante /*/")
   escreval( )
   escreval("   ************************")
   escreval("   *       Sobremesa      *")
   escreval("   ************************")
   escreval( )
   timer 500
   escreval("-> 1 Abacaxi................75cal.")
   escreval("-> 2 Sorvete Diet..........110cal.")
   escreval("-> 3 Mousse Diet...........170cal.")
   escreval("-> 4 Mousse de Chocolate...200cal.")
   timer 0
   escreval( )
   escreva("Escolha: ")
   leia(sobre)
   escolha sobre
      caso 1
         x <- x + 75
         c2 <- "Abacaxi 75cal."
      caso 2
         x <- x + 110
         c2 <- "Sorvete Diet 110cal."
      caso 3
         x <- x + 170
         c2 <- "Mousse Diet 170cal."
      caso 4
         x <- x + 200
         c2 <- "Mousse de Chocolate 200cal."
   fimescolha
   
   limpatela( )
   escreval("/*/ Bem-vindo ao Restaurante /*/")
   escreval( )
   escreval("   ************************")
   escreval("   *        Bebidas       *")
   escreval("   ************************")
   escreval( )
   timer 500
   escreval("-> 1 Chá.................20cal.")
   escreval("-> 2 Suco de Laranja.....70cal.")
   escreval("-> 3 Suco de Melão......100cal.")
   escreval("-> 4 Refrigerante Diet...65cal.")
   timer 0
   escreval( )
   escreva("Escolha: ")
   leia(bebida)
   escolha bebida
      caso 1
         x <- x + 20
         c3 <- "Chá 20cal."
      caso 2
         x <- x + 70
         c3 <- "Suco de Laranja 70cal."
      caso 3
         x <- x + 100
         c3 <- "Suco de Melão 100cal."
      caso 4
         x <- x + 65
         c3 <- "Refrigerante Diet 65cal."
   fimescolha
   
   limpatela( )
   escreval("/*/ Bem-vindo ao Restaurante /*/")
   escreval( )
   escreval("- Pratos escolhidos;")
   escreval( )
   timer 250
   escreval("Prato.......: ", maiusc(c1))
   escreval("Sobremesa...: ", maiusc(c2))
   escreval("Bebida......: ", maiusc(c3))
   escreval( )
   escreval("Calorias totáis: ", x, " calorias!")
   timer 0
   
Fimalgoritmo