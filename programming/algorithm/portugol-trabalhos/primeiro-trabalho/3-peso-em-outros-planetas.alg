Algoritmo "peso_em_outros_planetas"
// Disciplina  : L�gica e Fundamentos da Programa��o
// Professor   : Thiago Momesso de Menezes
// Descri��o   : Retorna ao usu�rio o peso relativo em outro planeta predefinido
// Autor(a)    : Rhuan Patriky de Sousa Ferrer - RA: 1028
// Data atual  : 17/03/2021
Var
// Se��o de Declara��es das vari�veis
   peso: real
   planeta: inteiro

Inicio
// Se��o de Comandos, procedimento, fun��es, operadores, etc...
   mudacor("Amarelo", "FRENTE")
   
   escreval("   *** Programa de Calculo de Gravidade ***")
   escreval( )
   escreval("- Bem-vindo usu�rio!")
   escreval("- Entre com seu peso")
   escreval("- Calcularemos o peso informado em outro Planeta")
   escreval("- Planetas j� predefinidos e escolhido por vo��!")
   escreval( )
   escreva("Peso: ")
   leia(peso)
   escreval( )
   escreval("- Muito bem!")
   escreval("- Agora escolha o planeta!")
   escreval(" ")
   escreval("1) Merc�rio")
   escreval("2) V�nus")
   escreval("3) Marte")
   escreval("4) J�piter")
   escreval("5) Saturno")
   escreval("6) Urano")
   escreval( )
   escreva("Planeta: ")
   leia(planeta)
   escreval( )
   
   escolha planeta
   caso 1
      escreval("Seu peso em 'Merc�rio' seria aproximadamente de ", ((peso / 10) * 0.37):3:3, " Kg")
   caso 2
      escreval("Seu peso em 'V�nus' seria aproximadamente de ", ((peso / 10) * 0.88):3:3, " Kg")
   caso 3
      escreval("Seu peso em 'Marte' seria aproximadamente de ", ((peso / 10) * 0.38):3:3, " Kg")
   caso 4
      escreval("Seu peso em 'J�piter' seria aproximadamente de ", ((peso / 10) * 2.64):3:3, " Kg")
   caso 5
      escreval("Seu peso em 'Saturno' seria aproximadamente de ", ((peso / 10) * 1.15):3:3, " Kg")
   caso 6
      escreval("Seu peso em 'Urano' seria aproximadamente de ", ((peso / 10) * 1.17):3:3, " Kg")
   outrocaso
      escreval("Sele��o Incorreta!")
   fimescolha

Fimalgoritmo