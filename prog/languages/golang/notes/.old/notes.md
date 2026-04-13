# Golang

## Introdução

- Características:

    - Compilada.
    - Limpa (pouco verbosa).
    - Pequena.
    - Segurança (memória).
    - Moderna.
    - C-like (sintaxe).
    - Oberon structure (estrutura/filosofia).
    - Fortemente tipada.
    - OO em Go: Métodos e Interfaces (únicas adições).

- Baseadas:

    - Algol60, Pascal, C, Modula, Modula-2, Oberon, Oberon-2.

- Criação:

    - Data: Novembro de 2009.
    - Criadores: Robert Griesemer, Rob Pike e Ken Thompson.
    - First release (1.0): Março 2012
    - Open Source.

- Motivações:

    - Alto envolvimento com C++.
    - Alta complexidade e multiplas forma da resolução do problema.

- Concorrência/Paralelismo:

    - Linguagem criada pela *Google*, um dos principais pontos fortes da linguagem é que ela é focada, em concorrência/paralelismo de tarefas/threads pois atua muito bem sobre os *cores* do computador (criada depois da época das multiplos processadores).

---

## Comandos (CLI)

- `go env [<variable_name>]`: faz a listagem das variáveis de ambiente do `go` (OBS: se passado o nome da variável para o `go env` ele retorna só o conteúdo da variável).

    - GOPATH: local aonde o `go` armazena seu *workspace*.
    - GOROOT: local aonde o `go` foi instalado.

- `go mod init <project_name>`: faz com que determinada pasta passe a ser um módulo (cada **módulo** é um *workspace*).

- `go help get`: baixa a biblioteca ou retorna infos da mesma.

- `go version`: retorna a versão da linguagem

- `godoc -http=:6060`: seta que no *localhost* na porta *6060* habilita a documentação *offline*.

- `go vet </path/to/file.go>`: comando para fazer checagem de sintaxe.

- `go build </path/to/file.go>`: constrói o binários a partir do fonte passado e gera o binário na mesma pasta que pode ser chamado sem o interpretador.

- `go run </path/to/file.go>`: *builda* e executa o arquivo *go* somente em memória? (OBS: Para rodar de a função main de outro arquivo, rode o comando com `*.go`)

Exemplos de builds com mudança e sistema, arquitetura e *build tag* (esteja dentro do diretório do módulo):

- -o: muda o path de saida do arquivo compilado.

```bash
# for windows
GOOS=windows GOARCH=amd64 go build -o /path/to/output/binary_amd64.exe -tags build_tag ./main.go
GOOS=windows GOARCH=386 go build -o /path/to/output/binary_i386.exe -tags build_tag ./main.go

# for macOS
GOOS=darwin GOARCH=amd64 go build -o /path/to/output/binary_amd64_darwin -tags build_tag ./main.go
GOOS=darwin GOARCH=386 go build -o /path/to/output/binary_i386_darwin -tags build_tag ./main.go

# for linux
GOOS=linux GOARCH=amd64 go build -o /path/to/output/binary_amd64_linux -tags build_tag ./main.go
GOOS=linux GOARCH=386 go build -o /path/to/output/binary_i386_linux -tags build_tag ./main.go
```

### Instalação/Atualização

Instalar módulos externos no projeto atual:

```go
go get -u golang.org/x/lint/golint@latest
```

Instalar alguma ferramenta na pasta de binários do go:

```go
go install -v golang.org/x/lint/golint@latest
```

---

## Estrutura de pastas (*modules*)

- **pkg**: guarda as bíbliotecas externas que foram importadas nos projetos com o decorrer do tempo em qualquer projeto golang (*shared library*).

- **src**: guarda os arquivos/projetos fontes.

- **bin**: guarda os arquivos compilado/buildados/binários.

Os *modules* em Golang são pacotes de bibliotecas que podem ser importadas em outros projetos. Logo o que você faz ser um modulo é um pasta inteira, logo, todos os seus projetos que estão dentro deste pacote são importados juntos.

