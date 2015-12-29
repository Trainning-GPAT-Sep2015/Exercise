package loadMarkdown

import (
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"path/filepath"
)

func FromFile(sourcePath string) ([]byte, error) {
	asbPath, err := filepath.Abs(sourcePath)
	if err != nil {
		return []byte{}, err
	}

	data, err := ioutil.ReadFile(asbPath)
	if err != nil {
		return []byte{}, err
	}

	unsafe := blackfriday.MarkdownCommon(data)
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	return html, nil
}
