# Use the official go docker image built on debian.
FROM golang

# Grab the source code and add it to the workspace.
ADD . /go/src/github.com/karesti/cm-voting

# Install revel and the revel CLI.
RUN go get github.com/revel/revel
RUN go get github.com/revel/cmd/revel

ENTRYPOINT revel run github.com/karesti/cm-voting dev 9000

# Open up the port where the app is running.
EXPOSE 9000

