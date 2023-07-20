# Caprover Setup (step-by-step)

Configuração Inicial do **CapRover**:

1. Tenha um subdomínio do tipo `A` configurado;
1. Suba uma instância de servidor (de preferência _debian based_?);
1. Instale o docker: `curl -fsSL https://get.docker.com/ | sudo sh -`
1. Libere as postas no _firewall_: `ufw allow 80,443,3000,996,7946,4789,2377/tcp; ufw allow 7946,4789,2377/udp`
1. Baixe/execute a instância do **CapRover**: `docker run -p 80:80 -p 443:443 -p 3000:3000 -v /var/run/docker.sock:/var/run/docker.sock -v /captain:/captain caprover/caprover`
	1. OBS: depois desse passo o **CapRover** já está disponível via web: `http://<ip_of_your_server>:3000`;
1. Faça o apontamento do subdomínio criado para o ip do servidor remoto no qual instalamos o **CapRover**;

A partir desse momento é a parte de setup do **CapRover** que será explicada via _cli_, porém, seus passos podem ser equivalentes pela interface web:

1. Instale o `npm` no servidor do **CapRover**: `sudo apt install -y npm`;
1. Instale o _CapRover CLI_ com: `npm install -g caprover`;
1. Execute o _CapRover CLI_ para iniciar a configuração do servidor: `caprover serversetup`;
	1. Caso já tenha configurado _https_ na instância do **CapRover** o comando de `serversetup` falhará e será necessário ir diretamente para o login com: `caprover login`;

Informações referentes ao deploy: <https://caprover.com/docs/get-started.html>;
