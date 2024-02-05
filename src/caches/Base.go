package caches

import "sync"

// 非共享内存
// 本部分缓存只存在于内存当中, 在各个应用程序之间不能共存

// AppCache 结构
type AppCache struct {
	Items  map[string]interface{}
	Locker sync.Mutex
}

// App 所有的缓存保存
var App = &AppCache{
	Items:  map[string]interface{}{},
	Locker: sync.Mutex{},
}

// Get 获取单个缓存记录
func (ths *AppCache) Get(platform string, key string, callback func() interface{}) interface{} {
	cacheKey := platform + "-" + key
	if val, exists := App.Items[cacheKey]; exists {
		return val
	}

	App.Locker.Lock()
	defer App.Locker.Unlock()
	val := callback()
	App.Items[cacheKey] = val
	return val
}

// Get 获取缓存
func Get(platform string, key string, callback func() interface{}) interface{} {
	cacheKey := platform + "-" + key
	if val, exists := App.Items[cacheKey]; exists {
		return val
	}

	App.Locker.Lock()
	defer App.Locker.Unlock()
	val := callback()
	App.Items[cacheKey] = val
	return val
}
