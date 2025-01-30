# Golang

## Setup's

### Inicial

1. Instalar a linguagem no sistema;

1. Mudar variável de ambiente `GOPATH`:
    1. Depois mover: `mv -v ~/{,.}go/`

1. No VSCode instalar a extensão do Go (`@id:golang.go`):
    1. Esperar a extensão carregar e executar a paleta de comandos Go para baixar as _tools_.

1. No VSCode definir nas configurações:
```json
    "go.lintTool": "golint",
    "go.lintOnSave": "file",
    "gopls": {
        "ui.semanticTokens": true
    },
    "go.toolsManagement.autoUpdate": true
```

1. No Linux, caso precise adicionar a pasta do binários do Go no `PATH`:
```sh
echo 'PATH+=":$HOME/.go/bin"' >> ~/.profile
```

### Repositórios Privados

1. Definir variável de ambiente `GOPRIVATE`:
    `go env -w GOPRIVATE='<remote>/<user>[/[*|<repo>[,...]]]'`

1. Caso queira usar SSH para manipular o _get_ das dependências execute:
    `git config --global url.'ssh://git@'.insteadOf 'https://'`

## Escopo de Objetos

Caso queira usar o recurso de um pacote em outro, além de precisar importa-lo, caso ele não esteja exportado (o nome do recurso está com a primeira letra minúscula) não será possível nem acessa-lo.

Caso seja estrutura que tenha atributos ou métodos e seus recursos não estejam exportados, da mesma forma você terá um acesso limitado.

## Parâmetros de Tipos Definidos

Caso queria que na criação de *structs* ou passagem de parâmetros de funções tenham valores *default* será precisa criar num determinado pacote um novo tipo que não seja exportável, as constantes (exportáveis) que serão os valores *default* desse novo tipo criado e um *struct* para que possa servir de ponte para acessar esses valores pois ela terá atributos que serão dos novos tipos privados.

Na instânciação dessa estrutura ou se passar essa estrutura como argumento de alguma função, na hora de popular seus atributos só conseguirá usar as constantes criadas anteriormentes pois elas terão os novos tipos privados e estarão disponíveis.

Também há a possibilidade de se criar uma interface com um ou outro tipo novo (colocando no corpo da interface somente o tipo e mantendo as constantes para conseguir usar os tipos).

## Clausula _select_

A clausúla `select` se assemelha a clausúla `switch` porém, ao invés de testar resultados *booleanos* cada "caso" da clausúla será um comado executado que dentro do seu corpo executa algum código.

## Generics

Links:

- <https://go.dev/doc/tutorial/generics>.

## Swagger

- mimetype: json,xml,plain,html,mpfd,png,jpeg,gif
- paramtype: query|path|header|body|formData|string|object
- datatype: string|integer|number|boolean|file|<struct>|[]|nil
- required: true|false

```go
// @summary     summary
// @description description
// @description description
// @tags        tag,tag
// @accept      #mimetype
// @produce     #mimetype
// @param       name #paramtype #datatype #required "description"
// @success     <code> {#paramtype} #datatype "description"
// @failure     <code> {#paramtype} #datatype "description"
// @router      /end/point/{@param} [<method>]
```

---

## Estrutura _http.ResponseWriter_

Quando a conexão é estabelecida com o `http.Handler` o *handler* fica aberto e você pode enviar quantos `http.ResponseWriter.Write()` quiser e a cada `Write()` é uma nova resposta que é enviada (em tempo real), é encerrado a conexão com a função *handler* encerrar.

---

## Comandos

Env:

- `go env`: Lista todas as variáveis de ambiente Go;
- `go env <ENV>`: Retorna o valor da variável de ambient especificada;
- `go env -w <ENV>=<value>`: Atualiza o valor de uma variável de ambiente Go e caso não exista, será criada;

Get:

- `[GOPROXY=direct] go get -u [-t all]`: Atualiza todos os pacotes do módulo.

Tidy:

- `[GIT_PROMPT=1] go mod tidy`: Atualização automática das dependencias.

---

_REFERENCELINKS_:

- [About errors](https://earthly.dev/blog/golang-errors/);
- [Printf types](https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/).
