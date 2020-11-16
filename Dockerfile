FROM golang:alpine AS build
WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ustream -a -ldflags "-s -w" main.go && \
    mv ustream /

FROM alpine
COPY --from=build /ustream /usr/local/bin/
CMD [ "ustream" ]

