Algoritmo "menu_opcoes"
// Disciplina  : L�gica e Fundamentos da Programa��o
// Professor   : Thiago Momesso de Menezes
// Descri��o   : Mostro como um menu de op��es funcionaria
// Autor(a)    : Rhuan Patriky de Sousa Ferrer - RA: 1028
// Data atual  : 19/03/2021
Var
// Se��o de Declara��es das vari�veis
   opcao: inteiro
   ponto: caractere
   
Inicio
// Se��o de Comandos, procedimento, fun��es, operadores, etc...
   mudacor("Amarelo", "FRENTE")
   
   opcao <- 1
   enquanto (opcao > 0) ou (opcao < 0) faca
      limpatela( )
      escreval(" *** Bem-vindo Usu�rio! ***")
      escreval( )
      escreval("Escolha uma das opc��es")
      escreval( )
      escreval("1. Cadastrar usu�rio;")
      escreval("2. Pesquisar usu�rio;")
      escreval("3. Listar usu�rio;")
      escreval("4. Excluir usu�rio;")
      escreval("0. Sair;")
      escreval( )
      escreva("Escolha: ")
      leia(opcao)
      escreval( )
      
      escolha opcao
      caso 1
         escreval("Op��o escolhida: 1. Cadastrar usu�rio;")
      caso 2
         escreval("Op��o escolhida: 2. Pesquisar usu�rio;")
      caso 3
         escreval("Op��o escolhida: 3. Listar usu�rio;")
      caso 4
         escreval("Op��o escolhida: 4. Excluir usu�rio;")
      caso 0
         escreva("'''SAIR'''")
      outrocaso
         escreval("Op��o invalida!")
         escreval(" ")
         escreval("'Enter' para voltar a escolher")
         escreva("'Esc' x2 para sair")
         leia(ponto)
      fimescolha
      escreval( )
      
      se((opcao > 0) e (opcao < 5))entao
         escreval("'Enter' para voltar a escolher")
         escreva("'Esc' x2 para sair")
         leia(ponto)
      fimse
   fimenquanto

Fimalgoritmo