package repository

type (
	topicRepositoryImpl struct{}
)

func newCommentRepository() TopicRepository {
	return &topicRepositoryImpl{}
}
