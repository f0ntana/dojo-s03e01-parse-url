
#### Coding Dojo s03e01 

### Problema: Analisando URLs

Dada uma URL, desenvolva um programa que indique se a URL é válida ou não e, caso seja válida, retorne as suas partes constituintes.

Exemplos:
* Entrada: http://www.google.com/mail/user=fulano
    * Saída:
        protocolo: http
        host: www
        domínio: google.com
        path: mail
        parâmetros: user=fulano
* Entrada: ssh://fulano%senha@git.com/
    * Saída:
        protocolo: ssh
        usuário: fulano
        senha: senha
        dominio: git.com
URL do problema: http://dojopuzzles.com/problemas/exibe/analisando-urls/

### Retrospectiva
Como foi o evento ?

#### Bons

* Local foda (Nuvem)
* Fez estudar o Tour do Golang 
* Linguagem foda
* Galero gostou de Golang
* Talk de Golang (Gabriel)
* Sinuca
* Snacks
* Muitos Snacks
* Refri e Cerveja
* Encontro de Brothers
* Muita gente nova (IFMT e SENAI)
* Participação feminina

#### Ruins 

* Sinuca (Galera mais jogou do que codou)
* Muitos foram embora cedo
* Barulho no inicio
* Mini panelinhas, tem que socializar mais
* Não fez a galera nova se apresentar
* Imprimir o problema
* Faltou o dojo helper 
* Muitas distrações no localhost
* O ambiente tava uma porra (não rodava teste direito)
* Teclado tava uma porra

#### Participantes (Os que foram embora, mandem PR)
* Alvaro Viebrantz
* Gabriel Pedro
* Thais Bueno
* Castrolol
* Andressa irmã do Thiago
* Josemar
* Guilherme
* Onias
* Ronny irmão do Rafael
* Rafael irmão do Ronny
* Amigo do Ronny
* João Vitor
* Nelson Junior
* Ziguifrido
* Weslley (Dono da porra toda)
* Carlos Rabelo
* Henrique Ribeiro
* Erlon Musk?
* Fabian Carlos
* Edipo (Popinho)
* Igor 
* Ronan
* Jossan
* Eduardo
* Victor Hugo
* Thiago irmão da Andressa

### Alguns pré-requisitos pra rodar o projeto

Este projeto asume que você está trabalhando com um -workspace-- padrão para Go como descrito em [Golang Code](http://golang.org/doc/code.html). Para este projeto foi utilizado Go 1.3.3, porém acredito que qualquer versão acima de 1.1 vá funcionar.

### Resolver dependências do projeto 

```shell
go get github.com/smartystreets/goconvey/convey
```

### Como rodar

```shell
goconvey ./...
# ou
go ./... -v
```
e acessar [locahost:8080](http://localhost:8080)
