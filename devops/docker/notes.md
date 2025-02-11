# Docker Notes

## Commands

### `up`

- `--build`: _rebuilda_ a imagem do _container_, quando o nosso serviço roda em cima de uma imagem que construímos por conta própria (a partir de um Dockerfile), sempre que qualquer coisa que faça parte do _build_ da imagem (ou seja, os recursos descritos dentro do Dockerfile) for alterado, precisamos usar `--build`, porque, por mais que reiniciamos/recriamos o container, ele usará a imagem já existente, que não contem as novas alterações.

- `--force-recreate`: recria o _container_ do zero (como se tivesse sendo criando pela primeira vez), quando usamos o comando `docker compose up`, por padrão, o Docker tenta reutilizar os _containers_ existentes se não houver alterações nos serviços ou nas imagens, no entanto, pode haver situações em que desejamos garantir que os containers sejam sempre recriados do zero, por exemplo, para garantir que os volumes sejam remontados, variáveis de ambiente sejam aplicadas novamente, ou simplesmente para garantir um estado limpo do container.

### `inspect`

Verificar _ip's_ do container:
```sh
docker container inspect <name/id> | grep -i 'ipaddress'
```

Verificar volumes do container:
```sh
docker container inspect <name/id> | grep -iA 11 'mounts'
```

### `volume`

Criar volume nomeado:
```sh
docker volume create <name>
```

## Volumes

> Volumes nada mais são do que "pastas dos containers" que salvamos ou mapeamos?

Basicamente existem dois tipos de volumes:

- _data volumes_ (volumes nomeados): volumes criados e controlados pelo _docker host_ (pelo próprio docker, seu _daemon_) que não tem ligação com a MÁQUINA _host_;
	- OBS: é como se fossem simplesmentes um ponto de montagem do próprio container?

- _bind volumes_ (volumes mapeados/conectados): volumes que mapeia algum diretório ou arquivo da máquina _host_ para dentro do container;

- _tmpfs volumes_ (volumes temporários): volumes que persitem dados em memória simplesmente (!).

### Data Volumes

Caso você crie um container no qual a sua imagem já cria algum volume por padrão, esse volume é criado com algum nome aleatório (`docker volume ls`).

Para uma melhor manutenção, é recomendado que se crie um volume manualmente (volume nomeado) e utilize esse volume nos seus containers para a persistencia, dessa forma garantindo até a reutilização dos dados gravados dos containers anteriores.

OBS: caso queira utulizar essa "tecnica" para criar seus volumes nomeados em imagens que já criam por padrão, deverá saber todos os volumes que esse imagem cria.

### Volumes na Criação de Containers

#### Opção `-v` || `--volume`

Sintaxe:
```sh
-v <named-volume>:<container-directory>[:<options>]
```

Exemplo:
```sh
docker container run -d --name mysql -v mysql:/var/lib/mysql:rw -e MYSQL_ROOT_PASSWORD=root mysql:latest
```

#### Opção `--mount`

Sintaxe:
```sh
--mount 'type=<volume-type>,src=<named-volume>,dst=<container-directory>[,<options>,...]'
```

Exemplo:

- `type`: tipo de volume utilizado;
	- `volume`: volume nomeado;
		- OBS: no caso de volumes nomeados caso ele não exista, ele será criado;
	- `bind`: volume mapeado;
	- `tmpfs`: volume temporário;
- `src` | `source`: nome do volume nomeado ou do diretório na maquina _host_ a ser mepeado:
- `dst` | `target`: nome do diretório dentro do container;

```sh
docker container run -d --name mysql --mount 'type=volume,source=mysql-db,target=/var/lib/mysql,readonly' -e MYSQL_ROOT_PASSWORD=root mysql:latest
```

_OBSERVATIONS_:
- Volumes não são excluídos por padrão quando seus containers morrem.

_REFERENCELINKS_:
- [Vídeo](https://www.youtube.com/watch?v=StQYXkFgeeA).

## Tips & Tricks

- Você pode referenciar algum container pelos 3 primeiros caracteres do seu _id_

### System Prune

Diferença entre os _system prunes_:
```
docker system prune, will remove:
- all stopped containers
- all unused networks
- all dangling images
- all unused build caches

docker system prune --all, will remove:
- all unassociated images
- all build cache

docker system prune --volumes, will remove:
- all unused anonymous volumes

docker system prune --all --volumes, will remove:
```

E como são suas formas completas:
```
docker system prune, will remove:
- all stopped containers
- all unused networks
- all dangling images
- all unused build caches

docker system prune --all, will remove:
- all stopped containers
- all unused networks
- all unassociated images
- all build cache

docker system prune --volumes, will remove:
- all stopped containers
- all unused networks
- all unused anonymous volumes
- all dangling images
- all unused build caches

docker system prune --all --volumes, will remove:
- all stopped containers
- all unused networks
- all unused anonymous volumes
- all unassociated images
- all build cache
```

## Warnings

- Cuidado com a utilização de _data volumes_ pois pode causar efeitos não esperados nos containers, no caso de uso de um container _MySQL_, quando um novo container utiliza um _data volume_, além de todos os dados (todos os _databases_ e tabelas), ele também puxará as informações "administrativas", logo a senha permanesse a mesma na primeira instâncias também.

## Compose

Remove todos os recursos referentes aquele compose:
```sh
docker compose down --rmi all
```

## Troubleshooting

"network (docker0): networks have same bridge name":
```sh
ip link delete docker0
rm -iv /var/lib/docker/network/files/local-kv.db
systemctl restart docker
docker rm -f `docker container ls -aq`
```

## Links

- [Docker Swarm/Service](https://www.cloudbees.com/blog/running-services-within-docker-swarm).

---

## Images

### MySQL

Variáveis de ambiente:
- `MYSQL_DATABASE`: Define o nome de um _database_ para ser criado junto com a criação do container, ou seja, caso o banco seja _dropado_ é necessário remover e recriar o container. Caso esteja usando volumes isso não resolverá, terá que remover e recriar o volume também.
