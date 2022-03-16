FROM golang:alpine
WORKDIR /go/src/app
COPY ../ .
ENV USER=go \
    UID=1000 \
    GID=1000 \
    CGO_ENABLED=0

RUN apk add --no-cache tzdata git && \
    go mod tidy && \
    go generate ./... && \
    go build -ldflags="-s -w" \
    -buildvcs=false \
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
    chown "$UID":"$GID" /go/src/app/witch

USER 1000
ENTRYPOINT ["/go/src/app/witch"]