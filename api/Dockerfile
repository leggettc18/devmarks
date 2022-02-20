FROM golang:1.16.5-alpine as base
WORKDIR /api
FROM aquasec/trivy:0.18.3 as trivy
RUN trivy --debug --timeout 4m golang:1.16.5-alpine && \
    echo "No image vulnerabilities" > result
FROM base as dev
COPY go.* ./
RUN go mod download
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN go env
RUN apk add build-base
RUN go get github.com/go-delve/delve/cmd/dlv && go get github.com/githubnemo/CompileDaemon
EXPOSE 4000 2345
FROM dev as test
COPY . .
RUN export CGO_ENABLED=0 && \
    go test -v ./...
FROM test as build-stage
RUN GOOS=linux go build -ldflags "-s -w" -o devmarks ./main.go
FROM base as prod
COPY --from=trivy result secure
COPY --from=build-stage /api/devmarks devmarks
RUN addgroup devmarks && adduser -D -G devmarks devmarks \
    && chown -R devmarks:devmarks /api && \
    chmod +x ./devmarks
USER devmarks
EXPOSE 4000
HEALTHCHECK CMD ["wget", "-q", "0.0.0.0:8080"]
CMD ["./devmarks", "serve"]