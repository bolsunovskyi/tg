package tg

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"gopkg.in/h2non/gock.v1"
)

const (
	chatID = 148901293
)

func init() {
	if err := godotenv.Load("./.env"); err != nil {
		log.Fatalln(err)
	}
}

func TestClient_SendMessage(t *testing.T) {
	cl := MakeClient(os.Getenv("tg_token"), http.DefaultClient)
	if err := cl.SendMessage(148901293, "hello world"); err != nil {
		t.Error(err)
		return
	}
}

func TestClient_SendPhotoUrlInlineKeyboard(t *testing.T) {
	cl := MakeClient(os.Getenv("tg_token"), http.DefaultClient)
	rsp, err := cl.SendPhotoUrlInlineKeyboard(&ImageInlineRequest{
		ChatID: chatID,
		Photo:  "https://d1.awsstatic.com/Digital%20Marketing/sitemerch/sign-in/en/Site-Merch_PAC_Backup-Restore_Sign-in_EN.7e859982944e5753420a056f5aefc1b14c07f39e.png",
		ReplyMarkup: InlineKeyboardMarkup{
			InlineKeyboard: [][]InlineKeyboardButton{
				{
					{
						Text:         "Like",
						CallbackData: "Like",
					},
					{
						Text:         "Dislike",
						CallbackData: "Dislike",
					},
				},
			},
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Log(rsp)
}

func TestClient_SendMessageFail(t *testing.T) {
	defer gock.Off()

	gock.New("https://api.telegram.org").
		Post("/bot/" + os.Getenv("tg_token") + "/sendMessage").
		Reply(400).
		JSON(map[string]string{"foo": "bar"})

	cl := MakeClient(os.Getenv("tg_token"), http.DefaultClient)
	if err := cl.SendMessage(148901293, "hello world"); err == nil {
		t.Error("no error on wrong response from telegram")
		return
	}
}

func TestClient_SendMessageFail2(t *testing.T) {
	defer gock.Off()

	gock.New("https://api.telegram.org").
		Post("/bot" + os.Getenv("tg_token") + "/sendMessage").
		Reply(400).
		JSON(map[string]string{"foo": "bar"})

	cl := MakeClient(os.Getenv("tg_token"), http.DefaultClient)
	if err := cl.SendMessage(148901293, "hello world"); err == nil {
		t.Error("no error on wrong response from telegram")
		return
	}
}

func TestClient_ChatAction(t *testing.T) {
	cl := MakeClient(os.Getenv("tg_token"), http.DefaultClient)
	if err := cl.ChatAction(148901293, ActionTyping); err != nil {
		t.Fatal(err)
	}
}
