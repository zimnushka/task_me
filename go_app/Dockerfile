FROM golang:1.19

WORKDIR /Users/kirillzimnuhov/dev/golang/taskMe/

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod main.go ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o main .

CMD ["./main"]