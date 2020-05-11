package autoupdate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestS3Release_Parse(t *testing.T) {
	a := assert.New(t)
	r := &s3Release{}

	err := r.Parse("releases/mfc_v2.0.0_darwin.tar.gz")
	a.NoError(err)
	a.Equal("releases/mfc_v2.0.0_darwin.tar.gz", r.Key)
	a.Equal("v2.0.0", r.Version)
	a.Equal("darwin", r.Platform)
}
