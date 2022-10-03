package service

import (
	"asletix_telegram/model"
	"asletix_telegram/pkg/repository"
	"os"
)

type Info interface {
	Registration() ([]*model.DayCount, []*model.DayCount, int64, error)
	OpenApp() ([]*model.DayCount, error)
	OpenAppDWM() (int, int, int, error)
	UniqueWorkout() ([]*model.DayCount, error)
	UniqueWorkoutDWM() (int, int, int, error)
	PlotDayCount(plotDataArray []*model.DayCount) (*os.File, error)
	PlotRegistration(dataHome, dataGym []*model.DayCount) (*os.File, error)
	Referal() ([]*model.RefCount, error)
	SelectedLanguage() ([]*model.LanguageCount, error)
	PlotLanguage(plotDataArray []*model.LanguageCount) (*os.File, error)
}

type Service struct {
	Info
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Info: NewInfoService(repos.User),
	}
}
