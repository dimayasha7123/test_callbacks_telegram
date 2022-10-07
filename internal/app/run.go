package app

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	url2 "net/url"
	"time"

	"github.com/dimayasha7123/test_callbacks_telegram/internal/models"
)

func (bs *bserver) Run(ctx context.Context) error {
	lastUpdateID := int64(0)

	// update router
	for {
		url := fmt.Sprintf(
			"https://api.telegram.org/bot%s/getUpdates?offset=%d",
			bs.apiKey,
			lastUpdateID+1,
		)

		resp, err := bs.httpClient.Get(url)
		if err != nil {
			return err
		}

		bytes, err := ioutil.ReadAll(resp.Body)

		updates := models.Updates{}
		err = json.Unmarshal(bytes, &updates)
		if err != nil {
			return err
		}

		if updates.Ok {
			for _, update := range updates.Result {

				postUrl, err := bs.updateRouter(ctx, update)

				// TODO do special handler
				if err != nil {
					postUrl = fmt.Sprintf(
						"https://api.telegram.org/bot%s/sendMessage?chat_id=%d&text=%s",
						bs.apiKey,
						update.Message.Chat.ID,
						url2.PathEscape(fmt.Sprintf("Ooops, something was wrong.\nError: %v", err)),
					)
				}

				_, err = bs.httpClient.Post(postUrl, "text/plain", nil)
				if err != nil {
					return err
				}

			}
		}

		if len(updates.Result) != 0 {
			lastUpdateID = updates.Result[len(updates.Result)-1].UpdateID
		}

		err = resp.Body.Close()
		if err != nil {
			return err
		}

		time.Sleep(50 * time.Millisecond)
	}
}
