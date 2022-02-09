FROM golang:1.16-alpine

COPY webPage/go.mod ./
COPY webPage/*.go ./

RUN go build -o /msstage

EXPOSE 8080

CMD [ "/msstage" ]