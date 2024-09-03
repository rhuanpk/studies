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
         # Ler duas notas do aluno: >= 7 ARPOVADO; >= 5 && < 7 RECUPERAÇÃO; < 5 REPROVADO

         # Declaração de variáveis
         $nota1 = (isset($_POST["nota1"]) && $_POST["nota1"] != null) ? $_POST["nota1"] : 0;
         $nota2 = (isset($_POST["nota2"]) && $_POST["nota2"] != null) ? $_POST["nota2"] : 0;
         $media = ($nota1 + $nota2) / 2;

         # Veificações
         if($media >= 7) {
            $situacao = "APROVADO";
         } elseif ($media >= 5) {
            $situacao = "RECUPERAÇÃO";
         } else {
            $situacao = "REPROVADO";
         }

         # Inicio
         echo "Média: $media</br>";
         echo "Situação: $situacao";
      ?>
      </br><a href="exercicio-03.html">Voltar</a>
   </div>
</body>
</html>