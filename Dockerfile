##
## Build Stage
##

FROM golang:latest as build

ENV GOOS=linux \
    GOARCH=amd64 \
    GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /workspace

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /workspace/rda ./cmd/main.go

##
## Final Stage
##

FROM scratch

WORKDIR /workspace

COPY --from=build /workspace/rda .

ENTRYPOINT ["./rda "]
