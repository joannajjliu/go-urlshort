package urlshort

import (
	"fmt"
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("HIPPOS", r.URL.Path)
		resolvedPath, ok := pathsToUrls[r.URL.Path]
		if ok {
			http.Redirect(w, r, resolvedPath, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

type PathObj struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	var p []map[string]string
	if err := yaml.Unmarshal(yml, &p); err != nil {
		return nil, err
	}
	pathsToUrls := make(map[string]string)
	for _, v := range p {
		pathsToUrls[v["path"]] = v["url"]

	}
	return MapHandler(pathsToUrls, fallback), nil

}
