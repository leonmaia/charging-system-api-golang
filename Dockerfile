FROM google/golang:stable

RUN go get github.com/tools/godep

RUN CGO_ENABLED=0 go install -a std

MAINTAINER Leon Maia <hi@leonmaia.com>
ENV APP_DIR $GOPATH/src/github.com/leonmaia/newmotion-golang

# Set the entrypoint
ENTRYPOINT ["/opt/app/newmotion-golang"]
ADD . $APP_DIR

# Compile the binary and statically link
RUN mkdir /opt/app
RUN cd $APP_DIR && godep restore
RUN cd $APP_DIR && CGO_ENABLED=0 go build -o /opt/app/newmotion-golang -ldflags '-d -w -s'

EXPOSE 6680
