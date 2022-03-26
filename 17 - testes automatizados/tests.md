Com o comando ```go test``` você consegue rodar os tests dos arquivos com _test no diretório atual.
Caso o arquivo de test não esteja no diretório atual, e sim dentro de um submodulo do pacote, é possivel rodar o comando:
```go test ./...```

o comando go test possui algumas flags

``` go test -v``` -> verbose, retorna o resultado do test mais detalhado.


é possível rodar mais de um test ao mesmo tempo, basta nas funções de test, chamar na primeira linha da função o:
    ```t.Parallel()```