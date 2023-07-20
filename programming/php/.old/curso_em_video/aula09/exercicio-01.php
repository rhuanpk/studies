<!DOCTYPE html>
<html>
<head>
   <meta charset="UTF-8"/>
   <link rel="stylesheet" href="../_css/estilo.css"/>
   <title>Project</title>
</head>
<body>
   <div>
      <?php
         # Declaração de variáveis
         $anoNascimento = (isset($_POST["anoNascimento"]) && $_POST["anoNascimento"] != null) ? $_POST["anoNascimento"] : 1900;
         $idade = date("Y") - $anoNascimento;

         # Veificações
         if($idade > 17) {
            $votar = true;
            $dirigir = true;
         } else {
            $votar = false;
            $dirigir = false;
         }

         # Inicio
         echo "<pre>";
         echo "Ano nascimento: $anoNascimento\n";
         echo "\nIdade: $idade anos\n";
         echo "\n\tVotar: ". ($votar ? "SIM" : "NÃO"). "\n";
         echo "\tDirigir: ". ($dirigir ? "SIM" : "NÃO");
         echo "</pre>";
      ?>
      </br><a href="exercicio-01.html">Voltar</a>
   </div>
</body>
</html>