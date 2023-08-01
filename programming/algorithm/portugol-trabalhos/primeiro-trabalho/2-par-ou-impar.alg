Algoritmo "par_ou_impar"
// Disciplina  : Lógica e Fundamentos da Programação
// Professor   : Thiago Momesso de Menezes 
// Descrição   : Diz se o número digitado é "par" ou "ímpar"
// Autor(a)    : Rhuan Patriky de Sousa Ferrer - RA: 1028
// Data atual  : 19/03/2021
Var
// Seção de Declarações das variáveis 
   nome: caractere
   numero, x: real

  //variavel "x" usada para fazer o calculo do "mod"

Inicio
// Seção de Comandos, procedimento, funções, operadores, etc... 
   mudacor("Amarelo", "FRENTE")
   
   escreval("           *** Par ou Ímpar ***")
   escreval( )
   escreva("Digite seu nome: ")
   leia(nome)
   limpatela( )
   escreval("           *** Par ou Ímpar ***")
   escreval( )
   escreval("- Bem-vindo > ", maiusc(nome), " <")
   escreval("- Entre com um número positivo")
   escreval("- Lhe retornaremos se o número é par ou ímpar")
   escreval( )
   escreva("Digite o número: ")
   leia(numero)
   escreval( )
   
   x <- numero mod 2
   se((x = 0) e (numero >= 0))entao
      escreval("O número escolhido é -> 'PAR'")
   senao
      se((x <> 0) e (numero >= 0))entao
         escreval("O número escolhido é -> 'ÍMPAR'")
      fimse
   fimse
   se(numero < 0)entao
      escreval("Número inválido para testes!")
   fimse

Fimalgoritmo