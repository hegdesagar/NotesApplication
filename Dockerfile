FROM golang:1.10

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . /app

#ENV GOPATH=/app

#RUN go build note
#RUN chmod u+s ./note

#RUN ./note read a




