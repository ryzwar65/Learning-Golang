package helpers

type AlamatSaya struct {
	City, Province string
}

func (alamatsaya AlamatSaya) Alamatku(alamatku map[string]string) string {
	return alamatsaya.City
}
