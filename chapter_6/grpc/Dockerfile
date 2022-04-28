FROM golang

RUN mkdir /opt/grpc_test/
WORKDIR /opt/grpc_test/

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
# ENV PATH "$PATH:$(go env GOPATH)/bin"

RUN git clone -b v1.46.0 --depth 1 https://github.com/grpc/grpc-go
WORKDIR /opt/grpc_test/grpc-go/examples/helloworld
# RUN python -m grpc_tools.protoc -I../../protos --python_out=. --grpc_python_out=. ../../protos/route_guide.proto

CMD bash