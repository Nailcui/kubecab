FROM alpine

WORKDIR /build
COPY kubecab .

CMD ["/build/kubecab"]
