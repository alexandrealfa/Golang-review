1) Com o comando ```go test``` você consegue rodar os tests dos arquivos com _test no diretório atual.
2) Caso o arquivo de test não esteja no diretório atual, e sim dentro de um submodulo do pacote, é possivel rodar o comando:
```go test ./...```

###o comando go test possui algumas flags, são elas:

- ``` go test -v``` -> verbose, retorna o resultado do test mais detalhado.
- ``` go test --cover``` -> Informa o coverage de teste da função.
- ``` go test --coverprofile file.txt``` -> gera um arquivo com os dados de cobertura do projeto.
- ```go tool cover --func=file.txt``` -> exibe os dados de coverage contidos no arquivo de coverage.
- ```go tool cover --html=file.txt``` -> exibe os dados de coverage contidos no arquivo de coverage, atraves de um arquivo html, 
com os dados detalhados, de coberturas, como das linhas que não foram cobertas.


é possível rodar mais de um test ao mesmo tempo, basta nas funções de test, chamar na primeira linha da função o:
    ```t.Parallel()```

