package kdb_test

import (
	"net"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"github.com/xuender/kdb"
)

// nolint: paralleltest
func TestGetIP(t *testing.T) {
	ass := assert.New(t)
	ips := kdb.GetIP()

	ass.Equal(4, len(ips))
	ass.NotEqual(uint8(0), ips[0])

	patch := gomonkey.ApplyFuncReturn(net.Dial, nil, net.ErrClosed)
	defer patch.Reset()

	ips = kdb.GetIP()

	ass.Equal(uint8(0), ips[0])
}
