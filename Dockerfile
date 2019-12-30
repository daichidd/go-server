FROM golang:1.12-alpine3.10

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
ENV GO111MODULE=on

WORKDIR $GOPATH/src/DaichiEndo
ADD . $GOPATH/src/DaichiEndo

# wait for other container
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.6.0/wait /wait 
RUN chmod +x /wait

RUN apk add --no-cache \
        alpine-sdk \
        git \
    && go get github.com/pilu/fresh

CMD ["fresh", "go", "run", "main.go"]
