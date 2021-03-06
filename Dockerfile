# https://github.com/yeasy/docker-hyperledger-fabric-peer
FROM yeasy/hyperledger-fabric-base:release-v1.4
WORKDIR $GOPATH/src/github.com/OCRVblockchain/fabric-certstore
COPY . .

RUN cd $FABRIC_ROOT \
    && git apply $GOPATH/src/github.com/OCRVblockchain/fabric-certstore/fabric.patch

COPY tools/protoc-gen-go /usr/local/bin/
ADD tools/gotools.tar.bz2 /usr/local/bin/

COPY --from=hyperledger/fabric-baseimage:amd64-0.4.22 /usr/bin/protoc /usr/local/bin/
COPY --from=hyperledger/fabric-baseimage:amd64-0.4.22 /usr/lib/ /usr/local/lib/

RUN export LD_LIBRARY_PATH=/usr/local/lib \
    && rm /go/bin/protoc-gen-go \
    && protoc --proto_path="/go/src/github.com/hyperledger/fabric/protos" --go_out=plugins=grpc:/go/src /go/src/github.com/hyperledger/fabric/protos/msp/identities.proto

RUN go get -u "github.com/patrickmn/go-cache"
RUN go get -u "github.com/syndtr/goleveldb/leveldb"

RUN cd $FABRIC_ROOT/peer \
    && go install -tags "experimental" -ldflags "$LD_FLAGS" \
    && go clean

EXPOSE 7051
CMD tail -f /dev/null