// Implement usage of minio client with tus to upload files to live file server

package files

import (
	"codenex.us/ralph/podcast-host"
	"codenex.us/ralph/podcast-host/system/minio"
	"github.com/eventials/go-tus"
	"github.com/tus/tusd"
	"io"
)

type Folder struct {
	Pdcst *podhost.Podcast
	Bkt   *podhost.Bucket
	P2B   *podhost.PodcastHasBucket
	Files map[string]*File
}

func NewFolder() (*Folder, error) {
	f := new(Folder)

	return f, nil
}

type File struct {
	Ep  *podhost.Episode
	Obj *podhost.Object
	E2O *podhost.EpisodeHasObject
}

// Uploader interface for routing uploads
type Uploader interface {
	Upload(obj string, filepath string, mc *minio.Minio)
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
