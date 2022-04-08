#!/bin/bash

version=v0.0.1

run:
	go run .

# docker build -t sgfoot/checkbox:v0.0.1 .

# docker run -it --rm --name checkbox e28a146bb891 /bin/bash
# docker run -itd --rm -p 8080:80 --name checkbox b1264b698ba4 /bin/bash
# docker run -itd --rm -p 80:80 0f898e7294a3 bash
docker-build:
	docker build -t sgfoot/checkbox:$(version) .


test-func: 
	go test -v -run 测试函数名称