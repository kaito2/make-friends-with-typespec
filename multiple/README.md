# Multiple API Server

複数の OpenAPI YAML によって定義された API をホスティングするサンプルコードです。

具体的には、以下2つをホスティングしています。

- 管理者向けの API サーバー (`/api/admin/v1/` ではじまる)
- ユーザー向けの API サーバー (`/api/user/v1/` ではじまる)

## API 定義の編集

[`schema/README.md`](./schema/README.md) を参照してください。

## API サーバーの起動

```
$ make run
go run ./cmd/...
2025/02/25 23:21:12 Starting server on :8080
```

管理者向けの API サーバーのパスは `/api/admin/v1/` です。

```
$ curl http://localhost:8080/api/admin/v1/tenant-clients
[{"name":"example_tenant_client_1"},{"name":"example_tenant_client_2"}]
```

ユーザー向けの API サーバーのパスは `/api/user/v1/` です。

```
$ curl http://localhost:8080/api/client/v1/stores
[{"name":"example_store_1","address":{"street":"","city":""}},{"name":"example_store_2","address":{"street":"","city":""}}]
```
