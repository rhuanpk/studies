Algoritmo "imc"

Var
   peso, altura: real
   imc: caractere
   x: real
   
Inicio
   escreval("   ** Cálculo de IMC **")
   escreval( )
   escreva("Peso: ")
   leia(peso)
   escreva("Altura: ")
   leia(altura)
   escreval( )
   x <- peso / (altura^2)
   se x < 18.5 entao
      imc <- "Abaixo do peso!"
   senao
      se (x >= 18.5) e (x < 25) entao
         imc <- "Normal!"
      senao
         se (x >= 25) e (x < 30) entao
            imc <- "Excesso de peso!"
         senao
            se (x >= 30) e (x < 40) entao
               imc <- "Obesidade!"
            senao
               se x > 39.9 entao
                  imc <- "Obesidade mórbida!"
               fimse
            fimse
         fimse
      fimse
   fimse
   escreval("Sua situação é: ", imc)
   
Fimalgoritmo