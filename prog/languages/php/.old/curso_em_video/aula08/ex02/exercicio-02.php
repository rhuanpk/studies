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
      /*
      Nessas comparações estou dizendo que:
      Caso a variável que está vindo via método GET não esteja setadade não seja nula,
      que entre no bloco "true", caso não, entre no else
      */
      $nome = (isset($_GET["nome"]) && $_GET["nome"] !== "") ? $_GET["nome"] : "Não informado!";
      $anoNascimento = (isset($_GET["anoNascimento"]) && $_GET["anoNascimento"] != null) ? $_GET["anoNascimento"] : "Não informado!";
      $sexo = (isset($_GET["sexo"]) && $_GET["sexo"] !== "") ? $_GET["sexo"] : "Não informado!";
      echo "Informações recolhidas";
      echo "<br/><br/>Nome.........................: $nome";
      echo "<br/>Ano de Nascimento...: $anoNascimento";
      echo "<br/>Sexo..........................: $sexo";
      echo "<br/><br/>Idade atual: ". (date("Y") - $anoNascimento);
   ?>
</div>
</body>
</html>