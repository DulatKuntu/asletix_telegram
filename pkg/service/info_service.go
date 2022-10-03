package service

import (
	"asletix_telegram/model"
	"asletix_telegram/pkg/repository"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/wcharczuk/go-chart/v2"
)

type InfoService struct {
	repoUser repository.User
}

func NewInfoService(repoUser repository.User) *InfoService {
	return &InfoService{repoUser: repoUser}
}

func (s *InfoService) Registration() ([]*model.DayCount, int64, error) {
	log.Print("Registration")
	data, err := s.repoUser.RegistrationLastMonthByDays()
	if err != nil {
		return nil, 0, nil
	}
	log.Print("Registration 2")
	totalCount, err := s.repoUser.TotalCount()
	if err != nil {
		return nil, 0, nil
	}
	log.Print("Registration 3")
	return data, totalCount, nil
}

func (s *InfoService) UniqueWorkout() ([]*model.DayCount, error) {
	data, err := s.repoUser.UniqueWorkoutLastMonthByDays()
	if err != nil {
		return nil, nil
	}

	return data, nil
}

func (s *InfoService) UniqueWorkoutDWM() (int, int, int, error) {
	day, err := s.repoUser.UniqueWorkoutLastDays(0)
	if err != nil {
		return 0, 0, 0, nil
	}
	week, err := s.repoUser.UniqueWorkoutLastDays(6)
	if err != nil {
		return 0, 0, 0, nil
	}
	month, err := s.repoUser.UniqueWorkoutLastDays(29)
	if err != nil {
		return 0, 0, 0, nil
	}

	return day, week, month, nil
}

func (s *InfoService) OpenApp() ([]*model.DayCount, error) {
	data, err := s.repoUser.OpenAppLastMonthByDays()
	if err != nil {
		return nil, nil
	}

	return data, nil
}

func (s *InfoService) OpenAppDWM() (int, int, int, error) {
	day, err := s.repoUser.OpenAppLastDays(0)
	if err != nil {
		return 0, 0, 0, nil
	}
	week, err := s.repoUser.OpenAppLastDays(6)
	if err != nil {
		return 0, 0, 0, nil
	}
	month, err := s.repoUser.OpenAppLastDays(29)
	if err != nil {
		return 0, 0, 0, nil
	}

	return day, week, month, nil
}

func (s *InfoService) PlotDayCount(plotDataArray []*model.DayCount) (*os.File, error) {
	log.Print(len(plotDataArray))

	times := []time.Time{}
	count := []float64{}

	for _, single := range plotDataArray {
		times = append(times, single.Time)
		count = append(count, float64(single.Count))
	}

	graph := chart.Chart{
		YAxis: chart.YAxis{
			ValueFormatter: chart.IntValueFormatter,
		},
		Series: []chart.Series{
			chart.TimeSeries{
				XValues: times,
				YValues: count,
			},
		},
	}

	f, err := ioutil.TempFile("/home/ubuntu/data", "temp-*.png")
	if err != nil {

		log.Print("PlotDayCount 1 error")
		return nil, err
	}

	err = graph.Render(chart.PNG, f)

	return f, err
}

func (s *InfoService) Referal() ([]*model.RefCount, error) {
	return s.repoUser.Referal()
}
