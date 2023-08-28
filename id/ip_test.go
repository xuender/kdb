package id_test

import (
	"net"
	"os"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"github.com/xuender/kdb/id"
)

// nolint: paralleltest
func TestGetIP(t *testing.T) {
	ass := assert.New(t)
	ips := id.GetIP()

	ass.Equal(4, len(ips))
	ass.NotEqual(uint8(0), ips[0])

	patch := gomonkey.ApplyFuncReturn(net.Dial, nil, net.ErrClosed)
	defer patch.Reset()

	ips = id.GetIP()

	ass.Equal(uint8(0), ips[0])
}

// nolint: paralleltest
func TestGetMachine(t *testing.T) {
	ass := assert.New(t)

	patch := gomonkey.ApplyFuncReturn(id.GetIP, nil)
	defer patch.Reset()

	ass.Equal(uint8(0), id.GetMachine())

	patch2 := gomonkey.ApplyFuncReturn(os.Getenv, "3")
	defer patch2.Reset()

	ass.Equal(uint8(3), id.GetMachine())
}
