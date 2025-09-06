# Serliv

Curso Web Frontend Fundamentos (HTML, CSS & JS).

## História

No princípio a internet foi criada com intúito de compartilhar artigos acadêmicos via Hipertexto (conexão entre documentos).

As empresas viram o potêncial dessa nova tecnologia e a utilizarem para outros fins, o que fez ser impossível ter manutenibilidade.

Além de todo o _workaround_ necessário para criar as páginas _web_ da época, um agravante era a briga entre os _browser_ da eṕoca: cada um tinha sua própria maneira de interpretar HTML. Isso fazia com que fosse necessário ter um cada arquivo HTML para cada _browser_.

Para resolver esse conflito, a W3C foi fundada a fim de implementar padrões para a _web_.

## WWW

A _World Wide Web_ é todo o conteúdo acessível na internet.

Nela existe a comunicação **Cliente X Servidor**: O **cliente solicita** alguma informação e o **servidor retorna** a informação solicitada. Um exemplo de cliente seria um _browser_ (Chrome) e um exemplo de servidor seria um _web server_ (Nginx).

Para que essa comunicação seja possível, é necessário protocolos como o HTTPS para que independente do Sistema Operacional do cliente e do servidor, as duas partes possam se entender.

## HTML

O _HyperText Markup Language_ é uma Linguagem de Marcação, um arquivo de texto comum com marcações que dão significa ao conteúdo.

Sintaxe:
- _Tag_ com fechamento: `<tag>` & `</tag>`
- _Tag_ sem fechamento: `<tag>`

### Tags

_Tags_ são os elementos/marcações do HTML.

As _tags_ podem ser aninhadas, ou seja, a primeira _tag_ que abre é a última que fecha.

### Declarações

- `doctype`: Informa ao _user agent_ o tipo de documento, se "html": HTML 5

### Atributos

Atrivutos são informações (metadados) de _tags_.

- `href`
	- `https://domain.com`: URL externa
	- `index.html`: Outra página do site
	- `#section`: Sessão dentro da mesma página
	- `mailto:user@domain.com`: Redirecionamento para o cliente de e-mail

### Semântica

Usamos _tags_ que dão importância ao conteúdo quando necessário.

Algumas _tags_ apenas mudam o estilo, outras, agregão carga semântica ao conteúdo. Isso é importante tanto para SEO (_Search Engine Optimization_) quanto para acessibilidade.

## CSS

O _Cascating Style Sheets_ é uma Linguagem de Estilização. Um arquivo com descrições de formatação dos elementos HTML.

Sintaxe: `selector { property: value; }`.

### Inclusão

Há algumas maneiras de incluir o CSS.

- _Inline_: Atributo de uma tag HTML: `<tag style="...">` & `</tag>`
- _Internal_: `<style>` & `</style>`
- _External_: `<link href="*.css" rel="stylesheet">`

### Box Model

Cada elemento do HTML se comporta como uma caixa.

_Sizing_: propriedade que define dimensões do elemento:
- `content-box`: As propriedades `padding` e `border` somam a largura (_width_) e altura (_height_) final do elemento (_default_)
- `border-box`: As propriedades `padding` e `border` NÃO somam a largura, ou seja, são imbutidas, na largura (_width_) e altura (_height_) final do elemento

_Display_: propriedade que define o tipo (comportamento) do elemento:
- `block`: _Width_ de 100% da área (_container_) e quebra de linha
- `inline`: _Width_ do tamanho do conteúdo e NÃO quebra linha
- `inlin-block`: Elemento continua como de linha, porém, com propriedades de elementos de bloco
- `none`: O elemento é removido (não somente escondido)
- `flex`: O elemento pai se torna o _flex container_ e seu filhos diretos _flex items_

_Margin_: propriedade que define o espaço externo de um elemento referente ao seu vizinho:
- `auto`: Empurra todo o espaço disponível de forma proporcional. Caso esteja esteja espaçando somente de um lado, empurrará todo o espaço disponível, porém, caso esteja empurrado dois lados opostos, distribuirá igual entre ambos.

_Width_: propriedade que define a largura do elemento:
- `fit-content`: A largura do elemento se ajustará automáticamente referente ao tamanho/quantidade do conteúdo dentro do elemento
- `min-content`: A largura do elemento será equivalente ao **menor** conteúdo dentro do elemento
- `max-content`: A largura do elemento será equivalente ao **maior** conteúdo dentro do elemento

