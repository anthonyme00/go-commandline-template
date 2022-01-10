package repository

type Repository struct {
	SQLRepository *ISQLRepository
}

func NewRepository(sqlRepo ISQLRepository) *Repository {
	return &Repository{
		SQLRepository: &sqlRepo,
	}
}

func (r *Repository) Clean() {
	(*r.SQLRepository).Clean()
}
