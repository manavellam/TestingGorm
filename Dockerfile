#Best Practice: Version Tag must be specified.
FROM golang:1.13
#Change to container's app directory
WORKDIR /go/src/app
#Copying the app
COPY . .
#Getting dependencies
RUN go get -d -v ./...
RUN go install -v ./...
#Exposing port
EXPOSE 8085
#Running the App
CMD ["go", "run", "main.go"]