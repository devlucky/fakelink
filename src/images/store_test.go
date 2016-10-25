package images

import (
	"testing"
	"os"
	"github.com/satori/go.uuid"
)

/*
	Generic test suite for stores
*/

func behavesLikeAStore(t *testing.T, store Store) {
	store.clear()
	testGetMissing(t, store)

	store.clear()
	testPutAndGet(t, store)
}

func testGetMissing(t *testing.T, store Store) {
	img := store.Get("missing")
	if img != nil {
		t.Error("Expected missing image to not be retrievable")
	}
}

func testPutAndGet(t *testing.T, store Store) {
	img := generateRandomImage()

	err := store.Put("some-image", img)
	if err != nil {
		t.Fatal("Unexpected error on image .Put", err)
	}

	retrievedImg := store.Get("some-image")
	if retrievedImg == nil {
		t.Error("Expected .Get image to retrieve the image we just saved")
	}

	if !imagesAreEqual(img, retrievedImg) {
		t.Error("Expected put and retrieved images to be equal")
	}
}

/*
	All implementations comply with the expected behavior
*/

func TestInMemoryStore(t *testing.T) {
	store := NewInMemoryStore()
	behavesLikeAStore(t, store)
}

func TestS3Store(t *testing.T) {
	store := NewS3Store(
		os.Getenv("MINIO_HOST"),
		os.Getenv("MINIO_PORT"),
		os.Getenv("MINIO_ACCESS_KEY"),
		os.Getenv("MINIO_SECRET_KEY"),
	)
	behavesLikeAStore(t, store)
}


/*
	BENCHMARKS
 */

func benchmarkStore(b *testing.B, store Store) {
	img := generateRandomImage()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		slug := uuid.NewV4().String()
		store.Put(slug, img)
		store.Get(slug)
	}
}

func BenchmarkInMemoryStore(b *testing.B) {
	store := NewInMemoryStore()
	benchmarkStore(b, store)
}

func BenchmarkS3Store(b *testing.B) {
	store := NewS3Store(
		os.Getenv("MINIO_HOST"),
		os.Getenv("MINIO_PORT"),
		os.Getenv("MINIO_ACCESS_KEY"),
		os.Getenv("MINIO_SECRET_KEY"),
	)
	benchmarkStore(b, store)
}