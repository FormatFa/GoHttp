package writer

import (
	"gohttp/listener"
	"os"
)

// TODO 写出http对象
func Write_https(https listener.Https, path string) {
	result := ""
	for _, http := range https {
		result = result + "# " + http.Name + " " + http.Url
		result += "\n"
		for _, head := range http.Headers {
			result += head.Key + ":" + head.Value
			result += "\n"
		}
		result += "\n"
	}
	os.WriteFile(path, []byte(result), 0644)
}
