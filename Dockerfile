FROM golang:latest

ENV TZ=Europe/Berlin
ARG VERSION=1.00
ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update -qq && apt-get -y upgrade
ENV TESSDATA_PREFIX=/usr/share/tesseract-ocr/5/tessdata/
RUN apt-get install -y -qq libtesseract-dev libleptonica-dev tesseract-ocr-deu cron && apt-get -y autoclean

RUN mkdir -p /build
ADD go.mod go.sum /build/

WORKDIR /build

RUN go mod download

ADD *.go /build

RUN mkdir -p /app/ && go build -o /app/blitzerbot .

WORKDIR /app

ADD crontab /etc/cron.d/blitzerbot-cron
RUN chmod 0644 /etc/cron.d/blitzerbot-cron
RUN crontab /etc/cron.d/blitzerbot-cron

ADD docker-entry.sh /docker-entry.sh
ENTRYPOINT ["/docker-entry.sh"]
CMD ["cron","-f", "-l", "2"]
