# api-salto-them

### 概要
albumに関してのAPIを叩けるプロダクトになります。

### 仕様技術
* インフラ環境(terraform化, IaC)
  * [terraform化したプロダクト](https://github.com/Shimizu1111/tf-api-salto-theme)
  * 使用技術
    * EC2
    * VPC
* コンテナ仮想化
  * Docker
  * Docker Compose
* CI/CD
  * CircleCI(自動ビルド・自動デプロイ)
* 言語
  * Go

### 動作方法
* PRを出す
```
shell/deploy.shファイルに IP_ADDRESSES=""を記載
EC2の起動(EC2を起動しないとCIが通らない)
PRを出す
CIでエラーが出ていないことを確認
マージ
```
* 起動(ビルド時のみ)
```
EC2の起動
shell/deploy.shファイルに IP_ADDRESSES=""を記載
docker-compose up -d (--build)
```
* 停止
```
docker-compose down
sudo chown -R ec2-user:ec2-user ./
```

### APIの叩き方
* 一覧取得
```
curl --request GET 'http://{ipアドレス}:8080/albums'
```
* 新規追加
```
curl --request POST 'http://{ipアドレス}:8080/albums' \
--header 'Content-Type: application/json' \
-d '{
    "id":"5", 
    "title":"cs",
    "artist":"ra",
    "price":55.5
}'
```
* 特定IDの値を取得 
```
curl --request GET 'http://{ipアドレス}:8080/albums/2'
```


### 今後学びたい技術
* makefileの作成
* Rundeckの導入
* Datadogの導入
* GitHubのサブモジュールを使用した開発(DBをサブモジュール化して再利用できて、ローカルでの動作確認もサブモジュールで気軽にテストできるようにする)
* redashの導入
* jenkinsの導入
* kibanaの導入
* アクセストークンを取得する対応
* umlの作成


