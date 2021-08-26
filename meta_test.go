package metax

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
)

func TestParseMeta(t *testing.T) {
	t.Run("parse id", func(t *testing.T) {
		meta := ParseMeta("xxxxxx")
		NewWithT(t).Expect(meta.String()).To(Equal("_id=xxxxxx"))
	})

	t.Run("parse meta", func(t *testing.T) {
		meta := ParseMeta("operator=1&operator=2&_id=xxx")
		NewWithT(t).Expect(meta.Get("operator")).To(Equal("1"))
		NewWithT(t).Expect(meta.String()).To(Equal("_id=xxx&operator=1&operator=2"))
	})
}

func TestMeta(t *testing.T) {

	t.Run("ContextConcat", func(t *testing.T) {
		ctx := ContextWith(context.Background(), "key", "1")
		ctx = ContextWithMeta(ctx, (Meta{}).With("key", "2", "3"))

		NewWithT(t).Expect(MetaFromContext(ctx)["key"]).To(Equal([]string{"1", "2", "3"}))
	})

	t.Run("ContextOverwrite", func(t *testing.T) {
		ctx := ContextWith(context.Background(), "_key", "1")
		ctx = ContextWithMeta(ctx, (Meta{}).With("_key", "2", "3"))

		NewWithT(t).Expect(MetaFromContext(ctx)["_key"]).To(Equal([]string{"2", "3"}))
	})

	t.Run("EmptyKeyIgnore", func(t *testing.T) {
		ctx := ContextWith(context.Background(), "", "1")
		NewWithT(t).Expect(MetaFromContext(ctx)).To(HaveLen(0))
	})
}
