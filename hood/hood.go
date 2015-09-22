package hood
import (
	"io/ioutil"
	"encoding/json"
	"os"
	"os/user"
	"time"
	"strconv"
	"text/template"
)

type Config struct {
	PhpExtension       string `json:"PHP_EXTENSION"`
	GitIgnoreExtension string `json:"GIT_IGNORE_EXTENSION"`
	GoExtension        string `json:"GO_EXTENSION"`
	JsExtension        string `json:"JS_EXTENSION"`
	CssExtension       string `json:"CSS_EXTENSION"`
	JavaExtension      string `json:"JAVA_EXTENSION"`
	PythonExtension    string `json:"PYTHON_EXTENSION"`
	HtmlExtension      string `json:"HTML_EXTENSION"`
	GitIgnoredList     []string `json:"GIT_IGNORED_LIST"`
	User               string `json:"USER"`
	Date               string `json:"DATE"`
}

var config Config
func init() {

	defaultConfig, _ := ioutil.ReadFile("./hood.json")
	json.Unmarshal(defaultConfig, &config)

	// user
	u, _ := user.Current()
	config.User = u.Username
	// date
	year, month, day := time.Now().Date()
	m := int(month)
	config.Date = strconv.Itoa(year) +"/"+ strconv.Itoa(m) +"/"+ strconv.Itoa(day)

	// merge config
	customConfig, err := ioutil.ReadFile(u.HomeDir + "/.hood.json")
	if err == nil {

		var ccfg Config
		json.Unmarshal(customConfig, &ccfg)
		if ccfg.GitIgnoredList != nil {
			config.GitIgnoredList = ccfg.GitIgnoredList
		}
		if ccfg.User != "" {
			config.User = ccfg.User
		}
	}

}

func GetConfig() (Config) {

	return config
}

func GenerateTemplate(t string, name string) (ok bool) {

	ok = true
	templ := template.New("hood template")
	type Meta struct {
		User string
		Date string
	}
	meta := Meta{
		config.User,
		config.Date,
	}
	switch t {
	case "php":
		content, _ := ioutil.ReadFile("./templates/.php.tpl")
		tpl, _ := templ.Parse(string(content))
		f, _ := os.Create(name + config.PhpExtension)
		tpl.Execute(f, meta)
	case "ignore":
		content, _ := ioutil.ReadFile("./templates/.gitignore.tpl")
		tpl, _ := templ.Parse(string(content))
		type Ignore struct {
			IgnoredList []string
		}
		ignore := Ignore{
			config.GitIgnoredList,
		}
		f, _ := os.Create(""+config.GitIgnoreExtension)
		tpl.Execute(f, ignore)
	case "go":
		content, _ := ioutil.ReadFile("./templates/.go.tpl")
		tpl, _ := templ.Parse(string(content))
		type Go struct {
			Package string
		}
		golang := Go{
			name,
		}
		f, _ := os.Create(name + config.GoExtension)
		tpl.Execute(f, golang)
	case "js":
		content, _ := ioutil.ReadFile("./templates/.js.tpl")
		tpl, _ := templ.Parse(string(content))
		f, _ := os.Create(name + config.JsExtension)
		tpl.Execute(f, meta)
	case "css":
		content, _ := ioutil.ReadFile("./templates/.css.tpl")
		tpl, _ := templ.Parse(string(content))
		f, _ := os.Create(name + config.CssExtension)
		tpl.Execute(f, meta)
	case "java":
		content, _ := ioutil.ReadFile("./templates/.java.tpl")
		tpl, _ := templ.Parse(string(content))
		f, _ := os.Create(name + config.JavaExtension)
		tpl.Execute(f, meta)
	case "python":
		content, _ := ioutil.ReadFile("./templates/.py.tpl")
		tpl, _ := templ.Parse(string(content))
		f, _ := os.Create(name + config.PythonExtension)
		tpl.Execute(f, meta)
	case "html":
		content, _ := ioutil.ReadFile("./templates/.html.tpl")
		tpl, _ := templ.Parse(string(content))
		type Html struct {
			Title string
		}
		html := Html{
			name,
		}
		f, _ := os.Create(name + config.HtmlExtension)
		tpl.Execute(f, html)
	default :
		ok = false
	}
	return
}