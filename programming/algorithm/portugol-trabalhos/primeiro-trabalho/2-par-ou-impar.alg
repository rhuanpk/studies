Algoritmo "par_ou_impar"
// Disciplina  : L�gica e Fundamentos da Programa��o
// Professor   : Thiago Momesso de Menezes 
// Descri��o   : Diz se o n�mero digitado � "par" ou "�mpar"
// Autor(a)    : Rhuan Patriky de Sousa Ferrer - RA: 1028
// Data atual  : 19/03/2021
Var
// Se��o de Declara��es das vari�veis 
   nome: caractere
   numero, x: real

  //variavel "x" usada para fazer o calculo do "mod"

Inicio
// Se��o de Comandos, procedimento, fun��es, operadores, etc... 
   mudacor("Amarelo", "FRENTE")
   
   escreval("           *** Par ou �mpar ***")
   escreval( )
   escreva("Digite seu nome: ")
   leia(nome)
   limpatela( )
   escreval("           *** Par ou �mpar ***")
   escreval( )
   escreval("- Bem-vindo > ", maiusc(nome), " <")
   escreval("- Entre com um n�mero positivo")
   escreval("- Lhe retornaremos se o n�mero � par ou �mpar")
   escreval( )
   escreva("Digite o n�mero: ")
   leia(numero)
   escreval( )
   
   x <- numero mod 2
   se((x = 0) e (numero >= 0))entao
      escreval("O n�mero escolhido � -> 'PAR'")
   senao
      se((x <> 0) e (numero >= 0))entao
         escreval("O n�mero escolhido � -> '�MPAR'")
      fimse
   fimse
   se(numero < 0)entao
      escreval("N�mero inv�lido para testes!")
   fimse

Fimalgoritmo