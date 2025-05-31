package repository

type (
	commentRepositoryImpl struct{}
)

func newCommentRepository() CommentRepository {
	return &commentRepositoryImpl{}
}
