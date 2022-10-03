package handler

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

// LSOFSinglePlotWithMessage used to send image and message
func (h *Handler) LSOFSinglePlotWithMessage(chatID int64, message string, file *os.File) error {
	nfile, err := os.Open("/home/ubuntu/data/" + filepath.Base(file.Name()))
	if err != nil {
		return err
	}
	defer nfile.Close()
	log.Print("LSOFSinglePlotWithMessage")

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	log.Print("LSOFSinglePlotWithMessage1")
	part, err := writer.CreateFormFile("photo", filepath.Base(nfile.Name()))
	log.Print("LSOFSinglePlotWithMessage2")
	if err != nil {
		return err
	}
	io.Copy(part, nfile)

	log.Print("LSOFSinglePlotWithMessage3")
	//Write additional parameters to multipart
	writer.WriteField("caption", message)
	writer.WriteField("chat_id", strconv.FormatInt(chatID, 10))
	writer.Close()

	log.Print("LSOFSinglePlotWithMessage4")
	r, err := http.NewRequest("POST", "https://api.telegram.org/bot"+os.Getenv("BotToken")+"/sendPhoto", body)
	if err != nil {
		return err
	}

	log.Print("LSOFSinglePlotWithMessage4")
	r.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	log.Print(bodyString)

	return nil
}

// LSOFSingleMessage used to send message
func (h *Handler) LSOFSingleMessage(chatID int64, message string) error {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	//Write additional parameters to multipart
	writer.WriteField("text", message)
	writer.WriteField("chat_id", strconv.FormatInt(chatID, 10))
	writer.Close()

	r, err := http.NewRequest("POST", "https://api.telegram.org/bot"+os.Getenv("BotToken")+"/sendMessage", body)
	if err != nil {
		return err
	}

	r.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return nil
}
