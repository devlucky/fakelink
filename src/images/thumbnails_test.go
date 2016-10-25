package images

import (
	"testing"
	"github.com/disintegration/imaging"
	"bytes"
	"image/jpeg"
	"fmt"
	"os"
)

const (
	testMaxWidth = 512
	testMaxHeight = 512
	testDumpThumbnails = true
)

func benchmarkThumbnailWithFiler(b *testing.B, filter imaging.ResampleFilter, filterName string) {
	img := getFixtureImage("sharknado.jpg")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		newImg := imaging.Fit(img, testMaxWidth, testMaxHeight, filter)

		if testDumpThumbnails && i == 0 {
			err := os.MkdirAll("../../tmp/", os.ModePerm)
			if err != nil {
				b.Fatalf("Unexpected error creating a temporary directory to hold the thumbnails: %s", err)
			}

			out, err := os.Create(fmt.Sprintf("../../tmp/thumbnail_%d_%s.jpg", testMaxWidth, filterName))
			if err != nil {
				b.Fatalf("Unexpected error dumping thumbnail to a file: %s", err)
			}

			jpeg.Encode(out, newImg, &jpeg.Options{
				Quality: 80,
			})
		}
	}
}

func BenchmarkNearestNeighbor(b *testing.B) {
	benchmarkThumbnailWithFiler(b, imaging.NearestNeighbor, "NearestNeighbor")
}

func BenchmarkBox(b *testing.B) {
	benchmarkThumbnailWithFiler(b, imaging.Box, "Box")
}

func BenchmarkLinear(b *testing.B) {
	benchmarkThumbnailWithFiler(b, imaging.Linear, "Linear")
}

func BenchmarkMitchellNetravali(b *testing.B) {
	benchmarkThumbnailWithFiler(b, imaging.MitchellNetravali, "MitchellNetravali")
}

func BenchmarkCatmullRom(b *testing.B) {
	benchmarkThumbnailWithFiler(b, imaging.CatmullRom, "CatmullRom")
}

func BenchmarkGaussian(b *testing.B) {
	benchmarkThumbnailWithFiler(b, imaging.Gaussian, "Gaussian")
}

func BenchmarkLanczos(b *testing.B) {
	benchmarkThumbnailWithFiler(b, imaging.Lanczos, "Lanczos")
}

// Test with actual image and see how it differs
func TestThumbnail(t *testing.T) {
	img := getFixtureImage("sharknado.jpg")

	sizes := []int{64, 128, 256, 512, 1024}
	for _, size := range sizes {
		thumbnail := Thumbnail(img, size, size)

		buf := new(bytes.Buffer)
		err := jpeg.Encode(buf, thumbnail, nil)
		if err != nil {
			t.Fatalf("Unexpected error encoding a thumbnail in JPEG: %s", err)
		}

		t.Logf("Thumbnail of size %dx%d takes %d bytes", size, size, buf.Len())
	}
}