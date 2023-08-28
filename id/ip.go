package id

import (
	"net"
	"os"

	"github.com/xuender/kit/types"
)

func GetIP() []byte {
	if conn, err := net.Dial("udp", "8.8.8.8:53"); err == nil {
		defer conn.Close()

		if localAddr, ok := conn.LocalAddr().(*net.UDPAddr); ok {
			return localAddr.IP
		}
	}

	return []byte{0, 0, 0, 0}
}

func GetMachine() uint8 {
	if machinKey := os.Getenv(MachineKey); machinKey != "" {
		if key, err := types.ParseInteger[uint8](machinKey); err == nil {
			return key
		}
	}

	bytes := GetIP()
	if len(bytes) == 0 {
		return 0
	}

	return bytes[len(bytes)-1]
}
