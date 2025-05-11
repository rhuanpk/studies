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
      $ano = $_GET["ano"];
      echo "Ano passado por parâmetro $ano!";
      echo "</br>Ano anterior a esse é ". --$ano. "!";
   ?>
</div>
</body>
</html>
