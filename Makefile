
COMPONENT := deis/mock-http-server
BUILD_IMAGE := $(COMPONENT)-build:latest
RELEASE_IMAGE := $(COMPONENT):latest
BINARY := mock-http-server

build:
	docker build -t $(BUILD_IMAGE) .
	docker cp `docker run -d $(BUILD_IMAGE)`:/go/bin/$(BINARY) image/
	docker build -t $(RELEASE_IMAGE) image
	rm -rf image/$(BINARY)

clean:
	docker rmi -f $(BUILD_IMAGE) $(RELEASE_IMAGE)
