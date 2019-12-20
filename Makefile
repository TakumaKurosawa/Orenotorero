# docker #
DOCKER = docker

DOCKER_BUILD_APP = $(DOCKER) build -t orenotorero-app client/
DOCKER_BUILD_API = $(DOCKER) build -t orenotorero-api server/API_server/

DOCKER_REMOVE_IMG = $(DOCKER) rmi busybox:1.28 orenotorero-api:latest orenotorero-app:latest

# kubectl #
KUBE = kubectl

KUBE_CREATE = $(KUBE) create -f Manifest/
KUBE_DELETE = $(KUBE) delete -f Manifest/

KUBE_APPLY_APP = $(KUBE) apply -f Manifest/app.yml
KUBE_APPLY_API = $(KUBE) apply -f Manifest/api.yml
KUBE_APPLY_DB = $(KUBE) apply -f Manifest/db.yml

KUBE_DELETE_APP = $(KUBE) delete -f Manifest/app.yml
KUBE_DELETE_API = $(KUBE) delete -f Manifest/api.yml
KUBE_DELETE_DB = $(KUBE) delete -f Manifest/db.yml

help: ## 使い方
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

start: ## アプリケーションを1発で立ち上げるコマンド
	$(DOCKER_BUILD_APP)
	$(DOCKER_BUILD_API)
	$(KUBE_CREATE)

stop: ## 立ち上がっているアプリケーションを1発でクローズするコマンド
	$(KUBE_DELETE)

restart: ## アプリケーションコンテナイメージから作り直すコマンド
	$(DOCKER_REMOVE_IMG)
	$(DOCKER_BUILD_APP)
	$(DOCKER_BUILD_API)
	$(KUBE_CREATE)

clear: ## 立ち上がっているアプリケーションをクローズし、関係するイメージも削除する
	$(KUBE_DELETE)
	$(DOCKER_REMOVE_IMG)

dev-db: ## 開発用にDBコンテナを立ち上げるコマンド
	$(DOCKER) build -t orenotorero-db-dev server/DB
	$(eval PASS = $(shell read password;echo $$password))
	$(eval DB = $(shell read db;echo $$db))
	$(DOCKER) run --rm --name dev-db -e MYSQL_ROOT_PASSWORD=$(PASS) -e MYSQL_DATABASE=$(DB) -d -p 30002:3306 orenotorero-db-dev

dev-db-stop: ## 立ち上がっている開発用DBコンテナを削除するコマンド
	$(DOCKER) stop dev-db
