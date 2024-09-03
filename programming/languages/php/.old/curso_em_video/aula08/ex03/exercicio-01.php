<!DOCTYPE html>
<html>
<head>
   <?php
      $txt = (isset($_GET["txt"]) && $_GET["txt"] !== "") ? $_GET["txt"] : "Hello World!";
      $tam = (isset($_GET["tam"]) && $_GET["tam"] !== "") ? $_GET["tam"] : "12pt";
      $cor = (isset($_GET["cor"]) && $_GET["cor"] !== "") ? $_GET["cor"] : "#000000";
   ?>
   <meta charset="UTF-8"/>
   <link rel="stylesheet" href="../../_css/estilo.css"/>
   <title>Project</title>
   <style>
      span.span {
         font-size: <?php echo "$tam"; ?>;
         color: <?php echo "$cor"; ?>;
      }
   </style>
</head>
<body>
   <div>
      <?php
         echo "<span class='span'>$txt</span>";
      ?>
   </div>
</body>
</html>