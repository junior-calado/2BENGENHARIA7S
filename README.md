# API de Gerenciamento de Cards MTG

Este é um projeto de API REST desenvolvido em Go (Golang) para gerenciamento de cards do jogo Magic: The Gathering. A API utiliza autenticação JWT e banco de dados PostgreSQL.

## 🚀 Tecnologias Utilizadas

- Go 1.23.0
- Gin (Framework Web)
- GORM (ORM para Go)
- PostgreSQL
- JWT para autenticação

## 📋 Pré-requisitos

Para executar este projeto, você precisará ter instalado:

1. Go (versão 1.23.0 ou superior)
2. PostgreSQL
3. Git

## 🔧 Configuração do Ambiente

1. Clone o repositório:
```bash
git clone [URL_DO_REPOSITÓRIO]
cd [NOME_DO_DIRETÓRIO]
```

2. Instale as dependências:
```bash
go mod download
```

3. Configure o banco de dados PostgreSQL:
- Crie um banco de dados chamado `mtg_db`
- As credenciais padrão são:
  - Host: localhost
  - Porta: 5432
  - Usuário: postgres
  - Senha: 1
  - Banco de dados: mtg_db

## 🏃‍♂️ Executando o Projeto

1. Inicie o servidor:
```bash
go run main.go
```

O servidor estará rodando em `http://localhost:8080`

## 🔐 Autenticação

A API utiliza autenticação JWT (JSON Web Token). Para acessar as rotas protegidas, você precisa:

1. Registrar um usuário:
```http
POST /register
Content-Type: application/json

{
    "username": "seu_usuario",
    "password": "sua_senha",
    "email": "seu_email@exemplo.com"
}
```

2. Fazer login:
```http
POST /login
Content-Type: application/json

{
    "username": "seu_usuario",
    "password": "sua_senha"
}
```

3. Usar o token retornado no header de todas as requisições subsequentes:
```
Authorization: Bearer seu_token_jwt
```

## 📚 Endpoints da API

### Rotas Públicas
- `POST /register` - Registro de novo usuário
- `POST /login` - Login de usuário

### Rotas Protegidas (requerem autenticação)
- `GET /cards` - Lista todos os cards
- `GET /cards/:id` - Obtém um card específico
- `POST /cards` - Cria um novo card
- `PUT /cards/:id` - Atualiza um card existente
- `DELETE /cards/:id` - Remove um card

## 📦 Estrutura do Projeto

```
.
├── controller/     # Controladores da API
├── middleware/     # Middlewares (autenticação)
├── model/         # Modelos de dados
├── service/       # Lógica de negócios
├── main.go        # Ponto de entrada da aplicação
├── go.mod         # Dependências do projeto
└── go.sum         # Checksums das dependências
```

## 🔄 Modelo de Dados

### Card
```go
type Card struct {
    ID          int     `json:"id"`
    Name        string  `json:"name"`
    ManaCost    string  `json:"manaCost"`
    Type        string  `json:"type"`
    Rarity      string  `json:"rarity"`
    Set         string  `json:"set"`
    Power       *int    `json:"power,omitempty"`
    Toughness   *int    `json:"toughness,omitempty"`
    Description string  `json:"description"`
    ImageURL    string  `json:"imageUrl"`
}
```

### User
```go
type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Password string `json:"-"`
    Email    string `json:"email"`
}
```

## 🔒 Segurança

- Senhas são armazenadas com hash usando bcrypt
- Autenticação JWT com expiração de 24 horas
- Middleware de autenticação para rotas protegidas
- Validação de dados de entrada