_Position_:
- `static`: A posição que os elementos são renderizados é a ordem que está no HTML (_default_)
- `relative`: A posição é relativa ao próprio elemento, a posição original do elemento é preservada e apenas sua visualização será afetada
- `absolute`: A posição do elemento é relacionada ao primeiro elemento pai posicionado
	- Um elemento posicionado é aquele que seu valor `position` é diferente `static`
	- Caso precisemos de um `position` "_fake_" no elemento pai do elemento que queremos manipular, podemos defini-lo apenas com `relative` sem passar nenhum dos valores `top`, `right`, `bottom` ou `left`
- `fixed`: A posição do elemento é relacionado a _viewport_
- `sticky`: É um híbrido entre `relative` e `fixed`, seu pai direto precisa ter algum mecanismo de _scroll_

### Float

Criado originalmente para colocar imagens ao lado de conteúdos, não sessões inteiras ao lado de outras, para isso, temos _CSS Grid_ ou _Flex Box_.

A propriedade `float` faz com que o elemento saia do fluxo natural do programa e os elementos posteriores tentam agora ocupar o espaço que foi vago pelo elemento que "saiu".

A propriedade `clear` limpa o `float` no elemento anterior, no caso, não a posição em que o elmento anterior com `float` se encontra, mas é como apesar dele ter flutado para a direita ou esquerda, o elemento pai continua enxergando o elemento com `float` e o próprio elemento com `clear` (o elemento posterior ao `float`) não tenta se encaixar pois ainda continua enxergando o elemento aterior.

A propriedade `overflow` paesar de não ter sido criado para esse fim, é usada no elemento pai quando os filhos tem `float` para que passem a serem enxergados novamente pelo pai para que ele possa determinar sua própria altura "corretamente".

### Overflow

A propriedade `overflow` foi originalmente pensada para trabalhar com _containers_ que possuam altura (_height_) definidas.

Essa propriedade serve para que possamos tratar o conteúdo que excede o tamanho do nosso elemento pois, caso o tamanho do conteúdo seja maior que o tamanho estipulado do próprio elemento, o conteúdo vazará.

### Z Index

A propriedade `z-index` serve para definir a sobreposição dos elementos do ponto de vista do usuário. O valor _default_ é **0** e quanto maior o seu valor, maior a sua prioridade na sobreposição. Valores negativos também são aceitos.

Dentro dessa propriedade também temos os contextos de empilhamentos, que são resetadas a cada novo _container_ (pai). Ou seja, caso tenhamos dois elementos pais posicionados que possuem filhos e um desses pais tem `z-index` superior ao outro pai, ele e todos os seus filhos se sobreporão, e mesmo que um elemento filho do pai que está sobreposto ter seu `z-index` maior ainda do que o próprio elemento pai que está se sobreponto ou qualquer um dos seus filhos, nada mudará, este filho se sobreporá apenas sobre os seus irmãos.

### Flexbox

O valor `flex` da propriedade `display` faz com que o elemento pai se torna um _flex container_ e seus filhos diretos _flex items_, ou seja, o que é afetado é apenas os elementos filhos.

Quando essa propriedade é aplicada, os elementos filhos dentro do _container_ passam a **não quebrar linha** e ficam **alinhados na horizontal**. Suas **alturas ocupam 100% do _container_** e suas **larguras são distribuídas igualmente** pela quantidade de filhos.

- `flex-wrap`
	- `nowrap`: Todos os _flex items_ são alinhados na horizontal (_default_)
	- `wrap`: Caso a soma da largura dos _flex items_ na linha seja maior que a largura do _flex container_, quebra a linha e os _flex items_ excedentes vão para a linha de baixo e assim sucessivamente

Os eixos do Flexbox (_Flexbox Axes_) são determinados pela propriedade `flex-direction` que impactará em todas as outras propriedades (pois usam os exios para seguiar).

A propriedade `flex-grow` e `flex-shrink` distribuem `width` ou `height` (com base no _main axis_) para os _flex items_. A distribuição é feita por partes/unidades de `flex-grow`: cada valor de `flex-grow` definido nos _flex items_ são somados para saber a proporção total que irá para cada.

A propriedade `order` altera a ordem dos _flex items_ dentro do _flex container_, parecido com o princípio da propriedade `z-index`. Valores negativos tem prioridade sobre positivos.

### Imagens

