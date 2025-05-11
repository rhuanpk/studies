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
      # Passe por URL "p=5368" por exemplo
      $preco = $_GET["p"];
      echo "O preço do produto é R$ $preco !";
      echo "</br>10% desse produto é R$ ". ($preco * 0.1). " !";
      $preco += $preco * 0.1;
      echo "</br>Produto com mais 10% é R$ " .number_format($preco, 2, ",", "."). " !";
      $preco = $_GET["p"];
      $preco -= $preco * 0.1;
      echo "</br>Produto com menos 10% é R$ " .number_format($preco, 2, ",", "."). " !";
   ?>
</div>
</body>
</html>
