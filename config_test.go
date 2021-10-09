package tower

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewConfig(t *testing.T) {
	tests := []struct {
		name string
		want *Config
	}{
		{"test", &Config{
			Name:             "Tower",
			IP:               "0.0.0.0",
			IPVersion:        "tcp4",
			Port:             8999,
			MaxPacketSize:    0,
			MaxConn:          1024,
			WorkerPoolSize:   0,
			MaxWorkerTaskLen: 0,
			MaxMsgChanLen:    1024,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewConfig())
		})
	}
}

func TestConfig_check(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Config{}
			c.check()
			assert.NotEqual(t, "", c.IP)
			assert.Equal(t, "0.0.0.0", c.IP)
			assert.NotEqual(t, 0, c.Port)
			assert.Equal(t, 8999, c.Port)
			assert.NotEqual(t, 0, c.MaxConn)
			assert.Equal(t, 1024, c.MaxConn)
			assert.NotEqual(t, 0, c.MaxPacketSize)
			assert.Equal(t, uint32(4096), c.MaxPacketSize)
			assert.NotEqual(t, 0, c.MaxMsgChanLen)
			assert.Equal(t, uint32(1024), c.MaxMsgChanLen)
		})
	}
}
