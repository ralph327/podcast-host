// Implement usage of minio client with tus to upload files to live file server

package files

import (
	"codenex.us/ralph/podcast-host"
	"github.com/eventials/go-tus"
	"github.com/tus/tusd"
	"io"
)

type Bucket struct {
	P       *podhost.Podcast
	Buckets map[string]string
}

type Object struct {
	P       *podhost.Episode
	Objects map[string]string
}

/* TUS
 */
type TusClient struct {
	Client *tus.Client
}

// Create TUS client for resumable uploads
func NewTusClient(url string, config *tus.Config) (*TusClient, error) {
	tsc := new(TusClient)

	client, err := tus.NewClient(url, config)

	if err != nil {
		return nil, err
	}

	tsc.Client = client

	return tsc, nil
}

/* FILE
 */
type FileStore struct {
	Name string
}

func (f *FileStore) NewUpload(info tusd.FileInfo) (id string, err error) {

	return
}

func (f *FileStore) GetInfo(id string) (tusd.FileInfo, error) {
	info := tusd.FileInfo{}

	return info, nil
}

func (f *FileStore) WriteChunk(id string, offset int64, src io.Reader) (int64, error) {
	var n int64
	var err error
	return n, err
}
