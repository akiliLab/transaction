############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git && apk add -qU openssh
COPY . /home/transaction
# for private repo pulling
ARG ssh_prv_key
ARG ssh_pub_key
# Authorize SSH Host
RUN mkdir -p /root/.ssh && \
    chmod 0700 /root/.ssh && \
    ssh-keyscan github.com > /root/.ssh/known_hosts

# Add the keys and set permissions
RUN echo "$ssh_prv_key" > /root/.ssh/id_rsa && \
    echo "$ssh_pub_key" > /root/.ssh/id_rsa.pub && \
    chmod 600 /root/.ssh/id_rsa && \
    chmod 600 /root/.ssh/id_rsa.pub

WORKDIR /home/transaction

# RUN git config --global http.sslBackend "openssl"
# RUN git config --global url.ssh://git@github.com/.insteadOf https://github.com/

# RUN go mod tidy

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/transaction

# Remove SSH keys
RUN rm -rf /root/.ssh/
############################
# STEP 2 build a small image
############################

FROM alpine:3.4

RUN apk --no-cache --update upgrade && apk add --no-cache ca-certificates && update-ca-certificates

EXPOSE 5001

WORKDIR /root
# Copy our static executable.
COPY --from=builder /go/bin/transaction .


# Run the payment binary.
ENTRYPOINT ["./transaction"]