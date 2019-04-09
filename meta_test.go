package metax

import (
	"testing"
)

func TestMeta(t *testing.T) {
	t.Run("parse id", func(t *testing.T) {
		meta := ParseMeta("xxxxxx")
		t.Log(meta)
	})

	t.Run("parse meta", func(t *testing.T) {
		meta := ParseMeta("operator=xxxxxx&_id=xxx")
		t.Log(meta)
	})
}
