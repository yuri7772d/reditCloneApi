package repository

type userRepositoryImpl struct{}

func newCommentRepository() UserRepository {
	return &userRepositoryImpl{}
}
