# 2020/10/14最新versionを取得
FROM golang:1.17.6-alpine
# アップデートとgitのインストール
RUN apk update && apk add git
# appディレクトリの作成
RUN mkdir /go/src/api-salto-theme
# ワーキングディレクトリの設定
WORKDIR /go/src/api-salto-theme
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/api-salto-theme

RUN GO111MODULE=off go get github.com/oxequa/realize
CMD ["realize", "start"]