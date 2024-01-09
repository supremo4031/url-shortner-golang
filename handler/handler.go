package handler

import (
	"net/http"
	"gopkg.in/yaml.v2"
)



func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusTemporaryRedirect) // status code: 307 temporary redirect
			return
		}
		fallback.ServeHTTP(w, r)
	}
}


func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// 1. Parse the yaml
	pathUrls, err := parseYaml(yamlBytes)
	if err != nil {
		return nil, err
	}

	// 2. Convert yaml array into map
	pathToUrls := buildMap(pathUrls)

	// 3. Return a map handler using the map
	return MapHandler(pathToUrls, fallback), nil
}

func parseYaml(data []byte) ([]pathUrl, error) {
	var pathUrls []pathUrl
	err := yaml.Unmarshal(data, &pathUrls)
	if err != nil {
		return nil, err
	}
	return pathUrls, err
}

func buildMap(pathUrls []pathUrl) map[string]string {
	pathToUrls := make(map[string]string)
	for _, pu := range pathUrls {
		pathToUrls[pu.Path] = pu.URL
	}
	return pathToUrls
}

type pathUrl struct {
	Path string `yaml:"path"`
	URL string `yaml:"url"`
}


