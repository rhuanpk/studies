# Vagrant Annotations

## Installation

1. Download key:

```sh
curl -fsSL https://apt.releases.hashicorp.com/gpg | sudo gpg --dearmor -o /etc/apt/trusted.gpg.d/hashicorp-archive-keyring.gpg
```

2. Criar _source-list_ arquivo:

```sh
sudo tee /etc/apt/sources.list.d/hashicorp.list <<< "deb https://apt.releases.hashicorp.com `lsb_release -cs` main"
```

3. Atualizar repos locais e instalar:

```sh
sudo apt update && sudo apt install -y vagrant
```

## Getting Started

1.  Pesquise a _box_ que desejará subir: <https://vagrantcloud.com/search>;

2. Crie um diretório e entre nele:

```sh
mkdir -pv /tmp/vagrant/ && cd /tmp/vagrant/
```

3. Inicie um projeto **vagrant** (isso criará um Vagrantfile que contem configurações comuns pré-definidas):

```sh
vagrant init hashicorp/bionic64
```

OBS: modifique o arquivo conforme achar necessário.

4. Depois disso basta subir a _VM_:

```sh
vagrant up
```

5. Caso deseja se conectar via `ssh` na _VN_:

```sh
vagrant ssh
```

6. Caso deseje remover a _VM_:

```sh
vagrant destroy
```

## Commands

- Listar _boxes_:

```sh
vagrant box list
```

- Baixar _box_:

```sh
vagrant box add <user>/<box>
```

---

_OBSERVATIONS_:

- Os comandos são referentes ao _workspace_ que você está (o lugar aonde estão o **Vagrantfile** & **.vagrant**), por esse motivo que não precisamos especificar qual _box_ que queremos conectar quando executamos `vagrant ssh` por exemplo ou também quando listamos as _boxes_ (`vagrant box list`), aparecerá as _boxes_ baixadas para o _workspace_ em que está.
