package file

import "io/ioutil"
import "fmt"
import "jaytaylor.com/html2text"

// HTML2Markdown converts HTML to Markdown-formatted text.

func HTML2Markdown(content []byte, HTMLFile string) ([]byte, error) {
        HTMLRaw, err := ioutil.ReadFile(HTMLFile)
        output, err := html2text.FromString(string(HTMLRaw), html2text.Options{PrettyTables: true})
        if err != nil {
                return nil, fmt.Errorf("HTML -> Markdown failed: %v", err)
        }
        return []byte(output), nil
}
