package models

type RegionPrimaryKey struct {
	Id string `json:"id"`
}

type CreateRegion struct {
	Name string `json:"name"`
}

type Region struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateRegion struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type RegionGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type RegionGetListResponse struct {
	Count   int       `json:"count"`
	Regions []*Region `json:"regions"`
}
