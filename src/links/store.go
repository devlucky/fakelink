package links

import (
	"encoding/json"
	"fmt"
	"gopkg.in/redis.v5"
	"log"
	"math/rand"
)

// Store allows saving and retrieving user-generated links.
type Store interface {
	Find(slug string) *Link
	FindRandom() (slug string)
	Create(link *Link) string
	clear()
}

// InMemoryStore is an in-memory implementation of a template store.
type InMemoryStore struct {
	public  map[string]*Link
	private map[string]*Link
}

// NewInMemoryStore creates a new in-memory store.
func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		public:  make(map[string]*Link),
		private: make(map[string]*Link),
	}
}

// Find retrieves a single Link from its slug.
func (store *InMemoryStore) Find(slug string) *Link {
	if hasFlag(slug, privateFlag) {
		return store.private[slug]
	}
	return store.public[slug]
}

// FindRandom retrieves a random Link slug.
func (store *InMemoryStore) FindRandom() (slug string) {
	if len(store.public) == 0 {
		return
	}

	i := 0
	n := rand.Int() % len(store.public)

	for s := range store.public {
		if i == n {
			slug = s
			break
		}
		i++
	}

	return
}

// Create creates a new Link.
func (store *InMemoryStore) Create(link *Link) string {
	slug := generateSlug(link)

	if link.Private {
		store.private[slug] = link
	} else {
		store.public[slug] = link
	}

	return slug
}

func (store *InMemoryStore) clear() {
	store.public = make(map[string]*Link)
	store.private = make(map[string]*Link)
}

// RedisStore is a redis based implementation of a link store.
type RedisStore struct {
	public  *redis.Client
	private *redis.Client
}

// NewRedisStore create a new in-memory store.
func NewRedisStore(host, port, password string) *RedisStore {
	return &RedisStore{
		public: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", host, port),
			Password: password,
			DB:       0,
		}),
		private: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", host, port),
			Password: password,
			DB:       1,
		}),
	}
}

// Find retrieves a single Link from its slug.
func (store *RedisStore) Find(slug string) *Link {
	var db *redis.Client

	if hasFlag(slug, privateFlag) {
		db = store.private
	} else {
		db = store.public
	}

	str, err := db.Get(slug).Result()
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

// FindRandom retrieves a random Link slug.
func (store *RedisStore) FindRandom() (slug string) {
	slug, err := store.public.RandomKey().Result()
	if err != nil {
		log.Printf("Getting link with slug %s failed with error %s", slug, err)
		return
	}

	return
}

// Create creates a new Link.
func (store *RedisStore) Create(link *Link) string {
	var db *redis.Client

	slug := generateSlug(link)
	if link.Private {
		db = store.private
	} else {
		db = store.public
	}

	bytes, err := json.Marshal(link)
	if err != nil {
		log.Printf("Unexpected error when marshaling a valid link: %s", err)
		return ""
	}

	err = db.Set(slug, string(bytes), 0).Err()
	if err != nil {
		log.Printf("Unexpected error when storing a link: %s", err)
		return ""
	}

	return slug
}

func (store *RedisStore) clear() {
	store.public.FlushDb()
	store.private.FlushDb()
}
