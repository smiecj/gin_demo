swagger:
	swag init

build: swagger
	go build -o main_gin main.go

stop:
	ps -ef | grep "main_gin" | grep -v grep | awk '{print $$2}' | xargs --no-run-if-empty kill -9

run:build stop
	nohup ./main_gin > test.log 2>&1 &

curl:
	curl http://localhost:8000/ping