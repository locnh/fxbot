package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	BOT_API_KEY string
	provided    bool
)

type masterFX struct {
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Date        string       `json:"date"`
	Data        masterFXData `json:"data"`
}

type masterFXData struct {
	ConversionRate float64 `json:"conversionRate"`
	CrdhldBillAmt  float64 `json:"crdhldBillAmt"`
	TransAmt       float64 `json:"transAmt"`
}

func init() {
	BOT_API_KEY, provided = os.LookupEnv("BOT_API_KEY")
	if !provided {
		log.Fatal("BOT_API_KEY is not set")
		os.Exit(128)
	}
}

func main() {
	bot, err := tgbotapi.NewBotAPI(BOT_API_KEY)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			msg.ReplyToMessageID = update.Message.MessageID
			msg.ParseMode = "markdown"

			switch update.Message.Command() {
			case "help", "start":
				msg.Text = "`1000 USD in EUR`"
			case "status":
				msg.Text = "I'm ok."
			default:
				msg.Text = "I don't know that command"
			}

			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}

		params, matched := parseMessage(update.Message.Text)
		if matched {
			resMsg := getFXRate(params)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, resMsg)
			msg.ParseMode = "markdown"
			msg.ReplyToMessageID = update.Message.MessageID

			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}
	}
}

func parseMessage(txtMsg string) (params []string, matched bool) {
	re := regexp.MustCompile(`(\d+).*(\w{3}).*(\w{3})`)
	matches := re.FindStringSubmatch(txtMsg)
	if len(matches) > 3 {
		params = matches[1:]
		return params, true
	}
	return []string{}, false
}

func getFXRate(params []string) (recv string) {
	recv = "Unknown error"

	url := "https://www.mastercard.us/settlement/currencyrate/fxDate=0000-00-00;transCurr=" + strings.ToUpper(params[1]) + ";crdhldBillCurr=" + strings.ToUpper(params[2]) + ";bankFee=0;transAmt=" + params[0] + "/conversion-rate"

	resp, err := http.Get(url)
	if err != nil {
		log.Print(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)

		var ex masterFX
		_ = json.Unmarshal([]byte(bodyString), &ex)

		p := message.NewPrinter(language.English)

		recv = p.Sprintf("\xF0\x9F\x92\xB8 `%.2f %s`   \xF0\x9F\x94\x84   `%.2f %s`\n\xF0\x9F\x92\xB1 Rate: `%.2f`",
			ex.Data.TransAmt,
			strings.ToUpper(params[1]),
			ex.Data.CrdhldBillAmt,
			strings.ToUpper(params[2]),
			ex.Data.ConversionRate)

	} else {
		recv = fmt.Sprintf("%d - %s", resp.StatusCode, resp.Status)
	}

	return recv
}
