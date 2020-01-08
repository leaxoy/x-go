package types

import (
	"encoding/json"
	"strings"

	"github.com/golang/glog"
	"github.com/speps/go-hashids"
	"golang.org/x/xerrors"
)

var seed *hashids.HashID

func init() {
	Init("iMaster")
}

func Init(salt string) {
	idData := hashids.NewData()
	idData.Salt = salt
	var err error
	seed, err = hashids.NewWithData(idData)
	if err != nil {
		glog.Fatalf("initialize hash id failed, %s", err)
	}
}

type ID int64

func (i ID) MarshalText() (text []byte, err error) {
	if i < 0 {
		return nil, xerrors.Errorf("err: ID(%d) must greater than 0", i)
	}
	str, err := seed.EncodeInt64([]int64{int64(i)})
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}

func (i *ID) UnmarshalText(text []byte) error {
	is, err := seed.DecodeInt64WithError(string(text))
	if err != nil {
		return xerrors.Errorf("unmarshal id failed, %w", err)
	}
	if len(is) != 1 {
		return xerrors.Errorf("bad unmarshal id length, %d", len(is))
	}
	*i = ID(is[0])
	return nil
}

func (i ID) MarshalJSON() ([]byte, error) {
	text, err := i.MarshalText()
	if err != nil {
		return nil, err
	}
	return json.Marshal(string(text))
}

func (i *ID) UnmarshalJSON(data []byte) error {
	var text string
	err := json.Unmarshal(data, &text)
	if err != nil {
		return err
	}
	return i.UnmarshalText([]byte(text))
}

func ParseIDList(str string) ([]ID, error) {
	items := strings.Split(str, ",")
	ids := make([]ID, 0, len(items))
	for _, item := range items {
		id := new(ID)
		err := id.UnmarshalText([]byte(item))
		if err != nil {
			continue
		}
		ids = append(ids, *id)
	}
	return ids, nil
}

func ParseInt64List(str string) ([]int64, error) {
	ids, err := ParseIDList(str)
	if err != nil {
		return nil, err
	}
	int64s := make([]int64, 0, len(ids))
	for _, id := range ids {
		int64s = append(int64s, int64(id))
	}
	return int64s, nil
}

func FormatIDList(ids []ID) (string, error) {
	items := make([]string, 0, len(ids))
	for _, id := range ids {
		item, err := id.MarshalJSON()
		if err != nil {
			continue
		}
		items = append(items, string(item))
	}
	return strings.Join(items, ","), nil
}
