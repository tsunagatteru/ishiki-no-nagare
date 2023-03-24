FROM golang:1.20.2-bullseye
RUN mkdir /app
ADD . /app 
WORKDIR /app
RUN mkdir data
RUN cp config.yml.default ./data/config.yml
CMD ["/app/inn --cfg data/config.yml --data data/"]