Os modulos sempre são importados pela *URL* que no caso o nome do pacote também será o final da *URL*:

```go
package main

// Nesse caso o pacote seria o "gin"
import "github.com/gin-gonic/gin"
// Nesse caso o pacote que importamos é só o "json" dentro do "gin"
import "github.com/gin-gonic/gin/internal/json"

func main() {
    // Acessando seus métodos com a sintaxe: <package>.{atribute|method()}
    gin.BindKey
    json.Marshal()
}
```

Para fazer com que um arquivo seja uma biblioteca, no início do arquivo deve ter o `package` declarado sendo que seu argumento é qualquer nome desejado, isso implicará basicamente em nada pois quando você importar o módulo, todos os grupos de pacotes dentro desse módulo serão importados, porém, na hora da compilação sim isso poderá fazer diferença (*tags de compilação*).

```go
package pathroot
```

Todo módulo (pasta com os arquivos *Golang*) tem que ter no mínimo um *main package* e uma *main function*.

### Criar o próprio móudulo

- Na pasta que você estiver que será o repositório local, dê o comando de inicialização de módulo apontando para o repositório remoto que hospedará o módulo:

```bash
go mod init <repo_url> # without "htps://"
```

O *Golang* então criará o arquivo de configuração do módulo.

- Depois das alteraçãoes dê o commit adicionando a tag que descreverá a versão do módulo, que por padrão sempre começa com *v* (e.g. `v0.1.0`) e então subindo para o repositório remoto.

- Para importarmos algum módulo desse tipo que seja externo (ou seja, não built-in do próprio *Golang*), rode o comado dentro da pasta que terá o projeto:

```bash
go get <domain>/<user>/<repo>@<version> // <version>: pode ser *latest* ou se preferir, não passe esse argumento.
```

OBS:

- caso precisa fazer o importe de um módulo que não seja diretamente o repo, terá que fazer o caminho completo do repo: `go get <domain>/<user>/<repo>/<module>@<version>`.

- caso o `go get` comando não esteja pegando o último commit, pode ser porque o *proxy* não o tenha reconhecido ainda, então: `GOPROXY=direct go get -u`.

Depois você consegue fazer o import desse módulo dentro do projeto com:

```go
import <module_repo_url>@<version>
```

OBS: caso você queira num projeto local importar um módulo externo, terá que iniciar um módulo no diretório do projeto com `go mod init <package_name>`, ou seja, ao ínves da *URL*, passe literalmente o nome do *root folder path*, depois, carregue os módulos importados com o comando `go mod tidy`.

### Workspaces

```
 workspace/ -> <workspace>: go work init ./module_1/; go work use ./module_2/
|
 \
  \__ module_1/ -> <module>: go mod init ./module_1/
  |
  |
  |__ module_2/ -> <module>: go mod init ./module_2/
```

Dessa forma você não precisa entrar diretamente na pasta que tem o `go.mod` *file* para que o go entendenda a hierarquia e também você conseguirá usar os packages de módulos locais, que estão nesse mesmo workspace.

### Nomenclatura dos membros de um módulo

- Os nomes que começarem com letra **maiúscula**, serão "exportados", públicos, ou seja, podem ser utilizado em outros projetos fazendo o importe do múdulo.
- Membros que iniciam com a letra **minúscula** serão "não exportados", privados, ou seja, mesmo que você importe o pacote, você não consegue acessar o membro.

### Documentação do módulo

- Para pacote:

Colocamos o comentário em cima da declaração `package`.

- Para método:

Colocamos o comentário em cima da declaração `func`.

- Para atributo:

Colocamos o comentário em cima da declaração `var`.

### Build tags

