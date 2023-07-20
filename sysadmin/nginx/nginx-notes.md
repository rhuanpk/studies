# NGINX Annotations

Possibilidades de uso:
- Webserver;
- Proxy reverso;
- Caching;
- Load Balancer.

## Como funciona?

Quando iniciado, ele cria um processo **master** (_single trhead_) e para cada requisição ele cria um _worker_ que é o processo dedicado para essa requisção e para cada nova requisição ele cria um novo _worker_ e como cada _worker_ é assíncrono entre as _threads_ da _CPU_, você pode ter o número máximo de _workers_ por vez pela quantidade de _threads_ que a _CPU_ do servidor tem.

Sendo assim, o **nginx** tem um processo principal (_nginx master process_) que gerencia cada processo subprocesso (_workers_).

## Comandos

_Nginx signals_:
- `stop`: fast shutdown;
- `quit`: graceful shutdown;
- `reload`: reloading the configuration file;
- `reopen`: reopening the log files.

Enviar _signals_ para o **nginx**:
```bash
nginx -s <signal>
```

## Arquvivos de Configuração

- Arquivo principal de configuração: `/etc/nginx/nginx.conf`;
- Arquivos de configuração de usuário: `/etc/nginx/conf.d/*.conf`.

Estrutura do arquivo:
- Diretiva simples: _statemants_ (declarações) de **única linha**, separadas por `;`;
- Diretiva de block: _statemants_ (declarações) de **múltiplas linha**, declarações simples dentro de blocos de `{ }`;
- Contexto: quando dentro de uma diretirava de bloco há outras.

### `location`

#### Expressões regulares (_RegEx_)

- `~`: case-sensitive (#regex);
- `~*`: case-insensitive (#regex);
- `/`: path normal (#fixedstring);
- `^~`: se casar com a correspondência e for o path mais longo, retorna imediatamente sem verificar as regex's (#fixedstring);
- `=`: pesquisa exata, retorna imediatamente (#fixedstring).

Os "locais" sem _RegEx_, ou seja com _string_ fixa (_prefix locations_) são olhados primeiro pelo **nginx** e caso tenham duas localizações que casam, a que tiver o caminho mais logo é aceita (`/any/path` vs `/long/any/path`).

Em contrapartida, dos "locais" especificados com _RegEx_, é pego o primeiro resultado que casar.

#### `/` ou `//`?

- `location /example { }`: corresponde a `http[s]://domain.any/example` e suas subpastas;
- `location /example/ { }`: corresponde **SOMENTE** a subpastas de `http[s]://domain.any/example`.

---

_REFERENCELINKS_:
- <https://nginx.org/en/docs/http/request_processing.html>;
- <https://nginx.org/en/docs/http/ngx_http_core_module.html#location>.
