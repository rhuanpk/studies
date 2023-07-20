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
         // Faça uma contagem de 1 até 10
         /*
         $foo = 1;
         while($foo <= 10) {
            echo "valor de foo: $foo<br>";
            ++$foo;
         }
         */
        // Faça uma contagem de 10 até 1 pulando de 2 em 2
        $foo = 10;
        while($foo >= 1) {
           echo "valor de foo: $foo<br>";
           $foo -= 2;
        }
      ?>
      </br><a href="javascript:history.go(-1)">Voltar</a>
   </div>
</body>
</html>