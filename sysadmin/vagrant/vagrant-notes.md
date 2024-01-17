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

- `-m`: template mínimalista;
- `--provider=<provider>`: especifica o provedor.

1.  Pesquise a _box_ que desejará subir: <https://vagrantcloud.com/search>;

2. Crie um diretório e entre nele:
```sh
mkdir -pv /tmp/vagrant/ && cd /tmp/vagrant/
```

3. Inicie um projeto **vagrant** (isso criará um Vagrantfile que contem configurações comuns pré-definidas):
```sh
vagrant init [-m] <repo/box>
```

OBS: modifique o arquivo conforme achar necessário.

4. Depois disso basta subir a _VM_:
```sh
vagrant [--provider=<provider>] up
```

5. Caso deseja se conectar via `ssh` na _VM_:
```sh
vagrant ssh
```

6. Caso deseje desligar a _VM_:
```sh
vagrant halt
```

6. Caso deseje remover a _VM_:
```sh
vagrant destroy
```

## Commands

Listar _boxes_:
```sh
vagrant box list
```

Baixar _box_:
```sh
vagrant box add <repo/box>
```

Reiniciar _box_ (também utilizado quando Vagrantfile é alterado):
```sh
vagrant reload [--provision]
```

### Plugins

Instalar plugins:
```sh
vagrant plugin install <plugin>
```

Atualizar plugins:
```sh
vagrant plugin update
```

## Providers

### VirtualBox

Se necessário:
```sh
vagrant plugin install vagrant-virtualbox
```

### Libvirt & Qemu/KVM

Sistemas operacionais:

- Debian:
    1. `apt install vagrant-libvirt libvirt-daemon-system`
    1. `usermod -aG libvirt "$USER"`

#### Libvirt

Se necessário:
```sh
vagrant plugin install vagrant-libvirt
```

#### Qemu/KVM

Se necessário:
```sh
vagrant plugin install vagrant-qemu
```

_REFERENCELINKS_:
- [Community Wiki](https://github.com/ppggff/vagrant-qemu).

## Vagrant Files

- <details>
	<summary>Debian (Qemu/KVM)</summary>

	```ruby
	Vagrant.configure("2") do |config|
	  config.vm.box = "debian/buster64"
	  config.vm.synced_folder ".", "/vagrant", disabled: true
	  #config.vm.network "forwarded_port", guest: 80, host: 80
	  config.vm.provider "qemu" do |qe|
	    qe.qemu_dir = "/usr/bin/"
	    qe.arch="x86_64"
	    qe.memory = "512"
	    qe.smp = "1"

	    # need for x86_64
	    qe.machine = "q35"
	    qe.cpu = "max"
	    qe.net_device = "virtio-net-pci"

	    # it seems this box need a VGA device (the debug serial port doesn't work... I don't know why)
	    qe.extra_qemu_args = %w(-vga std)

	    #config.vm.provision "shell", inline: <<-SHELL
	    #  apt update && apt upgrade -y && apt install -y nginx
	    #SHELL
	  end
	end
	```
</details>

---

_OBSERVATIONS_:

- Quase todos os comandos são referentes ao _workspace_ que você está (o lugar aonde estão o **Vagrantfile** & **.vagrant** pasta), por esse motivo que não precisamos especificar qual _box_ queremos conectar quando executamos `vagrant ssh` por exemplo;

- Caso alguma das _boxes_ precise do **VirtualBox**, no meu caso, um _debian base_ minimo, depois de instalar o _VB_, é necessário rodar o _script_ de configuração (`/sbin/vboxconfig`), posteriormente reiniciar a máquina para aplicar as configuração e por fim habilitar os módulos do _VB_ no _kernel_ (`sudo modprobe vboxdrv vboxnetflt vboxnetadp vboxpci`).

_REFERENCELINKS_:
- <https://wiki.debian.org/Vagrant>.
