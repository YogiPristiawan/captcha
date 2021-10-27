package handlers

import (
	"github.com/aldy505/asciitxt"
	"github.com/aldy505/decrr"
	tb "gopkg.in/tucnak/telebot.v2"
)

func (d *Dependencies) Ascii(m *tb.Message) {
	if m.Payload == "" {
		return
	}

	gen := asciitxt.New(m.Payload)

	_, err := d.Bot.Send(m.Chat, "<pre>"+gen+"</pre>", &tb.SendOptions{ParseMode: tb.ModeHTML, AllowWithoutReply: true})
	if err != nil {
		if err.Error() == "telegram: message must be non-empty (400)" {
			_, err := d.Bot.Send(
				m.Chat,
				"That text is not supported yet",
				&tb.SendOptions{
					ParseMode:         tb.ModeHTML,
					AllowWithoutReply: true,
					ReplyTo:           m,
				},
			)
			if err != nil {
				panic(decrr.Wrap(err))
			}
		} else {
			panic(decrr.Wrap(err))
		}
	}
}
