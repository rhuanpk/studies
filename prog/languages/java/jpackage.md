# jpackage

Empactoador nativo do Java.

## Opções

Windows:
- `--win-console`: Essencial para aplicativos CLI
- `--win-shortcut`: Cria atalho no Desktop
- `--win-menu`: Cria atalho no Menu
- `--win-dir-chooser`: Permite usuário escolher diretório de instalação
- ` --win-per-user-install`: Instala somente para o usuário corrente (não precisa de admin)
- `--win-upgrade-uuid`: Para controle de versão do _app_

Mac:
- `--mac-sign`: Essencial para aplicativos não provenientes da _Apple Store_

Linux:
- `--linux-shortcut`: Cria atalho no Desktop
- `--linux-menu-group`: Cria atalho no Menu
- `--linux-deb-maintainer=<email>`: Email da organização ou desenvolvedor

## Comando

Base:
```
jpackage --input /path/to/folder/project --main-jar relative/path/inside/project/to/file.jar --main-class <classname> --name <appname> --app-version x.y.z
```
