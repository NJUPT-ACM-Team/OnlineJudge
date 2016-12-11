protoc --proto_path=$GOPATH/src/OnlineJudge/protos --go_out=$GOPATH/src/OnlineJudge/pbgen/ $GOPATH/src/OnlineJudge/protos/api/*.proto
