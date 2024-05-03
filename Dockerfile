FROM golang:1.21

WORKDIR /src
COPY go.mod go.sum ./
RUN go install github.com/antithesishq/antithesis-sdk-go/tools/antithesis-go-instrumentor@latest
RUN go mod download

COPY . ./

RUN mkdir /src_inst
RUN /go/bin/antithesis-go-instrumentor . /src_inst

WORKDIR /src_inst/customer

RUN go build -o ./main

ENTRYPOINT ["./main"]