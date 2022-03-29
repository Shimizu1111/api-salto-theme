# RESTの設計

### RESTについて
* REST APIを作るときはシンプルであることが鉄則

### リソースの決定
* APIのユーザが必要としている情報資源(リソース)
  * APIユーザはどんな情報を必要としているか？ブログで例えると記事,コメント,タグ,カテゴリなどの情報
* リソースに対して必要な操作(CRUD)
  * そのリソースに対してどんな操作が必要か？ブログの記事は参照,作成,更新,削除が必要
* リソースの親子関係
  * それらのリソースにはどんな関係があるか。記事にはコメントが複数付き、記事はカテゴリに属す



### URLの設計（Cool URI）
* ひと目でAPIと分かるようなURLにする
  * http://api.example.com → URLにサブドメインを切る(example.comのサブドメインとしてのapi.example.com)
* URLにAPIのバージョンを含める
  * http://api.example.com/v1/article  → v1などのバージョンを記載
* URLに動詞を含めず、複数形の名詞のみで構成する
  * http://api.example.com/v1/createArticle → createなどの動詞を含めない
* アプリケーションや言語に依存する拡張子は含めない
  * http://api.example.com/v1/articles/126/comments → test.txtなどの拡張子を含めない
* リソースの関係性がひと目で分かるようにする
  * http://api.example.com/v1/articles/126/comments → URLを見ただけで、リソースがどこに属しているか分かるようにする
* URLを長くしすぎない
  * エンドポイントをなるべく短くする
* 大文字小文字が混在していないURI
* サーバ側のアーキテクチャの情報が記載されていないURI
* 単語をつなげる必要がある場合はハイフンを利用する
* スペースやエンコードを必要とする文字を使わない


### 参考記事
https://wp.tech-style.info/archives/683
https://qiita.com/NagaokaKenichi/items/6298eb8960570c7ad2e9
