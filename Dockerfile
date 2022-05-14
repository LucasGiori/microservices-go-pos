
FROM golang:1.18 as base

FROM base as dev

RUN apt-get update && apt-get install -y telnet iputils-ping;

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /app

CMD ["air"]