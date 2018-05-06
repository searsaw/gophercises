package adventure

import (
	"html/template"
	"net/http"
)

func NewHttpAdventure(jsonData []byte, templateFile string) (*http.ServeMux, error) {
	story, err := parseJSONData(jsonData)
	if err != nil {
		return nil, err
	}

	router := http.NewServeMux()
	tmpl := template.Must(template.ParseFiles(templateFile))

	for slug, storyArc := range story {
		func(slug string, storyArc StoryArc) {
			router.HandleFunc("/"+slug, func(w http.ResponseWriter, r *http.Request) {
				tmpl.Execute(w, storyArc)
			})
		}(slug, storyArc)
	}

	return router, nil
}
