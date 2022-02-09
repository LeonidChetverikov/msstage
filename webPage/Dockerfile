FROM golang:1.16-alpine


COPY *.go ./

RUN go build -o /msstage

EXPOSE 8080

CMD [ "/msstage" ]