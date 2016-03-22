.PHONY: test

CODENAME=garer
IMAGE=uisso/garer

#you can use with 6 (rasp A/B) or 7 (rasp2)
BGOARM=6

docker-build:
	docker build -t $(IMAGE_TAG) .

docker-clean:
	docker rm -f $(CODENAME) | true

docker-run: ops-clean
	docker run -d --name=$(CODENAME) --net host --env-file $(CODENAME).env $(IMAGE)
	docker logs -f --tail=100 $(CODENAME)

deps:
	go get -t ./...

clean:
	rm -rf $(CODENAME)

run: build
	./$(CODENAME)

test: deps
	go test -v ./...

doc:
	echo 'open browser with http://localhost:6060/pkg/github.com/uisso/garer/'
	godoc -http=:6060

build: clean
	GOOS=linux GOARCH=arm GOARM=$(BGOARM) go build -v -a -installsuffix cgo -o $(CODENAME)