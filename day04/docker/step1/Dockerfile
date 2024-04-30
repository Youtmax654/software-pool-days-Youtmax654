FROM golang:alpine

WORKDIR /go/src/app

EXPOSE 8080

COPY go.mod go.sum ./
RUN go mod download && go mod verify

ENV SERVER_PORT=8080
ENV SERVER_HOST=0.0.0.0
ENV HELLO_MESSAGE=world
ENV JWT_SECRET=e6054a44-6c8b-4a3e-8340-ccb4ea447452

COPY . .
RUN go build -v -o main .

ENTRYPOINT [ "./main" ]

CMD [ "go", "run", "./main" ]