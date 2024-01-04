package mem

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Database struct {
	Subjects  []string
	Resources []Resource
}

var Global Database

func ImportFiles() error {
	err := filepath.WalkDir("./files", func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && filepath.Ext(path) == ".json" {
			fileData, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			var resource Resource
			if err := json.Unmarshal(fileData, &resource); err != nil {
				return err
			}

			if !FindStringInArray(Global.Subjects, resource.Subject) {
				Global.Subjects = append(Global.Subjects, resource.Subject)
			}

			Global.Resources = append(Global.Resources, resource)
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
