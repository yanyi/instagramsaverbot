# ---------------------------
# Go builder
# ---------------------------

FROM golang:1.11.4-alpine3.8 as builder

ARG APP_HOME=/go/src/github.com/yanyi/instagramsaverbot/

RUN apk update && \
	apk add build-base \
		ca-certificates \
		git && \
	rm -rf /var/lib/apt/lists/*

WORKDIR $APP_HOME

COPY go.mod .
COPY go.sum .

RUN GO111MODULE=on go get -v -d ./...

COPY . .

WORKDIR $APP_HOME/cmd/instagramsaverbot
RUN env CGO_ENABLE=1 GOOS=linux GOARCH=386 go build -o /go/bin/instagramsaverbot

# ---------------------------
# Resulting image
# ---------------------------

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/instagramsaverbot /bin/instagramsaverbot
COPY --from=builder /go/src/github.com/yanyi/instagramsaverbot/configs/config.yml /configs/config.yml

ENTRYPOINT [ "/bin/instagramsaverbot", "-f", "/configs/config.yml" ]
