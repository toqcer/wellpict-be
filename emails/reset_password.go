package emails

import (
	"fmt"
	"wellpict-be/utils"
)

func GenerateResetPasswordEmail(link string) utils.MailMessage {
	msg := utils.NewMailMessage()

	msg.SetSubject("Reset Password")
	msg.SetBody(fmt.Sprintf("<p><a href=\"%s\">click here</a> to reset your password</p>", link))

	return *msg
}
