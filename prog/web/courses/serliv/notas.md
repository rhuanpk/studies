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
