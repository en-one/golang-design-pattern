package singleton

import "sync"

type Singleton struct{}

//饿汉式单例
// 定义s 返回值
//var s *Singleton
//func init() {
//	s = &Singleton{}
//}
//
//func GetInstance() *Singleton {
//	return s
//}

// 懒汉
var (
	lazyS *Singleton
	once  = sync.Once{}
)

func GetLazyInstance() *Singleton {
	if lazyS == nil {
		once.Do(func() {
			lazyS = &Singleton{}
		})
	}
	return lazyS
}
