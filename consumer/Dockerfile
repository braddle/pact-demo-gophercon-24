FROM --platform=linux/amd64 golang:1.22.4 as build

WORKDIR /code

# Install Pact Go Tools
RUN go install github.com/pact-foundation/pact-go/v2@latest
RUN pact-go -l DEBUG install

# Install UUID
RUN apt update -y
RUN apt install -y uuid

# Installing Pact Broker Cli Tool
RUN curl -LO https://github.com/pact-foundation/pact-ruby-standalone/releases/download/v2.4.6/pact-2.4.6-linux-x86_64.tar.gz
RUN tar xzf pact-2.4.6-linux-x86_64.tar.gz -C /
ENV PATH $PATH:/pact/bin

#Keep the docker container running
CMD tail -f /dev/null