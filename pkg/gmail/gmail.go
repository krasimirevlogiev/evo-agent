package gmail

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"

	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

// CreateService creates a new Gmail service client.
func CreateService(ctx context.Context, client *http.Client) (*gmail.Service, error) {
	srv, err := gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve Gmail client: %v", err)
	}
	return srv, nil
}

// SendEmail sends an email using the Gmail API.
func SendEmail(srv *gmail.Service, to string, subject string, body string) error {
	message := gmail.Message{
		Raw: base64.URLEncoding.EncodeToString([]byte(
			"To: " + to + "\r\n" +
				"Subject: " + subject + "\r\n" +
				"Content-Type: text/html; charset=utf-8\r\n" +
				"\r\n" +
				body,
		)),
	}

	_, err := srv.Users.Messages.Send("me", &message).Do()
	if err != nil {
		return fmt.Errorf("could not send email: %v", err)
	}
	return nil
}
