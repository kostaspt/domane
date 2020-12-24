package whois

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhois_IsAvailable(t *testing.T) {
	t.Run("available", func(t *testing.T) {
		w := NewWhois(NewMockParser())
		available, err := w.IsAvailable("available.com")
		assert.NoError(t, err)
		assert.True(t, available)
	})

	t.Run("not available", func(t *testing.T) {
		w := NewWhois(NewMockParser())
		available, err := w.IsAvailable("not-available.com")
		assert.NoError(t, err)
		assert.False(t, available)
	})
}
