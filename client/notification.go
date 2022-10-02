// Notify Service - © 2022 SouthWinds Tech Ltd - www.southwinds.io
// Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
// Contributors to this project, hereby assign copyright in this code to the project,
// to be licensed under the same terms as the rest of the code.

package notify

import "fmt"

type NotificationMsg struct {
	// Recipient of the notification if type is email
	Recipient string `yaml:"recipient" json:"recipient" example:"info@email.com"`
	// Type of the notification (e.g. email, snow, etc.)
	Type string `yaml:"type" json:"type" example:"email"`
	// Subject of the notification
	Subject string `yaml:"subject" json:"subject" example:"New Notification"`
	// Content of the template
	Content string `yaml:"content" json:"content" example:"A new event has been received."`
}

func (m NotificationMsg) Valid() error {
	if len(m.Subject) == 0 {
		return fmt.Errorf("subject is required")
	}
	if len(m.Content) == 0 {
		return fmt.Errorf("content is required")
	}
	if len(m.Recipient) == 0 {
		return fmt.Errorf("recipient is required")
	}
	return nil
}
