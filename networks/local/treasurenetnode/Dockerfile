FROM golang:stretch as build-env

# Install minimum necessary dependencies
ENV PACKAGES curl make git libc-dev bash gcc
ENV GOPROXY https://goproxy.cn,direct
RUN apt-get update && apt-get upgrade -y && \
    apt-get install -y $PACKAGES

# Set working directory for the build
# WORKDIR /go/src/github.com/tharsis/ethermint
WORKDIR /go/src/github.com/treasurenet

# Add source files
COPY . .

# build Ethermint
# build Treasurenet
RUN make build-linux

# Final image
FROM golang:1.15.14 as final

WORKDIR /

RUN apt-get update

# Copy over binaries from the build-env
# COPY --from=build-env /go/src/github.com/tharsis/ethermint/build/ethermintd /
# COPY --from=build-env /go/src/github.com/tharsis/ethermint/scripts/start-docker.sh /

COPY --from=build-env /go/src/github.com/treasurenet/build/treasurenetd /
COPY --from=build-env /go/src/github.com/treasurenet/scripts/start-docker.sh /

EXPOSE 26656 26657 1317 8545 8546

# Run treasurenetd by default, omit entrypoint to ease using container with treasurenetd
ENTRYPOINT ["/bin/bash", "-c"]