### Instruções

#### 1 - Executar projeto:
```shell
docker compose up -d --build
```

#### 2 - Criar Leilão
```shell
curl --location 'http://localhost:8080/auction' \
--header 'Content-Type: application/json' \
--data '{
    "product_name": "Nome do produto",
    "category": "Categoria",
    "description": "Descrição do produto",
    "condition": 0
}'
```

#### 3 - Aguarde o tempo configurado em AUCTION_INTERVAL no arquivo [.env](cmd/auction/.env)

#### 4 - Busque pelo status 1.
```shell
curl --location --request GET 'http://localhost:8080/auction?status=1' \
--header 'Content-Type: application/json'
```

Códigos de Status:

| Código | Status    |
|--------|-----------|
| 0      | Active    |
| 1      | Completed |
