FROM alpine:3.10.3
RUN apk add --no-cache cifs-utils ca-certificates \
    && adduser -D -u 1000 chart-registry
COPY bin/linux/amd64/chart-registry /chart-registry
USER 1000
ENTRYPOINT ["/chart-registry"]
