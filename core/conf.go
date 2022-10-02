// Notify Service - © 2022 SouthWinds Tech Ltd - www.southwinds.io
// Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
// Contributors to this project, hereby assign copyright in this code to the project,
// to be licensed under the same terms as the rest of the code.

package core

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

func init() {
	// load env vars from file if present
	godotenv.Load("proxy.env")
}

func IsLoggingEnabled() bool {
	value := os.Getenv("NOTIFY_LOGGING")
	return len(value) > 0
}

func getEmailFrom() (string, error) {
	value := os.Getenv("NOTIFY_EMAIL_FROM")
	if len(value) == 0 {
		return "", fmt.Errorf("variable NOTIFY_EMAIL_FROM is required and not defined")
	}
	return value, nil
}

func getSmtpUser() (string, error) {
	value := os.Getenv("NOTIFY_SMTP_USER")
	if len(value) == 0 {
		fmt.Printf("WARNING: NOTIFY_SMTP_USER not defined, using DPROXY_EMAIL_FROM instead\n")
		return getEmailFrom()
	}
	return value, nil
}

func getSmtpPwd() (string, error) {
	value := os.Getenv("NOTIFY_SMTP_PWD")
	if len(value) == 0 {
		return "", fmt.Errorf("variable NOTIFY_SMTP_PWD is required and not defined")
	}
	return value, nil
}

func getSmtpHost() (string, error) {
	value := os.Getenv("NOTIFY_SMTP_HOST")
	if len(value) == 0 {
		return "", fmt.Errorf("variable NOTIFY_SMTP_HOST is required and not defined")
	}
	return value, nil
}

func getSmtpPort() (int, error) {
	value := os.Getenv("NOTIFY_SMTP_PORT")
	if len(value) == 0 {
		return -1, fmt.Errorf("variable NOTIFY_SMTP_PORT is required and not defined")
	}
	port, err := strconv.Atoi(value)
	if err != nil {
		return -1, fmt.Errorf("NOTIFY_SMTP_PORT value should be numeric but its value was found to be '%s'; %s", value, err)
	}
	return port, nil
}
