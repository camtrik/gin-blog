# Gin Blog System (Updating)

## Language

- [English](#english)
- [日本語](#日本語)

---

## English

## Overview
A blog system based on [Gin](https://github.com/gin-gonic/gin).

## Features Completed
- Article release and modification
- Article tags management
- Upload images
- API Access Control by [JWT](https://github.com/dgrijalva/jwt-go)
- Opentracing by [Jaeger](http://github.com/uber/jaeger-client-go) 
- API [ratelimiter](https://github.com/juju/ratelimit)
- Hot reload


## TODO List
- [x] Upload image and files
- [x] API Access Control by [JWT](https://github.com/dgrijalva/jwt-go)
- [x] Integrate common middleware for enhanced logging and error handling.
- [x] [Ratelimiter](https://github.com/juju/ratelimit) & time out controller in middleware.
- [x] Opentracing by [Jaeger](http://github.com/uber/jaeger-client-go)
- [ ] Opentracing SQL?
- [x] Graceful shutdown
- [x] Hot reload configs


### JWT Authorization
Use JWT for API access control

Generate a token
```bash
curl -X POST \
  'http://127.0.0.1:8000/auth' \
  -d 'app_key=ebbi' \
  -d 'app_secret=gin-blog'
```

To access protected API endpoints, include the generated JWT token in the request header.

Use the `token` header to include the token: 
```bash
curl -X GET http://127.0.0.1:8000/api/v1/tags -H 'token: generated_token_here'
```

Replace `generated_token_here` with the actual token received from the authentication endpoint.

### Tag Management
1. **Create Tag**
    ```bash
    curl -X POST http://127.0.0.1:8000/api/v1/tags -F 'name=Go' -F created_by=ebbi
    ```

2. **Update Tag**
    ```bash
    curl -X PUT http://127.0.0.1:8000/api/v1/tags/1 -F name=cpp -F modified_by=ebbi
    ```

3. **Delete Tag**
    ```bash
    curl -X DELETE http://127.0.0.1:8000/api/v1/tags/1
    ```

4. **Get Tag List**
    ```bash
    curl -X GET 'http://127.0.0.1:8000/api/v1/tags?page=1&page_size=2'
    ```

### Article Management
1. **Create Article**
    ```bash
    curl -X POST http://127.0.0.1:8000/api/v1/articles -F 'title=Go Tutorial' -F 'desc=go test 1 desc' -F 'content=go test 1 content' -F created_by=ebbi
    ```

2. **Add Tag to Article**
    ```bash
    curl -X POST http://127.0.0.1:8000/api/v1/articles/:id/tags -F tag_id=3 -F created_by=ebbi
    ```

3. **Remove Tag from Article**
    ```bash
    curl -X DELETE http://127.0.0.1:8000/api/v1/articles/:id/tags -F tag_id=3
    ```

4. **Get Article List by Tag**
    ```bash
    curl -X GET 'http://127.0.0.1:8000/api/v1/articles?page=1&page_size=2' -F tag_id=3
    ```

5. **Get Article by ID**
    ```bash
    curl -X GET http://127.0.0.1:8000/api/v1/articles/:id
    ```

6. **Update Article**
    ```bash
    curl -X PUT http://127.0.0.1:8000/api/v1/articles/:id -F 'title=Go Tutorial 1' -F modified_by=ebbi
    ```

7. **Delete Article**
    ```bash
    curl -X DELETE http://127.0.0.1:8000/api/v1/articles/:id
    ```

### Update Image

```bash 
curl -X POST http://127.0.0.1:8000/upload/file -F file=@{file_path} -F type=1
```

## 日本語

## 概要
[Gin](https://github.com/gin-gonic/gin) に基づいたブログシステム。

## 完成済み機能
- 記事の管理
- 記事タグの管理
- 画像アップロード
- [JWT](https://github.com/dgrijalva/jwt-go)によるAPIアクセス制御
- [Jaeger](https://github.com/uber/jaeger-client-go)によるオープントレーシング
- API [レートリミッター](https://github.com/juju/ratelimit)
- ホットリロード

## TODO リスト
- [x] Upload image and files
- [x] API Access Control: Implement authentication to secure the API endpoints.
- [x] Integrate common middleware for enhanced logging and error handling.
- [x] [Ratelimiter](https://github.com/juju/ratelimit) & time out controller in middleware.
- [x] Opentracing by jaeger (logger)
- [ ] Opentracing SQL?
- [x] Graceful shutdown
- [x] Hot reload configs

### JWT 認証
API アクセス制御に JWT を使用

トークンの生成
```bash
curl -X POST \
  'http://127.0.0.1:8000/auth' \
  -d 'app_key=ebbi' \
  -d 'app_secret=gin-blog'
```

保護された API エンドポイントにアクセスするには、リクエストヘッダーに生成された JWT トークンを提示します。

`token` ヘッダーを使用します：
```bash
curl -X GET http://127.0.0.1:8000/api/v1/tags -H 'token: generated_token_here'
```

`generated_token_here` を認証エンドポイントから受け取った実際のトークンに置き換えてください。

### タグ管理
1. **タグの作成**
    ```bash
    curl -X POST http://127.0.0.1:8000/api/v1/tags -F 'name=Go' -F created_by=ebbi
    ```

2. **タグの更新**
    ```bash
    curl -X PUT http://127.0.0.1:8000/api/v1/tags/1 -F name=cpp -F modified_by=ebbi
    ```

3. **タグの削除**
    ```bash
    curl -X DELETE http://127.0.0.1:8000/api/v1/tags/1
    ```

4. **タグリストの取得**
    ```bash
    curl -X GET 'http://127.0.0.1:8000/api/v1/tags?page=1&page_size=2'
    ```

### 記事管理
1. **記事の作成**
    ```bash
    curl -X POST http://127.0.0.1:8000/api/v1/articles -F 'title=Go Tutorial' -F 'desc=go test 1 desc' -F 'content=go test 1 content' -F created_by=ebbi
    ```

2. **記事にタグを追加**
    ```bash
    curl -X POST http://127.0.0.1:8000/api/v1/articles/:id/tags -F tag_id=3 -F created_by=ebbi
    ```

3. **記事からタグを削除**
    ```bash
    curl -X DELETE http://127.0.0.1:8000/api/v1/articles/:id/tags -F tag_id=3
    ```

4. **タグによる記事リストの取得**
    ```bash
    curl -X GET 'http://127.0.0.1:8000/api/v1/articles?page=1&page_size=2' -F tag_id=3
    ```

5. **ID で記事を取得**
    ```bash
    curl -X GET http://127.0.0.1:8000/api/v1/articles/:id
    ```

6. **記事の更新**
    ```bash
    curl -X PUT http://127.0.0.1:8000/api/v1/articles/:id -F 'title=Go Tutorial 1' -F modified_by=ebbi
    ```

7. **記事の削除**
    ```bash
    curl -X DELETE http://127.0.0.1:8000/api/v1/articles/:id
    ```

### 画像のアップロード

```bash 
curl -X POST http://127.0.0.1:8000/upload/file -F file=@{file_path} -F type=1
```
