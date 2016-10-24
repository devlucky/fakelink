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
	"github.com/minio/minio-go"
)

type Store interface {
	Put(slug string, img image.Image) (url string, err error)
	Get(slug string) (img image.Image)
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

func (store *InMemoryStore) Put(slug string, img image.Image) (string, error) {
	store.images[slug] = img

	// Fake URL for testing purposes
	url := fmt.Sprintf("http://127.0.0.1/%s", slug)

	return url, nil
}

func (store *InMemoryStore) Get(slug string) (image.Image) {
	return store.images[slug]
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
}

func NewS3Store(host, port, accessKey, accessSecret string) (*S3Store) {
	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(accessKey, accessSecret, ""),
		Endpoint:         aws.String(fmt.Sprintf("http://%s:%s", host, port)),
		Region:           aws.String("us-east-1"),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	}
	store := &S3Store{
		client: s3.New(session.New(s3Config)),
	}

	store.createBucket()
	return store
}

func (store *S3Store) Put(slug string, img image.Image) (url string, err error) {
	bucket, key := aws.String(bucketName), aws.String(slug)

	// TODO: Test this using just the buffer
	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, img, nil)
	if err != nil {
		return
	}

	_, err = store.client.PutObject(&s3.PutObjectInput{
		Body:   bytes.NewReader(buf.Bytes()),
		Bucket: bucket,
		Key:    key,
	})
	if err != nil {
		return
	}

	// TODO: Find out how to return the actual URL (may need some steroids here)
	url = "some-url"
	return
}

func (store *S3Store) Get(slug string) (img image.Image) {
	out, err := store.client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(slug),
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





/*
	Implementation of a Store based on Minio's SDK
 */

type MinioStore struct {
	client *minio.Client
}

func NewMinioStore(host, port, accessKey, accessSecret string) (*MinioStore) {
	client, err := minio.New(fmt.Sprintf("%s:%s", host, port), accessKey, accessSecret, false)
	if err != nil {
		log.Fatalf("Unexpected error creating a Minio client: %s", err)
	}

	bucketExists, err := client.BucketExists(bucketName)
	if err != nil {
		log.Fatalf("Unexpected error checking whether a bucket exists: %s", err)
	}

	if !bucketExists {
		err = client.MakeBucket(bucketName, "us-east-1")
		if err != nil {
			log.Fatalf("Unexpected error creating a new bucket: %s", err)
		}
	}

	return &MinioStore{
		client: client,
	}
}

func (store *MinioStore) Put(slug string, img image.Image) (url string, err error) {
	// TODO: Test this using just the buffer
	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, img, nil)
	if err != nil {
		return
	}

	_, err = store.client.PutObject(bucketName, slug, bytes.NewReader(buf.Bytes()), "image/jpeg")
	if err != nil {
		return
	}

	// TODO: Find out how to return the actual URL (may need some steroids here)
	url = "some-url"
	return
}

func (store *MinioStore) Get(slug string) (img image.Image) {
	obj, err := store.client.GetObject(bucketName, slug)
	if err != nil {
		log.Print("Unexpected error retrieving image from Minio", err)
		return nil
	}

	img, err = jpeg.Decode(obj)
	if err != nil {
		log.Print("Unexpected error decoding image retrieved from Minio", err)
		return nil
	}

	return img
}

func (store *MinioStore) clear() {
	doneCh := make(chan struct{})
	defer close(doneCh)

	for object := range store.client.ListObjects(bucketName, "*", true, doneCh) {
		if object.Err != nil {
			log.Fatalf("Unexpected error listing objects from Minio: %s", object.Err)
		}

		err := store.client.RemoveObject(bucketName, object.Key)
		if err != nil {
			log.Fatalf("Unexpected error removing object from Minio: %s", object.Err)
		}
	}
}
