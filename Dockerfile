FROM golang:1.20.2-alpine3.16

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the app folder inside the container
COPY . /app

# Create a group and user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

#build the application
RUN go build note

#setUid
RUN chmod u+s ./note

# Tell docker that all future commands should run as the appuser user
USER appuser










