# Anotações (gerais) Sobre Programação

## Tipagens

### Forte e Fraca

Forte:

- O compilador/interpretador **NÃO faz conversão implícita** de tipos primitivos para operações possíveis.

Fraca:

- O compilador/interpretador **faz conversão implícita** de tipos primitivos para operações possíveis.

OBS: para ambos os casos serve o exemplo: `1 + "1"`.

### Estática e Dinâmica

Estática:

- **É necessário** informar o tipo de dado da variável na sua declaração.
- **É impossível** atribuir um tipo de dado diferente do seu tipo original.

Dinâmica:

- **Não é necessário** informar o tipo de dado da variável na sua declaração.
- **É possível** atribuir um tipo de dado diferente do seu tipo original.

## Classes Abstratas X Interfaces

**classes abstratas** e **interfaces** servem para definir padrões para outros classes herdarem e implementarem.

- ambas: **NÃO** pode ser intânciadas, somente herdadas ou implementadas.

- classe abstrata: pode contem código lógico.

- interface: contem somente os _statements_/definições dos atributos/métodos.

## Assinatura de Métodos

Nas linguagens mais antigas, cada função é chamada através de seu identificador (nome). Dentro do contexto onde ela é válida o identificador da função precisa ser único para que não haja ambiguidades no momento da chamada. Já nas linguagens mais modernas, existe o conceito de assinatura de funções, ou, no caso de linguagens orientadas a objetos, assinatura de métodos. Nesse caso, não é apenas o nome que identifica um método, oque o torna um método único dentro do contexto onde é válido (sua classe) é sua assinatura. A assinatura do método é formada pelo seu nome, pelo tipo, quantidade e ordem de seus parâmetros. Desta forma, é possível existirem na mesma classe, métodos com o mesmo nome, desde que tenham listas de parâmetros diferentes. Note que o nome dos parâmetros não faz parte da assinatura, apenas o tipo, quantidade e ordem deles. Também não faz parte da assinatura o tipo de retorno do método.

Exemplo:

Declaração: public void function(String string)
Assinatura: function(String)

Declaração: public void function(int value, String string)
Assinatura: function(int, String)
