package messaging

import (
	"drylang/core"
	"net/smtp"
	"strconv"
)

// BuiltinMail handles mail("send", host, port, user, pass, to, subject, body)
func BuiltinMail(v core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	if len(args) < 1 || args[0].Type != core.ValString {
		return core.UnknownValue, v.Errorf("want mail(method, ...args)")
	}
	method := args[0].Data.(string)

	switch method {
	case "send":
		if len(args) != 8 {
			return core.UnknownValue, v.Errorf("mail.send wants (host, port, user, pass, to, subject, body)")
		}
		
		for i := 1; i <= 7; i++ {
			if i == 2 {
				if args[i].Type != core.ValNumber {
					return core.UnknownValue, v.Errorf("mail.send port must be number")
				}
			} else {
				if args[i].Type != core.ValString {
					return core.UnknownValue, v.Errorf("mail.send args %d must be string", i)
				}
			}
		}

		host := args[1].Data.(string)
		port := int(args[2].Data.(float64))
		user := args[3].Data.(string)
		pass := args[4].Data.(string)
		to := args[5].Data.(string)
		subject := args[6].Data.(string)
		body := args[7].Data.(string)

		addr := host + ":" + strconv.Itoa(port)
		auth := smtp.PlainAuth("", user, pass, host)

		msg := []byte("To: " + to + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"\r\n" +
			body + "\r\n")

		err := smtp.SendMail(addr, auth, user, []string{to}, msg)
		if err != nil {
			return core.UnknownValue, v.Errorf("mail send error: %s", err)
		}

		return core.Value{Type: core.ValBool, Data: true}, nil

	default:
		return core.UnknownValue, v.Errorf("unknown mail method: %s", method)
	}
}
