package plant

type UrlCache struct {
	caches map[string]bool
}

func NewUrlCache() *UrlCache {
	return &UrlCache{caches: make(map[string]bool)}
}

func (c *UrlCache) Set(url string) bool {
	_, ok := c.caches[url]
	if !ok {
		c.caches[url] = true
		return true
	}
	return false
}

func (c *UrlCache) Len() int {
	return len(c.caches)
}
