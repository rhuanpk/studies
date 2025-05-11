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
      # Calcule a médida entre duas notas passadas por URL e utilizando o operador ternário, informe se o usuário está aprovado ou reprovado
      $nota1 = $_GET["nota1"];
      $nota2 = $_GET["nota2"];
      $media = ($nota1 + $nota2) / 2;
      echo "A média entre as notas $nota1 e $nota2: $media";
      echo "<br/>Situação do aluno: ". ($media >= 6 ? "APROVADO" : "REPROVADO");
   ?>
</div>
</body>
</html>