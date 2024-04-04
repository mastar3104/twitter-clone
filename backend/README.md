# Twitter Clone BackEnd

## 概要

Twitter CloneのバックエンドAPI

## ローカル実行方法

- golangのinstall(バージョンはgo.mod参照)

```
$ go version
go version go1.21.5 darwin/arm64
```

- docker, docker-composeのinstall

```
$ docker --version
Docker version 24.0.7
$ docker-compose --version
Docker Compose version 2.23.3
```

- `.env` のコピー
```
$ cp .env.example .env
```

- make dev(ローカル起動コマンド)

```
$ make dev
docker-compose -f ../db/docker-compose.yaml up -d
[+] Running 3/3
# dockerのMySqlが立ち上がるまで多少時間がかかります
go run main.go
⇨ http server started on [::]:8080
```

- http://localhost:8080 にアクセス

## API仕様書

*echo-swaggerにて自動生成のSwaggerを提供予定*

### tweet

#### タイムライン取得API

userIdに紐づくユーザのタイムラインを取得する。

| メソッド |         path         |
|:----:|:--------------------:|
| GET  | /v1/timeline/:userId |

- レスポンスボディ

```
{
    "tweets": [{
        "userId": "ユーザID",
        "userName": "ユーザ名",
        "tweetTime": 123456789,
        "content": "投稿内容"
    }]
}
```

#### ツイート取得API

userIdに紐づくユーザのツイートを取得する。

| メソッド |        path        |
|:----:|:------------------:|
| GET  | /v1/tweets/:userId |

- レスポンスボディ

```
{
    "tweets": [{
        "userId": "ユーザID",
        "userName": "ユーザ名",
        "tweetTime": 123456789,
        "content": "投稿内容"
    }]
}
```

#### ツイート投稿API

userIdに紐づくユーザのツイートを投稿する。

| メソッド |        path        |
|:----:|:------------------:|
| POST | /v1/tweets/:userId |

- リクエストボディ

```
{
    "content": "string"
}
```

- レスポンスボディ

```
{
        "userId": "ユーザID",
        "userName": "ユーザ名",
        "tweetTime": 123456789,
        "content": "投稿内容"
}
```

### user

#### ユーザ情報取得API

userIdに紐づくユーザ情報を取得する。

| メソッド |        path        |
|:----:|:------------------:|
| GET  | /v1/users/:userId/ |

- レスポンスボディ

```
{
        "userId": "xxxxx",
        "userName": "xxxxx",
}
```

#### ユーザ情報登録API

ユーザ情報を登録する。

| メソッド |    path    |
|:----:|:----------:|
| POST | /v1/users/ |

- リクエストボディ

```
{
        "userName": "xxxxx",
}
```

- レスポンスボディ

```
{
        "userId": "xxxxx",
        "userName": "xxxxx",
}
```

#### ユーザフォローAPI

ユーザIDに紐づくユーザが、リクエストされたユーザをフォローする。

| メソッド |           path           |
|:----:|:------------------------:|
| POST | /v1/users/:userId/follow |

- リクエストボディ

```
{
        "followUserId": "xxxxx",
}
```

#### ユーザ取得API

ユーザ情報の一覧を取得するAPI

| メソッド |   path    |
|:----:|:---------:|
| POST | /v1/users |

- レスポンスボディ

```
{
    "uers": [
      {
        "userId": "xxxxx",
        "userName": "xxxxx"
      }
    ]
}
```

## モジュール一覧

| モジュール名 |                 機能                  | 
|:------:|:-----------------------------------:|
| fanout | ユーザのつぶやき投稿に応じて、フォロワーに対してタイムラインを生成する |
| tweet  |    つぶやきの投稿・取得など、つぶやきに関する機能を提供する     | 
|  user  |   ユーザの取得・登録・フォローなど、ユーザに関する機能を提供する   |     

