# Arquitetura Hexagonal

Antigamente conhecida como _Ports and Adapters_. Tem o foco de organizar o código em camadas (como uma cebola).

Caracteristicas:

- Tem o foco em deixar a aplicação desaclopada do ambiente
- Facilita a integração com outros componentes
- Facilita a manutenção por longos períodos de tempo
- Cada camada é auto-contida e isolada do mundo externo
    - Cada camada se comunida uma com a outra pelas Portas/Adaptadores

Camadas:

- Interação com o banco de dados
- Interação com o usuário
