package service

import (
	"asletix_telegram/model"
	"asletix_telegram/pkg/repository"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"github.com/wcharczuk/go-chart/v2"
)

type InfoService struct {
	repoUser repository.User
}

func NewInfoService(repoUser repository.User) *InfoService {
	return &InfoService{repoUser: repoUser}
}

func (s *InfoService) SelectedLanguage() ([]*model.LanguageCount, error) {
	return s.repoUser.SelectedLanguage()
}

func (s *InfoService) Registration() ([]*model.DayCount, []*model.DayCount, int64, error) {
	log.Print("Registration")
	dataHome, err := s.repoUser.RegistrationLastMonthByDays("home")
	if err != nil {
		return nil, nil, 0, nil
	}
	dataGym, err := s.repoUser.RegistrationLastMonthByDays("gym")
	if err != nil {
		return nil, nil, 0, nil
	}
	log.Print("Registration 2")
	totalCount, err := s.repoUser.TotalCount()
	if err != nil {
		return nil, nil, 0, nil
	}
	log.Print("Registration 3")
	return dataHome, dataGym, totalCount, nil
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

func (s *InfoService) PlotRegistration(dataHome, dataGym []*model.DayCount) (*os.File, error) {
	times := []time.Time{}
	countHome := make([]opts.LineData, 0)
	countGym := make([]opts.LineData, 0)

	for i := 0; i < 30; i++ {
		times = append(times, dataHome[i].Time)
		countHome = append(countHome, opts.LineData{Value: dataHome[i].Count})
		countGym = append(countGym, opts.LineData{Value: dataGym[i].Count})

	}
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title: "User regsitration",
		}))

	// Put data into instance
	line.SetXAxis([]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}).
		AddSeries("Home", countHome).
		AddSeries("Gym", countGym).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	f, err := ioutil.TempFile("/home/ubuntu/data", "temp-*.png")
	if err != nil {

		log.Print("PlotDayCount 1 error")
		return nil, err
	}
	err = line.Render(f)
	return f, err
}

func (s *InfoService) PlotLanguage(plotDataArray []*model.LanguageCount) (*os.File, error) {

	languages := []string{}
	count := make([]opts.BarData, 0)

	for _, single := range plotDataArray {
		languages = append(languages, single.Language)
		count = append(count, opts.BarData{Value: single.Count})
	}

	bar := charts.NewBar()
	// set some global options like Title/Legend/ToolTip or anything else
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title: "User registration by languages",
	}))

	// Put data into instance
	bar.SetXAxis(languages).
		AddSeries("Category A", count)
	// Where the magic happens
	f, err := ioutil.TempFile("/home/ubuntu/data", "temp-*.png")
	if err != nil {

		log.Print("PlotDayCount 1 error")
		return nil, err
	}
	err = bar.Render(f)
	return f, err
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
