# URL Shortener

Este é um projeto de um encurtador de URLs, implementado em Go. Ele é constituído por uma API REST que recebe URLs longas e retorna um código curto único que pode ser usado para redirecionar para a URL original. A arquitetura do projeto é composta por várias camadas, incluindo uma camada de banco de dados, camadas de acesso aos dados, camadas de lógica de negócios e camadas de API.

## Funcionalidades

- Encurtar URLs: O usuário pode fornecer uma URL longa e obter um código curto único em troca.
- Redirecionar de acordo com o código: Quando um código é fornecido, o serviço redireciona para a URL original correspondente a esse código.
- Expiração de URLs: Os códigos de URL expiram após um período de tempo especificado.

## Instalação

Para rodar o projeto localmente, é necessário ter Go e Docker instalados. Siga os seguintes passos:

1. Clone este repositório:
    ```bash
    git clone https://github.com/dev81log/url_short.git
    ```
2. Navegue até o diretório do projeto:
    ```bash
    cd url_short
    ```
3. Rode o banco de dados com Docker:
    ```bash
    docker-compose up
    ```
4. Compile e rode a aplicação:
    ```bash
    go build
    ./url_short
    ```
A aplicação estará rodando na porta 8080.

## Uso

- Para encurtar uma URL, faça uma solicitação POST para `/` com a URL como JSON no corpo da solicitação. Exemplo:
    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"url":"https://www.example.com"}' http://localhost:8080/
    ```
- Para redirecionar para uma URL, faça uma solicitação GET para `/<code>`. Exemplo:
    ```bash
    curl http://localhost:8080/ABC123
    ```

## Testes

Para rodar os testes do projeto, use o seguinte comando:

```bash
go test ./...
```

## Contribuições

Contribuições são bem-vindas! Por favor, abra um issue ou faça um pull request.

## Licença

Este projeto está licenciado sob a [MIT License](LICENSE).
