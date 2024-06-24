FROM --platform=linux/amd64 golang:1.22.4 as build

WORKDIR /code
RUN go install github.com/pact-foundation/pact-go/v2@latest
RUN pact-go -l DEBUG install
RUN ls -l
CMD tail -f /dev/null