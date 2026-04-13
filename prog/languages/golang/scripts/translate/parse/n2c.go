package parse

import "github.com/nicksnyder/go-i18n/v2/i18n"

// N2C parses name to code
func N2C(name string) (string, error) {
	return localizer.Localize(
		&i18n.LocalizeConfig{
			MessageID: name,
		},
	)
	// translated, err := localizer.Localize(
	// 	&i18n.LocalizeConfig{
	// 		MessageID: name,
	// 	},
	// )
	// if err != nil {
	// 	notFound := new(i18n.MessageNotFoundErr)
	// 	if errors.As(err, &notFound) {
	// 		err = fmt.Errorf(
	// 			`%w: "%s"`,
	// 			ErrNotFound, name,
	// 		)
	// 	}
	// 	return "", err
	// }
	// return translated, nil
}
