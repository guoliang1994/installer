package cmd

import (
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func L(key string, lang string) string {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	// No need to load active.en.toml since we are providing default translations.
	// bundle.MustLoadMessageFile("active.en.toml")
	bundle.MustLoadMessageFile("lang\\en.toml")
	bundle.LoadMessageFile("lang\\zh.toml")
	localizer := i18n.NewLocalizer(bundle, lang)
	return localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: key, PluralCount: 0})
}
