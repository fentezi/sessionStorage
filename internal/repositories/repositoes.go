package repostirories

type Repositories struct {
	PostgreSQL
	Redis
}

func NewRepositories() *Repositories {
	return &Repositories{
		PostgreSQL: PostgreSQL{},
		Redis:      Redis{},
	}
}
