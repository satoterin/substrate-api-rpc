package substrate

import (
	"encoding/json"
	"fmt"
	"github.com/itering/scale.go/types"
	"github.com/itering/substrate-api-rpc/storage"
	"github.com/itering/substrate-api-rpc/util"
)

const (
	CidAura = 0x61727561
	CidBabe = 0x45424142
)

type PreRuntime struct {
	Data   string `json:"data"`
	Engine int64  `json:"engine"`
}

func (p *PreRuntime) isAura() bool {
	return CidAura == p.Engine
}
func (p *PreRuntime) isBabe() bool {
	return CidBabe == p.Engine
}

func (p *PreRuntime) getAuraAuthor(sessionValidators []string) string {
	if rawAuraPreDigestValue, err := storage.Decode(p.Data, "RawAuraPreDigest", nil); err == nil {
		modn := rawAuraPreDigestValue.ToRawAuraPreDigest().SlotNumber % int64(len(sessionValidators))
		return sessionValidators[modn]
	}
	return ""
}

func (p *PreRuntime) getBabeAuthor(sessionValidators []string) string {
	if rawBabePreDigestValue, err := storage.Decode(p.Data, "RawBabePreDigest", nil); err == nil {
		digest := rawBabePreDigestValue.ToRawBabePreDigest()
		if digest != nil {
			if digest.Primary != nil {
				return sessionValidators[digest.Primary.AuthorityIndex]
			} else if digest.Secondary != nil {
				return sessionValidators[digest.Secondary.AuthorityIndex]
			} else {
				return sessionValidators[digest.VRF.AuthorityIndex]
			}
		}
	}
	return ""
}

func ExtractAuthor(data []byte, sessionValidators []string) string {
	var p PreRuntime
	if len(sessionValidators) == 0 {
		return ""
	}
	if err := json.Unmarshal(data, &p); err != nil {
		return ""
	}
	if p.isAura() {
		return p.getAuraAuthor(sessionValidators)
	} else if p.isBabe() {
		return p.getBabeAuthor(sessionValidators)
	} else {
		return ""
	}
}

// LogDigest decode
func DecodeLogDigest(rawList []string) (r []storage.DecoderLog, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Recovering from panic in DecodeLogDigest error is: %v \n", r)
		}
	}()
	for _, logRaw := range rawList {
		m := types.ScaleDecoder{}
		m.Init(types.ScaleBytes{Data: util.HexToBytes(logRaw)}, nil)
		rb := m.ProcessAndUpdateData("LogDigest").(map[string]interface{})

		var log storage.DecoderLog
		util.UnmarshalToAnything(&log, rb)
		r = append(r, log)
	}
	return
}
