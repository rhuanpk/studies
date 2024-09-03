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
      # Permitir que o usuário escolha entre somar e multiplicar dois números
      $numero1 = $_GET["numero1"];
      $numero2 = $_GET["numero2"];
      $operador = $_GET["operador"];
      $operador = $operador == "soma" ? "+" : "*";
      echo 'Passe por URL dois números (pelas variáveis "numero1" e "numero2") e o operador (variável "operador") de adição ou multiplicação, passe como valor "soma" se quiser somar, se quiser multiplicar passe "multi"... para fazer a conta!';
      echo "<br/><br/>Primeiro número: $numero1";
      echo "<br/>Segundo número: $numero2";
      echo "<br/>Operador escolhido: $operador";
      $resultado = $operador == "+" ? $numero1 + $numero2 : $numero1 * $numero2;
      echo "<br/><br/>Resultado = ". $resultado;
   ?>
</div>
</body>
</html>