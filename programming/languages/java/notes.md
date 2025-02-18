# Java

## História

- Criado em 1990 pela _Sun Microsystems_
- James Gosling
- Comunicação entre diferente dispositivos (C++)
- Criado para ser realmente multiplataforma
    - C, compiladores para cada plataforma
    - C, peculiaridades para cada plataforma
- Criada com o nome de _GreenTalk_
- Com o projeto crescendo, foi renomeada para _Oak_
    - Criação do _Star7_ (o primeiro _touchscreen_) e do _Duke_ (mascote)
    - _Start7_ e _Oak_ "flopados"
- 1994, Tim Berners-Lee
    - Linguagem interativa para a WEB
    - Web Runner
- Java Coffe (Café Forte) -> Java
- HotJava Browser
- Netscape -> Javascript (Marketing)
- A volta das interatividades externas
    - Java Ring, dispositivo de autenticação setup
- NASA -> Robo (criado com Java) enviado para Marte
    - Populariazação da linguagem
- 2006, GPL 3 -> Introduzida ao Open Source
- 2009, Sun Microsystems vendida por bilhões para a Oracle
- Java em todos os lugares
    - Sistemas bancários
    - Chips em cartões de bancos
    - Smartphones (Apps)
    - Blue-ray
    - Kindle
    - Smartwatches

## Funcionamento

Linguagem compilada:
1. Código Fonte
1. Compilador
1. Executável
1. Máquina

OBS: um compilador para cada sistema

Linguagem Java:
1. Código Fonte
1. Compilador (JavaC)
1. Bytecode
1. "Interpretador" (JVM)
1. Máquina

OBS: uma JVM para cada sistema (respectivas versões)

Conceito WORA (Write Once, Run Anywhere) é o princípio da linguagem Java.

## Plataforma

### JRE

O JRE (Java Runtime Environment) sozinho é destinado apenas para **usuários** de qualquer sistema Java, pois, é nele que contem a JVM (sistema que executa o binário do Java (Bytecode)).

O JRE é composta por:
- JVM
    - Loader: carrega o bytecode na memória
    - Verifier: verificar se o programa pode ser executado
    - Interpreter: transforma o bytecode em código p/ máquina host
    - Manager: gerenciador de memória (da JVM)
    - JIT Compiler: otimiza o programa indentificando padrões
- Bibliotecas
    - Libs/APIs: que o programa usa

### JDK

O JDK (Java Development Kit) é literalmente o _kit_ com todas as coisas necessárias para **desenvolvedores** criarem apliações em Java, pois, além de dentro dele contem a JRE, possúi a própria linguagem Java (JavaLang) e ferramente de desenvolvimento como compiladores e _debuggers_.

O JDK é composto por:
- JRE
- JavaLang: a própria linguagem Java
- JavaTools
    - JavaC: compilador Java
    - Debugger: verificar o programa em tempo real

---

## Links

- [Custom Exceptions¹](https://www.scaler.com/topics/custom-exception-in-java/)
- [Custom Exceptions²](https://alvinalexander.com/java/java-custom-exception-create-throw-exception/)
- [The `super` _keyword_](https://alvinalexander.com/java/java-custom-exception-create-throw-exception/)
