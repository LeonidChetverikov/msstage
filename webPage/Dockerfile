FROM golang:1.16-alpine

WORKDIR /app

COPY webPage/go.mod ./
RUN go mod download

COPY webPage/*.go ./

RUN go build -o /msstage

EXPOSE 8080

CMD [ "/msstage" ]