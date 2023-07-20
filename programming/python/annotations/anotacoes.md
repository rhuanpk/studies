# Python

## Modos de executar um programa python

### Python interactive prompt

- Chame o interpretador do python na linha de comando com `python3`

- Uma vez estando dentro do interpretador pode-se fazer qualquer declaração python e executa-la com *return key*

### Source file

- Escreva um arquivo de qualquer nome com a extensão `.py` e salve-o

- Dê permissão de execução para o arquivo e rode-o chamando ele com o python:
	`python ./file.py`

---

## Módulos

Bibliotecas em Python são conhescidas como *módulos*, que básicamente pode ser qualquer arquivo que já vem copilado na linguagem ou qualquer arquivo externo que tenha extensão *.py* ou *.pyc*.

Para fazer o `import` de um módulo, o arquivo do programa Python deve estar no mesmo nível da pasta que está rodando o código príncipal ou em algum outro diretório da variável `sys.path` do Python.

OBS: não há necessidade de explanar a extesão do módulo na declaração do `import` comando.

### Compilação

Módulos podem ser compilados, porém, é uma tarefa cara (do que diz respeito ao processamento), por isso python disponibiliza os módulos com extensão *.pyc*, que são módulo pré compilados pelo criador e sua compilação final acontece no utilizador (e são agnósticos de plataforma).

### `from/import`

Se você deseja importar diretamente um atributo de uma bibliotéca/módulo sem precisar necessáriamente importar a bibliotéca/módulo inteira, você pode usar a declaração `from/import`, dizendo que do módulo (ou seja, `from`) quer importar (ou seja, `import`) tal atribudo:

```python
from sys import argv
print(argv)
```

Caso ainda queira importar mais de um atribudo de uma vez, pode-se usar o operador de separação para isso:

```python
from sys import argv, path
print(argv, path)
```

Caso ainda faça o `import` da forma recomendada, é necessário usar acessar os atrbutos pelo operador `.` passando pelo objeto:

```python
import sys
print(sys.argv)
```

OBS: a ressalva que essa técnica tem é que se caso você tenha variáveis declaradas com o mesmo nome, hávera o conflito.

### Atributo especial: `__name__`

Todos módulo tem um atributo especial imbutido que se chama `__name__` que **guarda o nome do módulo atual**, esse atributo é global e *spawnado* junto com o próprio módulo, portanto não há a necessidade de nenhuma instanciação extra.

### Atributo especial: `__version__`

Guarda o valor da versão do módulo caso ele seja declarado:

```python
__version__ = "0.0.0"
```

---

## Biblíotecas

### `sys`

Módulo do sístema operacional no qual está rodando que contem por exemplos os argumentos passados na linha de comando e os diretórios da variável *path* do Python.

#### Variáveis/Atributo

- `argv`: contém os argumentos passados na linha de comando.
- `path`: contém o caminho dos diretórios de importação de módulos.

### `os`

Módulo do sistema operacional que fornece funções (portáveis) para interação direto com o kernel.

#### Funções/Métodos

- `getcwd()`: retorna o path da onde o script foi invocado.
- `listdir()`: retorna a lista do diretório atual.

---

## Comentários

- One line:

```python
# Essa função printa na tela
print("hello world")
```

ou

```python
print("hello world") # Essa função printa na tela
```

---

## Tipagem dos dados?

```python
some_int: int = 1234
some_float: float = 1.234
some_string: str = "some string"
some_bool: bool = True|False
some_set: set = {1, 2, 3, 4}
some_list: list = []
some_tuple: tuple = 1, 2, 3, 4
some_dict: dict = {}
```

Tipos podem ou não ser informados nas declarações.

> "Python é dinâmicamente e fortemente tipada."

### Números

**Interios** (o tipo `int` serve para qualquer tamanho de inteiro sem a necessidade por exemplo de um `long`):

- 2

**Floats** (os dois tipos na verdade podem receber a notação *E*?):

- 52.3E-4 (que é equivalente a 52.3 * 10^-4^)

### Strings

**Aspas simples**:

String que tenham espaços, espaços em branco e tabs são preservados.

**Aspas duplas**:

Funcionam exatamente como as *aspas simples*, ainda conseguindo ter aspas simples na string (o mesmo vale para as aspas simples para terem duplas na string).

**Aspas triplas**

Você fica livre para escrever como um *text area*? Com essa variação de string você pode fazer coisas como:

```python
print("""This is a multi-line string. This is the first line.
This is the second line.
	"What's your name?," I asked.
He said "Bond, James Bond."
""")
```

E os tabs também serão interpretados.

**Strings brutas**

