# Golang

## Escopo de Objetos

Caso queira usar o recurso de um pacote em outro, além de precisar importa-lo, caso ele não esteja exportado (o nome do recurso está com a primeira letra minúscula) não será possível nem acessa-lo.

Caso seja estrutura que tenha atributos ou métodos e seus recursos não estejam exportados, da mesma forma você terá um acesso limitado.

## Parâmetros de tipos definidos com valores default.

Caso queria que na criação de *structs* ou passagem de parâmetros de funções tenham valores *default* será precisa criar num determinado pacote um novo tipo que não seja exportável, as constantes (exportáveis) que serão os valores *default* desse novo tipo criado e um *struct* para que possa servir de ponte para acessar esses valores pois ela terá atributos que serão dos novos tipos privados.

Na instânciação dessa estrutura ou se passar essa estrutura como argumento de alguma função, na hora de popular seus atributos só conseguirá usar as constantes criadas anteriormentes pois elas terão os novos tipos privados e estarão disponíveis.

Também há a possibilidade de se criar uma interface com um ou outro tipo novo (colocando no corpo da interface somente o tipo e mantendo as constantes para conseguir usar os tipos).

## Clausula *select*

A clausúla `select` se assemelha a clausúla `switch` porém, ao invés de testar resultados *booleanos* cada "caso" da clausúla será um comado executado que dentro do seu corpo executa algum código.

## Generics

Links:

- <https://go.dev/doc/tutorial/generics>.

## Estrutura *http.ResponseWriter*

Quando a conexão é estabelecida com o `http.Handler` o *handler* fica aberto e você pode enviar quantos `http.ResponseWriter.Write()` quiser e a cada `Write()` é uma nova resposta que é enviada (em tempo real), é encerrado a conexão com a função *handler* encerrar.

---

_REFERENCELINKS_:

- [About errors](https://earthly.dev/blog/golang-errors/);
- [Printf types](https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/).
