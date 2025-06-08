package credentials_service

import (
	"bytes"
	"context"
	"html/template"

	"github.com/GP-Hacks/auth/internal/models"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

const htmlTemplate = `
<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <title>Подтверждение регистрации</title>
</head>
<body style="font-family: Arial, sans-serif; background-color: #f9f9f9; padding: 20px;">
  <div style="max-width: 600px; margin: auto; background-color: white; padding: 30px; border-radius: 10px; box-shadow: 0 0 10px rgba(0,0,0,0.1);">
    <h2 style="color: #006400;">Добро пожаловать на сервис «Карта жителя Республики Татарстан»!</h2>
    <p>Вы начали процесс регистрации на нашем портале. Для завершения регистрации, пожалуйста, подтвердите ваш адрес электронной почты, нажав на кнопку ниже:</p>
    <div style="text-align: center; margin: 30px 0;">
      <a href="{{.URL}}" style="display: inline-block; padding: 12px 25px; background-color: #006400; color: white; text-decoration: none; border-radius: 5px; font-size: 16px;">Подтвердить почту</a>
    </div>
    <p>Если вы не регистрировались на сервисе «Карта жителя Республики Татарстан», просто проигнорируйте это письмо.</p>
    <hr style="margin-top: 40px;">
    <p style="font-size: 12px; color: #555;">С уважением,<br>Команда сервиса «Карта жителя Республики Татарстан»</p>
  </div>
</body>
</html>
`

func (s *CredentialsService) sendConfirmationEmail(credentials *models.Credentials) {
	token := uuid.New()
	confirmationURL := "https://tatarstan-card.ru/api/auth/confirm/" + token.String()

	tmpl, err := template.New("email").Parse(htmlTemplate)
	if err != nil {
		log.Error().Msg(err.Error())
	}

	var body bytes.Buffer
	err = tmpl.Execute(&body, map[string]string{
		"URL": confirmationURL,
	})
	if err != nil {
		log.Error().Msg(err.Error())
	}

	if err := s.emailTokensRepository.Save(context.Background(), token.String(), credentials.ID); err != nil {
		log.Error().Msg(err.Error())
	}

	m := models.Mail{
		To:     credentials.Email,
		Header: "Подтверждение почты",
		Body:   body.String(),
	}

	if err := s.notificationsAdapter.SendMail(&m); err != nil {
		log.Error().Msg(err.Error())
	}
}
