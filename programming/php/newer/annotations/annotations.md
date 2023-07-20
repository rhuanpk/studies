# PHP Annotations

## Tags

*Scripts* **PHP** podem ser escritos em qualquer arquivos e o php conseguirá entender pois, ele tentará ler somente o que estiver dentro das *tags*.

Nesses casos aonde você escreverá *scripts* dentro de outros documentos, será necessários usar as *tags* completas (abrir e fechar as *tags*), porém, a documentação oficial diz que em caso de arquivos com **PHP** puro é recomendável que apenes inicie as tags no final do arquivos sem ter o fechamento das mesmas no final do arquivos para evitar problemas de espaços em branco depois da *tag* de fechamente, que poderia acarretar em problemas na execução.

Exemplo de *tags*:

1. *Tag* normal: `<?php ?>`.

2. *Tag de echo* curta: `<?= ?>`.

3. *Tag* curta: `<? ?>`.

OBS: a *tag de echo* curta é usada para quando precisamos escrever algum conteúdo simplesmente, por exemplo:

```php
# This expression:
<?php echo "hello world!\n"; ?>

# Is the same that:
<?= "hello world!\n" ?>
```

---

## Types

Tipos primitivos:

- `string`: corda de caracteres.
- `int`: número inteiro.
- `float`: número com casa decímal.
- `bool`: tipo lógico.

Tipos compostos:

- `array`: lista de dados.
- `object`: asbstração de objeto real.
- `callable`.
- `iterable`.

Tipos especiais:

- `resource`.
- `NULL`: não existente (nulo).

Exemplos de utilização:

```php
$any_string = 'any string';
$any_int = 1;
$any_float = 1.5;
$any_booltrue = true;
$any_bool_false = false;
$any_array_empty = array();
$any_array_indexed = array(
	'hello world',
	1,
	2.5,
	true,
	false,
);
$any_array_associative = array(
	'first' => 'hello world',
	'second' => 1,
	'third' => 2.5,
	'fourth' => true,
	'fifth' => false,
);
$any_null = null;

echo gettype($any_string). "\n";
echo gettype($any_int). "\n";
echo gettype($any_float). "\n";
echo gettype($any_booltrue). "\n";
echo gettype($any_bool_false). "\n";
echo gettype($any_array_empty). "\n";
echo gettype($any_array_indexed). "\n";
echo gettype($any_array_associative). "\n";
echo gettype($any_null). "\n";
```

