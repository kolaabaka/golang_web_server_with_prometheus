FROM golang:latest
EXPOSE 8080
RUN mkdir golang_project
WORKDIR /golang_project
COPY . .

# CMD sleep infinity

RUN rm prometheus.yml && go mod tidy && go build cmd/main.go

CMD ["./main"]
