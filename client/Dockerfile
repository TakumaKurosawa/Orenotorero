# コンテナイメージの指定
FROM node:11.14.0
# ワーキングディレクトリの指定
WORKDIR /usr/src/app
# パッケージ管理用ファイルの転送
COPY ./package*.json ./
# コピーしたpackageファイルを用いてnpm installを走らせる　オプションは本番環境用
RUN yarn install 
# データを転送
COPY . .
# サービス開始用のシェルを実行
CMD ["bash","start.sh"]

