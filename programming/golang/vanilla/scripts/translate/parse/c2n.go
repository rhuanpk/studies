package parse

import "github.com/nicksnyder/go-i18n/v2/i18n"

// C2N parses code to name
func C2N(code string) (string, error) {
	return localizer.Localize(
		&i18n.LocalizeConfig{
			MessageID: code,
		},
	)
	// translated, err := localizer.Localize(
	// 	&i18n.LocalizeConfig{
	// 		MessageID: code,
	// 	},
	// )
	// if err != nil {
	// 	notFound := new(i18n.MessageNotFoundErr)
	// 	if errors.As(err, &notFound) {
	// 		err = fmt.Errorf(
	// 			`%w: "%s"`,
	// 			ErrNotFound, code,
	// 		)
	// 	}
	// 	return "", err
	// }
	// return translated, nil
}
