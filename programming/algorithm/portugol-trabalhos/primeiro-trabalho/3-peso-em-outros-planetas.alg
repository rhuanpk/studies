Algoritmo "peso_em_outros_planetas"
// Disciplina  : Lógica e Fundamentos da Programação
// Professor   : Thiago Momesso de Menezes
// Descrição   : Retorna ao usuário o peso relativo em outro planeta predefinido
// Autor(a)    : Rhuan Patriky de Sousa Ferrer - RA: 1028
// Data atual  : 17/03/2021
Var
// Seção de Declarações das variáveis
   peso: real
   planeta: inteiro

Inicio
// Seção de Comandos, procedimento, funções, operadores, etc...
   mudacor("Amarelo", "FRENTE")
   
   escreval("   *** Programa de Calculo de Gravidade ***")
   escreval( )
   escreval("- Bem-vindo usuário!")
   escreval("- Entre com seu peso")
   escreval("- Calcularemos o peso informado em outro Planeta")
   escreval("- Planetas já predefinidos e escolhido por voçê!")
   escreval( )
   escreva("Peso: ")
   leia(peso)
   escreval( )
   escreval("- Muito bem!")
   escreval("- Agora escolha o planeta!")
   escreval(" ")
   escreval("1) Mercúrio")
   escreval("2) Vênus")
   escreval("3) Marte")
   escreval("4) Júpiter")
   escreval("5) Saturno")
   escreval("6) Urano")
   escreval( )
   escreva("Planeta: ")
   leia(planeta)
   escreval( )
   
   escolha planeta
   caso 1
      escreval("Seu peso em 'Mercúrio' seria aproximadamente de ", ((peso / 10) * 0.37):3:3, " Kg")
   caso 2
      escreval("Seu peso em 'Vênus' seria aproximadamente de ", ((peso / 10) * 0.88):3:3, " Kg")
   caso 3
      escreval("Seu peso em 'Marte' seria aproximadamente de ", ((peso / 10) * 0.38):3:3, " Kg")
   caso 4
      escreval("Seu peso em 'Júpiter' seria aproximadamente de ", ((peso / 10) * 2.64):3:3, " Kg")
   caso 5
      escreval("Seu peso em 'Saturno' seria aproximadamente de ", ((peso / 10) * 1.15):3:3, " Kg")
   caso 6
      escreval("Seu peso em 'Urano' seria aproximadamente de ", ((peso / 10) * 1.17):3:3, " Kg")
   outrocaso
      escreval("Seleção Incorreta!")
   fimescolha

Fimalgoritmo