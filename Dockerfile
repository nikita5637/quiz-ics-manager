# Builder
ARG BASE_IMAGE=golang:1.18.10-alpine3.17
FROM ${BASE_IMAGE} AS builder

COPY ./ /go/src/quiz-ics-manager-api

WORKDIR /go/src/quiz-ics-manager-api
RUN go build -buildvcs=auto -o ics-manager-api ./cmd/quiz-ics-manager-api

# ics-manager-api

FROM alpine:3.17 as server

COPY --from=builder /go/src/quiz-ics-manager-api/ics-manager-api /bin/
COPY --from=builder /go/src/quiz-ics-manager-api/config.toml /etc/

EXPOSE 8082

CMD ["/bin/ics-manager-api"]