- `background-attachment`
	- `scroll`: A imagem é "scrollada" junto com a página (_default_)
	- `fixed`: A imagem NÃO é "scrollada" junto com a página
- `background-repeat`
	- `no-repeat`: A imagem NÃO se repete
	- `repeat`: A imagem se repete sempre
- `background-size`
	- `contain`: Size 100% mantendo a proporção da imagem (não corta a imagem e a repete)
	- `cover`: Size 100% NÃO mantendo a proporção da imagem (corta a imagem e NÃO a repete)
- `background-position`
	- `right bottom`: Fixa a imagem repectivamente: horizontal, vertical

### Unidades

Unidades de medidas utilizadas pelas propriedades.

#### `px`

Unidade de medida em _pixels_. Não é interessante que seja utilizada em fontes pois pode sobrescrever preferências do SO, por exemplo.

Normalmente não é utilizado para:
- `width`
- `height`
- `font-size`

Normalmente poderia ser utilizado para:
- `margin`
- `padding`
- `text-shadow`
- `box-shadow`

**1 px** é igual a **1 pixel** da resolução da tela.

#### `em`

Unidade de medida **relativa** ao `font-size` (_default_) do próprio elemento, ou seja, se uma _tag_ `p` tem seu `font-size` padrão de `16px`, se definimo-os para `2em`, agora seu `font-size` será de `32px`.

Caso o **elemento pai** tenha um `font-size` definido, então a unidade de medida `em` será relativa a este valor, e não a do próprio elemento.

#### `rem`

Unidade de medida **relativa** ao `font-size` do elemento _root_. Seu funcionamento é igual ao da unidade de medida `em`.

#### `ch`

Unidade de medida **relativa** a largura do caractere "0" (zero). Se dizemos que a largura de um _container_ é de `100ch`, a largura desse _container_ será equivalente a de 100 zeros em sequência.

É como dizer que a largura é de "100 caracteres" (no caso do "0" porque é o maior caractere possível em largura?).

### Tips

Posicionamento: conceito de que cada propriedade se relaciona com o elemento pai ou vizinho:
- `width`: Referente ao elemento **pai**
- `height`: Referente ao elemento **pai**

- `margin`: Referente ao elemento **vizinho**
- `border`: Referente ao próprio **elemento**
- `padding`: Referente ao próprio **elemento**

- `block`: Referente ao elemento **pai**
- `inline`: Referente ao **conteúdo**
- `inlin-block`: Referente ao **conteúdo**

Alinhamento:
- `margin: auto;`: Centralizamos o próprios **elemento** (_container_)
- `text-align: center;` Centralizamos o **conteúdo** dentro do elemento (_container_)

Eixos:
- `x`: Eixo **x** significa a **horizontal**
- `y`: Eixo **y** significa a **vertical**

Margin: quando uma "conflita" com a outra, elas se sobrepõem. Digamos que temos um elemento e seu irmão em baixo, caso o irmão de cima tenha uma `margin-bottom` de `25px` e o irmão de baixo uma `margin-top` de também `25px`, a distância entre os dois elementos será de `25px`, não de `50px` pois, a maior margin que sempre irá se sobrepor/empurrar.

Unidades: unidades de medidas usadas pelas propriedades:
- `px`: sempre use **números pares**, **múltiplos de 4**
- `%`: para `height` só funcionará (como o esperado) caso o elemento pai tenha um `height` definido
- `%`: para `margin` e `padding`, a porcentagem será sempre referente ao `width` do **elemento pai**

---

## Misc

Conteúdo complementar referente a este.

### Requisição

O que acontece quando entramos num site pela _web_?

O _browser_ requisita a página HTML para o servidor:
```
Site --request--> ISP --request--> Servidor
Site <---hrml---- ISP <---html---- Servidor
```

Para o _browser_ requisitar o site corretamente o nome de domínio deve ser traduzido num endereço IP:
```
Domain <--ip--> DNS
```

### Imagens

Quais são os tipos de imagens mais comuns?

- _Raster_ (Rasterizada): A imagem é montada por _pixels_ e cada pixels tem uma tonalidade de cor, o que faz com que caso a imagem seja expandida, sua qualidade seja comprometida. Também é conhecida por _Bitmap_
- _Vector_ (Vetorial): A imagem é montada por cálculos matemáticos que dizem onde cada cor deve estar, o que faz com que caso a imagem seja expandida, sua qualidade NÃO seja comprometida
