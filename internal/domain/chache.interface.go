package domain

type Cache interface {
	
	GetCache(key string) (string, error)
	
	SetCache(key string, value string) error
	
	DeleteCache(key string) error
}
