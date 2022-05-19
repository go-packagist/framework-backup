package support

import (
	"github.com/golang-module/carbon/v2"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCarbon(t *testing.T) {
	assert.Equal(t, time.Now().Format("2006-01-02 15:04:05"), carbon.Now().Format("Y-m-d H:i:s"))
}
