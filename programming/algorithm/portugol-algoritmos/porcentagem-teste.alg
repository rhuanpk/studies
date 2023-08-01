Algoritmo "porcentagem_teste"

Var
   y: real
   z: inteiro
   count: inteiro
   x1, x2, x3, x4, x5: real
   p1, p2, p3, p4, p5: real
   
Inicio
   count <- 500
   para z de 1 ate count faca
      y <- randi(3) + randi(3)
      escreval(y)
      se y = 0 entao
         x1 <- x1 + 1
      fimse
      se y = 1 entao
         x2 <- x2 + 1
      fimse
      se y = 2 entao
         x3 <- x3 + 1
      fimse
      se y = 3 entao
         x4 <- x4 + 1
      fimse
      se y = 4 entao
         x5 <- x5 + 1
      fimse
   fimpara
   
   p1 <- (x1 * 100) / count
   p2 <- (x2 * 100) / count
   p3 <- (x3 * 100) / count
   p4 <- (x4 * 100) / count
   p5 <- (x5 * 100) / count
   
   escreval( )
   escreval("0)", x1, "  = ", p1, "%")
   escreval("1)", x2, "  = ", p2, "%")
   escreval("2)", x3, "  = ", p3, "%")
   escreval("3)", x4, "  = ", p4, "%")
   escreval("4)", x5, "  = ", p5, "%")
   
Fimalgoritmo
