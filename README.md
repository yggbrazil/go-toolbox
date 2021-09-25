[![Test](https://github.com/yggbrazil/go-toolbox/actions/workflows/test.yml/badge.svg)](https://github.com/yggbrazil/go-toolbox/actions/workflows/test.yml)

# Go-ToolBox #
Repositório com o objetivo de compartilhar a "caixa de ferramentas" que utilizamos na empresa, para agilizar o nosso processo de desenvolvimento de soluções tecnológicas

```sh
go get -u github.com/yggbrazil/go-toolbox
```

## api ##

[api](https://github.com/yggbrazil/go-toolbox/tree/master/api) é um wrapper do [Echo](https://github.com/labstack/echo) com a configurações básicas para criar uma API REST em poucas linhas

## cache ##

[cache](https://github.com/yggbrazil/go-toolbox/tree/master/cache) É um wrapper do [go-cache](https://github.com/patrickmn/go-cache) uma lib de cache em memória com tempo de expiração, básicamente tem somente o metodo Set e Get

## check ##

[check](https://github.com/yggbrazil/go-toolbox/tree/master/check) É solução para poder utilizar operador ternário em Golang

## cookie ##

[cookie](https://github.com/yggbrazil/go-toolbox/tree/master/cookie) É um lib para adicionar e deletar cookie no framework [Echo](https://github.com/labstack/echo)

## crypt ##

[crypt](https://github.com/yggbrazil/go-toolbox/tree/master/crypt) É um lib para ajudar na geração de hashs no Golang

## database ##

[database](https://github.com/yggbrazil/go-toolbox/tree/master/database) É um wrapper do [SQLx](https://github.com/jmoiron/sqlx) com o objetivo de entrar uma conexão com banco de dados (MySQL, Postgres ou SQLite) somente lhe indicando o arquivo env com os paramentros de conexão

## format ##

[format](https://github.com/yggbrazil/go-toolbox/tree/master/format) É um lib com funções de formatação para diversos tipos
## gcp_pubsub ##

[gcp_pubsub](https://github.com/yggbrazil/go-toolbox/tree/master/gcp_pubsub) É um wrapper sobre a lib de PubSub da GCP

## grpc ##

[grpc](https://github.com/yggbrazil/go-toolbox/tree/master/grpc) É um lib que ajuda na criação de um server ou client GRPC

## handler ##

[handler](https://github.com/yggbrazil/go-toolbox/tree/master/handler) É um lib para criar funções utilizadas em diversos handlers no framework [Echo](https://github.com/labstack/echo) como a BindAndValidade para fazer o bind na struct e validar ela

## ibge ##
[ibge](https://github.com/yggbrazil/go-toolbox/tree/master/ibge) É uma lib responsável por retornar informações sobre localidade vindas do IBGE

## json2env ##

[json2env](https://github.com/yggbrazil/go-toolbox/tree/master/json2env) é uma lib que le um arquivo json e coloca os valores no enviroment

## jwt ##

[jwt](https://github.com/yggbrazil/go-toolbox/tree/master/jwt) É um wrapper do [jwt-go](https://github.com/dgrijalva/jwt-go) para facilitar a utilização de jwt nos projetos

## log ##

[log](https://github.com/yggbrazil/go-toolbox/tree/master/log) É uma lib para lidar com log, para log em arquivo ou no terminal com a linha do arquivo com o erro

## path ##
[path](https://github.com/yggbrazil/go-toolbox/tree/master/path) É uma lib com funções relacionadas à diretórios

## signature ##
[signature](https://github.com/yggbrazil/go-toolbox/tree/master/signature) É uma lib para criar a struct de Signature para SOAP do XML

## template ##
[template](https://github.com/yggbrazil/go-toolbox/tree/master/template) É uma lib onde se manda o template e a struct e ele retorna o template compilado com as váriavéis

## text ##
[text](https://github.com/yggbrazil/go-toolbox/tree/master/text) É uma lib para manipulação de STRINGS

## validator ##
[validator](https://github.com/yggbrazil/go-toolbox/tree/master/validator) É uma lib para validação dos dados de uma struct, com várias validações, como CPF, CNPJ, email e etc

## viacep ##
[viacep](https://github.com/yggbrazil/go-toolbox/tree/master/viacep) É uma lib que retorna informações de localidade ao informar o CEP
