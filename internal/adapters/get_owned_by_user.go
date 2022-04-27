package adapters

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"text/template"
	"time"

	"github.com/MalukiMuthusi/mintbase/internal/models"
	"github.com/MalukiMuthusi/mintbase/logger"
)

func GetOwnedByUser(user models.UserIDParameter) (*map[string]interface{}, error) {

	// create the query
	queryTemplate := `
	query MyQuery {
		thing(where: {tokens: {ownerId: {_eq: "{{.User}}" }}}) {
		  id
		  metadata {
			title
			media
			description
		  }
		}
	  }
`

	tmpl, err := template.New("queryTemplate").Parse(queryTemplate)
	if err != nil {
		logger.Log.Info(err)
		return nil, err
	}

	var b bytes.Buffer

	err = tmpl.Execute(&b, user)
	if err != nil {
		logger.Log.Info(err)
		return nil, err
	}

	query := b.String()

	jsonData := map[string]string{
		"query": query,
	}

	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		logger.Log.Info(err)
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, "https://mintbase-mainnet.hasura.app/v1/graphql", bytes.NewBuffer(jsonValue))
	if err != nil {
		logger.Log.Info(err)
		return nil, err
	}

	client := &http.Client{Timeout: time.Second * 100}

	response, err := client.Do(request)
	if err != nil {
		logger.Log.Info(err)
		return nil, err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.Log.Info(err)
		return nil, err
	}

	// var res models.Resp

	var res2 map[string]interface{}

	err = json.Unmarshal(data, &res2)
	if err != nil {
		logger.Log.Info(err)
		return nil, err
	}

	return &res2, nil
}
