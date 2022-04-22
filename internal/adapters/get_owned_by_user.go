package adapters

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"text/template"
	"time"

	"github.com/MalukiMuthusi/mintbase/internal/models"
)

func GetOwnedByUser(user *models.UserIDParameter) ([]byte, error) {

	// create the query
	queryTemplate := `
	{ 
		thing	(
			where: {
				tokens: {
					ownerId: {_eq: " + {{.User}} + "}
				}
			}
		) { 
						id, 
						metadata { 
							title, 
							media 
						}
					}
				}
`

	tmpl, err := template.New("query").Parse(queryTemplate)
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer

	err = tmpl.Execute(&b, user)
	if err != nil {
		return nil, err
	}

	query := b.String()

	jsonData := map[string]string{
		"query": query,
	}

	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", "https://mintbase-mainnet.hasura.app/v1/graphql", bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}

	client := &http.Client{Timeout: time.Second * 100}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
