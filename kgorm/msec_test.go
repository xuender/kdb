package kgorm_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/xuender/kdb/kgorm"
)

type Mod struct {
	Msec kgorm.Msec `json:"msec"`
}

func TestMsec_MarshalJSON(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	mod := &Mod{Msec: kgorm.Msec(lo.Must1(time.Parse("2006-01-02", "2023-08-28")))}
	data := lo.Must1(json.Marshal(mod))

	ass.Equal(`{"msec":1693180800000}`, string(data))
	data = lo.Must1(json.Marshal(&Mod{}))
	ass.Equal(`{"msec":null}`, string(data))
}

func TestMsec_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	mod := &Mod{}
	_ = json.Unmarshal([]byte(`{"msec":1693180800000}`), mod)

	ass.Equal(2023, time.Time(mod.Msec).Year())
}
