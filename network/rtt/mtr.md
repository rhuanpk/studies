# Round-Trip TIme

Round-Trip Time (RTT) é o tempo total do percurso do pacote.

- A cada rodada é enviado um novo pacote ICMP individual para CADA _hope_ até o destino
- O RTT é o tempo total de ida do pacote ICMP até chegar ao destino e voltar a origem
- A rota que o pacote ICMP toma para fazer o percurso de ida pode não ser a mesma que fará na volta

Analisando até o _hope_ 3 de um percurso de 10 _hopes_:

1. Hope
	1. Pacote ICMP Request com TTL=1 sai da **origem**
	1. Vai para o **1º roteador** e morre
	1. 1º roteador envia pacote ICMP Exceeded para a **origem** e morre

	Total: 2ms

1. Hope
	1. Pacote ICMP Request com TTL=2 sai da **origem**
	1. Vai para o **1º roteador** e decrementa TTL=1
	1. Vai para o **2º roteador** e morre
	1. 2º roteador envia pacote ICMP Exceeded para o **1º roteador**
	1. Vai para a **origem** e morre

	Total: 8ms

1. Hope
	1. Pacote ICMP Request com TTL=3 sai da **origem**
	1. Vai para o **1º roteador** e decrementa TTL=2
	1. Vai para o **2º roteador** e decrementa TTL=1
	1. Vai para o **3º roteador** e morre
	1. 3º roteador envia pacote ICMP Exceeded para o **2º roteador**
	1. Vai para a **1º roteador**
	1. Vai para a **origem** e morre

	Total: 5ms

Como podemos observar, na rodada em questão (42, por exemplo), o pacote ICMP sequencial 42 que é enviado para o _hope_ 1 é diferente do pacote 42 que é enviado para o _hope_ 2, que é diferente para o _hope_ 3 e assim por diante. Sendo assim, é por esse motivo que um _hope_ mais a frente pode ter uma latência menor do que _hopes_ anteriores, uma vez que o pacote ICMP anterior a ele, por ser independente, tenha traçado uma rota diferente.

Em outras palavras, podemos dizer que, o pacote ICMP enviado para o _hope_ 3, passou por:
```
origem > roteador 1 > roteador 2 > roteador 3 > roteador 2 > roteador 1 > origem
```
Porém, sua passada pelo _hope_ 2 nada ter a ver com RTT do _hope_ 2 gerado anteriormente.

Da mesma maneira, o pacote ICMP enviado para o _hope_ 2 gerado anteriormente, passou por:
```
origem > roteador 1 > roteador 2 > roteador 1 > origem
```
E da mesma maneira também, sua passada pelo _hope_ 1 nada ter a ver com RTT do _hope_ 1 gerado anteriormente.