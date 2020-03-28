FROM golang:latest-alpine

#set working dir
WORKDIR /app

#copy go mod and go sum file to workdir
COPY go.mod go.sum ./

#download dependencies
RUN go mod download

#copy source to workdir
COPY . .

#build the app
RUN go build -o app .

#expose port 8000
EXPOSE 8000

#run the executable
CMD ["./app"]