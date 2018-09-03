package request

import "fmt"

type Path struct {
	base string

	GetMultiple string
	GetOne      string
	CreateOne   string
	UpdateOne   string
	DeleteOne   string
}

func (p *Path) SetPath(path string) {
	p.base = path
	idPath := fmt.Sprintf("%s%s", path, "{id:[0-9]+}")

	p.GetMultiple = path
	p.GetOne = idPath
	p.CreateOne = path
	p.UpdateOne = idPath
	p.DeleteOne = idPath
}

func (p *Path) Endpoint(endpoint string) string {
	return fmt.Sprintf("%s.%s/", p.base, endpoint)
}