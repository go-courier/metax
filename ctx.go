package metax

import "context"

type Ctx struct {
	ctx context.Context
}

func (c Ctx) Context() context.Context {
	if c.ctx == nil {
		return context.Background()
	}
	return c.ctx
}
