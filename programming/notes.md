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

## Classes Abstratas x Interfaces

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

## Arquitetura x Design de Software

Quando falamos a respeito desse assunto, deveras temos muitos termos as vezes para referênciar a mesma coisa. Aqui vou abordar somente os 3 três termos que estão de fato mais presentes na vida do programador que são eles:

- Arquitetura de Software/Projeto;
- Design de Software/Projeto;
- Design Patterns/Padrões de Projeto.

### Arquitetura de Software

Arquitetura de Software diz respeito a **organização/separação do _software_** e como essas partes interagem entre si. Ou seja, como você vai particionar o seu código. Fala diretamente a respeito de escalabilidade, segurança e distribuição do código.

> Arquitetura de Software: A arquitetura de software é um nível acima do design de software e envolve a organização de componentes maiores e suas interações, como módulos, camadas, serviços e subsistemas. A arquitetura lida com a estrutura geral do sistema, suas decisões de design de alto nível, escalabilidade, segurança e distribuição.

Na prática, é como geralmente organizamos as pastas do nosso código e pensamos na estrura dele (na parte física) como um todo. Agluns exemplos de Padrões Arquiteturais são:

- MVC;
- Hexagonal;
- MVVM;
- Layers;
- Singleton.

### Design de Software

Dentro da área do Design de Software há outras vertentes que é o que geralmente estamos procurando, por exemplo, os Design Patterns, ou seja, na verdade Design Patterns FAZ PARTE do Design de Software. O Design de Software de fato se preocupado mais com o "micro" do que com o "macro" que é a função da Arquitetura de Software.

Na prática, é quando pensamos em como escrever o nosso código, o nome que damos as variáveis, funções e estruturas daddos. Se vamos aplicar algum padrão de boas práticas de programação ou algum padrão de projeto (Design Pattern), como exemplo:

- Clean Code;
- SOLID;
- Factory's;
- Builder;
- Obverver.
