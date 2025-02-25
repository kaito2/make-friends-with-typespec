# Schema

このディレクトリでは、 TypeSpec を用いて API のスキーマを管理しています。

## ディレクトリ構造

- `admin`: 管理者向けのAPIのスキーマ
- `client`: ユーザー向けのAPIのスキーマ

## OpenAPI YAML の生成

```bash
npm run gen:all
```

## TypeSpec の編集

以下のコマンドを実行しつつ、 TypeSpec を編集すると、 OpenAPI YAML が自動で生成されます。

```bash
# 管理者向けのAPIのスキーマを編集
npm run gen:admin:watch

# ユーザー向けのAPIのスキーマを編集
npm run gen:client:watch
```
