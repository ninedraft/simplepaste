package simplepaste

import (
	"net/url"
	"net/http"
	"bytes"
)

//look defines.go

type Paste struct {
	Text       string
	Name       string
	Privacy    string
	ExpireDate string
	//You have to get it by .GetUserKey method, if you want to create private paste
	UserKey string
}

func NewPaste(name string, text string) Paste {
	return Paste{
		Text: text,
		Name: name,
		Privacy: Public,
		ExpireDate: Never,
	}
}

type API struct {
	APIKey  string
}

func NewAPI(api_key string) *API {
	return &API{
		APIKey:  api_key,
	}
}

//Returns paste link string and nil if everything is ok
func (api * API) SendPaste(paste Paste) (string, error) {
	if paste.UserKey == "" && paste.Privacy == "2" {
		return "", PrivacyModError
	}
	values := url.Values{}
	values.Set("api_dev_key", api.APIKey)
	values.Set("api_user_key", paste.UserKey)
	values.Set("api_option", "paste")
	values.Set("api_paste_code", paste.Text)
	values.Set("api_paste_name", paste.Name)
	values.Set("api_paste_private", paste.Privacy)
	values.Set("api_paste_expire_date", paste.ExpireDate)
	response, err := http.PostForm("http://pastebin.com/api/api_post.php", values)
	defer response.Body.Close()
	if err != nil {
		return "", err
	}
	if response.StatusCode != 200 {
		return "", PastePostingError
	}
	buf := bytes.Buffer{}
	_, err = buf.ReadFrom(response.Body)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

//Returns raw paste text
func (api * API) GetPasteTextById(paste_id string) (string, error) {
	response, err := http.Get("http://pastebin.com/raw.php?i=" + paste_id)
	defer response.Body.Close()
	if err != nil {
		return "", err
	}
	if response.StatusCode != 200 {
		return "", PasteGetError
	}
	buf := bytes.Buffer{}
	_, err = buf.ReadFrom(response.Body)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (api * API) GetUserKey(username string, password string)(string, error) {
	values := url.Values{}
	values.Set("api_dev_key", api.APIKey)
	values.Set("api_user_name", username)
	values.Set("api_user_password", password)
    req, err := http.NewRequest("POST", "http://pastebin.com/api/api_login.php", bytes.NewBufferString(values.Encode()))
    client := &http.Client{}
    response, err := client.Do(req)
    if err != nil {
		return "", err
    }
    defer response.Body.Close()
	buf := bytes.Buffer{}
   	_, err = buf.ReadFrom(response.Body)
    return buf.String(), err
}

