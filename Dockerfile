FROM golang:1.25

WORKDIR /app

COPY app /app

RUN go build -o GHStalker .

ENTRYPOINT ["./GHStalker"]
