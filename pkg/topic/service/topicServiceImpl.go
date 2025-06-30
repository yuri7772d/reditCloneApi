package service

import (
	"github.com/yuri7772d/reditCloneApi/entities"
	profileModel "github.com/yuri7772d/reditCloneApi/pkg/profile/model"
	profileService "github.com/yuri7772d/reditCloneApi/pkg/profile/service"
	"github.com/yuri7772d/reditCloneApi/pkg/topic/model"
	"github.com/yuri7772d/reditCloneApi/pkg/topic/repository"
)

type topicImpl struct {
	repository repository.Topic
}

func New(repository repository.Topic) Topic {
	return &topicImpl{repository}
}

func (s *topicImpl) Post(profileID int, title string, images ...string) {

	if len(images) != 0 {
		Topic := &entities.Topic{ProfileID: profileID, Title: title}
		for _, img := range images {
			Topic.Images = append(Topic.Images, entities.TopicImage{Image: img})
		}
		tx, _ := s.repository.Begin()
		if err := s.repository.Create(Topic); err != nil {
			s.repository.Rollback(tx)
		}
		s.repository.Commit(tx)
	}

}

func (s *topicImpl) Listing(profileService profileService.Profile) []*model.Topic {

	topicEntities, err := s.repository.Listing()
	if err != nil {
		return nil
	}

	var profileIDs []int
	for _, topicEntity := range topicEntities {
		profileIDs = append(profileIDs, topicEntity.ProfileID)
	}

	profileBriefs := profileService.GetBrief(profileIDs...)
	if len(profileBriefs) == 0 {
		return nil
	}

	profileBriefsMap := make(map[int]profileModel.Brief, len(profileBriefs))
	for _, profileBrief := range profileBriefs {
		profileBriefsMap[profileBrief.ID] = profileBrief
	}

	topicModel := make([]*model.Topic, len(topicEntities))

	for i, topicEntity := range topicEntities {
		var images []string

		for _, Img := range topicEntity.Images {
			images = append(images, Img.Image)
		}

		topicModel[i] = &model.Topic{
			ID:           topicEntity.ID,
			Tiltle:       topicEntity.Title,
			Images:       images,
			LikeCount:    topicEntity.Like.Count,
			ProfileID:    topicEntity.ProfileID,
			ProfileName:  profileBriefsMap[topicEntity.ProfileID].Name,
			ProfileImage: profileBriefsMap[topicEntity.ProfileID].Image,
		}
	}

	return topicModel
}
