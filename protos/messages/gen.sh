protoc \
--proto_path=$GOPATH/src/OnlineJudge/protos \
--go_out=plugins=grpc:$GOPATH/src/OnlineJudge/pbgen \
$GOPATH/src/OnlineJudge/protos/messages/*.proto
