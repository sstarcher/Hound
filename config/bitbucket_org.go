package config

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GenerateBitbucket(name string, token string, exclude string) map[string]*Repo {
	repos := make(map[string]*Repo)

	client := http.Client{}
	req, err := http.NewRequest("GET", "https://api.bitbucket.org/1.0/users/"+name, nil)
	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(name+":"+token)))

	resp, err := client.Do(req)
	check(err)

	data, err := ioutil.ReadAll(resp.Body)
	check(err)
	resp.Body.Close()

	var jsonData map[string]interface{}
	err = json.Unmarshal(data, &jsonData)
	check(err)

	for _, element := range jsonData["repositories"].([]interface{}) {
		repoName := element.(map[string]interface{})["name"].(string)
		url := fmt.Sprintf("git@bitbucket.org:%s/%s.git", name, repoName)
		repos[repoName] = &Repo{
			Url: url,
		}

	}

	return repos
}
