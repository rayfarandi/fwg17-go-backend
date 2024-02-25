FROM golang:latest

WORKDIR /app
COPY . .


# dijalankan saat build image
RUN go mod tidy



EXPOSE 8080

# dijalankan setiap kali docker container di jalankan 
CMD go run .            

# komentar