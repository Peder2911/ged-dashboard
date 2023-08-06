.DEFAULT_TARGET: build
.PHONY: build

build:
	docker build -t reg.rackness.net/peder2911/ged-dashboard -f build/package/Dockerfile .
