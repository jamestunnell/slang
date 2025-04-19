package archives

import (
	"github.com/jamestunnell/slang"
)

type Repo struct {
	archives map[string]*TarGz
}

func NewRepo() *Repo {
	return &Repo{
		archives: map[string]*TarGz{},
	}
}

func (r *Repo) Add(a *TarGz) {
	r.archives[a.GetMeta().String()] = a
}

func (r *Repo) Remove(meta slang.PackageMeta) bool {
	key := meta.String()

	if _, found := r.archives[key]; !found {
		return false
	}

	delete(r.archives, key)

	return true
}

func (r *Repo) GetAllMeta() []slang.PackageMeta {
	metas := make([]slang.PackageMeta, len(r.archives))
	i := 0

	for _, archive := range r.archives {
		metas[i] = archive.GetMeta()

		i++
	}

	return metas
}

func (r *Repo) Get(meta slang.PackageMeta) (*TarGz, bool) {
	a, found := r.archives[meta.String()]

	return a, found
}
