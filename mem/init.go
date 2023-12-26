package mem

import (
	"encoding/json"
	"os"
	"path/filepath"
)

var Subjects []string
var Resources []Resource

func ImportFiles() error {
	err := filepath.Walk("./files", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".json" {
			fileData, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			var resource Resource
			if err := json.Unmarshal(fileData, &resource); err != nil {
				return err
			}

			if !FindStringInArray(Subjects, resource.Subject) {
				Subjects = append(Subjects, resource.Subject)
			}

			Resources = append(Resources, resource)
		}
		return nil
	})

	return err
}

func FindStringInArray(arr []string, target string) bool {
	for _, str := range arr {
		if str == target {
			return true
		}
	}
	return false
}
