# Toplip

## Opções do comando

Algumas funcionalidades desse comando incrível!

### Encriptar

- Mostra a encriptação e sai
	~~~shell
	$ toplip file.txt
	~~~

- Gera o arquivo encriptado
	~~~shell
	$ toplip file.txt > file.encrypted
	~~~

### Decriptar

- Mostra o arquivo descriptografado e sai
	~~~shell
	$ toplip -d file.encrypted
	~~~

- Restaura o arquivo criptografado
	~~~shell
	$ toplip -d file.encrypted > file.txt
	~~~

### Encriptar com múltiplas senhas

- Encripta o arquivo com "n" senhas 
	~~~shell
	$ toplip -c 2 file.txt > file.encrypted
	~~~

- Para **decriptar um arquivo com múltiplas senhas** passe junto com o parâmetro "-d" o parâmetro "-c" com o respectivo número de senhas, exemplo:
	~~~shell
	$ toplip -c 2 -d file.encrypted > file.txt 
	~~~

### Encriptar arquivo numa foto

~~~shell
$ toplip -m image.jpg file.txt > image-encrypted.jpg
~~~

- Para **descriptar um arquivo escondido numa imagem basta estar passando a imagem criptografada** com o parâmetro "-d", exemplo:
	~~~shell
	$ toplip -d image.jpg > file.txt
	~~~

- *O parâmetro "-c" é suportado com esse se passado primeiro, exemplo:*
	
	~~~shell
	$ toplip -c 7 -m image.jpg file.txt > image-encrypted.jpg
	~~~
