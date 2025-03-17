# Exercício 6

## Slides

O slide está disponível no arquivo `slides.pdf`.

## Instruções para rodar

Para rodar a imagem Docker e, assim, **rodar o exercício** cliente-servidor, execute o seguinte comando:

```bash
PROTOCOL=<protocol> docker compose up --build --scale client=<n_clients>
```

Observações:

- `<protocol>` deve ser `rabbitmq` ou `mqtt`.
- A flag `--build` apenas é necessária na primeira execução ou, ainda, em caso de mudança no código.
- `<n_clients>` é o número de clientes a serem executados. Caso seja 1, a flag `--scale client=1` não é necessária.
- Ao final da execução com sucesso de todos os clientes, serão exibidos os resultados no terminal.