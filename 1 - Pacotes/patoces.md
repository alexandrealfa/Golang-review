# Packages GO 

- criação de um módulo em go, onde será centralizado todos os pacotes do projeto: ```Comand: go mod init mod_name```
- em golang o modulo criado será o gestor de depêndencias do projeto like package.json
- com um arquivo de .mod em seu projeto é possivel realizar o build do projeto. ```Comand: go build ```
- é possivel ter varios pacotes dentro de um único projeto, basta ir criando pastas.
- para importar funcões de um pacote que não está no mesmo nível é necessário definir a função com a letra inicial maiuscula.
- caso a função seja utilizada em um pacote compartilhado, sua definição pode ser com uma letra minúscula.


## Pacotes externos Go

para importar um pacote externo basta usar o ``` Comand: go get **url** ```

para remover todas as dependências que não estão sendo utilizadas basta utilizar ``` Comand: go mod tidy ```
