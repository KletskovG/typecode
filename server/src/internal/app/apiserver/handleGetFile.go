package apiserver

import (
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
	var lines [][]string = *new([][]string)

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

	lines := splitRawByLines(file.Code)
	result := map[string]interface{}{
		"data": lines,
		"file": file.Name,
		"owner": file.Owner,
	}
	jsonValue, _ := json.Marshal(result)
	return jsonValue
}