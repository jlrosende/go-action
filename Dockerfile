FROM --platform=${BUILDPLATFORM} golang:1.21-alpine as build

ARG TARGETOS="linux"
ARG TARGETARCH="amd64"

WORKDIR /src

ENV CGO_ENABLED=0

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -a -o /out/sisu .

FROM alpine

COPY --from=build /out/sisu /sisu
COPY entrypoint.sh /entrypoint.sh

ENTRYPOINT [ "/entrypoint.sh" ]
