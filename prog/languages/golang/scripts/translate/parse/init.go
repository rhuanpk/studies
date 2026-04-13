package parse

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func init() {
	const folder = "langs"
	var files = []string{
		folder + "/c2n.pt.toml",
		folder + "/n2c.pt.toml",
	}

	bundle := i18n.NewBundle(language.Portuguese)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	for _, file := range files {
		if _, err := bundle.LoadMessageFile(file); err != nil {
			log.Println("error in load message file:", err)
			return
		}
	}

	localizer = i18n.NewLocalizer(bundle)
}
