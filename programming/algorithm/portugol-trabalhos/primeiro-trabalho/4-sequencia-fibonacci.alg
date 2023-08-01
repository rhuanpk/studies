Algoritmo "sequencia_fibonacci"
// Disciplina  : Lógica e Fundamentos da Programação
// Professor   : Thiago Momesso de Menezes 
// Descrição   : Mostra a sequência de Fibonacci até aonde o usuário desejar
// Autor(a)    : Rhuan Patriky de Sousa Ferrer - RA: 1028
// Data atual  : 22/03/2021
Var
// Seção de Declarações das variáveis
   nome: caractere

   x, y, resultado: inteiro
   // Variáveis usadas para realização da conta

   numero, controle: inteiro
   // Variável "controle": - Utilizada para encargo da função "para"
   //                      e usada tambem como contador
   
Inicio
// Seção de Comandos, procedimento, funções, operadores, etc... 
   mudacor("Amarelo", "FRENTE")

   x <- 1
   y <- -1
   escreval("         Bem-vindo Usuário!")
   escreval( )
   escreva("Digite seu nome: ")
   leia(nome)
   limpatela( )
   escreval("     *** Sequência de Fibonacci ***")
   escreval( )
   escreval("- Saudações '", maiusc(nome), "', digite um número")
   escreval("- lhe retornameros a sucessão de Fibonacci até tal;")
   escreval( )
   escreva("Número: ")
   leia(numero)
   escreval( )
   
   para controle de 1 ate numero faca
        resultado <- x + y
        y <- x
        x <- resultado
        escreval(controle, "°) ", resultado)
   fimpara
   
   escreval(" ")
   escreval("Obrigado, até a proxima!")
   escreval(" ")
   escreval("'Esc' para sair!")
   
Fimalgoritmo