Você pode ter vários arquivos com métodos `main()` do mesmo módulo se eles tiverem as *build tags*. Fazendo isto, na hora da compilação, terá que informar para qual *build tag* quer fazer a compilação caso todos arquivos tenham *build tags* (se não, o arquivo que não tiver *build tag* e tiver o método `main()` será invocado).

Na sintaxe da linguagem, ela deve ser declarada na primeira linha do arquivo:

```go
//go:build tag_name
```

Para fazer o build de um arquivo específico do pacote que tenha as build tags, terá então que utilizar a seguinte flag:

```bash
go build -o /path/to/save/binary -tags tag_name ./
```

---

## Comentários

- *Single line*:

```
// <comment>
```

- *Block line*:

```
/*
    <comments>
*/
```

---

## Aspas

- Duplas: `""`: Aspas duplas são usadas para `strings`, ou seja, textos literáis.
- Simples: `''`:

---

## Ponto e vírgula

O uso do `;` em *Golang* não é obrigatório, o compilador identifica uma nova linha lógica apenas com a *new line* (\n).

---

## Estrutra básica mínima

1. Todo programa `go` roda dentro de um *main package*:

```go
package main
```

2. Todo programa `go` precisa de uma *main function* para rodar sobre (C-like):

```go
func main () {
    // <commands>
}
```

---

## Constantes/Variaveis

### Tipagem

- String (array de caracteres (*standard indexing*))

```go
var any_float string
```

OBS: strings com tamanho 0 são strings vazias como: `""` e é ilegal acessar o endereço de uma string: `&"any_string"[<index>]`.

- Inteiro

```go
var any_float int64
```

- Inteiro (unsigned)

```go
var any_float uint64
```

- Float

```go
var any_float float64
```

- Booleano

```go
var any_bool bool
```

- Unicode

```go
var any_unicode rune = "a"
```

### Auto atribuição

```go
int any_int // = 0

float any_float // = 0.0

bool any_bool // = False

string any_string // = ""

pointer any_pointer // = nil
```

### Declaração

```go
// Primeiro define a variável:
var any_var string

// Segundo inicializa ela com algum valor:
any_var = "hello world!"

// Definir e iniciar a variável com a sintaxe clássica
var xpto string = "hello world!"

// Pode-se fazer os dois com a seguinte sintaxe:
other_var := "hello world!"
```

### Exemplos

```go
// Import também suporta estrutura de bloco e esquema de alias.
import (
    "fmt"
    m "math"
)

// Constante (tipo declarado de forma explicita).
const pi float64 = 3.14

// Variável (tipo (float64) inferido pelo compilador).
var raio = 5.29

// Criar variável já inicializando e logo depois utilizando a mesma.
area := pi * m.Pow(raio, 2)
fmt.Println("A área da circunferência é", area)

// Constantes e Variáveis podem ser criadas dentro de blocos também:
const (
    a = 1.0
    b = 2.0
)
var (
    c = 3
    d = 4
)
fmt.Println(a, b, c, d)

// Declaração múltipla em única linha fazendo o recebimento de forma respectiva:
var e, f bool = true, false
fmt.Println(e, f)

g, h, i := 2, false, "epa!"
fmt.Println(g, h, i)

/*
    Formas de printar em várias linhas:
        - Múltiplos `Println()`.
        - `Println()` com \n`.
        - Bloco de string.
*/

// Múltiplos `Println()`:
fmt.Println("I'm leaning Golang!")
fmt.Println("I'm leaning Golang!")

// `Println()` com \n`:
fmt.Println("I'm leaning\nGolang!")

// Bloco de string:
any_string := `
    Hello World!
    Olá mundo...
`
fmt.Println(any_string)

/*
    Formas de printar/pegar um carater específico da string:
        - convertendo com `string()`
        - percorrendo com `[n:m]`
        - formatando com `fmt.Printf()`
*/

// A partir da posição 0 quero 1 caracter:
fmt.Println("any_string"[0:1])

// Converter o código ASCII da posição 0 em caracter:
fmt.Println(string("any_string"[0]))

