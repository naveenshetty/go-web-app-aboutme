FROM golang:1.23.6 as base

WORKDIR /app

COPY go.mod .

RUN go mod download

#copy source code onto docker image
COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o main .

#######################################################
# Reduce the image size using multi-stage builds
# distroless image to run the application
FROM gcr.io/distroless/base

# Copy the binary from the previous stage
COPY --from=base /app/main .

# Copy the static files from the previous stage
COPY --from=base /app/static ./static

# Expose the port on which the application will run
EXPOSE 8080

# Command to run the application with binery file we created in stage 1
CMD ["./main"]