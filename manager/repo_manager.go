package manager

import (
	"warung_makan_bahari/repository"
)

type RepoManager interface {
	MenuRepo() repository.MenuRepo
}

type repoManager struct {
	infra Infra
}

func (r *repoManager) MenuRepo() repository.MenuRepo {
	return repository.NewMenuRepo(r.infra.SqlDb())
}

func NewRepoManager(infra Infra) RepoManager {
	return &repoManager{
		infra: infra,
	}
}
