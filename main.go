package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	BOT_API_KEY string
	BOT_DEBUG   string
	provided    bool
	httpClient  *http.Client
)

type RevolutQuote struct {
	Sender    AmountCurrency `json:"sender"`
	Recipient AmountCurrency `json:"recipient"`
	Rate      RateInfo       `json:"rate"`
}

type AmountCurrency struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type RateInfo struct {
	Rate float64 `json:"rate"`
}

func init() {
	BOT_API_KEY, provided = os.LookupEnv("BOT_API_KEY")
	if !provided {
		log.Print("BOT_API_KEY is not set")
		os.Exit(128)
	}
	BOT_DEBUG, provided = os.LookupEnv("BOT_DEBUG")
	if !provided {
		log.Print("BOT_DEBUG is not set, default is false")
	}
	httpClient = &http.Client{}
}

func main() {
	bot, err := tgbotapi.NewBotAPI(BOT_API_KEY)
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	if BOT_DEBUG != "" {
		bot.Debug = true
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
				msg.Text = "`0.88EUR/USD or USD/EUR`"
			case "status":
				msg.Text = "I'm ok."
			default:
				msg.Text = ""
			}

			if msg.Text != "" {
				if _, err := bot.Send(msg); err != nil {
					log.Printf("Error sending command reply: %v", err)
				}
			}
		}

		params, matched := parseMessage(update.Message.Text)
		if matched {
			if resMsg := getFXRate(params); resMsg != "" {
				log.Printf("%v", params)

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, resMsg)
				msg.ParseMode = "markdown"
				msg.ReplyToMessageID = update.Message.MessageID

				if _, err := bot.Send(msg); err != nil {
					log.Printf("Error sending message: %v", err)
				}
			} else {
				log.Print("Error, no message")
			}
		}
	}
}

func parseMessage(txtMsg string) (params []string, matched bool) {
	re := regexp.MustCompile(`^([0-9]*[.]?[0-9]+)?[ ]?(\w{3})[ ]?/?[ ]?(\w{3})$`)
	matches := re.FindStringSubmatch(txtMsg)
	if len(matches) > 3 {
		params = matches[1:]
		if params[0] == "" {
			params[0] = "1"
		}
		return params, true
	}
	return []string{}, false
}

func getFXRate(params []string) (recv string) {
	recv = ""

	url := fmt.Sprintf("https://www.revolut.com/api/exchange/quote?amount=%s&country=DE&fromCurrency=%s&isRecipientAmount=false&toCurrency=%s",
		params[0],
		strings.ToUpper(params[1]),
		strings.ToUpper(params[2]))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Print(err)
		return
	}
	req.Header.Add("accept-language", "en")

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Print(err)
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Print(err)
			return
		}

		var quote RevolutQuote
		if err := json.Unmarshal(bodyBytes, &quote); err != nil {
			log.Printf("Error unmarshalling Revolut response: %v", err)
			return
		}

		senderAmount, err := strconv.ParseFloat(params[0], 64)
		if err != nil {
			log.Printf("Error parsing amount '%s': %v", params[0], err)
			return
		}

		p := message.NewPrinter(language.English)

		if quote.Rate.Rate > 0 {
			recipientAmount := senderAmount * quote.Rate.Rate

			recv = p.Sprintf("\xF0\x9F\x92\xB8 `%.2f %s`   \xF0\x9F\x94\x84   `%.2f %s`\n\xF0\x9F\x92\xB1 Rate: `%.6f`",
				senderAmount,
				quote.Sender.Currency,
				recipientAmount,
				quote.Recipient.Currency,
				quote.Rate.Rate)
		}

	} else {
		recv = fmt.Sprintf("%d - %s", resp.StatusCode, resp.Status)
	}

	return recv
}
