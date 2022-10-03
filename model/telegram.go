package model

import "time"

// RequestMessageTelegram is data that telegram sends to us
type RequestMessageTelegram struct {
	Message       *Message       `json:"message"`
	CallBackQuery *CallBackQuery `json:"callback_query"`
}

// CallBackQuery callback query
type CallBackQuery struct {
	ID      string   `json:"id"`
	Message *Message `json:"message"`
	Data    string   `json:"data"`
}

// Message gen message
type Message struct {
	MessageID int64  `json:"message_id"`
	Text      string `json:"text"`
	Chat      struct {
		ID int64 `json:"id"`
	} `json:"chat"`
}

//SendMessageTelegram is data that we should send to telegram
type SendMessageTelegram struct {
	ChatID      int64        `json:"chat_id"`
	Text        string       `json:"text"`
	ReplyMarkup *ReplyMarkup `json:"reply_markup"`
	MessageID   int64        `json:"message_id"`
}

type ReplyMarkup struct {
	InlineKeyboards [][]*InlineKeyboard `json:"inline_keyboard" bson:"inline_keyboard"`
}

type InlineKeyboard struct {
	Text         string `json:"text" bson:"text"`
	CallbackData string `json:"callback_data" bson:"callback_data"`
}

//TelegramSubscriber in database
type TelegramSubscriber struct {
	ChatID int64  `json:"chatID" bson:"chatID"`
	Status string `json:"status" bson:"status"`
}

// UserCount in db
type UserCount struct {
	Count int64 `json:"count" bson:"count"`
}

// ResponseMessage res message
type ResponseMessage struct {
	Ok     bool     `json:"ok"`
	Result *Message `json:"result"`
}

// ResponseMessageArray res messageArray
type ResponseMessageArray struct {
	Ok     bool       `json:"ok"`
	Result []*Message `json:"result"`
}

type LastData struct {
	Total int `json:"total" bson:"total"`
	Month int `json:"month" bson:"month"`
	Week  int `json:"week" bson:"week"`
	Day   int `json:"day" bson:"day"`
}

type DayCount struct {
	Count int       `json:"count" bson:"count"`
	Time  time.Time `json:"time" bson:"time"`
}

type RefCount struct {
	Count int    `json:"count" bson:"count"`
	Name  string `json:"name" bson:"name"`
}

type ByDayCount struct {
	Count int `json:"count" bson:"count"`
	Year  int `json:"year" bson:"year"`
	Month int `json:"month" bson:"month"`
	Day   int `json:"day" bson:"day"`
}

// GeneralInfo
type GeneralInfo struct {
	LastData  *LastData   `json:"lastData" bson:"lastData"`
	LastMonth []*DayCount `json:"lastMonth" bson:"lastMonth"`
}

// PieChartInfo
type PieChartInfo struct {
	NoFriends   int `json:"noFriends" bson:"noFriends"`
	FiveFriends int `json:"fiveFriends" bson:"fiveFriends"`
	MoreFriends int `json:"moreFriends" bson:"moreFriends"`
}
