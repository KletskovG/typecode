package github

import (
	"github.com/kletskovg/typecode/server/src/internal/consts"
	log "github.com/sirupsen/logrus"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"github.com/google/go-github/github"
	"github.com/kletskovg/typecode/server/src/internal/utils"
	"regexp"
)

var (
	MaxFileArraySize = 20
)

type GhFile struct {
	Name string
	Owner string
	Code 	string
}

// GetFile(language string) find random file from certain repository with specified language
func GetFile(language string, languageExtension string) (GhFile, error) {
	log.Info("GETTIG REPO WITH LANGUAGE " + language)
	repo, err := GetRandomRepository(language)

	if err != nil {
		log.Error("GETTING REPO ERROR")
		log.Error(err)
		return GhFile{}, err
	}

	contents := repo.GetContentsURL();
	contentsURL := contents[:len(contents) - 8] // Remove /{+path} from the end of API url
	resp, err := http.Get(contentsURL)

	if err != nil {
		log.Error(err)
		return GhFile{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if resp.Body  != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		log.Error(err)
		return GhFile{}, err
	}
	

	var data []github.RepositoryContent 
	err = json.Unmarshal(body, &data)

	if err != nil {
		log.Error(err)
		return GhFile{}, err
	}
	
	utils.ShuffleRepos(data)
	

	files := createFilesArray(data, languageExtension)

	if files != nil {
		// utils.ShuffleRepos(files)
		indexes := utils.Shuffle(len(files))

		for i := range files {
			index := indexes[i]
			tmp := files[index]
			tmp1 := files[i]
			files[i] = tmp
			files[index] = tmp1
		}

		randIndex := utils.GetRandomElement(len(files))

		
		return files[randIndex], nil
	}

	return GhFile{}, nil
}

func checkFileForExtension (findedFileExtension string, file github.RepositoryContent) bool {
	filename := file.GetName()
	var re = regexp.MustCompile(`(?m)(?:\.([^.]+))?$`) // RegExp which finds file extension
	extension := string(re.Find([]byte(filename)))
	// If length of extension if less then 0 it means no matches
	if len(extension) > 0 {
    return findedFileExtension == extension[1:]
  } 
	
	return false
}

func GetRawFile (url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		log.WithFields(log.Fields{
			"url": url,
			"error": err.Error(),
		}).Error("Error during http request in getRawFile")
		return "", err
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	body, readErr := ioutil.ReadAll(resp.Body)

	if readErr != nil {
		log.WithFields(log.Fields{
			"body": resp.Body,
			"error": readErr.Error(),
		}).Error("Error during reading response body in getRawFile")
		return "", err
	}

	return string(body), nil
}

func createFilesArray (repo []github.RepositoryContent, fileExtension string) []GhFile {
	var files []GhFile
	for i := range repo {
		log.WithFields(log.Fields{
			"type": repo[i].GetType(),
		}).Info("Check type of repo")
		
		if len(files) > 20 {
			break
		}

		match := checkFileForExtension(fileExtension, repo[i])

		

		if match && repo[i].GetSize() > consts.MinCodeSize {
			code,_ := GetRawFile(repo[i].GetDownloadURL())
		
			file := GhFile {
				Name: repo[i].GetName(),
				Owner: "Some",
				Code: code,
			}
			files = append(files, file)
		} else if repo[i].GetSize() == 0 {
			extractedFiles,_ := processDir(repo[i].GetURL(), files, fileExtension)
			
			for i := range  extractedFiles {
				files = append(files, extractedFiles[i])
			}
		}
	}

	return files
}

func processDir (dirUrl string, files []GhFile, fileExtension string) ([]GhFile, error)  {
	// Get Contents of Directory
	resp, err := http.Get(dirUrl)

	if err != nil {
		log.Error(err)
	}

	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)

	if readErr != nil {
		log.Error(err)
	}

	var data []github.RepositoryContent 
	err = json.Unmarshal(body, &data)

	if err != nil {
		log.Error(err)
	}

	for i := range data {
		if data[i].GetSize() > consts.MinCodeSize && len(files) <  MaxFileArraySize {
			match := checkFileForExtension(fileExtension, data[i])
			
			if match {
				code := data[i].GetDownloadURL()
				raw,_ := GetRawFile(code)
				file := GhFile {
					Name: data[i].GetName(),
					Owner: "Some",
					Code: raw,
				}

				files = append(files, file)
			}
		} else if data[i].GetSize() == 0 {
			return processDir(data[i].GetURL(), files, fileExtension)
		}
	}

	return files, nil
}

