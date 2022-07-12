FROM golang:alpine
ENV CGO_ENABLED=0

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o main ./cmd/auth
ADD ./cmd/auth/ /usr/local/monogo/pem

EXPOSE 8080

CMD [ "./main" ]