# 使用するGoのイメージを指定
FROM --platform=linux/amd64 golang:1.22

# 作業ディレクトリを設定
WORKDIR /go/src

# 必要なツールと依存関係をインストール
RUN apt-get update && apt-get install -y git curl

# air をダウンロードしてインストール
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b /go/bin

# プロジェクトルートのgo.modとgo.sumをコピー
COPY src/go.mod src/go.sum ./

# モジュールを初期化し、依存関係を整理
RUN go mod tidy

# ソースコードと設定ファイルをコピー
COPY src/ .

# .air.toml ファイルをコピー
COPY src/.air.toml /go/src/.air.toml

# デフォルトコマンドを設定
CMD ["air"]
