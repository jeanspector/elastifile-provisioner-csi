package emanage

import (
	"github.com/pborman/uuid"

	"eurl"
	"rest"
)

const tenantsUri = "api/tenants"

type tenants struct {
	conn *rest.Session
}

type tenantRes struct {
	Id   int       `json:"id"`
	Name string    `json:"name"`
	Uuid uuid.UUID `json:"uuid"`
	Url  eurl.URL  `json:"url"`
}

func (t *tenants) GetAll(opt *GetAllOpts) ([]tenantRes, error) {
	if opt == nil {
		opt = &GetAllOpts{}
	}

	var result []tenantRes
	if err := t.conn.Request(rest.MethodGet, tenantsUri, opt, &result); err != nil {
		return nil, err
	}

	return result, nil
}
