# Kubernetes

O Kubernetes, de forma semelhante ao Docker Swarm, é um orquestrador de containers, porém, extremamente robusto.

Para ter um _cluster_ Kubernetes, basicamente precisamos configurar o _adm_ (que seria o _node_ principal) que gerencia os _workers_ (_nodes_ secundários) adicionados a ele, ou seja, literalmente um cluster, várias máquinas secundárias acopladas a uma principal. No caso, tanto o _adm_ quanto os _workers_ serão meros softwares rodando em alguma máquina, seja **física ou virtual**, num **_PC_ ou servidor**, _on-premisse_ ou _cloud_.

Para gerenciar o _cluster_ usamos o `kubectl`, ou seja, poderemos ver informações ou enviar comandos para nosso _cluster_. Para fazer o _setup_ do _cluster_ usamos as ferramentas padrões `kubeadm` e `kubelet` em produção, agora, para testes locais podemos usar o `kind` ou o `minikube` que pode simular um ambiente real criando containers Docker para usar como se fossem as máquinas (mencionada no parágrafo anterior).

O **pod** é a menor "partícula" do Kubernetes. Dentro de um _node_ do Kubernetes (que lembrando é uma máquina), rodará pelo menos um **pod** e dentro dele rodará pelo menos um container Docker.
