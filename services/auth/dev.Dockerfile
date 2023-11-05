FROM golang:1.19-alpine

WORKDIR /app

COPY ../shared ./shared
COPY ../auth ./service

RUN cd /app/service && go mod download
RUN cd /app/shared && go mod download

RUN cd /app/service && go get github.com/githubnemo/CompileDaemon
RUN cd /app/service && go install github.com/githubnemo/CompileDaemon

ENTRYPOINT cd /app/service && CompileDaemon -polling -log-prefix=false -build="go build -o ./build/auth" -command="./build/auth APP_ENV=DEVELOPMENT" -directory="./"
