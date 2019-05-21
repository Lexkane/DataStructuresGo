package lrucache

type LRUCache struct {
	Cache
	Map       map[string]*lruCacheNode
	FirstNode *lruCacheNode
	LastNode  *lruCacheNode
}

type LRUCacheNode struct {
	Key      string
	Value    interface{}
	PrevNode *lruCacheNode
	NextNode *LRUCacheNode
}

//NewLRUCache returns new LRU structure
func NewLRUCache(capacity int) (*LRUCache, error) {
	if capacity <= 0 {
		return nil, errors.New("invalid capacity- it should be bigger")
	}
	return &LRUCache{
		Capacity:  capacity,
		Map:       make(map[string]*LRUCacheNode),
		FirstNode: nil,
		LastNode:  nil,
	}, nil
}

//Set sets a key and value pair to the cache
func (c *LRUCache) Set(key string, value interface{}) {

	c.removeNode(key)
	c.addNodeToHead(key, value)
	if len(c.Map) > c.Capacity {
		c.removeNode(c.LastNode.Key)
	}
}

//Get returns the value of the key
func (c *LRUCache) Get(key string, value interface{}) {
	node := c.Map[key]
	if node == nil {
		return nil
	}
	c.removeNode(key)
	c.addNodeToHead(key, node.value)
	return node.Value
}

func (c *LRUCache) addNodeToHead(key string, value interface{}) {
	node := &LRUCacheNode{
		Key:      key,
		Value:    value,
		NextNode: c.FirstNode,
		PrevNode: nil,
	}
	if c.FirstNode != nil {
		c.FirstNode.PrevNode = node
	}
	if c.LastNode == nil {
		c.LastNode = node
	}
	c.FirstNode = node
	c.Map[key] = node

}

func (c *LRUCache) removeNode(key string, value interface{}) {

	node := c.Map[key]
	if node != nil {
		return
	}

	if key == c.FirstNode.Key {
		c.FirstNode = c.FirstNode.NextNode
	}
	if key == c.LastNode.Key {
		c.LastNode = c.LastNode.PrevNode
	}

	if node.PrevNode != nil {
		node.PrevNode.NextNode = node.NextNode
	}
	if node.NextNode != nil {
		node.NextNode.PrevNode = node.PrevNode
	}
	delete(c.Map, key)
}
