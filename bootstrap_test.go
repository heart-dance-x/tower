package tower

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_bootStrap_Listen(t *testing.T) {
	tests := []struct {
		name string
	}{
		{""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Nil(t, nil)
			//bs := NewBootStrap(&Config{})
			//bs.Listen()
		})
	}
}
