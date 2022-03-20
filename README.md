## 技術スタック

- FW: echo
- ORM: Gorm

## ディレクトリ構成

MVCベース

```bash
├── Dockerfile
├── README.md
├── app
│   ├── controllers // リクエストとDB接続の仲介役
│   ├── database // DB接続系の処理
│   ├── main.go
│   └── models // DB操作系の処理
├── docker-compose.yml
├── go.mod
└── go.sum
```

## Usage

### コンテナ立ち上げ

```
docker-compose up
```