OBS: **PHP** também suporta [*heredoc*](https://www.php.net/manual/pt_BR/language.types.string.php).

### Casting

- `(string)`: converte para string.
- `(int)`: converte para inteiro.
- `(float)`: converte para número de ponto flutuante.
- `(bool)`: converte para booleano.
- `(array)`: converte para array.
- `(object)`: converte para um objeto.

Exemplo de utilização:

```php
$xpto = 10;            // $xpto é um interio
$str = "$xpto";        // $str é uma string
$fst = (string) $xpto; // $fst tambem é uma string
```

OBS: outra forma de definir um tipo a uma variável já verificar se o *casting* foi possível, é com a função [settype](#function_settype).

---

## Double and Single Quotes

O **PHP** faz a distinção de aspas simples e duplas da mesma forma que o *bash*, sendo que as **aspas simples interpretam literalmente as strings** e **aspas duplas expandem conteúdos dinâmicos**.

Exemplo com aspas simples:

```php
$xpto = 'any string';
echo 'output: {$xpto}!\noutput: exitings...';
# output: {$xpto}!\noutput: exitings...
echo "output: {$xpto}!\noutput: exitings...";
# output: any string!
# output: exitings...
```

---

## Classes

### Inherits

Classes podem ser extendidas com a cláusula `extends` na declaração da classe.

Exemplo de utilização:

```php
class Employee extends Person {}
```

### Access Operators

Para acessar os membros de um objeto é utilizado o operador `->`.

Exemplo de utilização:

```php
car->sound();
```

Para acessar os membros `static` de uma classe é utilizado o operador `::`.

Exemplo de utilização:

```php
math::sum();
```

### Object Property

Tipos de propriedades:

```php
public $variable_1 = 1;
protected $variable_2 = 2;
private $variable_3 = 3;
static $variable_4 = 4;
```

- `private`: somente a classe atual terá acesso ao campo ou método.

- `protected`: apenas classe e subclasses atuais (e às vezes também classes do mesmo pacote) desta classe terão acesso ao campo ou método.

- `public`: qualquer classe pode se referir ao campo ou chamar o método.

- `static`: pode utilizar o objeto sem nenhuma instância ou herança, chamando o objeto diretamente pela classe.

---

## *include* and *require*

As duas notações é para fazer o *import* de classes externas para dentro da sua classe.

A diferença entre os dois é que:

- `include`: caso o arquivo da classe não exista, não estoura erro ao início do tempo de execução, somente quando já estiver em execução e no meio do programa for chamado alguma função deste arquivo faltante.

- `require`: caso o arquivo da classe não exista, ele para a execução do script ao início do tempo de execução.

Exemplo de utilização:

```php
include 'script.php';
require 'script.php';
```

Quando o import é feito, e ele tem o mesmo efeito que o `source` no *bash*. Ele carrega para dentro do arquivo que está importando todo o conteúdo do arquivo importado.

### *once* Notation

Caso os *include's* ou *require's* sejam chamado mais de uma vez no mesmo arquivo, o arquivo importado será executado mais de uma vez. Nesses casos temos a chamados dessas notações com a cláusula *once*.

Exemplo de utilização:

```php
include_once 'script.php';
require_once 'script.php';
```

### *@* Operator

Caso tenha algum erro que venha dos arquivos importados, esses erros podem ser **suprimidos** com o operador `@` que irá no ínicio da chamada do comando.

Exemplo de utilização:

```php
@include_once 'script.php';
@require_once 'script.php';
```

### Observations

- Para passar o caminho do arquivo, o interpretador entendende hierarquia padrão de SO.

---

## Operators

Lista de operadores já mostrando as precedências:

| Associação      | Operadores                                               |
| --------------- | -------------------------------------------------------- |
| não associativo | clone new                                                |
| esquerda        | [                                                        |
| direita         | **                                                       |
| direita         | ++ -- ~ (int) (float) (string) (array) (object) (bool) @ |
| não associativo | instanceof                                               |
| direita         | !                                                        |
| esquerda        | * / %                                                    |
| esquerda        | + - .                                                    |
| esquerda        | << >>                                                    |
| não associativo | < <= > >=                                                |
| não associativo | == != === !== <> <=>                                     |
| esquerda        | &                                                        |
| esquerda        | ^                                                        |
| esquerda        | |                                                        |
| esquerda        | &&                                                       |
| esquerda        | ||                                                       |
| direita         | ??                                                       |
| esquerda        | ? :                                                      |
| direita         | = += -= *= **= /= .= %= &= |= ^= <<= >>=                 |
| esquerda        | and                                                      |
| esquerda        | xor                                                      |
| esquerda        | or                                                       |

---

## Built-in

- `var_dump(<object>)`: print informações sobre o objeto passado como argumento (e.g. `var_dump($variable)`).

- `gettype(<object>)`: printa o tipo do objeto passado como argumento (e.g. `gettype($variable)`) (usado para *debug* e não para o código).

- `is_<type>(<object>)`: retorna true ou false verificando se o objeto passado é do tipo verificado (e.g. `is_string($variable)`).

<a id="function_settype"></a>
- `settype(<variable>, <type>)`: seta o tipo passado como segundo argumento na variável passada como primeiro argumento e retorna `true` ou `false` caso a operação seja bem ou mal sucedida, respectivamente (e.g. `settype($variable, 'bool')`).

- `isset(<variables>)`: verifica se uma ou mais variáveis passadas estão definidas (se existem) e retorna `true` caso *é definida* ou `false` caso não seja definida (e.g. `isset($array['first'])`).

- `unset(<variable>)`: destroi uma ou mais variáveis passada como argumento e caso esteja numa função, o objeto excluido é somente referente ao escopo daquela função (e.g. `unset($array['first'])`).
