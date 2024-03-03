FROM golang:1.21.6-alpine
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN mkdir uploads
COPY . .
RUN go build -o ./out/dist .

EXPOSE 8000

#CMD ["wait-for-it", "mysql:3306", "--", "your-application-command"]
CMD ./out/dist
