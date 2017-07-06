// Inform to gaurun server.
package notification

import(
	"net/http"
	"log"
	"encoding/json"
)

func InformClients(clients []Client, msg string) {

				token_list := []string
				for _, cli_info := clients {
								token_list = append(token_list, cli_info.Token)
				}
				json_str := `
					{"notifications":
						"token":[` + token_list
				json_map := map[string]interface{
								"notifications": token_list,
								"platform": 1,
								"message": msg,
								"badge": 1,
								"sound": default,
								"content_available": false,
								"expiry": 10
				}
				if json_str, err := json.Marshal(json_str); err == nil {
								json_byte := []byte( json_str )
								http.Post("127.0.0.1:1056/push", "application/json", &json_byte)
				}
}
