# go modules

## step-by-step

Nas pasta que será o módulo:

1. Tenha um repositório remoto criado.
1. `go mod init <domain/path>` dentro da pasta que será o módulo.
1. Commite e pushe com a tag de versão.

Nas pasta que terá o main:

1. `go get -u <domain/path>`.
1. importe o modulo dentro do código com `import <domain/path>`.

Observations:

- Preferêncialmente colocar sempre o nome do pacote como o nome da pasta pai (pasta aonde está o arquivo).
- O que é colocado no import é a URL do **módulo** porém, quando é feito a chamada de algum método desse módulo, ele passa pelo nome do **pacote** (*package*) que podem não ser o mesmo nome.
- Caso o método do pacote que está tetando acessar, não ficará disponível caso o membro esteja privado.
