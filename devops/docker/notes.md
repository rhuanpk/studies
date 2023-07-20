# docker annotations

## commands

### `inspect`

verificar _ip's_ do container:

```bash
docker container inspect <container_name/id> | grep -i ipaddress
```

verificar volumes do container:

```bash
docker container inspect hungry_raman | grep -iA 11 mounts
```

### `volume`

criar volume nomeado:

```bash
docker volume create <volume_name>
```

## volumes

volumes nada mais são "pastas dos containers que salvamos ou mapeamos"?

basicamente existem dois tipos de volumes:

- _data volumes_ (volumes nomeados): volumes criados e controlados pelo _docker host_ (pelo próprio docker, seu _daemon_) que não tem ligação com a MÁQUINA _host_;
	- OBS: é como se fossem simplesmentes um ponto de montagem do próprio container?

- _bind volumes_ (volumes mapeados/conectados): volumes que mapeia algum diretório ou arquivo da máquina _host_ para dentro do container;

- _tmpfs volumes_ (volumes temporários): volumes que persitem dados em memória simplesmente (!);

### _data volumes_

caso você crie um container no qual a sua imagem já cria algum volume por padrão, esse volume é criado com algum nome aleatório (`docker volume ls`).

para uma melhor manutenção, é recomendado que se crie um volume manualmente (volume nomeado) e utilize esse volume nos seus containers para a persistencia, dessa forma garantindo até a reutilização dos dados gravados dos containers anteriores.

OBS: caso queira utulizar essa "tecnica" para criar seus volumes nomeados em imagens que já criam por padrão, deverá saber todos os volumes que esse imagem cria.

### usar volume na criação do container

#### opção `-v | --volume`

syntax:

```bash
-v <named_volume>:<container_directory>[:<options>]
```

comando:

```bash
docker container run -d --name mysql-db -v mysql-db:/var/lib/mysql:ro -e MYSQL_ROOT_PASSWORD=password mysql:latest
```

#### opção `--mount`

syntax:

```bash
--mount 'type=<volume_type>,src=<named_volume>,dst=<container_directory>[,<options>,...]'
```

comando:

- `type`: tipo de volume utilizado;
	- `volume`: volume nomeado;
		- OBS: no caso de volumes nomeados caso ele não exista, ele será criado;
	- `bind`: volume mapeado;
	- `tmpfs`: volume temporário;
- `src` | `source`: nome do volume nomeado ou do diretório na maquina _host_ a ser mepeado:
- `dst` | `target`: nome do diretório dentro do container;

```bash
docker container run -d --name mysql-db --mount 'type=volume,source=mysql-db,target=/var/lib/mysql,readonly' -e MYSQL_ROOT_PASSWORD=password mysql:latest
```

**OBSERVATIONS**:

- volumes não são excluídos por padrão quando seus containers morrem;

**REFERENCELINKS**:

- [vídeo](https://www.youtube.com/watch?v=StQYXkFgeeA);

## tips

- você pode referenciar algum container pelos 3 primeiros caracteres do seu _id_;

## warnings

- cuidado com a utilização de _data volumes_ pois pode causar efeitos não esperados nos containers, no caso de uso de um container _MySQL_, quando um novo container utiliza um _data volume_, além de todos os dados (todos os _databases_ e tabelas), ele também puxará as informações "administrativas", logo a senha permanesse a mesma na primeira instâncias também;

## links

- [docker swarm/service](https://www.cloudbees.com/blog/running-services-within-docker-swarm);
