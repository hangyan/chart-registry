FROM alpine:3.11.0
RUN apk add --no-cache cifs-utils ca-certificates \
    && adduser -D -u 1000 chart-registry 
COPY bin/linux/amd64/chart-registry /chart-registry
COPY scripts/certs/domain.crt /usr/local/share/ca-certificates/
RUN update-ca-certificates
USER 1000
ENTRYPOINT ["/chart-registry"]
