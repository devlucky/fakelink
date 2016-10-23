package links

import (
	"github.com/satori/go.uuid"
	"gopkg.in/redis.v5"
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
)

// A Store allows saving and retrieving user-generated links
type Store interface {
	Find(slug string) *Link
	FindRandom() (slug string)
	Create(link *Link) string
	clear()
}

/*
	In-memory
*/

// In-memory implementation of a template store
type InMemoryStore struct {
	links map[string]*Link
}

// Create a new in-memory store
func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{make(map[string]*Link)}
}

func (store *InMemoryStore) Find(slug string) *Link {
	return store.links[slug]
}

func (store *InMemoryStore) FindRandom() (slug string) {
	if len(store.links) == 0 {
		return
	}

	i := 0
	n := rand.Int() % len(store.links)

	for s := range store.links {
		if i == n {
			slug = s
			break
		}
		i++
	}

	return
}

func (store *InMemoryStore) Create(link *Link) string {
	slug := uuid.NewV4().String()
	store.links[slug] = link

	return slug
}

func (store *InMemoryStore) clear() {
	store.links = make(map[string]*Link)
}


/*
	Redis
*/

// Redis implementation of a template store
type RedisStore struct {
	client *redis.Client
}

// Create a new in-memory store
func NewRedisStore(host, port, password string, db int) *RedisStore {
	return &RedisStore{
		client: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", host, port),
			Password: password,
			DB:       db,
		}),
	}
}

func (store *RedisStore) Find(slug string) *Link {
	str, err := store.client.Get(slug).Result()
	if err != nil {
		log.Printf("Getting link with slug %s failed with error %s", slug, err)
		return nil
	}

	link := &Link{}
	if err = json.Unmarshal([]byte(str), link); err != nil {
		log.Printf("Unexpected error when unmarshaling a previously stored link: %s", err)
		return nil
	}

	return link
}

func (store *RedisStore) FindRandom() (slug string) {
	slug, err := store.client.RandomKey().Result()
	if err != nil {
		log.Printf("Getting link with slug %s failed with error %s", slug, err)
		return
	}

	return
}

func (store *RedisStore) Create(link *Link) string {
	slug := uuid.NewV4().String()

	bytes, err := json.Marshal(link)
	if err != nil {
		log.Printf("Unexpected error when marshaling a valid link: %s", err)
		return ""
	}

	err = store.client.Set(slug, string(bytes), 0).Err()
	if err != nil {
		log.Printf("Unexpected error when storing a link: %s", err)
		return ""
	}

	return slug
}

func (store *RedisStore) clear() {
	store.client.FlushDb()
}