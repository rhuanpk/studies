Algoritmo "menu_opcoes"
// Disciplina  : Lógica e Fundamentos da Programação
// Professor   : Thiago Momesso de Menezes
// Descrição   : Mostro como um menu de opções funcionaria
// Autor(a)    : Rhuan Patriky de Sousa Ferrer - RA: 1028
// Data atual  : 19/03/2021
Var
// Seção de Declarações das variáveis
   opcao: inteiro
   ponto: caractere
   
Inicio
// Seção de Comandos, procedimento, funções, operadores, etc...
   mudacor("Amarelo", "FRENTE")
   
   opcao <- 1
   enquanto (opcao > 0) ou (opcao < 0) faca
      limpatela( )
      escreval(" *** Bem-vindo Usuário! ***")
      escreval( )
      escreval("Escolha uma das opcções")
      escreval( )
      escreval("1. Cadastrar usuário;")
      escreval("2. Pesquisar usuário;")
      escreval("3. Listar usuário;")
      escreval("4. Excluir usuário;")
      escreval("0. Sair;")
      escreval( )
      escreva("Escolha: ")
      leia(opcao)
      escreval( )
      
      escolha opcao
      caso 1
         escreval("Opção escolhida: 1. Cadastrar usuário;")
      caso 2
         escreval("Opção escolhida: 2. Pesquisar usuário;")
      caso 3
         escreval("Opção escolhida: 3. Listar usuário;")
      caso 4
         escreval("Opção escolhida: 4. Excluir usuário;")
      caso 0
         escreva("'''SAIR'''")
      outrocaso
         escreval("Opção invalida!")
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