# N+1問題 検証用リポジトリ

Qiita の記事「[[解説] SQLクエリのN+1問題](https://qiita.com/muroya2355/items/d4eecbe722a8ddb2568b)」のために作成したリポジトリ

# 構成
- Docker
- Golang
- PostgreSQL
- テーブルについては記事を参照してください

# 環境構築方法

- ホストOS: リポジトリのクローン
```
   > git clone https://github.com/muroya2355/go_n_1
   > cd ~/go_n_1
```

- ホストOS: Docker-Compose の起動
```
    > docker-compose up -d
```

- ホストOS: golang コンテナにログイン
```
    > docker container exec -it app /bin/bash
```

# 検証方法

- golangコンテナ: 各手法を実行してみる
```
    # go run n_1.go
    # go run join.go
    # go run map.go
```

- golangコンテナ: テストデータの追加
```
    書籍データを20万件、利用者データを10万件追加する
    (10分くらいかかります)
    # go run testdata.go
```

- golangコンテナ: もう一度、各手法を実行する
```
    # go run n_1.go
    # go run join.go
    # go run map.go
```