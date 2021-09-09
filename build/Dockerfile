FROM golang:alpine
WORKDIR /go/src/app
COPY ../ .
ENV USER=go \
    UID=1000 \
    GID=1000 \
    CGO_ENABLED=0

RUN go mod tidy && \
    go generate ./... && \
    go build -ldflags="-s -w" \
    -o witch && \
    addgroup --gid "$GID" "$USER" && \
    adduser \
    --disabled-password \
    --gecos "" \
    --home "$(pwd)" \
    --ingroup "$USER" \
    --no-create-home \
    --uid "$UID" \
    "$USER" && \
    chown "$UID":"$GID" /go/src/app/witch && \
    apk add --no-cache tzdata

USER 1000
ENTRYPOINT ["/go/src/app/witch"]