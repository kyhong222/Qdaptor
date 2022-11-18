from golang:1.19.2

copy  . /root/Qdaptor

WORKDIR /root/Qdaptor/src
RUN go build
CMD ["./Qdaptor"]