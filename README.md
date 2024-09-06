# API de Fast Food

Esta é uma API RESTful para um sistema de fast food, desenvolvida em Go usando o framework Gin e GORM como ORM.

## Funcionalidades

- Gerenciamento de itens do menu (CRUD)
- Integração com banco de dados SQLite
- Estrutura de projeto organizada e escalável

## Requisitos

- Go 1.16 ou superior
- SQLite

## Instalação

1. Clone o repositório:
   ```
   git clone https://github.com/seu-usuario/fastfood-api.git
   cd fastfood-api
   ```

2. Instale as dependências:
   ```
   go mod tidy
   ```

3. Execute a aplicação:
   ```
   go run cmd/main.go
   ```

A API estará disponível em `http://localhost:8080`.

## Estrutura do Projeto

```
fastfood-api/
├── cmd/
│   └── main.go                 # Ponto de entrada da aplicação
├── internal/
│   ├── handlers/
│   │   └── menu_handler.go     # Manipuladores de requisições
│   ├── models/
│   │   └── menu_item.go        # Modelos de dados
│   └── repository/
│       └── menu_repository.go  # Camada de acesso a dados
├── pkg/
│   └── database/
│       └── database.go         # Configuração do banco de dados
├── config/
│   └── config.go               # Configurações da aplicação
├── go.mod
├── go.sum
└── README.md
```

## Endpoints

- `GET /menu`: Lista todos os itens do menu
- `GET /menu/:id`: Retorna um item específico do menu
- `POST /menu`: Cria um novo item no menu
- `PUT /menu/:id`: Atualiza um item existente no menu
- `DELETE /menu/:id`: Remove um item do menu

## Desenvolvimento

Para adicionar novas funcionalidades ou modificar as existentes, siga estas etapas:

1. Crie um novo branch: `git checkout -b minha-nova-funcionalidade`
2. Faça suas alterações e commit: `git commit -am 'Adiciona nova funcionalidade'`
3. Push para o branch: `git push origin minha-nova-funcionalidade`
4. Crie um novo Pull Request

## Contribuição

Contribuições são bem-vindas! Por favor, leia o arquivo CONTRIBUTING.md para detalhes sobre nosso código de conduta e o processo para enviar pull requests.

## Licença

Este projeto está licenciado sob a Licença MIT - veja o arquivo [LICENSE.md](LICENSE.md) para detalhes.
