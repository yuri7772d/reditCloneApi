package server
topicRepository := _topicRepository.New(db, echo.New().Logger)
	topicService := _topicService.New(topicRepository)
	topicController := _topicController.New(topicService)