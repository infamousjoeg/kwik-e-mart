FROM golang as builder

RUN apk update && apk add --no-cache git

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go get -d -v

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o kwikemart .

FROM scratch
COPY --from=builder /app/kwikemart /app/
EXPOSE 8080
ENTRYPOINT ["/app/kwikemart"]