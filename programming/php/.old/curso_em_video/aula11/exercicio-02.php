<!DOCTYPE html>
<html>
<head>
   <meta charset="UTF-8"/>
   <link rel="stylesheet" href="../_css/estilo.css"/>
   <title>Project</title>
</head>
<body>
   <div>
      <form method="get" action="exercicio-02.php">
         <?php
            $cont = 0;
            $itens = 4;
            while($cont <= $itens) {
               echo "Valor ". ($cont + 1). ": <input type='number' name='valor$cont' value='0' min='0' max='100'><br>";
               ++$cont;
            }
            $valor1 = $_GET["valor0"];
            echo "Valor do input valor1: $valor1<br>";
         ?>
         <input type="submit" value="Enviar">
         <!-- </br><a href="javascript:history.go(-1)">Voltar</a> -->
      </form>
   </div>
</body>
</html>