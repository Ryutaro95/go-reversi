# reversi webapp
* フロントエンドはHTML,JS,CSSで作成する
* サーバサイドはGoでAPIを作成
* フロントエンドのJSからGoのAPIを叩いて結果を表示する

## サーバサイド
* Go
  * Framework: Gin
* DB: MySQL
* 以下エンドポイントを用意
  * /api/games
    * GET
    * POST
  * /api/games/latest/turns/{turnCount}
    * GET
      * 最新の盤面やターン数、次はどちらのターンなのかなどの情報がレスポンスとして返ってくる
        ```json
        {
            "turn_count": 1,
            "board": [
                [0, 0, 0, 0, 0, 0, 0, 0],
                [0, 0, 0, 0, 0, 0, 0, 0],
                [0, 0, 0, 0, 0, 0, 0, 0],
                [0, 0, 0, 1, 2, 0, 0, 0],
                [0, 0, 0, 2, 1, 0, 0, 0],
                [0, 0, 0, 0, 0, 0, 0, 0],
                [0, 0, 0, 0, 0, 0, 0, 0],
                [0, 0, 0, 0, 0, 0, 0, 0],
            ],
            "next_disc": 1,
            "winner_disc": 1
        }
        ```
  * /api/games/latest/turns
    * POST
