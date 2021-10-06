package repository

type Space interface {
	Fetch() (error)
}
