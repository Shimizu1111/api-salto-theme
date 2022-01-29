# api-salto-them

### インフラ環境
* インフラ環境(terraform化, IaC)
https://github.com/Shimizu1111/tf-api-salto-theme

### 動作方法
* 起動(ビルド時のみ)
```
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
