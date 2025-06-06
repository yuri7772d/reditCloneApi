package entities

type (
	TopicLike struct {
		topicId      string
		likedUserIDs map[string]struct{} // เก็บ userId ของผู้ที่ like topic นี้ไปแล้ว
	}
	TopicLikeDatabaseSgyma      struct{} //โครงสร้างตารางไว้ migrate db
	TopicLikeCountDatabaseSgyma struct{}
)

func NewTopicLike(topicId string) TopicLike {
	return TopicLike{
		topicId:      topicId,
		likedUserIDs: make(map[string]struct{}),
	}
}

func (t *TopicLike) Update(userIDs []string) error {
	t.likedUserIDs = make(map[string]struct{}, len(userIDs))
	for _, id := range userIDs {
		t.likedUserIDs[id] = struct{}{}
	}
	return nil
}

func (t *TopicLike) Add(userID string) error {
	t.likedUserIDs[userID] = struct{}{}
	return nil
}

func (t *TopicLike) Remove(userID string) error {
	delete(t.likedUserIDs, userID)
	return nil
}

func (t *TopicLike) Tracked() []string {
	users := make([]string, 0, len(t.likedUserIDs))
	for id := range t.likedUserIDs {
		users = append(users, id)
	}
	return users
}

func (t *TopicLike) Total() int {
	return len(t.likedUserIDs)
}
