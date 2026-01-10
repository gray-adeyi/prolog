FROM golang:1.25-alpine AS build

ARG GRPC_HEALTH_PROBE_VERSION=v0.4.43

WORKDIR /prolog
COPY . .
RUN CGO_ENABLED=0 go build -o /bin/prolog ./cmd/prolog
RUN wget -qO/bin/grpc_health_probe \
    "https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/\
${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64" && \
    chmod +x /bin/grpc_health_probe

FROM scratch
COPY --from=build /bin/prolog /bin/prolog
COPY --from=build /bin/grpc_health_probe /bin/grpc_health_probe
ENTRYPOINT [ "/bin/prolog" ]