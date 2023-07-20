<!DOCTYPE html>
<head>
    <meta charset="UTF-8">
    <link rel="stylesheet" href="../_css/estilo.css"/>
    <title>Projects</title>
</head>
<body>
<div>
    <?php
        $n1 = $_GET["a"];
        $n2 = $_GET["b"];

        $soma = $n1 + $n2;
        $media = ($n1 + $n2) / 2;

        echo "<h2>Valores recebidos $n1 e $n2</h2>";
        echo "A soma vale $soma";
        # Da para colocar ou não os parênteses na hora de fazer um cálculo simples na concatenação
        echo "</br>A subtração vale ". $n1 - $n2;
        echo "</br>A multiplicação vale ". ($n1 * $n2);
        echo "</br>A divisão vale ". ($n1 / $n2);
        echo "</br>O modulo vale ". ($n1 % $n2);
        echo "</br>Media vale $media"
    ?>
</div>
</body>
</html>
