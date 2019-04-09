package metax

import (
	"context"
	"net/url"
	"strings"
)

var (
	contextKey = &struct{}{}
)

func ContextWithMeta(ctx context.Context, meta Meta) context.Context {
	return context.WithValue(ctx, contextKey, meta)
}

func MetaFromContext(ctx context.Context) Meta {
	if m, ok := ctx.Value(contextKey).(Meta); ok {
		return m
	}
	return Meta{}
}

func ParseMeta(query string) Meta {
	if strings.Index(query, "=") == -1 {
		return Meta{
			"_id": []string{query},
		}
	}
	values, err := url.ParseQuery(query)
	if err == nil {
		return Meta(values)
	}
	return Meta{}
}

type Meta map[string][]string

func (m Meta) Merge(metas ...Meta) Meta {
	meta := Meta{}

	for _, me := range append([]Meta{m}, metas...) {
		for k, v := range me {
			meta[k] = v
		}
	}

	return meta
}

func (m Meta) Clone() Meta {
	meta := Meta{}
	for k, v := range m {
		meta[k] = v
	}
	return meta
}

func (m Meta) Add(key string, value string) {
	m[key] = append(m[key], value)
}

func (m Meta) Get(key string) string {
	if m == nil {
		return ""
	}
	vs := m[key]
	if len(vs) == 0 {
		return ""
	}
	return vs[0]
}

func (m Meta) With(key string, values ...string) Meta {
	meta := m.Clone()
	meta[key] = values
	return meta
}

func (m Meta) String() string {
	return url.Values(m).Encode()
}
