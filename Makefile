# 定義變數
DOCKER_COMPOSE_FILE = "./docker/docker-compose.yml"

# 編譯並構建應用程序
image-build:
	@docker build -f ./docker/Dockerfile -t mattwang13/fliqt:latest .

image-push:
	@docker push mattwang13/fliqt:latest

image-rmi:
	@docker image rm mattwang13/fliqt:latest

# 啟動服務 (Docker Compose)
run:
	@docker-compose -f $(DOCKER_COMPOSE_FILE) up --build

# 停止服務
stop:
	@docker-compose -f $(DOCKER_COMPOSE_FILE) down --volumes --remove-orphans

# 查看服務狀態
status:
	@docker-compose -f $(DOCKER_COMPOSE_FILE) ps

# 查看服務日誌
logs:
	@docker-compose -f $(DOCKER_COMPOSE_FILE) logs -f

# 重新啟動服務
restart: stop run
