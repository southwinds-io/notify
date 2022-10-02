// Notify Service - © 2022 SouthWinds Tech Ltd - www.southwinds.io
// Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
// Contributors to this project, hereby assign copyright in this code to the project,
// to be licensed under the same terms as the rest of the code.

package main

import (
	"fmt"
	"net/http"
	util "southwinds.dev/http"
	"southwinds.dev/notify/core"
	_ "southwinds.dev/notify/docs"
	"strings"
)

// @title Notify Service
// @version 1.0.0
// @description Notification Service
// @contact.name SouthWinds Tech ltd
// @contact.url https://www.southwinds.io/
// @contact.email info@southwinds.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @Summary Sends a new notification
// @Description sends a notification of the specified type
// @Tags Notifications
// @Router /notify [post]
// @Param notification body core.NotificationMsg true "the notification information to send"
// @Accept application/yaml, application/json
// @Produce plain
// @Failure 400 {string} bad request: the server cannot or will not process the request due to something that is perceived to be a client error (e.g., malformed request syntax, invalid request message framing, or deceptive request routing)
// @Failure 500 {string} internal server error: the server encountered an unexpected condition that prevented it from fulfilling the request.
// @Success 201 {string} notification has been sent
func notifyHandler(w http.ResponseWriter, r *http.Request) {
	notification := new(core.NotificationMsg)
	err := util.Unmarshal(r, notification)
	if util.IsErr(w, err, http.StatusBadRequest, "cannot unmarshal notification") {
		return
	}
	if util.IsErr(w, notification.Valid(), http.StatusBadRequest, "invalid notification") {
		return
	}
	switch strings.ToUpper(notification.Type) {
	case "EMAIL":
		err = core.SendMail(*notification)
		if util.IsErr(w, err, http.StatusBadRequest, "cannot email notification") {
			return
		}
	default:
		util.Err(w, http.StatusBadRequest, fmt.Sprintf("notification type '%s' is not supported", strings.ToUpper(notification.Type)))
	}
}
