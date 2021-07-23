package certstore

import (
	"github.com/OCRVblockchain/fabric/protos/common"
	"github.com/OCRVblockchain/fabric/protos/msp"
	"github.com/OCRVblockchain/fabric/protos/peer"
	"github.com/golang/protobuf/proto"
)

func deserializeIdentity(identity []byte) (*msp.SerializedIdentity, error) {
	sID := &msp.SerializedIdentity{}
	err := proto.Unmarshal(identity, sID)
	return sID, err
}

func identityFromSigHeader(header []byte) (*msp.SerializedIdentity, error) {
	sh := &common.SignatureHeader{}
	if err := proto.Unmarshal(header, sh); err != nil {
		return nil, err
	}
	return deserializeIdentity(sh.Creator)
}

func extractEndorsements(p []byte) []*peer.Endorsement {
	pp := &peer.ChaincodeActionPayload{}
	if err := proto.Unmarshal(p, pp); err != nil {
		logger.Warn(err)
		return nil
	}
	return pp.Action.Endorsements
}

func extractPayload(payload []byte) (*msp.SerializedIdentity, *peer.Transaction, error) {
	pld := &common.Payload{}
	if err := proto.Unmarshal(payload, pld); err != nil {
		return nil, nil, err
	}
	tx := &peer.Transaction{}
	if err := proto.Unmarshal(pld.Data, tx); err != nil {
		return nil, nil, err
	}
	sID, err := identityFromSigHeader(pld.Header.SignatureHeader)
	return sID, tx, err
}

func marshalIdentity(sID *msp.SerializedIdentity, oldIdentity []byte) []byte {
	newIdentity, err := proto.Marshal(sID)
	if err != nil {
		logger.Warn(err)
		return oldIdentity
	}
	return newIdentity
}
