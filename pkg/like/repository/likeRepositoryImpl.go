package repository

type (
	likeRepositoryImpl struct{}
)

func newCommentRepository() LikeRepository {
	return &likeRepositoryImpl{}
}
