package cache

import (
	"encoding/json"
	"errors"
	"os"
	"time"

	jsonutils "github.com/equinor/radix-cli/pkg/utils/json"
	log "github.com/sirupsen/logrus"
)

type Cache struct {
	filename  string
	namespace string
}

type namespaceItems map[string]Item
type fileContent map[string]namespaceItems

type Item struct {
	UpdatedAt time.Time `json:"updated_at"`
	ExpiresAt time.Time `json:"expires_at"`
	Content   string    `json:"content"`
}

func New(filename, namespace string) Cache {
	return Cache{
		filename:  filename,
		namespace: namespace,
	}
}

func (c *Cache) load() (fileContent, error) {
	_, err := os.Stat(c.filename)
	if errors.Is(err, os.ErrNotExist) {
		return make(fileContent), nil
	}
	if err != nil {
		return nil, err
	}

	content := fileContent{}
	err = jsonutils.Load(c.filename, &content)
	if err != nil && err.Error() == "EOF" { // if the file is empty
		return make(fileContent), nil
	}
	if err != nil {
		return nil, err
	}

	if content == nil {
		return make(fileContent), nil
	}

	return content, nil
}

func (c *Cache) read() (namespaceItems, error) {
	content, err := c.load()
	if err != nil {
		return nil, err
	}

	items := content[c.namespace]
	if items == nil {
		items = make(namespaceItems)
	}

	return items, nil
}

func (c *Cache) GetItem(key string) (string, bool) {
	items, err := c.read()
	if err != nil {
		return "", false
	}

	item, ok := items[key]
	if !ok {
		return "", false
	}

	if time.Now().After(item.ExpiresAt) {
		return "", false
	}

	return item.Content, ok
}
func (c *Cache) GetItemUnmarsjalJson(key string, v any) bool {
	blob, ok := c.GetItem(key)
	if !ok {
		return false
	}

	err := json.Unmarshal([]byte(blob), v)
	if err != nil {
		return false
	}
	return true
}

func (c *Cache) ClearItem(key string) error {
	items, err := c.read()
	if err != nil {
		return err
	}

	delete(items, key)
	return c.save(items)
}
func (c *Cache) SetItem(key string, content string, ttl time.Duration) {
	items, err := c.read()
	if err != nil {
		log.Printf("Error loading file %s: %s", c.filename, err)
		return
	}

	items[key] = Item{
		UpdatedAt: time.Now(),
		ExpiresAt: time.Now().Add(ttl),
		Content:   content,
	}

	err = c.save(items)
	if err != nil {
		log.Printf("Error saving file %s: %s", c.filename, err)
	}
}

func (c *Cache) ClearAllItemsInNamespace() error {
	content := fileContent{}
	err := jsonutils.Load(c.filename, &content)
	if err != nil {
		return err
	}
	delete(content, c.namespace)

	return jsonutils.Save(c.filename, content)
}

func (c *Cache) save(items namespaceItems) error {
	content, err := c.load()
	if err != nil {
		return err
	}

	content[c.namespace] = items

	return jsonutils.Save(c.filename, content)
}
