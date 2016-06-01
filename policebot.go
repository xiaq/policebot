package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/xiaq/tg"
)

var emojis = []rune{
	'\U0001F46E', // police officer
	'\U0001F693', // police car
	'\U0001F694', // oncoming police car
	'\U0001F6A8', // police cars revolving light
}

const (
	minLen = 10
	maxLen = 100
	police = "\U0001F6A8"
)

func main() {
	buf, err := ioutil.ReadFile("token.txt")
	if err != nil {
		log.Fatalln("cannot read token file:", err)
	}
	token := strings.TrimSpace(string(buf))

	rand.Seed(time.Now().UnixNano())

	bot := tg.NewCommandBot(token)
	bot.OnCommand("callpolice", callpolice)
	bot.Main()
}

func callpolice(b *tg.CommandBot, text string, msg *tg.Message) {
	b.Get("/sendMessage", tg.Query{
		"chat_id": msg.Chat.ID,
		"text":    makeReply(),
	}, nil)
}

func makeReply() string {
	n := randrange(minLen, maxLen)
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		buf.WriteRune(emojis[rand.Intn(len(emojis))])
	}
	return buf.String()
}

func randrange(l, u int) int {
	return l + rand.Intn(u-l)
}