Strings brutas são equivalentes as aspas simples no shell do Linux e para ser válida a string tem que possuir um caracter `r` antes da mesma:

```python
print(r"as new line\n")
```

**Sequências de escape**

Elas funcionam praticamente da mesma forma que nos shell's do Linux...

- Você pode escapar qualquer caracter com *backslash* (barra invertida: "\"):

```python
print('Um copo d\'água.')
```

- Para uma nova linha use o caracter de escape `backslash + n`:

```python
print("Primeira linha\nSegunda linha\nTerceira linha")
```

- Para um tab use o caracter de escape `backslash + t`:

```python
print("\tInício do parágrafo.")
```

- Para suprimir a *new line* use *backslash* no final da *string*:

```python
print("""
hello \
world
""")
```

### Booleanos

- `True`: Informa que o valor é verdadeiro (código de retorno 1): equivalente a **true**

- `False`: Informa que o valor é falso (código de retorno 0): equivalente a **false**

### Tipos exepcionais (especiais)

- `None`: Diz que o conteúdo da variável é vazio, ou seja, não tem conteúdo: equivalente a **null**

### Casting

- `int(some_variable|some_expression|"<some_string>")`: converte um único valor passado como argumento para inteiro (pega somente a parte inteira da expressão e se for possível converter)

---

## Estrutura de dados

São literalmente estruturas de dados (tipos de dados) que podem armazenar um grupo de informaçẽos relacionadas.

### *list*

A estrutura de dados do tipo **lista** é equivalente a um array {uni|multi}dimensional indexado que aceita incerções, remoções e pesquisa de seus itens. Sua sintaxe é proposta com `[]` (colchetes).

Exemplo de declaração:

```python
any_list = ["hello", 0, "world", 1.5, True]
``` 

#### Métodos:

- `append({"any_string"|any_variable})`: add o argumento passado por parâmetro no final da lista.

- `sort()`: simplesmente ordena a lista sem nenhum explícito (ou seja, afeta o próprio objeto).

### *tuple*

A estrutura de dados do tipo **tupla** é equivalente a um array {uni|multi}dimensional indexado, porém, não aceita *appends* ou alterações *itself*. Essa estrutura de dados é mais usada simplesmente para guardar valores e não para ser manipulada, por isso, seu uso é mais comum quando sabemos exatamente o que iremos guardar nela. Os meios de acessala são iguais aos da estrutura de dados `list` (por meio de colchetes) e podemos guardar praticamente qualquer tipo de valor dentro.

Exemplo de declaração:

```python
any_tuple = ("first", "second", "third")
```

### *dictionary*

### *set*

---

## Objetos/Classes

> "*Quando usamos uma variável i e atribuímos um valor a ela, digamos inteiro 5, você pode pensar nisso como a criação de um **objeto** (ou seja, instância) i de **classe** (ou seja, tipo) int.*"
>
> By: Swaroopch.



---

## Funções

- `help("<function>")`: printa na tela informações sobre alguma decralação python

- `exit()`: exita do prompt do python

- `print("<some_string>", some_variable, some_expression)`: printa na tela os argumentos passados

	- Parâmetro especial `end` altera o IFS (*Internal Field Separator*) da função print que é passada:
		`print("anything", end="")`

- `format(foo, bar)`: usada geralmente junto com alguma string e é mais performático do que concatenação direta? troca os parâmetros da string pelos argumentos da função, exemplos:

	- Trocando os parâmetros respectivamente informando o índice de cada um:
		`print("Me chamo {0} e tenho {1} anos de idade!".format(nome, idade))`
	- Trocando os parâmetros respectivamente informando o índice de cada um:
		`print("Me chamo {} e tenho {} anos de idade!".format(nome, idade))`
	- Trocando os parâmetros informando o índice personalizado de cada um independente da ordem em que são passados:
		`print("Me chamo {nome} e tenho {idade} anos de idade!".format(idade=idade, nome=nome))`
	- Formatando a quatidade de casas flutuantes de um parâmetro float:
		`print("{:.2f}".format(1.0/3))`
	- Informe a quantidade de caracteres, o que faltar da string passada, complete com determinado caracter:
		`print("{:_^11}".format("hello"))`
	- Pode-se declarar os parâmetros nomeados diretamente na passagem:
		`print("{name} wrote {book}".format(name="Swaroop", book="A Byte of Python"))`

- `input("enter to the value: ")`: função que retorna a entrada de dados do usuário no tipo string.

- `range(start_number, end_number, step_count)`: retorna uma sequência de números definida (como um objeto?).

- `len(string_variabel|"string")`: retorna o tamanho da string passada.

- `strip(string_variabel|"string")`: retorna a mesma string só que com os espaços em branco no seu início e final cortados.

- `lstrip(string_variabel|"string")`: retorna a mesma string só que com os espaços em branco no seu início (*left*) cortados.

- `rstrip(string_variabel|"string")`: retorna a mesma string só que com os espaços em branco no seu final (*right*) cortados.

- `dir([object_name])`: retorna toda a declaração de um de terminado objeto, caso não passado argumentos, retorna todos o membros do módulo atual, caso passado o nome do módulo, retorna os membros do mesmo.

- `vars([object_name])`: retorna todos os membros do objetos além dos seus valores e caso não passado o nome do objeto como parâmetro, retornará os membros do objeto pela qual está sendo invocada. 

### Comando return

O comando/declaração `return` é usado para retornar o valor de uma função. Caso você use a declaração sem passar nenhum valor para ela, implicitamente a função retorna com o tipo de dado especial `None`.

Mesmo que você não faça nenhuma declaração explica do comando `return`, implicitamente no final de cada função há uma chamada do comando.

### DocStrings (documentation strings)

É a forma de documentarmos os nossos métodos/calsses. Uma `string` na primeira linha lógica dessa função é sua *DocString*. 

Por convenção:

1. A primeira linha é iniciada com **primeira letra maiúscula** e termina com um **ponto final**.

1. A segunda linha é uma linha em branco.

1. Da terceira linha em diante é a explicação do objeto. 

Para poder chamar a *DocString*, se usa a sintaxe `<object_name><access_operator><special_atribute>`:

como:

```python
def hello():
	"""Say a simple "hello world"."""
	print("hello world!")
print(hello.__doc__)
```

---

## Comandos

- `pass`: serve para indificar um bloco de código vazio sem instruções.
- `del`: serve para *unsetar* algum objeto (seja módulo importado, método ou atributo declarado)

---

## Parâmetros nomeado

É quando passamos parâmetros com nomes ao invés de seu índice, por exemplo com a função `format()`. No python >= 3.6 foi introduzido as "*f-strings*", que é um caminho mais curto para os parâmetros nomeados:

```python
nome = "rhuan"
idade = 22

print(f"Me chamo {nome} e tenho {idade} anos de idade!")
```

OBS: eles também podem invocar métodos.

---

## Linha Lógica e Física

Uma linha física é o que você vê quando escreve o programa. Uma linha lógica é o que Python vê como uma única instrução. O Python assume implicitamente que cada linha física corresponde a uma linha lógica. Um exemplo de uma linha lógica é uma declaração como `print("hello world")` se isso estivesse em uma linha por si só (como você vê em um editor), isso também corresponde a uma linha física. Implicitamente, o Python incentiva o uso de uma única instrução por linha, o que torna o código mais legível. Se você deseja especificar mais de uma linha lógica em uma única linha física, você deve especificar isso explicitamente usando um ponto e vírgula (;) que indica o final de uma linha/instrução lógica. Por exemplo:

```python
i = 5
print(i)
```

é efetivamente o mesmo que

```python
i = 5;
print(i);
```

que também é igual

```python
i = 5; print(i);
```

e igual

```python
i = 5; print(i)
```

---

## Operadores

- **+** (mais) adiciona dois objetos:

	- `3 + 5 = 8`
	- `"a" + "b" = "ab"`

- **-** (menos) dá a subtração de um número do outro e se o primeiro operando estiver ausente, assume-se que é zero:

	- `- 5.2` = *um número negativo*
	- `50 - 24 = 26`

- **\*** (multiplicar) dá a multiplicação dos dois números ou retorna a string repetida tantas vezes:

	- `2 * 3 =  6`
	- `"la" * 3 = "lalala"`

- **\*\*** (potência) retorna x elevado a y:

	- `3 ** 4 = 81` *(ou seja: 3 * 3 * 3 * 3)*

- **/** (dividir) divida x por y:

	- `13 / 3 = 4.333333333333333`

- **//** (dividir e andar) divida x por y e arredonde a resposta para o valor inteiro mais próximo. Observe que se um dos valores for um float, você receberá um float de volta:

	- `13 // 3 = 4`
	- `-13 // 3 = -5`
	- `9 // 1.81 = 4.0`

- **%** (módulo) retorna o resto da divisão:

	- `13 % 3 = 1`
	- `-25.5 % 2.25 = 1.5`

- **<<** (desvio à esquerda) desloca os bits do número para a esquerda pelo número de bits especificado (cada número é representado na memória por bits ou dígitos binários, ou seja, 0 e 1):

	- `2 << 2 = 8` *(2 é representado 10em bits, deslocamento à esquerda de 2 bits dá 1000 o que representa o decimal 8)*

- **>>** (deslocamento para a direita) desloca os bits do número para a direita pelo número de bits especificado:

	- `11 >> 1 = 5` *(11 é representado em bits pelo 1011 qual quando deslocado para a direita em 1 bit dá 101 qual é o decimal 5)*

- **&** (E bit a bit) E bit a bit dos números: se ambos os bits forem 1, o resultado será 1. Caso contrário, é 0:

	- `5 & 3 = 1` *(0101 & 0011 = 0001)*

- **|** (OU bit a bit) OU bit a bit dos números: se ambos os bits forem 0, o resultado será 0. Caso contrário, é 1:

	- `5 | 3 = 7` *(0101 | 0011 = 0111)*

- **^** (XOR bit a bit) XOR bit a bit dos números: se ambos os bits ( 1 or 0) forem iguais, o resultado será 0. Caso contrário, é 1:

	- `5 ^ 3 = 6` *(O101 ^ 0011 = 0110)*

- **~** (inversão bit a bit) a inversão bit a bit de x é `-(x+1)`:

	- `~5 = -6` *([mais detalhes](http://stackoverflow.com/a/11810203))*

- **<** (menor que) retorna se x é menor que y. Todos os operadores de comparação retornam *true* ou *false*. Observe a capitalização desses nomes:

	- `5 < 3 = false`
	- `3 < 5 = true`
	- `3 < 5 < 7 = true`

- **>** (maior que) retorna se x é maior que y. Se ambos os operandos forem números, eles serão primeiro convertidos em um tipo comum. Caso contrário, ele sempre retorna *false*:

	- `5 > 3 = true`

- **<=** (menos que ou igual a) retorna se x é menor ou igual a y:

	- `x = 3; y = 6; x <= y = true`

- **>=** (maior que ou igual a) retorna se x é maior ou igual a y:

	- `x = 4; y = 3; x >= y = true`

- **==** (igual a) compara se os objetos são iguais:

	- `x = 2; y = 2; x == y = true`
	- `x = "str"; y = "stR"; x == y = false`
	- `x = "str"; y = "str"; x == y = true`

- **!=** (não igual a) compara se os objetos não são iguais:

	- `x = 2; y = 3; x != y true`

- **not** (booleano NÃO) se x for *true*, ele retornará *false*. Se x for *false*, ele retornará *true*.

	- `x = true; not x = false`.

- **and** (booleano AND) x and y retorna *false* se x é *false*, senão retorna a avaliação de y:

	- `x = false; y = true; x and y = false`

Desde que x é false. Nesse caso, o Python não avaliará y, pois sabe que o lado esquerdo da expressão 'e' é false o que implica que toda a expressão será false independente dos outros valores. Isso é chamado de avaliação de curto-circuito.

- **or** (booleano OR) se x for *true*, ele retornará *true*, senão retornará avaliação de y:

	- `x = true; y = false; x or y = true`

A avaliação de curto-circuito também se aplica aqui.

### Operadores de atribuição

Adição:

```python
foobar = 2
foobar += 3
```

Subtração:

```python
foobar = 2
foobar -= 3
```

Multiplicação:

```python
foobar = 2
foobar *= 3
```

Divisão:

```python
foobar = 2
foobar /= 3
```

### Ordem de precedência

- `lambda` : Lambda Expression

- `if - else` : Conditional expression

- `or` : Boolean OR

- `and` : Boolean AND

- `not x` : Boolean NOT

- `in, not in, is, is not, <, <=, >, >=, !=, ==` : Comparisons, 
including membership tests and identity tests

- `|` : Bitwise OR

- `^` : Bitwise XOR

- `&` : Bitwise AND

- `<<, >>` : Shifts

- `+, -` : Addition and subtraction

- `*, /, //, %` : Multiplication, Division, Floor Division and Remainder

- `+x, -x, ~x` : Positive, Negative, bitwise NOT

- `**` : Exponentiation

- `x[index], x[index:index], x(arguments...), x.attribute` : Subscription, slicing, call, attribute reference

- `(expressions...), [expressions...], {key: value...}, {expressions...}` : Binding or tuple display, list display, dictionary display, set display

*OBS: bottom up*

A ordem de precedência natural pode ser mudada com parênteses (*()*), a expressão que estiver dentro do parênteses será executada com prioridade sobre todos e se tiver mais de um na linha lógica, executará na ordem em que é achado (da esquerda para a direita).

---

## Fluxo de controle

Sintaxes:

- `if`

```python
if <expression>:
	<commands>
elif <expression>:
	<commands>
elif <expression>:
	<commands>
else:
	<commands>
```

```python
if <expression>:
	<commands>
```

- `for`

```python
for <interaction_variable> in <expression>:
    <commands>
```

```python
for <interaction_variable> in <expression>:
    <commands>
else:
    <commands>
```

- `while`

```python
while <expression>:
	<commands>
```

```python
while <exoression>:
    <commands>
else:
    <commands>
```

OBS: a cláusula especial `else` para o *loop* `for` e para o *loop* `while` sempre será executada quando a condição do *loop* for falsa, e logo depois de executada, ele sai do laço enfim por completo. Para evitar que o fluxo caia na cláusula `else` pode utilizar-se do comando `break`.

### Declarações/Comandos *break* && *continue*

- `break`: sai do loop atual

- `continue`: pula o restante das instruções no bloco de loop atual e continua para a próxima iteração do loop

---

## Escopo de variáveis

**Variáveis locais** são declaradas dentro de sub rotinas definindo elas normalmente (como se estivesse definindo fora)

Para manipular **variáveis globais** dentro de funções pe preciso usar a instrução `global`, dessa forma:

```python
foo = "something"
def function():
	global foo
	foo += " :P"
function(); print(foo)
# something :P
```

OBS: pode-se passar para a instrução `global` mais de uma variável como exemplo:

```python
def function():
	global foo, bar, xpto
```

---

## Parâmetros e argumentos (de funções)

**Parâmetros**: São os valores que colocamos nos parênteses (*()*) das funções em suas **declarações**.

**Argumentos**: São os valores que colocamos nos parênteses (*()*) das funções en suas **chamadas**.

### Valores padrões de parâmetros

São valores *default* de parãmetros (ou seja, são definidos na declaração da função) para que caso não sejam passados como argumentos, eles ainda existam e com determinado valor padrão dentro da função.

A sua sintaxe é explicitada pelo operador de atribuição `=`:

```python
def function(first_param, second_param = "foo", third_param = 12.5):
```

Outra característica dos valores *default* é que na assinatura dos métodos, os parâmetros que tiverem valroes padrões tem que ser os últimos parâmetros.

#### *Keywords* (paralavra-chave) de argumentos

Na hora de passar argumentos na chamada da função você pode específicar o nome do parâmetros que irá receber o valor, nesse caso, você não precisa passar os argumentos nas repectivas ordens.

### Parâmetros *VarArgs*

Às vezes você pode querer definir uma função que pode ter qualquer número de parâmetros, ou seja, um número variável de argumentos, isso pode ser feito usando os asteriscos:

```python
def total(a=5, *numbers, **phonebook):
    print('a', a)

    #iterate through all the items in tuple
    for single_item in numbers:
        print('single_item', single_item)

    #iterate through all the items in dictionary
    for first_part, second_part in phonebook.items():
        print(first_part,second_part)

total(10,1,2,3,Jack=1123,John=2231,Inge=1560)
```

Retorno:

```python
a 10
single_item 1
single_item 2
single_item 3
Inge 1560
John 2231
Jack 1123
```

Como funciona:

Quando declaramos um parâmetro com asterisco como `*param`, todos os argumentos posicionais daquele ponto até o final são coletados como uma **tupla** chamada 'param'.

Da mesma forma, quando declaramos um parâmetro com estrela dupla, como `**param`, todos os argumentos de palavras-chave desse ponto até o final são coletados como um **dicionário** chamado 'param'.

---

## Pacotes

Pacotes são pastas de módulos com um arquivo `__init__.py` especial que indica ao Python que esta pasta é especial porque contém módulos Python.

Para criar um pacote chamado *world* com subpacotes *asia*, *africa*, etc. e esses subpacotes por sua vez contêm módulos como *india*, *madagascar*, etc.

A estrutura das pastas seria:

```
- <some folder present in the sys.path>/
    - world/
        - __init__.py
        - asia/
            - __init__.py
            - india/
                - __init__.py
                - foo.py
        - africa/
            - __init__.py
            - madagascar/
                - __init__.py
                - bar.py
```

Os pacotes são apenas uma conveniência para organizar os módulos hierarquicamente.

---

## Anotações aleatórias

- "*keycode*" da tecla **enter**?

- Por padrão não se utiliza "*()*" para fazer expressões singulares, por exemplo em *loops*, porém se tiver expressões mais complexas ou quiser mudar a ordem de precedência o uso é completamente válido

```
kbd:[enter]
```

## Type Union

```python
def soma(x: int, y: float | None = None) -> float:
	if isinstance(y, float | int):
		return x + y
	return x

print(soma(1, 10))
```

---

## Links de estudos

- <https://python.swaroopch.com>
