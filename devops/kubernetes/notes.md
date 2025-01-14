# Kubernetes

O Kubernetes, de forma semelhante ao Docker Swarm, é um orquestrador de containers, porém, extremamente robusto.

Para ter um _cluster_ Kubernetes, basicamente precisamos configurar o _adm_ (que seria o _node_ principal) que gerencia os _workers_ (_nodes_ secundários) adicionados a ele, ou seja, literalmente um cluster, várias máquinas secundárias acopladas a uma principal. No caso, tanto o _adm_ quanto os _workers_ serão meros softwares rodando em alguma máquina, seja **física ou virtual**, num **_PC_ ou servidor**, _on-premisse_ ou _cloud_. Você também pode ter o _setup_ no qual todos são simplesmente _nodes_!

Para gerenciar o _cluster_ usamos o `kubectl`, ou seja, poderemos ver informações ou enviar comandos para nosso _cluster_. Para fazer o _setup_ do _cluster_ usamos as ferramentas padrões `kubeadm` e `kubelet` em produção, agora, para testes locais podemos usar o `kind` ou o `minikube` que pode simular um ambiente real de _cluster_ criando containers Docker para usar como se fossem as máquinas (mencionada no parágrafo anterior).

O **pod** é a menor "partícula" do Kubernetes. Dentro de um _node_ do Kubernetes (que lembrando é uma máquina), rodará pelo menos um **pod** e dentro dele rodará pelo menos um container Docker.

## Componentes

Esses são alguns dos componentes do Kubernetes:

- **cluster**: junção do nó _node master_ com os _nodes workers_
    - **node**: máquina física ou virtual, _on-premises_ ou _cloud_ que comporta um _master_ ou um _worker_
        - **control-plane**: _node master_ que gerência todo o cluster, mantendo o estado correto do _cluster_
            - **kube-apiserver**: a interface de comunicação (CLI/UI) com o _cluster_ (_control-plane_ > _data-plane_)
            - **scheduler**: gerencia em qual _node_ os pods serão alocados
            - **controller-manager**: gerencia o estado do _cluster_ garantindo que tudo está de pé como o configurado
                - **hpa**: gerencia o número de réplicas com base no uso de RAM e CPU especificados
                - **pdb**: número mínimo de _pods_ para interrupções voluntárias (manutenção do _node_ ou atualizações da aplicação)
                - **cronjobs**: executa _jobs_ de tempos em tempos
            - **etcd**: banco de dados do Kubernetes, armazena configurações e informações do _cluster_ (todos os manifestos?)
                - **configmap**: guarda informações NÃO sigilosas como variáveis de ambiente
                - **secrets**: guarda informações sigilosas como senhas e _strings_ de conexão
        - **data-plane**: _node worker_ que mantem os recursos das aplicações que são gerenciadas pelo _control-plane_
            - **kubelet**: gerencia os _containers_ (e seu estado) dentro de cada _node_
            - **kube-proxy**: gerencia a **rede** e o **balanceamento de carga** dentro de cada _node_
            - **namespaces**: **divide logicamente** o _node_ possibilitando gestão separada de recursos em ambientes isolados
                - **deployment**: gerencia versões e **replicasets** da aplicação sem _downtime_
                    - **replicaset**: gerencia **réplicas** dos _pods_
                        - **pods**: a menor unidade do Kubernetes, contém **um ou vários containers**
                            - **container**: _Docker_ com aplicação
            - **probes**: _healthchecker_ que verifica a saúdo dos _pods/containers_
            - **service**: abstração para um conjunto lógico de _pods_ com regras de acesso (**portas**)
- **volumes**: volumes temporários (em tempo de _container_) ou persistentes para dados
    - **persistent-volume**: os PVs são recursos reais de armazenamento, seja um disco físico ou um dispositivo virtual de armazenamento, é anexado ao _cluster_ mas que não serve para ser utilizado diretamente e, quando criado, aloca um tamanho virtual dentro do dispositivo vinculado, seja esse tamanho o próprio tamanho do disco ou menor
    - **persistent-volume-claim**: os PVCs são volumes de tamanho fixo ou dinâmicos (se vinculado a um _storage-class_), anexados a pods que utilizam o espaço disponibilizado pelos PVs
    - **storage-class**: os StorageClasses são configurações para fazer o alocamento de espaço dinâmico dos PVCs criando novos PVs se necessário

## Ferramentas

Há ferramentas externas e do próprio Kubernetes para trabalharmos.

### Kind

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

### KubeCTL

O `kubectl` procura pelo o arquivo `~/.kube/config` que é o arquivo de conexão com o _cluster_. Arquivo esse que é gerado automáticamente pelo `kind` por exemplo quando o _cluster_ é criado.

