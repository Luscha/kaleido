package datasource

type Result struct {
	ID     string
	Status int
	Body   []byte
	Err    error
}

type DataSource struct {
	Provider   string        `json:"provider"`
	Repository string        `json:"repository"`
	Name       string        `json:"name"`
	Manifest   any           `json:"manifest"`
	Depends    []DataDepends `json:"depends"`
}

func (d DataSource) DependsOnSomething() bool {
	return 0 != len(d.Depends)
}

type DataDepends struct {
	Value    string `json:"value"`
	Template string `json:"template"`
}
