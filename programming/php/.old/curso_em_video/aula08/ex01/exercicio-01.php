<!DOCTYPE html>
<html>
<head>
   <meta charset="UTF-8"/>
   <link rel="stylesheet" href="../../_css/estilo.css"/>
   <title>Project</title>
</head>
<body>
   <div>
      <?php
         $valor = $_GET["v"];
         $raizQuadrada = sqrt($valor);
         echo "<br/>Raiz quadrada de $valor = $raizQuadrada";
      ?>
      <a href="exercicio-01.html"><br/><br/>Voltar</a>
   </div>
</body>
</html>