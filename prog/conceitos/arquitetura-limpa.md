# Clean Architecture

- Independência de **framework**;

- Independência de **interface de usuário**;

- Independência de **banco de dados**;

- Regras de negócio nunca devem ser implementadas na interface do usuário (no _front_ da aplicação ?);

- Testabilidade.

`drivers -> adapters -> services -> entities`:

- Entities: As abstrações das tabelas do banco em nosso código, as entidades são geralmente classes, aonde definimos as propriedades da entity, tipos e acessos ("_getters and setters_");

- Services: Regra de negócio da aplicação (lógica), ela que se comunicara com a entidade e toda e qualquer informação passará por ela;

- Adapters: Tradução para a comunicação com os elementos externos (converter uma entidade processamento em uma de saída). Aqui costumamos ter **controllers**, **presents** e **repositories**;

- Drivers: Libs e qualquer outra coisa externa que tenha que se comunicar com o _core_ do nosso sistema.

## O Cerne do Conceito

As dependências sempre devem apontar para as camadas mais internas da aplicação. Ou seja, a camada de **entities** nunca sabera de nada a não ser sobre ela mesma (pois é a mais interna). A camada de **adapters** poderá saber sobre **services** -> **entities** mas, em contrapartida, nunca sabera sobre **drivers**.

## As Pastas na Prática

Cada entidade/contexto terá a sua própria arvore base de diretórios, por exemplo, se tivermos no nosso sistema as entidades `user`, `book` e `card`, todas elas teriam dentro de si (ou algo assim):

```
internal/
└── <resource>
    ├── adapter
    ├── controller
    ├── entity
    └── service
```