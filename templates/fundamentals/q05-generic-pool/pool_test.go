package pool

import "testing"

func BenchmarkPool(b *testing.B) {
	p := New(func() int { return 0 })
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			v := p.Get()
			p.Put(v)
		}
	})
}
