# Go Studies

Repositório para registrar estudos, exercícios e projetos práticos em Go.

## Estrutura

```txt
.
├── fundamentos/          # exemplos pequenos e isolados
├── projetos/             # aplicações completas ou estudos maiores
├── anotacoes/            # rascunhos, listas de estudo e observações
└── README.md             # índice principal do repositório
```

## Como usar este repositório

- Use `fundamentos/` para praticar conceitos específicos de Go.
- Use `projetos/` para APIs, CLIs, serviços e aplicações com mais estrutura.
- Use `anotacoes/` para registrar dúvidas, decisões técnicas e próximos temas.
- Cada projeto maior deve ter seu próprio `README.md`, `go.mod` e instruções de execução.

## Fundamentos

| Tema | Pasta | Status |
| --- | --- | --- |
| Olá mundo | [`fundamentos/01-ola-mundo`](fundamentos/01-ola-mundo) | iniciado |

Para executar um exemplo:

```bash
cd fundamentos/01-ola-mundo
go run main.go
```

## Projetos

| Projeto | Descrição | Tecnologias |
| --- | --- | --- |
| [`REST-API-1`](projetos/REST-API-1) | Primeira API REST para estudo de handlers, models, repositories e use cases. | Go |
| [`REST-API-2`](projetos/REST-API-2) | API REST de produtos com PostgreSQL e Docker. | Go, Gin, PostgreSQL, Docker |
| [`go-link-shortener`](projetos/go-link-shortener) | API de encurtamento de URLs. | Go, Chi, PostgreSQL, Redis, Docker |

## Tópicos estudados

- Variáveis
- Structs
- Ponteiros
- Concorrência
- APIs REST
- Organização em camadas
- Repositories e use cases
- PostgreSQL
- Docker

## Próximos tópicos

### Banco de dados e ORMs

1. `database/sql`
2. `sqlx`
3. Escolher e comparar:
   - `GORM`
   - `Bun`
4. Estudar mais à frente:
   - `Ent`
   - `sqlc`

### Boas práticas em Go

- Testes com `testing`
- Interfaces pequenas
- Tratamento de erros
- Context propagation com `context.Context`
- Estrutura de projeto com `cmd/`, `internal/` e `pkg/`
- Configuração por variáveis de ambiente

## Convenção para novos estudos

Para exemplos pequenos:

```txt
fundamentos/02-nome-do-tema/
└── main.go
```

Para projetos:

```txt
projetos/nome-do-projeto/
├── cmd/
├── internal/
├── go.mod
└── README.md
```
