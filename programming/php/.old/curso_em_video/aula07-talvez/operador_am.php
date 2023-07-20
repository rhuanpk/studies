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
      $an = $_GET["an"];
      $idade = 2021 - $an;
      echo "Pessoas com $idade anos ". (($idade > 17 && $idade < 65) ? "SÃO" : "NÃO SÃO"). " obrigados a votar!";
   ?>
</div>
</body>
</html>