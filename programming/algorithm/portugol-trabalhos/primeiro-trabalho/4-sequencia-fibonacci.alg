Algoritmo "sequencia_fibonacci"
// Disciplina  : L�gica e Fundamentos da Programa��o
// Professor   : Thiago Momesso de Menezes 
// Descri��o   : Mostra a sequ�ncia de Fibonacci at� aonde o usu�rio desejar
// Autor(a)    : Rhuan Patriky de Sousa Ferrer - RA: 1028
// Data atual  : 22/03/2021
Var
// Se��o de Declara��es das vari�veis
   nome: caractere

   x, y, resultado: inteiro
   // Vari�veis usadas para realiza��o da conta

   numero, controle: inteiro
   // Vari�vel "controle": - Utilizada para encargo da fun��o "para"
   //                      e usada tambem como contador
   
Inicio
// Se��o de Comandos, procedimento, fun��es, operadores, etc... 
   mudacor("Amarelo", "FRENTE")

   x <- 1
   y <- -1
   escreval("         Bem-vindo Usu�rio!")
   escreval( )
   escreva("Digite seu nome: ")
   leia(nome)
   limpatela( )
   escreval("     *** Sequ�ncia de Fibonacci ***")
   escreval( )
   escreval("- Sauda��es '", maiusc(nome), "', digite um n�mero")
   escreval("- lhe retornameros a sucess�o de Fibonacci at� tal;")
   escreval( )
   escreva("N�mero: ")
   leia(numero)
   escreval( )
   
   para controle de 1 ate numero faca
        resultado <- x + y
        y <- x
        x <- resultado
        escreval(controle, "�) ", resultado)
   fimpara
   
   escreval(" ")
   escreval("Obrigado, at� a proxima!")
   escreval(" ")
   escreval("'Esc' para sair!")
   
Fimalgoritmo