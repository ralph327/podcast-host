// Handles data passed to templates

package payload

import (
	"github.com/spf13/viper"
)

type Payload struct {
	Data map[string]interface{}
}

type PayloadService interface {
	GetData() map[string]interface{}
}

func New() (*Payload, error) {
	p := new(Payload)

	p.Data = make(map[string]interface{})

	return p, nil
}

func (p *Payload) GetData() map[string]interface{} {
	return p.Data
}

func (p *Payload) Init(conf *viper.Viper) error {
	p.Data["SiteDetails"] = struct {
		BaseURL  string
		SiteName string
		Title    string
	}{
		conf.GetString(conf.GetString("Env") + "BaseURL"),
		conf.GetString(conf.GetString("Env") + "SiteName"),
		conf.GetString(conf.GetString("Env") + "SiteName"),
	}

	return nil
}

// Delete all contents of the payload
func (p *Payload) Clear() error {
	for k := range p.Data {
		delete(p.Data, k)
	}

	return nil
}
