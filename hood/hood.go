package hood

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
	"strconv"
	"text/template"
	"time"
)

type Config struct {
	PhpExtension       string   `json:"PHP_EXTENSION"`
	GitIgnoreExtension string   `json:"GIT_IGNORE_EXTENSION"`
	GoExtension        string   `json:"GO_EXTENSION"`
	JsExtension        string   `json:"JS_EXTENSION"`
	CssExtension       string   `json:"CSS_EXTENSION"`
	JavaExtension      string   `json:"JAVA_EXTENSION"`
	PythonExtension    string   `json:"PYTHON_EXTENSION"`
	HtmlExtension      string   `json:"HTML_EXTENSION"`
	CExtension         string   `json:"C_EXTENSION"`
	GitIgnoredList     []string `json:"GIT_IGNORED_LIST"`
	User               string   `json:"USER"`
	Date               string   `json:"DATE"`
}

var config Config

const (
	dirMode = 0755
)

func init() {

	defaultConfig, _ := Asset("../config/hood.json")
	json.Unmarshal(defaultConfig, &config)

	// user
	u, _ := user.Current()
	config.User = u.Username
	// date
	year, month, day := time.Now().Date()
	m := int(month)
	config.Date = strconv.Itoa(year) + "/" + strconv.Itoa(m) + "/" + strconv.Itoa(day)

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

func GetConfig() Config {

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
	// For single file
	case "php":
		content, _ := Asset("../templates/.php.tpl")
		tpl, _ := templ.Parse(string(content))
		f, _ := os.Create(name + config.PhpExtension)
		tpl.Execute(f, meta)
	case "ignore":
		content, _ := Asset("../templates/.gitignore.tpl")
		tpl, _ := templ.Parse(string(content))
		type Ignore struct {
			IgnoredList []string
		}
		ignore := Ignore{
			config.GitIgnoredList,
		}
		f, _ := os.Create("" + config.GitIgnoreExtension)
		tpl.Execute(f, ignore)
	case "go":
		content, _ := Asset("../templates/.go.tpl")
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
		content, _ := Asset("../templates/.js.tpl")
		tpl, _ := templ.Parse(string(content))
		f, _ := os.Create(name + config.JsExtension)
		tpl.Execute(f, meta)
	case "css":
		content, _ := Asset("../templates/.css.tpl")
		tpl, _ := templ.Parse(string(content))
		f, _ := os.Create(name + config.CssExtension)
		tpl.Execute(f, meta)
	case "java":
		content, _ := Asset("../templates/.java.tpl")
		tpl, _ := templ.Parse(string(content))
		f, _ := os.Create(name + config.JavaExtension)
		tpl.Execute(f, meta)
	case "python":
		content, _ := Asset("../templates/.py.tpl")
		tpl, _ := templ.Parse(string(content))
		f, _ := os.Create(name + config.PythonExtension)
		tpl.Execute(f, meta)
	case "html":
		content, _ := Asset("../templates/.html.tpl")
		tpl, _ := templ.Parse(string(content))
		type Html struct {
			Title string
		}
		html := Html{
			name,
		}
		f, _ := os.Create(name + config.HtmlExtension)
		tpl.Execute(f, html)
	case "c":
		content, _ := Asset("../templates/.c.tpl")
		tpl, _ := templ.Parse(string(content))
		f, _ := os.Create(name + config.CExtension)
		tpl.Execute(f, meta)
		// For project
	case "ansible":
		/*
				  - cfg
				  - inventory,requirements.yml
			      - files,templates,roles,group_vars,host_vars
				  - setup,build,deploy
				  - setup.yml,build.yml,deploy.yml
		*/
		// ansible.cfg
		ansibleCfgContent, _ := Asset("../templates/ansible/ansible.cfg")
		ansibleCfgFile, _ := os.Create("ansible.cfg")
		ansibleCfgFile.Write(ansibleCfgContent)

		// inventory
		os.Create("inventory")
		// requirements.yml
		os.Create("requirements.yml")

		// files,templates,roles,group_vars,host_vars
		os.MkdirAll("./files", dirMode)
		os.MkdirAll("./templates", dirMode)
		os.MkdirAll("./roles", dirMode)
		os.MkdirAll("./group_vars/all", dirMode)
		os.MkdirAll("./host_vars", dirMode)

		// setup,build,deploy
		os.MkdirAll("./setup", dirMode)
		os.MkdirAll("./build", dirMode)
		os.MkdirAll("./deploy", dirMode)
		// soft link
		os.Symlink("../files", "./setup/files")
		os.Symlink("../templates", "./setup/templates")
		os.Symlink("../roles", "./setup/roles")

		os.Symlink("../files", "./build/files")
		os.Symlink("../templates", "./build/templates")
		os.Symlink("../roles", "./build/roles")

		os.Symlink("../files", "./deploy/files")
		os.Symlink("../templates", "./deploy/templates")
		os.Symlink("../roles", "./deploy/roles")

		// setup.yml,build.yml,deploy.yml
		setupContent, _ := Asset("../templates/ansible/setup.yml")
		setupFile, _ := os.Create("setup.yml")
		setupFile.Write(setupContent)

		buildContent, _ := Asset("../templates/ansible/build.yml")
		buildFile, _ := os.Create("build.yml")
		buildFile.Write(buildContent)

		deployContent, _ := Asset("../templates/ansible/deploy.yml")
		deployFile, _ := os.Create("deploy.yml")
		deployFile.Write(deployContent)

		defer func() {
			setupFile.Close()
			buildFile.Close()
			deployFile.Close()
		}()

	default:
		ok = false
	}
	return
}
