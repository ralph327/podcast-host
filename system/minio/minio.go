// Implements the Minio Client

package minio

import (
	"github.com/minio/minio-go"
	"log"
	"net/url"
)

// The type of minio instances that can be expected
type MinioType struct {
	Type        string
	Description string
}

// A map of types
var Types = map[string]*MinioType{
	"live": &MinioType{
		Type:        "live",
		Description: "Live episodes on block storage",
	},
	"archive": &MinioType{
		Type:        "archive",
		Description: "Archived episodes on IA S3",
	},
	"backup": &MinioType{
		Type:        "backup",
		Description: "Episodes backed up on Glacier S3",
	},
}

type Minio struct {
	Client *minio.Client
	Conf   *Conf
}

func New(c *Conf) (*Minio, error) {
	client, err := minio.New(c.Endpoint,
		c.AccessKeyID,
		c.SecretAccessKey,
		isSecure(c.Endpoint),
	)
	if err != nil {
		return nil, err
	}

	conf := c

	m := &Minio{Client: client, Conf: conf}

	return m, nil
}

type Conf struct {
	Type            *MinioType
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
}

func NewConf(typ *MinioType, url string, accessKey string, secretKey string) (*Conf, error) {
	c := &Conf{
		Type:            typ,
		Endpoint:        findHost(url),
		AccessKeyID:     accessKey,
		SecretAccessKey: secretKey,
		UseSSL:          isSecure(url),
	}

	log.Println("MinioConf:", c.Endpoint)

	return c, nil
}

func (c *Conf) GetType() *MinioType {
	return c.Type
}

// Find the Host of the given url.
func findHost(urlStr string) string {
	u, err := url.Parse(urlStr)
	log.Println("findHost", u)
	if err != nil {
		panic(err)
	}
	return u.Host
}

// Finds out whether the url is http(insecure) or https(secure).
func isSecure(urlStr string) bool {
	u, err := url.Parse("//" + urlStr)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	return u.Scheme == "https"
}
