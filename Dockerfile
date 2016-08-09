FROM golang:1.6-wheezy

RUN apt-get update -y
RUN apt-get install -y tesseract-ocr libleptonica-dev libtesseract-dev wamerican-large

ADD . /go/src/github.com/bittersweet/imgfiltrate/

RUN ls /go/src/github.com/bittersweet/imgfiltrate/
RUN cd /go/src/github.com/bittersweet/imgfiltrate/ && go get -v
RUN cd /go/src/github.com/bittersweet/imgfiltrate/ && go install -v
