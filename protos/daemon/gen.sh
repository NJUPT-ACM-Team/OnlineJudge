protoc \
--proto_path=$GOPATH/src/OnlineJudge/protos/daemon \
--go_out=plugins=grpc:$GOPATH/src/OnlineJudge/Daemon/pb \
$GOPATH/src/OnlineJudge/protos/daemon/backend.proto
