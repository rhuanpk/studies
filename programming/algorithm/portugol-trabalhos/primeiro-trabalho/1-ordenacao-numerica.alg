Algoritmo "ordena��o_numerica"
// Disciplina  : L�gica e Fundamentos da Programa��o
// Professor   : Thiago Momesso de Menezes
// Descri��o   : Ordena os n�meros em ordem crescente
// Autor(a)    : Rhuan Patriky de Sousa Ferrer - RA: 1028
// Data atual  : 24/03/2021
Var
// Se��o de Declara��es das vari�veis
   nome: caractere
   n1, n2, n3, aux: inteiro

Inicio
// Se��o de Comandos, procedimento, fun��es, operadores, etc...
   mudacor("Amarelo", "FRENTE")
   
   escreval("   === Ol� Usu�rio, digite seu nome! ====")
   escreval( )
   escreva("Nome: ")
   leia(nome)
   limpatela( )
   escreval("   *** Ordem Crescente dos N�meros ***")
   escreval(" ")
   escreval("- Bem-vindo ", maiusc(nome))
   escreval("- Entre com tr�s n�meros separando-os por 'ENTER'")
   escreval("- Lhe retornaremos os mesmos n�meros em ordem crescente!")
   escreval( )
   escreva("N�mero 1: ")
   leia(n1)
   escreva("N�mero 2: ")
   leia(n2)
   escreva("N�mero 3: ")
   leia(n3)
   escreval( )

   se (n1 > n2) entao
      aux <- n1
      n1 <- n2
      n2 <- aux
   fimse
   se (n2 > n3) entao
      aux <- n2
      n2 <- n3
      n3 <- aux
      se (n1 > n2) entao
         aux <- n1
         n1 <- n2
         n2 <- aux
      fimse
   fimse
   
   escreval("Em ordem crescente:", n1, ",", n2, ",", n3)

Fimalgoritmo