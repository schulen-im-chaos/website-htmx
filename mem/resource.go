package mem

import "strings"

type Resource struct {
	Title    string `json:"title"`
	Summary  string `json:"summary"`
	Comment  string `json:"comment"`
	Subject  string `json:"subject"`
	Grade    string `json:"grade"`
	FileName string `json:"file_name"`
	Author   string `json:"author"`
}

func ResourcesBySubjectAndGradeMap(subjectName string, query string) map[string][]Resource {
	query = strings.ToLower(query)

	filteredResources := make(map[string][]Resource)

	for _, resource := range Global.Resources {
		if resource.Subject == subjectName {
			if query == "" {
				filteredResources[resource.Grade] = append(filteredResources[resource.Grade], resource)
			} else {
				if strings.Contains(strings.ToLower(resource.Grade), query) ||
					strings.Contains(strings.ToLower(resource.Title), query) ||
					strings.Contains(strings.ToLower(resource.Author), query) ||
					strings.Contains(strings.ToLower(resource.Summary), query) {
					filteredResources[resource.Grade] = append(filteredResources[resource.Grade], resource)
				}
			}
		}
	}

	return filteredResources
}
