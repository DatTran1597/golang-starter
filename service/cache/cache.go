package cache

type Cacher interface {
	Get(key string) (interface{}, error)
	Set(key string, val interface{}) error
}
