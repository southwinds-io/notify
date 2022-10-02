# Notify Service

Sends notifications using various communication technologies.

__NOTE__: currently email only is implemented

## Configuration

| variable            | description                                | example                                      |
|---------------------|--------------------------------------------|----------------------------------------------|
| `NOTIFY_LOGGING`    | enable http server request logging         | empty for not logging, any value for logging |
| `NOTIFY_EMAIL_FROM` | the sender of the email                    | admin@example.com                            |
| `NOTIFY_SMTP_USER`  | the username to connect to smtp the server | joe                                          |
| `NOTIFY_SMTP_PWD`   | the password to connect to smtp the server | !SWh47xGhR                                   |
| `NOTIFY_SMTP_HOST`  | the smtp hostname                          | example.smtp.com                             |
| `NOTIFY_SMTP_PORT`  | the smtp port                              | 587                                          |