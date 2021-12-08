# Imersão FullCycle 5

## Desafio GO

### Container de Desenvolvimento

Para acessar o container de desenvolvimento execute os seguintes comandos:

```shell
docker-compose up -d

docker-compose exec app bash
```

Dentro do container de desenvolvimento execute o seguinte comando para executar a aplicação

```shell
go run cmd/main.go
```

### Container de produção

Para rodar o container de produção execute os seguintes comandos:

```shell
docker-compose -f docker-compose.prod.yaml up -d
```

Para acessar o container de produção execute o seguinte comando:

```shell
docker-compose exec app bash
```