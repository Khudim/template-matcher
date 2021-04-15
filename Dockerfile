FROM gocv/opencv:latest

WORKDIR $GOPATH/src/template-matcher

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/template-matcher .

EXPOSE 8080

CMD ["./out/template-matcher"]