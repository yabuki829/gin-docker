以下は、Gin を使って `/api/hello` にアクセスした際に「Hello World」と返すハンズオンのサンプルになります。

---

### 1. Gin のインストール

まず、Go のプロジェクトディレクトリを作成し、Go モジュールを初期化します。

```bash
mkdir gin-handson
cd gin-handson
go mod init example.com/gin-handson
```

次に、Gin をインストールします。

```bash
go get -u github.com/gin-gonic/gin
```

---

### 2. コードの作成

`main.go` というファイルを作成し、以下のコードを記述する

```go
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    // Gin のデフォルトルーターを作成
    r := gin.Default()

    // GET /api/hello に対するルートハンドラーを定義
    r.GET("/api/hello", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello World")
    })

    // サーバーをポート8080で起動
    r.Run(":8080")
}
```

#### 説明

- **gin.Default()**  
  Gin のデフォルトのミドルウェア（ロギング、リカバリ）が設定されたルーターを生成します。

- **r.GET("/api/hello", func(c *gin.Context) {...})**  
  `/api/hello` への GET リクエストを受け付けるルートを定義し、リクエストがあった際に「Hello World」という文字列を返します。

- **r.Run(":8080")**  
  サーバーをポート8080で起動します。これで `http://localhost:8080/api/hello` にアクセスすると「Hello World」が返ります。

---

### 3. サーバーの起動と動作確認


```bash
go run main.go
```

```bash
http://localhost:8080/api/hello
```

