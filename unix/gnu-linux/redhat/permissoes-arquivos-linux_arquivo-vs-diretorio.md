# Permissões de arquivo Linux (_FILE_ vs _DIRECTORY_)

## O que as permissões de arquivo do Linux realmente fazem?

Já falei sobre como visualizar as permissões de arquivo, a quem elas se aplicam e como ler quais permissões estão habilitadas ou desabilitadas. Mas o que essas permissões realmente fazem na prática?

### Leitura (r)

A permissão de leitura é usada para acessar o conteúdo do arquivo. Você pode usar uma ferramenta como `cat` ou `less` no arquivo para exibir o conteúdo do arquivo. Você também pode usar um editor de texto como o `vi` ou visualizar o arquivo para exibir o conteúdo do arquivo. A permissão de leitura é necessária para fazer cópias de um arquivo, porque você precisa acessar o conteúdo do arquivo para duplicá-lo.

### Escrita (w)

A permissão de gravação permite modificar ou alterar o conteúdo de um arquivo. A permissão de gravação também permite usar os operadores de redirecionamento ou acréscimo no shell (`>` ou `>>`) para alterar o conteúdo de um arquivo. Sem permissão de gravação, não são permitidas alterações no conteúdo do arquivo.

### Execução (x)

A permissão de execução permite executar o conteúdo de um arquivo. Normalmente, os executáveis seriam coisas como comandos ou aplicativos binários compilados. No entanto, a permissão de execução também permite que alguém execute scripts de shell Bash, programas Python e uma variedade de linguagens interpretadas.

Existem outras maneiras de executar o conteúdo de um arquivo sem permissão de execução. Por exemplo, você pode usar um intérprete que tenha permissão de execução para ler um arquivo com instruções para o intérprete executar. Um exemplo seria invocar um script de shell Bash dessa forma:

```sh
$ script bash.sh
```

O executável que está sendo executado é o bash. O arquivo **script.sh** é lido pelo interpretador **Bash** e seus comandos são executados. O conteúdo deste artigo é de uso geral, mas no Linux geralmente existem maneiras adicionais de realizar tarefas.

## Como funcionam as permissões de diretório?

Os tipos de arquivo de diretório são indicados com `d`. Conceitualmente, as permissões funcionam da mesma maneira, mas os diretórios interpretam essas operações de maneira diferente.

### Leitura (r)

Como os arquivos normais, esta permissão permite ler o conteúdo do diretório. No entanto, isso significa que você pode visualizar o conteúdo (ou arquivos) armazenado no diretório. Essa permissão é necessária para que coisas como o comando `ls` funcionem.

### Escrita (w)

Tal como acontece com os arquivos normais, isso permite que alguém modifique o conteúdo do diretório. Ao alterar o conteúdo do diretório, você está adicionando arquivos ao diretório ou removendo arquivos do diretório. Como tal, você deve ter permissão de gravação em um diretório para mover (`mv`) ou remover (`rm`) arquivos dele. Você também precisa de permissão de gravação para criar novos arquivos (usando `touch` ou um operador de redirecionamento de arquivo `>` e `>>`) ou copiar arquivos (`cp`) para o diretório.

### Execução (x)

Essa permissão é muito diferente em diretórios em comparação com arquivos. Essencialmente, você pode pensar nisso como um fornecimento de acesso ao diretório. Ter permissão de execução em um diretório autoriza você a ver informações estendidas sobre os arquivos no diretório (usando `ls -l`, por exemplo), mas também permite que você altere seu diretório de trabalho (usando `cd`, ou seja, ir para esse diretório) ou passe por este diretório no caminho para um subdiretório abaixo.

A falta de permissão de execução em um diretório pode limitar as outras permissões de maneiras interessantes. Por exemplo, como você pode adicionar um novo arquivo a um diretório (aproveitando a permissão de gravação) se não consegue acessar os metadados do diretório para armazenar as informações de um arquivo novo e adicional? Você não pode. É por esse motivo que os arquivos do tipo diretório geralmente oferecem permissão de execução para um ou mais proprietários do usuário, proprietário do grupo ou outros.

---

_REFERENCELINKS_:

- De: <https://www.redhat.com/sysadmin/linux-file-permissions-explained>.
