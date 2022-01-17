FROM golang:1.17

RUN mkdir /app
ADD . /app
WORKDIR /app

# Download necessary Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /csv-reader

EXPOSE 3000

CMD [ "/csv-reader" ]