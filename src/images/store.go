package images

import (
	"image"
	"github.com/aws/aws-sdk-go/aws/session"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"bytes"
	"image/jpeg"
)

type Store interface {
	Put(key string, img image.Image) (url string, err error)
	Get(key string) (img image.Image)
	clear()
}

/*
	In-memory implementation of a Store. Used for testing
 */

type InMemoryStore struct {
	images map[string]image.Image
}

func NewInMemoryStore() (*InMemoryStore) {
	return &InMemoryStore{
		images: make(map[string]image.Image),
	}
}

func (store *InMemoryStore) Put(key string, img image.Image) (url string, err error) {
	store.images[key] = img
	url = fmt.Sprintf("http://127.0.0.1/%s", key)
	return
}

func (store *InMemoryStore) Get(key string) (image.Image) {
	return store.images[key]
}

func (store *InMemoryStore) clear() {
	store.images = make(map[string]image.Image)
}

/*
	Implementation of a Store based on AWS S3's API and SDK
 */

const bucketName = "link-images"

type S3Store struct {
	client *s3.S3
	urlPattern string
}

func NewS3Store(host, port, accessKey, accessSecret, publicUrl string) (*S3Store) {
	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(accessKey, accessSecret, ""),
		Endpoint:         aws.String(fmt.Sprintf("http://%s:%s", host, port)),
		Region:           aws.String("us-east-1"),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	}
	store := &S3Store{
		client: s3.New(session.New(s3Config)),
		urlPattern: publicUrl + "/" + bucketName + "/%s",
	}

	store.createBucket()
	return store
}

func (store *S3Store) Put(key string, img image.Image) (url string, err error) {
	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, img, nil)
	if err != nil {
		return
	}

	_, err = store.client.PutObject(&s3.PutObjectInput{
		Body:   bytes.NewReader(buf.Bytes()),
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		return
	}

	url = fmt.Sprintf(store.urlPattern, key)
	return
}

func (store *S3Store) Get(key string) (img image.Image) {
	out, err := store.client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		log.Print("Unexpected error retrieving image from S3", err)
		return nil
	}

	img, err = jpeg.Decode(out.Body)
	if err != nil {
		log.Print("Unexpected error decoding image retrieved from S3", err)
		return nil
	}

	return img
}

func (store *S3Store) clear() {
	out, err := store.client.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		log.Fatalf("Unexpected error listing all objects: %s", err)
	}

	objects := make([]*s3.ObjectIdentifier, len(out.Contents))

	for _, obj := range out.Contents {
		objects = append(objects, &s3.ObjectIdentifier{Key: obj.Key})
	}

	_, err = store.client.DeleteObjects(&s3.DeleteObjectsInput{
		Bucket: aws.String(bucketName),
		Delete: &s3.Delete{Objects: objects},
	})
	if err != nil {
		log.Fatalf("Unexpected error deleting all objects: %s", err)
	}
}

func (store *S3Store) createBucket() {
	_, err := store.client.HeadBucket(&s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})

	// If the bucket does not exist, we create it
	if err != nil {
		_, err = store.client.CreateBucket(&s3.CreateBucketInput{
			Bucket: aws.String(bucketName),
		})
		if err != nil {
			log.Fatal("Unexpected error creating an S3 bucket", err)
		}
	}
}


