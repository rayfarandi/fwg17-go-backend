FROM golang

WORKDIR /app
COPY . .


RUN go mod tidy

EXPOSE 7773


CMD go run .            

