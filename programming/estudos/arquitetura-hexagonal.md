# Arquitetura Hexagonal

- **adapter**: Responsável somente por entrada e saída de dados (métodos que recebem os dados e os que devolvem?):

	- **input**: Sub camada do **adapter** referente ao que diz respeito a entrada de dados:

		- **controller**: ?

			- **middlewares**: Como os _middlwares_ estão dentro dos _controllers_ e validam ou tratam a entrada dos dados, faz sentido serem _input adapter's_;

			- **routes**: Tem a declaração das rotas para acessar a minha aplicação, chama os _controllers_ que estão um nível a cima (falando dos arquivos que estão dentro da pasta de **routes** para os arquivos que estão no mesmo nível da pasta) e que usa os **middlewares** que está no mesmo nível que ele;

			- `<c|r|u|d>_<resource>_controller`.go: São quem vão receber receber os dados, valida-los (validar tipos de dados, etc) e então chamar o **service** que por sua vez chama o **repository** (nesse arquivo também é definida a _interface_ do resource?).

		- **converter**: ?

		- **model**: Mais outro nível de sub pasta para organização:

			- **request**: Dentro desta vai o model (**struct**) referente a essa camada, por exemplo, se é o **adapter** de entrada, então é a estrutura com as tags do _json_ por exemplo;

			- **response**: Vale o mesmo para a explicação de cima.

		- OBS: antes de ir diretamente para esse nível de pasta, caso você tenha _inputs_ para **kafka**, **sqs** e **http** por exemplo, você criaria essas como sub pastas primeiro doque esse nível (esse do exemplo).

	- **output**:

		- **repository**: Funções que se conectaram com o banco de dados e executaram as ações necessárias (são chamados pelos **service's** e devolvem para eles também);

		- **model**: Aqui vai os modelos (**structs**) para as entidades (DTO's para comunicação com o banco? que é o que as função do **repository** desse escopo usa):

			- **Entity**: Os arquivos que teram aqui dentro seguem a regra do **model** dos _input adapter's_.

- **application**: Responsável pelas regras de negócio, código que contém a lógica:

	- **constants**: ?

	- **domain**: Estruturas puras que servem somente para transitar os dados entre as camadas (DTO's?). No caso.

	- **port**: Interfaces?

		- **input**: Interfaces de entrada?

		- **output**: Interfaces de saída?

	- **services**: Códigos com as regras de negócio e que chamam os **repository's**. Dentro do _service_ também terá validações mas, validações referentes ao banco de dados (verificar se o e-mail já é existente, criptografar a senha e etc) e por vim chama o método do **repository** para salvar o recurso (ou outras operações que está fazendo).

- **configuration**: Responsável pelas configurações.

OBS: nenhuma dessas camadas/domínios devem iniciar a aplicação (no máximo configurar rotas da aplicação no **adapter** e, mesmo que configuradas no **adapter**, que sejam carregadas na função `main()`).

### In Code

- A função `main()`: Deve carregar tudo que precisa ser carregado, variáveis de ambiente, dependências, conexão com banco de dados, interfaces, _logger_ e etc.
