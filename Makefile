godep:
	go get github.com/tools/godep
	godep restore ./...

docker-build:
	docker build -t newmotion .

docker-run:
	docker run --publish 6680:6680 --name newmotion --rm newmotion

setup: godep

test: godep
	godep go test ./...

