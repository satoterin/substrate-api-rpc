package storage

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/itering/scale.go/types"
	"github.com/itering/substrate-api-rpc/util"
	"github.com/shopspring/decimal"
)

func Decode(raw string, decodeType string, metadata *types.MetadataStruct) (s StateStorage, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Recovering from panic in Decode error is: %v \n", r)
		}
	}()
	m := types.ScaleDecoder{}

	option := types.ScaleDecoderOption{}
	if metadata != nil {
		metadataStruct := types.MetadataStruct(*metadata)
		option.Metadata = &metadataStruct
	}
	m.Init(types.ScaleBytes{Data: util.HexToBytes(raw)}, &option)
	return StateStorage(util.InterfaceToString(m.ProcessAndUpdateData(decodeType))), nil
}

type StateStorage string

func (s *StateStorage) bytes() []byte {
	return []byte(string(*s))
}

func (s *StateStorage) string() string {
	return string(*s)
}

func (s *StateStorage) ToStringSlice() (r []string) {
	_ = json.Unmarshal(s.bytes(), &r)
	return
}

func (s *StateStorage) ToString() (r string) {
	if err := json.Unmarshal(s.bytes(), &r); err != nil {
		return s.string()
	}
	return
}

func (s *StateStorage) ToInt() (r int) {
	if r, err := strconv.Atoi(s.string()); err == nil {
		return r
	}
	return 0
}

func (s *StateStorage) ToInt64() (r int64) {
	i, _ := strconv.ParseInt(s.string(), 10, 64)
	return i
}

func (s *StateStorage) ToStakingLedgers() (r *StakingLedgers) {
	_ = json.Unmarshal(s.bytes(), &r)
	return
}

func (s *StateStorage) ToMapString() (r map[string]string) {
	_ = json.Unmarshal(s.bytes(), &r)
	return
}
func (s *StateStorage) ToMapInterface() (r map[string]interface{}) {
	_ = json.Unmarshal(s.bytes(), &r)
	return
}

func (s *StateStorage) ToExposures() (r *Exposures) {
	_ = json.Unmarshal(s.bytes(), &r)
	return
}

func (s *StateStorage) ToRawAuraPreDigest() (r *RawAuraPreDigest) {
	_ = json.Unmarshal(s.bytes(), &r)
	return
}

func (s *StateStorage) ToRawBabePreDigest() (r *RawBabePreDigest) {
	_ = json.Unmarshal(s.bytes(), &r)
	return
}

func (s *StateStorage) ToValidatorPrefsLegacy() (r *ValidatorPrefsLegacy) {
	_ = json.Unmarshal(s.bytes(), &r)
	return
}

func (s *StateStorage) ToU32FromCodec() (r uint32) {
	if s.string() == "" {
		return 0
	}
	return uint32(binary.LittleEndian.Uint32(util.HexToBytes(s.string())[0:4]))
}

// ToDecimal
// Python GRPC return balance type is String, when grpc return json, balance string will return "balance"
func (s *StateStorage) ToDecimal() (r decimal.Decimal) {
	if s.string() == "" {
		return decimal.Zero
	}
	return decimal.RequireFromString(strings.ReplaceAll(s.string(), "\"", ""))
}

func (s *StateStorage) ToBalanceLock() (r []BalanceLock) {
	_ = json.Unmarshal(s.bytes(), &r)
	return
}

func (s *StateStorage) ToEraPoints() (r *EraPoints) {
	_ = json.Unmarshal(s.bytes(), &r)
	return
}

func (s *StateStorage) ToRegistration() (r *Registration) {
	_ = json.Unmarshal(s.bytes(), &r)
	return
}

func (s *StateStorage) ToAccountData() (r *AccountData) {
	_ = json.Unmarshal(s.bytes(), &r)
	return
}

func (s *StateStorage) ToAccountInfo() (r *AccountInfo) {
	_ = json.Unmarshal(s.bytes(), &r)
	return
}

func (s *StateStorage) ToActiveEraInfo() (r *ActiveEraInfo) {
	_ = json.Unmarshal(s.bytes(), &r)
	return
}

func (s *StateStorage) ToProposal() (r *Proposal) {
	_ = json.Unmarshal(s.bytes(), &r)
	return
}
func (s *StateStorage) ToReferendumInfo() (r *ReferendumInfo) {
	_ = json.Unmarshal(s.bytes(), &r)
	return
}
