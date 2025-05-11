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
         # Ler um número e qual operação realizar: dobro, cubo ou raiz quadrada
         $numero = (isset($_GET["numero"]) && $_GET["numero"] != null) ? $_GET["numero"] : 0;
         $op = (isset($_GET["op"]) && $_GET["op"] !== "") ? $_GET["op"] : "undefined";
         switch($op) {
            case "dobro":
               $ress = $numero * 2;
               break;
            case "cubo";
               # $ress = $numero ^ 3; 
               $ress = pow($numero, 3);
               break;
            case "raiz";
               # $ress = $numero ^ (1/2);
               $ress = sqrt($numero);
               break;
            default:
               echo "Erro! Operador não definido.</br></br>";
         }
         echo "Valor recebido: $numero</br>";
         echo "Operação: $op</br>";
         echo "</br>Resultado = $ress</br>";
      ?>
      </br><a href="exercicio-01.html">Voltar</a>
   </div>
</body>
</html>