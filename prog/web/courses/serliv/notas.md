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

### Atributos

Atrivutos são informações (metadados) de _tags_.

- `href`:
	- `https://domain.com`
	- `index.html`
	- `#section`
	- `mailto:user@domain.com`

### Semântica

Usamos _tags_ que dão importância ao conteúdo quando necessário.

Algumas _tags_ apenas mudam o estilo, outras, agregão carga semântica ao conteúdo. Isso é importante tanto para SEO (_Search Engine Optimization_) quanto para acessibilidade.

## CSS

O _Cascating Style Sheets_ é uma Linguagem de Estilização. Um arquivo com descrições de formatação dos elementos HTML.

Sintaxe: `selector { property: value; }`

### Inclusão

Há algumas maneiras de incluir o CSS.

- _Inline_: Atributo de uma tag HTML: `<tag style="...">` & `</tag>`
- _Internal_: `<style>` & `</style>`
- _External_: `<link href="*.css" rel="stylesheet">`

### Box Model

Cada elemento do HTML se comporta como uma caixa.

_Display_: tipo (comportamento) do elemento:
- `block`: _Width_ de 100% da área (container) e quebra de linha
- `inline`: _Width_ do tamanho do conteúdo e NÃO quebra linha
- `inlin-block`: Elemento continua como de linha, porém, com propriedades de elementos de bloco
- `none`: O elemento é removido (não somente escondido)

Posicionamento: cada propriedade se relaciona com o elemento pai ou vizinho:
- `width`: Referente ao elemento **pai**

_Sizing_: definição de dimensões do elemento:
- `content-box`: As propriedades `padding` e `border` somam a largura (_width_) e altura (_height_) final do elemento (_default_)
- `border-box`: As propriedades `padding` e `border` são imbutidas na largura (_width_) e altura (_height_) final do elemento (_default_)
