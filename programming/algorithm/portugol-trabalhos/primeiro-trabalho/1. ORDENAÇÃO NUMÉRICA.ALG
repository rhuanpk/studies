Algoritmo "ordenação_numerica"
// Disciplina  : Lógica e Fundamentos da Programação
// Professor   : Thiago Momesso de Menezes
// Descrição   : Ordena os números em ordem crescente
// Autor(a)    : Rhuan Patriky de Sousa Ferrer - RA: 1028
// Data atual  : 24/03/2021
Var
// Seção de Declarações das variáveis
   nome: caractere
   n1, n2, n3, aux: inteiro

Inicio
// Seção de Comandos, procedimento, funções, operadores, etc...
   mudacor("Amarelo", "FRENTE")
   
   escreval("   === Olá Usuário, digite seu nome! ====")
   escreval( )
   escreva("Nome: ")
   leia(nome)
   limpatela( )
   escreval("   *** Ordem Crescente dos Números ***")
   escreval(" ")
   escreval("- Bem-vindo ", maiusc(nome))
   escreval("- Entre com três números separando-os por 'ENTER'")
   escreval("- Lhe retornaremos os mesmos números em ordem crescente!")
   escreval( )
   escreva("Número 1: ")
   leia(n1)
   escreva("Número 2: ")
   leia(n2)
   escreva("Número 3: ")
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