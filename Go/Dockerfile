# Start from base image
FROM golang:latest

# Set the current working directory inside the container
WORKDIR /app

RUN go get "github.com/go-sql-driver/mysql"
RUN go get "github.com/gin-gonic/gin"

# Copy source from current directory to working directory
COPY . .

# Build the application
#RUN go build -o cmd/app/main.go .

# Expose necessary port
#EXPOSE 8080

# Run the created binary executable after wait for mysql container to be up
#CMD ["./wait-for.sh" , "db:3307" , "--timeout=300" , "--" , "./main"]