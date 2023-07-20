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
         # Ler o dia da semana (2-7) e mostrar se precisa ou não ir para a escola
         $diaSemana = (isset($_GET["diaSemana"]) && $_GET["diaSemana"] != null) ? $_GET["diaSemana"] : 1;
         switch($diaSemana) {
            case 2:
            case 3:
            case 4:
            case 5:
            case 6:
               $estuda = "ESTUDA";
               break;
            case 7:
               $estuda = "NÃO ESTUDA";
               break;
            case 1:
               echo "1 Não aceito como parâmetro!";
               break;
         }
         if($diaSemana != 1) {
            echo "Você $estuda nesse dia!";
         }
      ?>
      <!-- Função JS que faz voltar um no histórico de navegação --> 
      </br><a href="javascript:history.go(-1)">Voltar</a>
   </div>
</body>
</html>