FROM alpine:latest as alpine
RUN apk add -U --no-cache ca-certificates

FROM scratch
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY ./azure-demo-k8s-go azure-demo-k8s-go
ENV PATH /usr/local/bin:/azure-demo-k8s-go
ENTRYPOINT ["/azure-demo-k8s-go"]