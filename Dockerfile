FROM golang:latest

WORKDIR /app
COPY . .

# dijalankan saat build image
RUN go mod tidy




EXPOSE 8814

# dijalankan saat docker container di jalankan
CMD go run .            

# komentar