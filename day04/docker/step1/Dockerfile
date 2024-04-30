FROM golang:alpine

WORKDIR /go/src/app

EXPOSE 8080

COPY go.mod go.sum ./
RUN go mod download && go mod verify

ENV SERVER_PORT=8080
ENV SERVER_HOST=0.0.0.0
ENV HELLO_MESSAGE=world

ARG JWT_SECRET

COPY . .
RUN go build -v -o main .

ENTRYPOINT [ "./main" ]

CMD [ "go", "run", "./main" ]