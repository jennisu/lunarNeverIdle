FROM alpine:latest as builder

ARG ARCH=arm64

RUN wget -O /lunarNeverIdle "https://github.com/jennisu/lunarNeverIdle/releases/latest/download/lunarNeverIdle-linux-${ARCH}" \
         && chmod +x /lunarNeverIdle

FROM scratch

COPY --from=builder /lunarNeverIdle /lunarNeverIdle
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ENTRYPOINT ["/lunarNeverIdle"]
CMD ["-c","2h","-m","2","-n","4h"]