build:
	docker build -t nfmsjoeg/kwik-e-mart:latest .

run:
	summon -p summon-conjur go run main.go

docker_run:
	summon -p summon-conjur docker run --name kwik-e-mart -d \
		--restart unless-stopped \
		--env-file @SUMMONENVFILE \
		-p 8080:8080 \
		nfmsjoeg/kwik-e-mart:latest

compile:
	echo "Compiling for every OS and Platform"
	go build -o bin/kwikemart-darwin-amd64 main.go
	GOOS=darwin GOARCH=arm64 go build -o bin/kwikemart-darwin-arm main.go
	GOOS=linux GOARCH=amd64 go build -o bin/kwikemart-linux-amd64 main.go
	GOOS=windows GOARCH=amd64 go build -o bin/kwikemart-windows-amd64 main.go

create_table:
	./run.sh create_table

drop_table:
	./run.sh drop_table

insert_customers:
	./run.sh insert_customers