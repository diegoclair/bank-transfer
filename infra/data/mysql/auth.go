package mysql

type authRepo struct {
	db connection
}

// newAuthRepo returns a instance of dbrepo
func newAuthRepo(db connection) *authRepo {
	return &authRepo{
		db: db,
	}
}
