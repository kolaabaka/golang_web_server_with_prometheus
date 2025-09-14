FROM golang:latest
EXPOSE 8080
RUN mkdir golang_project
WORKDIR /golang_project
COPY . .

# CMD ["go", "mod", "tidy", ".", "&&", "go", "run", "main.go"]
