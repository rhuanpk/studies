Algoritmo "programa_secretaria"
// Disciplina  : L�gica e Fundamentos da Programa��o
// Professor   : Thiago Momesso de Menezes 
// Descri��o   : Programa que simula o sistema de uma secretaria
// Autor(a)    : Rhuan Patriky de Sousa Ferrer
// Data atual  : 19/06/2021

tipo
   aluno = registro
      ra: caractere
      nome: caractere
      dtNascimento: caractere
      turma: caractere
      n1, n2: real
      notas: logico
   fimregistro

Var
// Se��o de Declara��es das vari�veis 

// --- argumentos dos procedimento "inserir_aluno" e "inserir_notas" ---

   aux1: caractere // aux para 'nome'
   aux2: caractere // aux para 'data de nascimento'
   aux3: caractere // aux para 'turma'
   aux4: real // aux para 'nota 1'
   aux5: real // aux para 'nota 2'

// ---------------------------------------------------------------------

   ficha: vetor[0..49] de aluno   // vet que recebe os registros do aluno
   contador: inteiro // conta o n�mero de alunos cadastrados
   
   readkey: caracter // var auxiliar para "esperar" at� que o usu�rio digite algo
   
   opcao: inteiro // escolhe a opcao no menu
   flag: logico // flag usada para controle do la�o do menu
   
   x: caractere // var aux para receber o ra para busca
   indice: inteiro // var aux que recebe o valor da fun��o "buscar_aluno"
   
// Se��o de declara��o de procedimentos
   procedimento inserir_aluno (x, y, z: caractere)
   inicio
      contador <- contador + 1
      ficha[contador].ra <- "21" + numpcarac(contador)
      ficha[contador].nome <- x
      ficha[contador].dtNascimento <- y
      ficha[contador].turma <- z
      ficha[contador].notas <- falso
   fimprocedimento
   
   procedimento exibir_alunos
   var
      status: caractere // var aux�liar que recebe o status das notas
      media: real // var aux�liar que recebe o calculo da m�dia das notas
      i: inteiro // var de controle do "para"
   inicio
      limpatela
      escreval("***** LISTA DE ALUNOS *****")
      se contador = - 1 entao
         escreva(">>> Nenhum Aluno inserido !")
      fimse
      para i de 0 ate contador faca
         escreva(i + 1, " - ")
         escreva("RA: ", ficha[i].ra, ", ")
         escreva("Nome: ", ficha[i].nome, ", ")
         escreva("Nascimento: ", ficha[i].dtNascimento, ", ")
         escreva("Turma: ", ficha[i].turma, ", ")
         escreva("M�dia 1: ", ficha[i].n1, ", ")
         escreva("M�dia 2: ", ficha[i].n2, ", ")
         media <- (ficha[i].n1 + ficha[i].n2) / 2
         escreva("M�dia Final: ", media, ", ")
         se (media < 7) e (ficha[i].notas = falso) entao
            status <- "NOTAS N�O LAN�ADAS"
         senao
            se media < 7 entao
               status <- "REPROVADO"
            senao
               status <- "APROVADO"
            fimse
         fimse
         escreval(status)
      fimpara
      escreval
      escreval("ALUNOS EXIBIDOS: ", contador + 1)
      escreva("Volte ao menu! <enter>")
      leia(readkey)
   fimprocedimento

   procedimento mostrar_aluno (x: inteiro)
   var
      media: real // var aux�liar que recebe o calculo da m�dia das notas
   inicio
      se x = -1 entao
         escreval("Aluno n�o encontrado!")
      senao
         escreval("***** FICHA DE ALUNO *****")
         escreva("RA: ", ficha[x].ra)
         escreval(" Nome: ", ficha[x].nome)
         escreva("Data de Nascimento: ", ficha[x].dtNascimento)
         escreval(" Turma: ", ficha[x].turma)
         escreva("Nota 1:", ficha[x].n1)
         escreval(" Nota 2:", ficha[x].n2)
         media <- (ficha[x].n1 + ficha[x].n2) / 2
         escreva("M�dia Final: ", media)
         escreva(" Situa��o: ")
         se nao ficha[x].notas entao
            escreval("PENDENTE")
         senao
            se media < 7 entao
               escreval("REPROVADO")
            senao
               escreval("APROVADO")
            fimse
         fimse
      fimse
      escreva("Volte ao menu! <enter>")
      leia(readkey)
   fimprocedimento
   
   procedimento inserir_notas (x: inteiro; y, z: real)
   inicio
         ficha[x].n1 <- y
         ficha[x].n2 <- z
         ficha[x].notas <- verdadeiro
   fimprocedimento

// Se��o de Declara��es de fun��es

   funcao buscar_aluno (x: caractere): inteiro
   var
      i: inteiro // var de controle do para
   inicio
      se contador < 0 entao
         retorne -1
      senao
         para i de 0 ate contador faca
            se x = ficha[i].ra entao
               retorne i
               interrompa
            senao
               retorne -1
            fimse
         fimpara
      fimse
   fimfuncao

Inicio
// Se��o de Comandos, procedimento, fun��es, operadores, etc...
   contador <- -1
   flag <- falso
   enquanto nao flag faca
      limpatela
      escreval("Bem vindo ao programa da secretaria da FACEC!")
      escreval("Selecione o que deseja fazer atrav�s do menu abaixo:")
      escreval("********** MENU **********")
      escreval("1. Cadastrar novo aluno")
      escreval("2. Exibir lista de alunos")
      escreval("3. Exibir ficha de aluno")
      escreval("4. Inserir notas de aluno")
      escreval("5. Sair do programa")
      escreval("**************************")
      escreva("Digite o n�mero da op��o desejada: ")
      leia(opcao)
      escolha opcao
         caso 1
            limpatela
            escreval("INSERIR ALUNO:")
            escreva("Nome: ")
            leia(aux1)
            escreva("Data de Nascimento: ")
            leia(aux2)
            escreva("Turma: ")
            leia(aux3)
            inserir_aluno(aux1, aux2, aux3)
            escreval("Inser��o bem sucedida!")
            escreva("Volte ao menu! <enter>")
            leia(readkey)
         caso 2
            exibir_alunos
         caso 3
            limpatela
            escreval("***** BUSCA DE ALUNO *****")
            escreva("RA do aluno: ")
            leia(x)
            indice <- buscar_aluno(x)
            mostrar_aluno(indice)
         caso 4
            limpatela
            escreval("***** INSERIR NOTAS *****")
            escreva("RA do aluno: ")
            leia(x)
            indice <- buscar_aluno(x)
            se indice = -1 entao
               escreval("Aluno n�o encontrado!")
            senao
               escreval("Inserindo notas para o aluno ", ficha[indice].nome, ":")
               escreva("Nota 1: ")
               leia(aux4)
               escreva("Nota 2: ")
               leia(aux5)
               inserir_notas(indice, aux4, aux5)
               escreval("Notas inseridas com sucesso!")
            fimse
            escreva("Volte ao menu! <enter>")
            leia(readkey)
         caso 5
            flag <- verdadeiro
         outrocaso
            escreval
            escreva("Op��o inv�lida! <enter>")
            leia(readkey)
      fimescolha
   fimenquanto

Fimalgoritmo