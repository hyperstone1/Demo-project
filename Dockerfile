FROM golang:latest

RUN go version
WORKDIR /go/src/
COPY ./ ./

RUN chmod +x wait-for-postgres.sh

# RUN cat /flyway/run_migration

RUN go mod download
RUN go build -o ./main.go .
CMD [ "./main" ]
