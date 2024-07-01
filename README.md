# domainmodel

## 概要
ドメインモデルをgolangで記述する際のテンプレート

## 実装内容
- コントローラーの実装：gin
- リポジトリの実装 : gorm

## DB
今のところMySQLサーバーのみdockerに対応。あらかじめ以下のコンテナを起動して利用する
``` sh
docker run --name my-mysql -e MYSQL_ROOT_PASSWORD=my-secret-pw -e MYSQL_DATABASE=testdb -e MYSQL_USER=testuser -e MYSQL_PASSWORD=testpw -p 3306:3306 -d mysql:8.0
```

## TODO
- Docker対応
- テスト
