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
         if($idade < 16) {
            $tipoVoto = "NÃO VOTA";
         } elseif (($idade < 18) || ($idade > 64)) {
            $tipoVoto = "OPCIONAL";
         } else {
            $tipoVoto = "OBRIGATÓRIO";
         }

         # Inicio
         echo "<pre>";
         echo "Ano nascimento: $anoNascimento\n";
         echo "\nIdade: $idade anos\n";
         echo "\n\tVoto: $tipoVoto";
         echo "</pre>";
      ?>
      </br><a href="exercicio-02.html">Voltar</a>
   </div>
</body>
</html>