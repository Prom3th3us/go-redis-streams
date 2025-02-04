docker:
	@docker-compose down
	@docker-compose build
	@docker-compose up -d

dockerdown:
	@docker-compose down

build:
	@echo "---- Building Application ----"
	@go build -o consumer consumer/*.go
	@go build -o producer producer/*.go

consume:
	@echo "---- Running Consumer ----"
	@export REDIS_HOST=localhost
	@export STREAM=events
	@export GROUP=GroupOne
	@go run consumer/*.go

run:
	@echo "---- Running Producer ----"
	@export REDIS_HOST=localhost
	@export STREAM=events
	@go run producer/*.go

docker-build:
	@docker login --username miguelemos --password Alatriste007
	@echo "---- Building docker images ----"
	@echo "---- -building: redis_consumer:latest ----"
	@docker build . -t miguelemos/redis_consumer:latest --network=host -f consumer.docker
	@echo "---- -building: redis_producer:latest ----"
	@docker build . -t miguelemos/redis_producer:latest --network=host -f producer.docker
	@docker push miguelemos/redis_producer:latest
	@docker push miguelemos/redis_consumer:latest
