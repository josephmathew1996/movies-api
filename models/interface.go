package models

//UserService interface
type UserService interface {
	GetUserByEmail(string) (User, error)
}

//MovieService interface
type MovieService interface {
	GetAllMovies(MovieFilter) ([]Movie, int, error)
	GetMovieByID(int) (Movie, error)
	Create(Movie) (Movie, error)
	Update(Movie) error
	Delete(int) error
}
