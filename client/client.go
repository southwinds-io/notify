// Notify Service - © 2022 SouthWinds Tech Ltd - www.southwinds.io
// Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
// Contributors to this project, hereby assign copyright in this code to the project,
// to be licensed under the same terms as the rest of the code.

package notify

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	*http.Client
	host, token string
}

func New(host, user, pwd string) Client {
	c := Client{
		host:  host,
		token: basicToken(user, pwd),
	}
	c.Timeout = 60 * time.Second
	return c
}

func (c *Client) Notify(msg NotificationMsg) error {
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	request, err := http.NewRequest(http.MethodPut, c.url("/notify"), bytes.NewReader(msgBytes))
	if err != nil {
		return err
	}
	request.Header.Set("Authorization", c.token)
	request.Header.Set("User-Agent", fmt.Sprintf("SW-NOTIFY-CLIENT-%s", Version))
	resp, reqErr := c.Do(request)
	if reqErr != nil {
		return reqErr
	}
	if resp.StatusCode > 299 {
		return fmt.Errorf("cannot set type, notify server responded with: %s", resp.Status)
	}
	return nil
}

func (c *Client) url(format string, args ...any) string {
	v := fmt.Sprintf("%s%s", c.host, fmt.Sprintf(format, args...))
	return v
}

func basicToken(user string, pwd string) string {
	return fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", user, pwd))))
}
