package certstore

import (
	"bytes"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/common/flogging"
	"github.com/hyperledger/fabric/protos/msp"
	"github.com/hyperledger/fabric/protos/peer"
	"github.com/patrickmn/go-cache"
)

var logger = flogging.MustGetLogger("certstore")

// StoreCertsFromEnvelope extracts all certificates from envelope and stores them
func StoreCertsFromEnvelope(payload []byte) {

	fmt.Println("************ DEBUG ************")

	sID, tx, err := extractPayload(payload)
	if err != nil {
		logger.Warn(err)
		return
	}
	if sID.IdBytes != nil {
		storeCert(sID.IdBytes)
	}

	for _, a := range tx.Actions {

		fmt.Println("ACTION:", a)

		sID, err := identityFromSigHeader(a.Header)
		if err != nil {
			logger.Warn(err)
			continue
		}
		if sID.IdBytes != nil {
			storeCert(sID.IdBytes)
		}
		for _, x := range extractEndorsements(a.Payload) {

			fmt.Println("ENDORSER:", x.Endorser)

			sID, err := deserializeIdentity(x.Endorser)
			if err != nil {
				logger.Warn(err)
				continue
			}
			if sID.IdBytes != nil {
				storeCert(sID.IdBytes)
			}
		}
	}
}

// RemoveCertIfCached removes certificate from identity if it exists in store
func RemoveCertIfCached(identityBytes []byte) []byte {
	sID := &msp.SerializedIdentity{}
	if err := proto.Unmarshal(identityBytes, sID); err != nil {
		logger.Warn(err)
		return identityBytes
	}
	sID.IdRef = makeID(sID.IdBytes)

	if cert, err := getCert(sID.IdRef); err != nil {
		logger.Warn(err)
		return identityBytes
	} else if !bytes.Equal(cert, sID.IdBytes) {
		return identityBytes
	}

	sID.IdBytes = nil
	return marshalIdentity(sID, identityBytes)
}

// StoreFromTransientMap stores certificate from TransientMap to temporary cache
func StoreFromTransientMap(p []byte) {

	fmt.Println("************ DEBUG ************")

	ppp := &peer.ChaincodeProposalPayload{}
	if err := proto.Unmarshal(p, ppp); err != nil {
		return
	}

	//fmt.Println("ChaincodeProposalPayload:", ppp)
	fmt.Println("TransientMap CERT:", ppp.TransientMap["cert"])

	if ppp.TransientMap == nil {
		return
	}
	if cert, ok := ppp.TransientMap["cert"]; ok {
		certCache.Lock()
		defer certCache.Unlock()
		id := makeID(cert)
		certCache.transientCache.Add(string(id), cert, cache.DefaultExpiration)
	}
}

// GetCertIfNeeded sets certificate to identity unless it exists
func GetCertIfNeeded(s *msp.SerializedIdentity) *msp.SerializedIdentity {
	if s.IdBytes == nil {
		var err error
		if s.IdBytes, err = getCert(s.IdRef); err != nil {
			logger.Warn(err)
		}
	}
	return s
}