Instalação:
```sh
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl" && chmod +x ./kubectl && sudo mv -v ./kubectl /usr/local/bin/
```

Comandos:
- `kubectl cluster-info [dump]`: verifica informações do _cluster_
- `kubectl proxy`: lista os _endpoints_ do próprio _cluster_
- `kubectl apply [-f] {config.yml|folder}`: sobe a configuração pelo manifesto (arquivo de configuração)
- `kubectl scale deployment <deployment> --replicas <count>`: define a quantidade de replicas para determinado _pod_
- `kubectl delete [--all] <resource> [<id>]`: daleta o recurso especificado (pode acabar sendo recursivo)
- `kubectl describe <resource> [<id>]`: exibe informações/logs sobre algum recurso/tipo de recurso
- `kubectl rollout history deployment <deployment>`: lista o histórico de mudanças do _deployment_
- `kubectl rollout undo deployment <deployment> --to-revision=<id>`: faz o _rollback_ da aplicação
- `kubectl logs <pod>`: retorna os logs do _pod_ especificado
- `kubectl get [-A] [-v9] [-owide] [-ojson] [-ojsonpath='{.data}'] [--watch] [--show-labels] {all|nodes|namespaces|deployments|replicasets|pods|services|endpoints|hpa|secrets} [label-name]`
    - lista os _pods/nodes/services/endpoints_
    - `-A`: todos os _namespaces_
    - `-v9`: full verbose
    - `-owide`: mais campos na saída
    - `-ojson`: formata a saída em JSON
    - `-ojsonpath`: extrai informaçãos como JSON
    - `--watch`: fica monitorando o comando executado
    - `--show-labels`: mostra as _labels_ dos _pods_
- `kubectl run [-it] [--dry-run -oyaml] --image <image> <pod> [<command>] [> ./pod.yaml]`
    - cria um novo _pod_
    - `-i`: executa de forma interativa
    - `-t`: aloca um _tty_
    - `--dry-run`: apenas executa o comando sem de fato subir o _pod_
    - `-oyaml`: _printa_ a configuração do comando executado em YAML
- `kubectl create [--dry-run -oyaml] deployment --image <image> <deployment> [> ./deployment.yaml]`
    - cria um novo _deployment_ > _replicaset_ > _pod_
    - `--dry-run`: apenas executa o comando sem de fato subir o _deployment_
    - `-oyaml`: _printa_ a configuração do comando executado em YAML
- `kubectl expose deployment <deployment> --port <incoming-port> --target-port <outgoing-port>`
    - cria um _service_ na frente de um _deployment_
    - expõe a porta do _deployment_ para redirecionar as requisições a algum _pod_ que está escutando nessa mesma porta
    - `--port`: porta de entrada da requisição
    - `--target-port`: porta de saída do _pod_

Quando um pod é criado com `run`, além de todas as coisas, ele recebe também uma _label_, que é basicamente um grupo no qual está insirido. Quando criamos um _deployment_, basicamente criamos vários pods com as mesma _label_, ou seja, estarão dentro do mesmo grupo, dessa forma, o _load balancer_ do Kubernetes começa a atuar pois, dessa forma, por de baixo dos panos, teremos um único _endpoit_, que será a _label_, então, o próprio Kubernetes cuidará de fazer os redirecionamento para os pods dentro desse grupo. A um nível maior na hierarquia do Kubernetes, os _namespaces_ também atuam como agrupadores.

Por conta do _DNS_ do Kubernetes, a _label_ dos pods também serviram com seus nomes de domínios, ou seja, possibilitando coisas como `curl -iL <label-name>`. Para sistema fora da rede dos pods, podemos encontra-los como `http://<label-name>.<namespace>.svc[.cluster.local]`.

## Helm

Helm é um "gerenciador de pacotes" Kubernetes, no caso, ele atuará como se fosse um receita para todo um _cluster_, como se fosse um Dockerfile, que no caso do Helm, é chamado de _Chart_, porém, diferentemente de um Dockerfile, o Chart não é apenas uma arquivo... seria mais correto afirmar que é um _bundle_, que na prática será uma pasta contentdo todos os arquivos de configurações e manifestos Kubernetes para que o Helm crie o pacote.

Versionadores remotos para pacotes Helm:

- [Artifact Hub](https://artifacthub.io)
- [Helm Charts](https://bitnami.com/stacks/helm)

Helm CLI:

- `helm install <helm> <url>`: cria um novo Helm
- `helm uninstall <helm>`: desinstala um pacote
- `helm create <helm>`: cria um novo chart
- `helm package <helm>`: empacota o Helm criado
- `helm list`: lista os Helm's instalados
