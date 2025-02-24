# Exercício 5

## Slides

O slide está disponível no arquivo `slides.pdf`.

## Instruções para rodar

Para rodar a imagem Docker e, assim, **rodar o exercício** cliente-servidor, execute o seguinte comando:

```bash
PROTOCOL=<protocol> docker compose up --build --scale rpc-client=<n_clients>
```

Observações:

- `<protocol>` deve ser `go-rpc` ou `grpc`.
- A flag `--build` apenas é necessária na primeira execução ou, ainda, em caso de mudança no código.
- `<n_clients>` é o número de clientes a serem executados. Caso seja 1, a flag `--scale rpc-client=1` não é necessária.

Para conferir um **resumo dos resultados**, com valores de média, mediana e desvio padrão de 10.000 execuções, execute o seguinte comando:

```bash
go run main.go <protocol> results
```

Observações:

- `<protocol>` deve ser `go-rpc` ou `grpc`.

Para acessar o **tempo de cada execução**, basta ir à pasta `shared-volume` e abrir o arquivo `grpc-results.txt` ou `go-rpc-results.txt`, a depender do protocolo utilizado.