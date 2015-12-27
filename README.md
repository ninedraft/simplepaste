# Golang bindings for the [Pastebin.com](pastebin.com) API

Not all methods were implemented. but they are coming soon!

All methods are fairly self explanatory, and reading the godoc page should explain everything. If something isn't clear, open an issue or submit a pull request.

The scope of this project is just to provide a wrapper around the API without any additional features.

## Examples

```go 
package main

import(
	"fmt"
	"github.com/ninedraft/simplepaste"
)

func main(){
	api = simplepaste.NewAPI("Your API key")
	
	paste := simplepaste.NewPaste("Paste name", "Paste text")
	paste.ExpireDate = simplepaste.Month //your paste will be available for one month
	link, err := api.SendPaste(paste) //returns link to the paste and nil, if everything is ok
	if err != nil {
		panic(err)
	}
	fmt.Println(link)	
}

```


