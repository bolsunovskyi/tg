package tg

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const baseURL = "https://api.telegram.org/bot"

type Client struct {
	httpClient *http.Client
	token      string
}

func MakeClient(token string, client *http.Client) Client {
	return Client{
		httpClient: client,
		token:      token,
	}
}

func (c Client) SendMessage(chatID int, text string) error {
	return c.sendRequest("sendMessage", map[string]string{
		"text":    text,
		"chat_id": strconv.Itoa(chatID),
	})
}

func (c Client) SendPhotoUrlInlineKeyboard(rq *ImageInlineRequest) (*Photo, error) {
	rsp, err := c.sendRequestJSON("sendPhoto", rq)
	if err != nil {
		return nil, err
	}

	var pRsp PhotoResponse
	if err := json.Unmarshal(rsp, &pRsp); err != nil {
		return nil, err
	}

	var big Photo
	for _, ph := range pRsp.Result.Photo {
		if ph.FileSize > big.FileSize {
			big = ph
		}
	}

	return &big, nil
}

func (c Client) EditMessageInlineKeyboard(chatID, messageID int, rq *InlineKeyboardMarkup) error {
	_, err := c.sendRequestJSON("editMessageReplyMarkup", map[string]interface{}{
		"chat_id":      chatID,
		"message_id":   messageID,
		"reply_markup": rq,
	})

	return err
}

func (c Client) SendPhotoUrl(chatID int, imageURL string) error {
	return c.sendRequest("sendPhoto", map[string]string{
		"photo":   imageURL,
		"chat_id": strconv.Itoa(chatID),
	})
}

func (c Client) sendRequestJSON(method string, rq interface{}) ([]byte, error) {
	bts, err := json.Marshal(rq)
	if err != nil {
		return nil, err
	}

	hrq, err := http.NewRequest("POST", fmt.Sprintf("%s%s/%s", baseURL, c.token, method),
		bytes.NewReader(bts))
	if err != nil {
		return nil, err
	}

	hrq.Header.Add("Content-Type", "application/json")
	rsp, err := c.httpClient.Do(hrq)
	if err != nil {
		return nil, err
	}

	bts, err = ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}

	if rsp.StatusCode != http.StatusOK {
		return nil, errors.New(string(bts))
	}

	return bts, nil
}

func (c Client) sendRequest(method string, values map[string]string) error {
	var uValues url.Values = make(map[string][]string)
	for k, v := range values {
		uValues[k] = []string{v}
	}

	rq, _ := http.NewRequest("POST", fmt.Sprintf("%s%s/%s", baseURL, c.token, method),
		strings.NewReader(uValues.Encode()))
	rq.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	rsp, err := c.httpClient.Do(rq)
	if err != nil {
		return err
	}

	if rsp.StatusCode != http.StatusOK {
		bts, _ := ioutil.ReadAll(rsp.Body)
		return errors.New(string(bts))
	}

	return nil
}

func (c Client) ChatAction(chatID int, action string) error {
	return c.sendRequest("sendChatAction", map[string]string{
		"action":  action,
		"chat_id": strconv.Itoa(chatID),
	})
}
