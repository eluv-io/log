package multi_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/eluv-io/apexlog-go"
	"github.com/eluv-io/apexlog-go/handlers/memory"
	"github.com/eluv-io/apexlog-go/handlers/multi"
)

func init() {
	log.Now = func() time.Time {
		return time.Unix(0, 0)
	}
}

func Test(t *testing.T) {
	a := memory.New()
	b := memory.New()

	log.SetHandler(multi.New(a, b))
	log.WithField("user", "tj").WithField("id", "123").Info("hello")
	log.Info("world")
	log.Error("boom")

	assert.Len(t, a.Entries, 3)
	assert.Len(t, b.Entries, 3)
}
