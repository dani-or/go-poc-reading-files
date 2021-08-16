#docker image build -t new .
#docker run --network host -d new
FROM golang:alpine AS build
WORKDIR /go/src/myapp
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/myapp cmd/main.go
RUN apk add -U --no-cache ca-certificates

FROM scratch
ADD resources/s3.crt /etc/ssl/certs/
ENV NEQUI_BUCKET_NAME="nequi-s3-select-tmp-2"
ENV NEQUI_FILE_KEY="resource/FINACLE_NEQUICARTERA_20200508_VENCIDOS.csv"

COPY --from=build /go/bin/myapp /go/bin/myapp
EXPOSE 8080
ENTRYPOINT ["/go/bin/myapp"]