FROM golang:alpine AS builder

#RUN mkdir -p /go/src/github.com/tyybbi/guiching

WORKDIR /guiching

COPY . /guiching

RUN CGO_ENABLED=0 GO111MODULE=off GOOS=linux GOARCH=amd64 go build -ldflags="-w -s"

FROM scratch

WORKDIR /guiching

COPY --from=builder /guiching /guiching

EXPOSE 3000

ENTRYPOINT ["/guiching/guiching"]

