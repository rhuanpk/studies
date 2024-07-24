# Kubernetes

O Kubernetes, de forma semelhante ao Docker Swarm, é um orquestrador de containers, porém, extremamente robusto.

Para ter um _cluster_ Kubernetes, basicamente precisamos configurar o _adm_ (que seria o _node_ principal) que gerencia os _workers_ (_nodes_ secundários) adicionados a ele, ou seja, literalmente um cluster, várias máquinas secundárias acopladas a uma principal. No caso, tanto o _adm_ quanto os _workers_ serão meros softwares rodando em alguma máquina, seja **física ou virtual**, num **_PC_ ou servidor**, _on-premisse_ ou _cloud_. Você também pode ter o _setup_ no qual todos são simplesmente _nodes_!

Para gerenciar o _cluster_ usamos o `kubectl`, ou seja, poderemos ver informações ou enviar comandos para nosso _cluster_. Para fazer o _setup_ do _cluster_ usamos as ferramentas padrões `kubeadm` e `kubelet` em produção, agora, para testes locais podemos usar o `kind` ou o `minikube` que pode simular um ambiente real de _cluster_ criando containers Docker para usar como se fossem as máquinas (mencionada no parágrafo anterior).

O **pod** é a menor "partícula" do Kubernetes. Dentro de um _node_ do Kubernetes (que lembrando é uma máquina), rodará pelo menos um **pod** e dentro dele rodará pelo menos um container Docker.

## Kind

Instalação:
```sh
curl -Lo ./kind 'https://github.com/kubernetes-sigs/kind/releases/latest/download/kind-linux-amd64' && chmod +x ./kind && sudo mv -v ./kind /usr/local/bin/
```

_Cluster_:
1. Crei o arquivo de configuração:
    ```sh
    tee ./config.yaml <<- eof
    kind: Cluster
    apiVersion: kind.x-k8s.io/v1alpha4
    nodes:
    - role: control-plane
    - role: worker
    - role: worker
    - role: worker
    eof
    ```
1. "Suba" o _cluster_:
    `kind create cluster --config ./config.yaml`

## KubeCTL

O `kubectl` procura pelo o arquivo `~/.kube/config` que é o arquivo de conexão com o _cluster_. Arquivo esse que é gerado automáticamente pelo `kind` por exemplo quando o _cluster_ é criado.

Instalação:
```sh
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl" && chmod +x ./kubectl && sudo mv -v ./kubectl /usr/local/bin/
```

Comandos:
- `kubectl get pod [-A] [-owide]`: lista os pods
    - `-A`: todos os _namespaces_
    - `-owide`: mais verbosidade na saída
    - `--show-labels`: mostra as labels dos pods
- `kubectl get node [-v9]`: lista os _nodes_
    - `-v9`: _full verbose_
- `kubectl run --image <docker-hub-image> [-it] [--dry-run -oyaml] <pod-name> [<command>] [> ./pod.yaml]`: cria um novo pod
    - `-i`: executa de forma interativa
    - `-t`: aloca um tty
    - `--dry-run`: apenas executa o comando sem de fato subir o pod
    - `-oyaml`: _printa_ a configuração do comando executado em YAML
- `kubectl delete pod <pod-{name|id}>`: daleta um pod
- `kubectl apply -f ./{pod|deployment}.yaml`: sobe a configuração pelo manifesto (arquivo de configuração)
- `kubectl create deployment --image <docker-hub-image> [--dry-run -oyaml] <deployment-name> [> ./deployment.yaml]`: cria um novo _deployment_ > _replicaset_ > _pod_
    - `--dry-run`: apenas executa o comando sem de fato subir o _deployment_
    - `-oyaml`: _printa_ a configuração do comando executado em YAML
- `kubectl get deployment`: lista os _deployments_
- `kubectl get replicaset`: lista os _replicaset's_
- `kubectl scale deployment <deployment-name> --replicas <count>`: define a quantidade de replica para determinado pod (e todos os containers dentro dele?)
- `kubectl delete deployment <deployment-name>`: delete todo um _deployment_, suas replicas e pods
- `kubectl expose deployment <deployment-name> --port <incoming-request-port> --target-port <outgoing-pod-port>`: cria um _service_ na frente de um deployment, ou seja, expõe a porta do _deployment_ para redirecionar a requisição a algum pod que está escutando nessa mesma porta
- `kubectl get service [<service-name>] [-oyaml]`: lista os _services_
    - `-oyaml`: _printa_ a configuração do comando executado em YAML
- `kubectl get endpoints <label-name>`: lista todos os pods com a label, ou seja, caso seja feito um _deployment_, listará todo o _deployment_ (ou seja, os pods por trás dele)

Quando um pod é criado com `run`, além de todas as coisas, ele recebe também uma _label_, que é basicamente um grupo no qual está insirido. Quando criamos um _deployment_, basicamente criamos vários pods com as mesma _label_, ou seja, estarão dentro do mesmo grupo, dessa forma, o _load balancer_ do Kubernetes começa a atuar pois, dessa forma, por de baixo dos panos, teremos um único _endpoit_, que será a _label_, então, o próprio Kubernetes cuidará de fazer os redirecionamento para os pods dentro desse grupo. A um nível maior na hierarquia do Kubernetes, os _namespaces_ também atuam como agrupadores.

Por conta do _DNS_ do Kubernetes, a _label_ dos pods também serviram com seus nomes de domínios, ou seja, possibilitando coisas como `curl -iL <label-name>`. Para sistema fora da rede dos pods, podemos encontra-los como `http://<label-name>.<namespace>.svc[.cluster.local]`.
