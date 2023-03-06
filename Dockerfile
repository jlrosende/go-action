FROM golang:1.19 as build

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o app .

FROM alpine

COPY --from=build /app/app /app

ENTRYPOINT [ "/app" ]
