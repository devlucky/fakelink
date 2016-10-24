package images

import (
	"testing"
	"image"
	"image/color"
	"math/rand"
	"net/url"
	"os"
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

	imgUrlStr, err := store.Put("some-image", img)
	if err != nil {
		t.Fatal("Unexpected error on image .Put", err)
	}

	_, err = url.Parse(imgUrlStr)
	if err != nil {
		t.Errorf("Expected %s to be an actual URL", imgUrlStr)
	}


	retrievedImg := store.Get("some-image")
	if retrievedImg == nil {
		t.Error("Expected .Get image to retrieve the image we just saved")
	}

	if !imagesAreEqual(img, retrievedImg) {
		t.Error("Expected put and retrieved images to be equal")
	}
}

func imagesAreEqual(a, b image.Image) (bool) {
	return a.Bounds().Eq(b.Bounds())
}

func generateRandomImage() (image.Image){
	bounds := image.Rect(0,0,20,20)
	img := image.NewRGBA(bounds)
	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X - 1; x++ {
			a0 := uint8(rand.Float32() * 255)
			rgb0 := uint8(rand.Float32() * 255)
			rgb0 = rgb0 * a0
			img.SetRGBA(x, y, color.RGBA{rgb0,rgb0,rgb0,a0})
		}
	}

	return img
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