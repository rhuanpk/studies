Algoritmo "13_salario"

Var
   x, s, m: real
   p1, p2: real

Inicio
   repita
      escreval("    ** Programa de cálculo de 13 Salário! **")
      escreval( )
      escreva("Salário: ")
      leia(s)
      escreva("Meses trabalhados: ")
      leia(m)
      p1 <- s / 12 * m / 2
      p2 <- p1 - (p1 * 0.075)
      escreval( )
      escreval("1°)Primeira Parcela.. = R$ ", p1:5:2)
      escreval("2°)Segunda Parcela... = R$ ", p2:5:2)
      escreval( )
      escreval("-> Décimo terceiro bruto..... = R$ ", s / 12 * m:5:2)
      escreval( )
      escreval("-> INSS proporcional......... = R$ ", p1 * 0.075:5:2)
      escreval( )
      escreval("-> Décimo terceiro líquido... = R$ ", p1 + p2:5:2)
      escreval( )
      escreval("Novamente 0")
      escreval("Encerrar 1")
      escreval( )
      leia(x)
      limpatela( )
   ate x = 1

Fimalgoritmo