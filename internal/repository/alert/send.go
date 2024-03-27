package alert

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	alerrRepository "github.com/rudianto-dev/gotemp-api-service/internal/domain/alert/repository"
)

func (s *AlertRepository) Send(ctx context.Context, alertEntity *alerrRepository.AlertEntity) (err error) {
	client := &http.Client{
		Timeout: 30 * time.Minute,
	}

	message := fmt.Sprintf("%s receiver: %s, sender : %s.", alertEntity.Message, alertEntity.Receiver, alertEntity.Sender)
	url := fmt.Sprintf("%s/bot%s/sendMessage?chat_id=%s&text=%s", s.telegram.URL, s.telegram.Token, s.telegram.ChannelID, message)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t, _ := ioutil.ReadAll(resp.Body)
		err = errors.New(string(t))
		return
	}
	return
}
