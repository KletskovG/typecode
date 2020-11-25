package apiserver

import (
	"bytes"
	"github.com/kletskovg/typecode/server/src/internal/consts"
	"io/ioutil"
	"net/http"
	"github.com/kletskovg/typecode/server/src/internal/github"
	"strings"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/kletskovg/typecode/server/src/internal/utils"
)

func (server *APIServer) HandleGetFile () http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		utils.EnableCors(&w)
		params := mux.Vars(r)
		language := params["language"]
		extension := params["extension"]

		log.WithFields(log.Fields{
			"language": language,
		}).Warn("SEE QUERY PARAM")

		file := processFile(language, extension)
		w.Write(file)
	}
}

func splitRawByLines (raw string) [][]string {
	linesSplit := strings.Split(raw, "\n")
	lines := *new([][]string)

	for i := range linesSplit {
    lineSymbols := strings.Split(linesSplit[i], "")
		lines = append(lines, lineSymbols)
	}

  return lines
}

func processFile(language string, extension string) []byte {
	var file = github.GhFile{}
	
	ghFile, getFileErr := github.GetFile(language, extension)
	
	if getFileErr != nil {
		//
		//log.Info("GIHUB FILE IN PROCESS FILE")
		//log.Info(cachedFile)
		//
		//htmlStr, _ := cachedFile["htmlUrl"].(string)
		//raw, _ := github.GetRawFile(string(htmlStr))
		//file.Code = raw
		//
		//if name, ok := cachedFile["language"].(string); ok {
		//	file.Name = name
		//}
		//
		//if owner, ok := cachedFile["rawUrl"].(string); ok {
		//	file.Owner = owner
		//}
		
		log.Info(file)
	} else {
		file.Code = ghFile.Code
	}

	log.Info("LOOK AT THE FILE")
	log.Info(ghFile)
	cacheJson,_ := json.Marshal(ghFile)
	requestUrl := []string{consts.DB_SERVICE_URL, "/insert"}
	res, err := http.Post(strings.Join(requestUrl, ""), "application/json", bytes.NewBuffer(cacheJson))

	if err != nil {
		log.Error(err)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, _ := ioutil.ReadAll(res.Body)
	log.Info("CHECK DB RESPONSE")
	log.Info(string(body))
	lines := splitRawByLines(file.Code)
	result := map[string]interface{}{
		"data": lines,
		"file": file.Name,
		"owner": file.Owner,
	}
	jsonValue, _ := json.Marshal(result)
	return jsonValue
}