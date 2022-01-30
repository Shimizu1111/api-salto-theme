# エンドポイントの設計
/albums の場合
GET – Get a list of all albums, returned as JSON.(アルバムの一覧を取得)
POST – Add a new album from request data sent as JSON.(アルバムに追加)

/albums/:id の場合
GET – Get an album by its ID, returning the album data as JSON.(特定のアルバムを取得)


# 今後やりたいこと
* ログ出力(/var/logへの出力)
* jenkinsかcircleCIを導入して自動ビルドと自動デプロイをAWS環境に行う
  * start.shとdeploy.shの中身を実装
* Makefileの中身を作成
* AWS環境の構築
  * EC2
  * VPC
  * ECS(fargate)
  * ECR
  * ELB
  * Aurora
* AWS環境をterraformで構築
* 


