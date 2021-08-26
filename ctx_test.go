package metax

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
)

type SomeString struct {
	Ctx
}

func (s *SomeString) WithContext(ctx context.Context) *SomeString {
	return &SomeString{
		Ctx: s.Ctx.WithContext(ctx),
	}
}

func TestCtx(t *testing.T) {
	s := &SomeString{}
	s2 := s.WithContext(ContextWith(context.Background(), "k", "1"))

	NewWithT(t).Expect(MetaFromContext(s.Context()).Get("k")).To(Equal(""))
	NewWithT(t).Expect(MetaFromContext(s2.Context()).Get("k")).To(Equal("1"))
}
