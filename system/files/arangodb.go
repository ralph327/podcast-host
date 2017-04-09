// Implementation of files/data structures backed by arangodb

package files

import (
	"codenex.us/ralph/podcast-host/system/db"
	"codenex.us/ralph/podcast-host/system/minio"
	"io"
	//"gopkg.in/gin-gonic/gin.v1"
	"log"
)

type ArangoFiles struct {
	Fldr *Folder
	DB   *db.ArangoDB
}

func NewArangoFiles(db *db.ArangoDB) (*ArangoFiles, error) {
	var err error
	af := new(ArangoFiles)

	// Create new folder to be instantiated
	af.Fldr, err = NewFolder()
	if err != nil {
		return nil, err
	}

	// set database
	af.DB = db

	log.Println("returning from NewArangoFiles()")

	return af, nil
}

func (af *ArangoFiles) Upload(obj string, filepath io.Reader, mc *minio.Minio) {
	log.Println("calling MC client")
	// Upload the zip file with FPutObject
	//_, err := mc.Client.FPutObject(af.Fldr.Bkt.Name, obj, filepath, "")
	n, err := mc.Client.PutObject("docs", obj, filepath, "application/octet-stream")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Uploaded %i bytes", n)
}
