## Write Go GIN RESTful API Service include below Condition (Push to Github)

HR System Backend API (沒有限定格式 目的是看看 Candidate 會怎麼思考)

```text
route:
- POST /employees
- GET /employees
- GET /employees/:id
- PUT /employees/:id
- PATCH /employees/:id
```

MySQL as  Database

```text
go run port: 3306
docker-compose 13306 

account and password in config.yaml
```

Redis as Cache

```text
go run port: 6379
docker-compose 16379 
```

GORM Migration

```text
server on start will auto migrate
```

GORM MySQL SEED data

```text
server on start will auto create seed data
```

Unit Test

```text
internal/service/employee_test.go
```

Makefile for build and deploy

```text
make image-build

# 推到 docker hub
make image-push

沒辦法 deploy 沒有架服務
```

Run all service using docker-compose

```text
啟動服務
make run

停止服務
make stop

重新啟動服務
make restart

查看服務狀態
make status

查看服務 log
make logs
```
