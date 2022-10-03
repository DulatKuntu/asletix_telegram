package handler

import (
	"asletix_telegram/model"
	"asletix_telegram/pkg/service"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

const (
	TokenTelegramStat = "1847491332:AAHFMcJ_8wPCqG3yxPwSyMgLQCLS2Vv2Igc"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.POST("/info", h.Info)

	return router
}

func (h *Handler) Info(c *gin.Context) {
	log.Print("1")
	body := &model.RequestMessageTelegram{}
	if err := c.ShouldBindWith(&body, binding.JSON); err != nil {
		fmt.Println("could not decode request body", err)
		return
	}

	if body.Message != nil {
		switch strings.ToLower(body.Message.Text) {
		case "/registration":
			go h.RegistrationCount(body.Message.Chat.ID, c)
		// case "/activeUsers":
		case "/refers":
			go h.RefersCount(body.Message.Chat.ID, c)
		case "/openapp":
			go h.OpenAppCount(body.Message.Chat.ID, c)
		// case "/activeUsers":
		case "/uniqueworkout":
			go h.UniqueWorkout(body.Message.Chat.ID, c)
			// case "/activeUsers":
		case "/language":
			go h.SelectedLanguage(body.Message.Chat.ID, c)
		}
	}
}

func (h *Handler) SelectedLanguage(chatID int64, c *gin.Context) {
	data, err := h.services.SelectedLanguage()
	if err != nil {
		log.Print(err.Error())
		return
	}
	img, err := h.services.PlotLanguage(data)

	if err != nil {
		log.Print(err.Error())
		return
	}

	defer img.Close()
	defer os.Remove(img.Name()) // clean up

	var mes strings.Builder
	for _, item := range data {
		mes.WriteString(item.Language)
		mes.WriteString(": ")
		mes.WriteByte(byte(item.Count))
		mes.WriteByte('\n')
	}
	if err = h.LSOFSinglePlotWithMessage(chatID, mes.String(), img); err != nil {
		log.Print(err.Error())
	}
}

func (h *Handler) UniqueWorkout(chatID int64, c *gin.Context) {
	data, err := h.services.UniqueWorkout()
	if err != nil {
		log.Print(err.Error())
		return
	}
	img, err := h.services.PlotDayCount(data)

	if err != nil {
		log.Print(err.Error())
		return
	}

	defer img.Close()
	defer os.Remove(img.Name()) // clean up

	lastDay, lastWeek, lastMonth, err := h.services.UniqueWorkoutDWM()

	if err != nil {
		log.Print(err.Error())
		return
	}

	mes := "Month: " + strconv.Itoa(lastMonth) + "\nWeek: " + strconv.Itoa(lastWeek) + "\nDay: " + strconv.Itoa(lastDay)

	if err = h.LSOFSinglePlotWithMessage(chatID, mes, img); err != nil {
		log.Print(err.Error())
	}
}

func (h *Handler) OpenAppCount(chatID int64, c *gin.Context) {
	data, err := h.services.OpenApp()
	if err != nil {
		log.Print(err.Error())
		return
	}
	img, err := h.services.PlotDayCount(data)

	if err != nil {
		log.Print(err.Error())
		return
	}

	defer img.Close()
	defer os.Remove(img.Name()) // clean up

	lastDay, lastWeek, lastMonth, err := h.services.OpenAppDWM()

	if err != nil {
		log.Print(err.Error())
		return
	}

	mes := "Month: " + strconv.Itoa(lastMonth) + "\nWeek: " + strconv.Itoa(lastWeek) + "\nDay: " + strconv.Itoa(lastDay)

	if err = h.LSOFSinglePlotWithMessage(chatID, mes, img); err != nil {
		log.Print(err.Error())
	}
}

func (h *Handler) RegistrationCount(chatID int64, c *gin.Context) {
	log.Print("RegistrationCount")
	dataHome, dataGym, total, err := h.services.Registration()
	if err != nil {
		log.Print(err.Error())
		return
	}
	img, err := h.services.PlotRegistration(dataHome, dataGym)

	if err != nil {
		log.Print(err.Error())
		return
	}
	log.Print("RegistrationCount2")

	defer img.Close()
	defer os.Remove(img.Name()) // clean up

	lastDay := dataHome[len(dataHome)-1].Count
	lastDay = lastDay + dataGym[len(dataGym)-1].Count
	lastWeek := 0
	lastMonth := 0
	for i := 0; i < len(dataHome); i++ {
		lastMonth += dataHome[i].Count
		lastMonth += dataGym[i].Count

	}

	for i := len(dataHome) - 1; i >= len(dataHome)-7; i-- {
		lastWeek += dataHome[i].Count
		lastWeek += dataGym[i].Count
	}

	mes := "Total: " + strconv.Itoa(int(total)) + "\nMonth: " + strconv.Itoa(lastMonth) + "\nWeek: " + strconv.Itoa(lastWeek) + "\nDay: " + strconv.Itoa(lastDay)

	if err = h.LSOFSinglePlotWithMessage(chatID, mes, img); err != nil {
		log.Print(err.Error())
	}
}

func (h *Handler) RefersCount(chatID int64, c *gin.Context) {
	data, err := h.services.Referal()
	if err != nil {
		log.Print(err.Error())
		return
	}
	mes := ""

	for _, single := range data {
		mes += single.Name + " : " + strconv.Itoa(single.Count) + "\n"
	}

	if err = h.LSOFSingleMessage(chatID, mes); err != nil {
		log.Print(err.Error())
	}
}
