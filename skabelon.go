package skabelon

type Skabelon struct {
	*skabelonInternals
}

type skabelonInternals struct {
	inited      bool
	environmets map[string]string
}

func New() *Skabelon {
	skabelonInstance := new(Skabelon)
	skabelonInstance.skabelonInternals = new(skabelonInternals)
	return skabelonInstance
}
