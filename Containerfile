FROM golang:1.25-alpine AS build
WORKDIR /prolog
COPY . .
RUN CGO_ENABLED=0 go build -o /bin/prolog ./cmd/prolog

FROM scratch
COPY --from=build /bin/prolog /bin/prolog
ENTRYPOINT [ "/bin/prolog" ]