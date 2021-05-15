package singleton

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetInstance(t *testing.T) {
	//hungry
	//ins1 := GetInstance()
	//ins2 := GetInstance()
	//lazy
	ins1 := GetLazyInstance()
	ins2 := GetLazyInstance()

	assert.Equal(t, ins1, ins2)
}

//压力测试 使用Benchmark开头
//执行压力测试需要带上参数-test.bench
//使用test.B
func BenchmarkGetInstanceParallel(b *testing.B) {
	//hungry
	//ins1 := GetInstance()
	//ins2 := GetInstance()
	//lazy
	ins1 := GetLazyInstance()
	ins2 := GetLazyInstance()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if ins1 != ins2 {
				b.Errorf("test fail")
			}
		}
	})
}
