FROM golang:alpine
RUN mkdir /app
ADD . /app/
WORKDIR /appRUN go build -o main .
RUN adduser -S -D -H -h /app appuser
USER appuser
CMD [ "./main" ]