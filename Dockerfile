FROM golang:1.20.2-bullseye
RUN mkdir /app
ADD . /app 
WORKDIR /app
RUN mkdir data
RUN cp examples/config.yml ./data/config.yml
RUN go build -o inn cmd/inn.go
ENV GIN_MODE=release
EXPOSE 8080
CMD /app/inn --cfg data/config.yml --data data/
