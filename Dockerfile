FROM golang:1.20.2-bullseye
RUN mkdir /app
ADD . /app 
WORKDIR /app
RUN mkdir data
RUN cp config.yml.default ./data/config.yml
RUN go build -o inn cmd/inn.go
CMD /app/inn --cfg data/config.yml --data data/
