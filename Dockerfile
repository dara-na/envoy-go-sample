FROM golang:1.16

WORKDIR /go/src/app
COPY app .

RUN ls
RUN CGO_ENABLED=0 go install 
ENTRYPOINT ["sample-api"]
