package entities

// TopicLike ใช้เก็บข้อมูลการกด like ของผู้ใช้ในแต่ละ topic ในหน่วยความจำ (in-memory)
type TopicLike struct {
	topicID      int
	likedUserIDs map[int]struct{} // userID ที่เคยกด like topic นี้
}

// TopicLikeRecord ใช้สำหรับบันทึก user ที่เคยกด like ลงฐานข้อมูล
type TopicLikeRecord struct {
	UserID           int `gorm:"not null;"`
	TopicLikeCountID int `gorm:"not null;"`

	TopicLikeCount TopicLikeCountRecord `gorm:"foreignKey:TopicLikeCountID;references:ID;constraint:OnDelete:CASCADE"`
}

// TopicLikeCountRecord ใช้เก็บยอดรวมการกด like ในแต่ละ topic
type TopicLikeCountRecord struct {
	ID      int `gorm:"primaryKey;autoIncrement;"`
	TopicID int `gorm:"not null;"`
	Count   int `gorm:"not null;"`
}

func (TopicLikeRecord) TableName() string {
	return "topic_like"
}

func (TopicLikeCountRecord) TableName() string {
	return "topic_like_count"
}

// NewTopicLike สร้าง TopicLike instance ใหม่
func NewTopicLike(topicID int) TopicLike {
	return TopicLike{
		topicID:      topicID,
		likedUserIDs: make(map[int]struct{}),
	}
}

// GetTopicID คืนค่า topic ID ของ instance นี้
func (t *TopicLike) GetTopicID() int {
	return t.topicID
}

// Update แทนที่ userID ทั้งหมดที่เคยกด like ด้วยรายการใหม่
func (t *TopicLike) Update(userIDs []int) {
	t.likedUserIDs = make(map[int]struct{}, len(userIDs))
	for _, id := range userIDs {
		t.likedUserIDs[id] = struct{}{}
	}
}

// Add เพิ่ม userID เข้าไปในรายการ like
func (t *TopicLike) Add(userID int) {
	t.likedUserIDs[userID] = struct{}{}
}

// Remove ลบ userID ออกจากรายการ like
func (t *TopicLike) Remove(userID int) {
	delete(t.likedUserIDs, userID)
}

// Tracked คืนค่า slice ของ userID ที่เคยกด like
func (t *TopicLike) Tracked() []int {
	users := make([]int, 0, len(t.likedUserIDs))
	for id := range t.likedUserIDs {
		users = append(users, id)
	}
	return users
}

// Total คืนค่าจำนวนผู้ใช้ที่เคยกด like
func (t *TopicLike) Total() int {
	return len(t.likedUserIDs)
}