// Formatar o código ASCII da posição 0 em caracter (%d para pegar o código ASCII):
fmt.Printf("%c", "any_string"[0])
```

OBS: variáveis não utilizadas geram erro de compilação (exceção verificada?).

### Tipos de dados especiais

- `nil`: tipo de dado **null**

- `1`: qualquer retorno diferente de 1 é **false**

### Casting

*float*:

```go
float64(any)
```

*int*:

```go
// Ele não arredonda, pega só a parte inteira
int(any)
```

*string*:

```go
// Passe o código ascii para retornar o caracter
string(97)
```

*int para string*:

```go
strconv.Itoa(97)
```

*string para int*:

```go
strconv.Atoi(97)
```

*string para booleano*:

```go
strconv.ParseBool("true")
```

### Globais e locais

Exemplo prático:

```go
package main

import "fmt"

func main() {  
    
    any_int := 1
    
    fmt.Println(any_int)     // prints 1
    
    {
        fmt.Println(any_int) // prints 1
        any_int := 2
        fmt.Println(any_int) // prints 2
    }
    
    fmt.Println(any_int)     // prints 1 
}
```

---

## Funções

Sintaxe:

```go
func <name>(<param_name> <param_type>, ...) <return_type> {
    return a + b
}
```

Exemplo:

```go
func somar(a int, b int) int {
    return a + b
}
```

---

## Controle de Fluxo

### *if/else*

Características:

- Na sua expressão não aceita parênteses
- Mesmo blocos de linha única precisam da abertura e fechamento de chaves

Sintaxe:

```go
if <expression> {
    <commands>
} else if <expression> {
    <commands>
} else {
    <commands>
}
```

### *for*

Sintaxe:

```go
for [<expression>] {
    <commands>
}
```

### *switch*

Sintaxe:

```go
switch [<expression>] {
    case <expression>:
        <commands>
    ...
    default:
        <commands>
}
```

---

## Estruturas de Dados

### *Arrays*

Estrutura homogênea (mesmo tipo), estática (fixo) e indexada (índices numéricos que começam por 0).

Declarando e inicializando com tamanho explícito e em momento separados:

```go
var any_array [2]string
xpto[0], xpto[1] = "hello", "world"
fmt.Println(xpto)
```

Declarando e inicializando de forma reduzida com tamanho explícito:

```go
xpto := [5] int {1, 2, 3, 4, 5}
fmt.Println(xpto)
```

Declarando e inicializando de forma reduzida com tamanho implícito:

```go
xpto := [...] int {1, 2, 3, 4, 5}
fmt.Println(xpto)
```

Declarando e inicializando e especificando qual índice irá receber:

```go
var xpto [5] string = [5] string {3:"rhuan", 4:"patriky"}
fmt.Println(xpto)
```

### *Slice*

Essa estrutura é basicamente igual a um `array` porém, ela apenas faz referência na memória.

```go
// Pegue do índice 1 até o índice 3.
any_slice := any_array [1:3]
```

```go
// Cria um slice a partir de um "array virtual" (que não existe).
any_slice = make([]string, 10, 20)
```

#### Atributos/Métodos

- `append()`: appenda todos os argumento passados (a partir do segundo) no slice passado como primeiro argumento.

### *Map*

Estrutura de dados igual a um array associativo que deve ser inicializado e não declarado.

- Declarar já passando os valores:

```go
any_map := map[string]float64{
    "first": 123.45,
    "second": 678.90,
    "third": 159.32,
}
```

- Meio de somente declarar sem passar já os valores de início, isso porque o make aloca memória:

```go
any_map := make(map[int]string)
```

- Deletar um índice do mapa:

```go
delete(map_name, index)
```

---

## Bibliotecas

### *built-in*

Objetos padrão que já são compilados da linguagem.

#### Atributos/Métodos

- `len()`: retorna o tamanho do objeto passada como argumento.

- `cap()`: retorna a capacidade do objeto passado como argumento (em alguns casos será diferente de `len()`).

- `defer <command>`: *defer command* executa o que é passado para ele quando a função for encerrada e caso queira declarar mais de uma ação a ser executada quando a função encerrar, precisa-se definir de um defer para cada ação, todos os defer's declarados são executados do último declarado para o primeiro.

### *fmt*

Biblioteca *formatar* que é usadas para finalizadades de exibição e formatação da saida.

#### Atributos/Métodos

- `Print()`: printa na STDOUT os argumentos passados separados por vírgula (o sinal de "+" pode ser usado para concatenar valores).

- `Println()`: igual a função `Print()` porém, adicioanando uma *new line* no final.

- `Sprint()`: retorna o valor do argumento passado como string (*casting* ?).

- `Printf()`: printf c-like: os "argumentos de troca" são trocado pelos respectivos valores na sequência.

    - `%s`: string.
    - `%t`: booleano.
    - `%d`: inteiro.
    - `%f`: float arredondando.
        - `%.2f`: float limitado a 2 casas decimáis.
    - `%v`: genérico, serve para vários tipos diferentes.

### *math*

Biblioteca *matemática* que tem métodos para funções matemáticas.

#### Atributos/Métodos

- `MaxInt64`: (#***static***) guarda o valor máximo de um inteiro de 64bits.

- `Pow()`: (#***static***) retorna a potência da expressão passado como primeiro argumento sob a segunda expressão passada como argumento.

- `Max()`: retorna o maior valor entre os dois argumentos passados.

- `Min()`: retorna o menor valor entre os dois argumentos passados.

### *reflect*

Biblioteca de *reflexão* para manipulação de tipos de objetos.

#### Atributos/Métodos

- `TypeOf()`: (#***static***) retorna o tipo de dado passado como argumento.

### *time*

Biblioteca para fazer manipulações de *data*.

#### Atributos/Métodos

- `Second`: retorna o valor de `1s`.

- `Unix()`: retorna a data com base nos parâmetros passados.

- `Equal()`: retorna boleano comparando as duas datas passadas como parâmetro.

- `Sleep()`: espera o tempo em segundos passado como argumento.

- `Now()`: retorna todas as informações a respeito da data e hora do momento que o código é executado.

    - `Hour()`: retorna somente a parte da hora.

### *string*

Biblioteca para fazer manipulação de *string*.

#### Atributos/Métodos

- `Join()`: retona a concatenação da primeira string passada como argumento com a segunda passada como argumento.

### *strings*

Biblioteca para fazer manipulações de string's.

#### Atributos/Métodos

- `Replace(<variable>, "\n", "", -1)`: em cima do conteúdo passado como primeiro argumento da *replace* no segundo argumento com o terceiro argumento na quantidade de vezes passada no quarto argumento (caso seja -1, será todas as ocorrências).

### *os*

Biblioteca para fazer a manipulação de elemnto do SO (*operacional system*)

#### Atributos/Métodos

- `Exit(<exit_code>)`: Encerra o programa retornando o código de saida passado como argumento.

- `Create("/path/to/file.any")`: Cria um novo arquivo no caminho passado. Seu primeiro valor de retorno é a instãncia de um `*os.File` e o segundo um `error`.

- `Open("/path/to/file")`: Abre o arquivo passado como argumento para leitura caso ele exista e retorna uma instância de `*os.File`.

- `(file *File)`:
    
    - `WriteString()`: Escreve no arquivo aberto.
    
    - `Close()`: Fecha o arquivo criado.

- `Stat("/path/to/file.any")`: verificar se o arquivo passado como argumento existe, caso não existe, o segundo retorno da função é um `error`.

- `IsNotExist(<error_variable>)`: passe uma variável que pode conter um `error` para quando se verifica a existência de um arquivo ou diretório. Caso a varável de possível `error` esteja populada com o erro esperado, retorna `bool` `false`.

- `Rename("/path/to/old/file.any", "/path/to/new/file.any")`: renomeia o arquivo passado como primeiro argumento pelo arquivo passado como segundo argumento.

### *bufio*

Biblioteca usada para fazer manipulação de de I/O.

#### Atributos/Métodos

- `NewReader(os.Stdin)`: retorna a instância de objeto para manipulação da STDIN.

    - `read.ReadString('\n')`: chama um campo de input que lerá da STDIN, retorna o input até o delimitador (incluindo o delimitador) passado como argumento e caso aconteça algum erro na leitura, retorna o input até o erro e também retorna um `error` como segundo output.

- `NewScanner(<*os.File_variable>)`: retonar um `*bufio.Scanner` (estrutura para manipulação de leitura de arquivo?).

- `(scanner *bufio.Scanner)`:

    - `Scan()`: percorre todas as linhas do objeto instânciado.

    - `Text()`: tem o conteúdo da linha atual que o `bufio.Scanner.Scan()` está.

### *io/ioutil*

Biblioteca usada para fazer manipulação de de I/O.

#### Atributos/Métodos

- `ReadFile()`: tenta ler o arquivo e o primeiro retorno é o conteúdo desse arquivo em bytes e o segundo um `error` para caso a operação falhe (se a leitura for bem sucedida, basta converter os `bytes` em `string` com `string(<return_variable>)`).

---

## Ponteiros

- `&`: o **&** notação retorna o endereço na memória de algum objeto.
- `*`: o **\*** notação diz que a variável em questão é um ponteiro.

Printar o endereço de memória com `fmt.Printf()`:

```go
fmt.Printf("%x", &any_object)
```

Example:

```go
// variable
var any_int int = 5

