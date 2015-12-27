package simplepaste

import(
	"errors"
)

var(
	PastePostingError = errors.New("Error while posting: pastebin.com is not available")
	PasteGetError = errors.New("Error while getting: pastebin.com is not available")
	PrivacyModError = errors.New("You can't send private paste without user_key!")
)

const (
	Never       =  "N"
	TenMinutues =  "10M"
	Hour        =  "1H"
	Day         =  "1D"
	Week        =  "1W"
	TwoWeeks    =  "2W"
	Month       =  "1M"

	Public   = "0"
	Unlisted = "1"
	//Only allowed in combination with api_user_key, as you have to be logged into your account to access the paste
	Private = "2"
)