// create pointer
var int_pointer *int

// store the address of any_int in pointer variable
int_pointer = &any_int

// display info
fmt.Printf("Memory address of variable any_int: %x\n", &any_int)
fmt.Printf("Memory address stored in int_pointer variable: %x\n", int_pointer)
fmt.Printf("Contents of *int_pointer variable: %d\n", *int_pointer)
```

---

## Operadores ternários?

Possibilidades:

- mapa ternário:

```go
any_var := map[bool]string{true: "true", false: "false"} [1 > 0]
println(any_var)
```

- função anônima com `if`:

```go
any_var := func() string {if 1 > 0 {return "true"} else {return "false"}}()
println(any_var)
```

---

## Anotações diversas

### HTTP

Toda a parte que de alguma forma envolve protocolo HTTP, o Golang tem um pacote *built-in* que trás ferramentas para
fazer essa manipulação.

- Por exemplo, se fizermos um método para sarvir como API, podemos fazer com que o programa que executamos utilize
um método do pacote `http` que abre um determinada porta no *SO* e a fica escutando (`http.ListenAndServe()`)
e caso alguma informação entre nessa porta, ele analiza para saber se é uma requisição com determinado corpo,
caso tenha determinado parâmetro nesse corpo, ele procura num mapa de funções interno do pacote o valor desse
parâmetro, caso encontre, ele a executa.

---

## Links

- [Cooder (GitHub)](https://github.com/cod3rcursos/curso-go/)
- [Operadores Lógicos](https://dicasdeprogramacao.com.br/operadores-logicos/)
- [Go Mod](https://aprendagolang.com.br/2022/06/17/entenda-o-que-sao-os-arquivos-go-mod-e-go-sum/)
- [Teste local de API](https://aprendagolang.com.br/2021/11/08/como-testar-apis-e-integracoes-localmente/)
- [CRUD](https://www.youtube.com/watch?v=vxDqv6BKZCw&list=PLm-xZWCprwYQ3gyCxJ8TR1L2ZnUOPvOpr)

---

## Code Runner

Executador de programas em várias linguagens:

- Executar o código:
    `ctrl+al`
