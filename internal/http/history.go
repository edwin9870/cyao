package http

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"reflect"
	"regexp"
	"strings"

	"github.com/edwin/cyoa/internal/util"
)

type ArcWithName struct {
	Name string
	Arc
}

type Chronicle struct {
	Intro     Arc `json:"intro"`
	NewYork   Arc `json:"new-york"`
	Debate    Arc `json:"debate"`
	SeanKelly Arc `json:"sean-kelly"`
	MarkBates Arc `json:"mark-bates"`
	Denver    Arc `json:"denver"`
	Home      Arc `json:"home"`
}
type Options struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
type Arc struct {
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []Options `json:"options"`
}

var templateFunctions = template.FuncMap{
	"hastag": func(input string) string {
		input = strings.ToLower(input)
		reg, err := regexp.Compile("[^a-zA-Z0-9]+")
		util.CheckIfError(err)
		input = reg.ReplaceAllString(input, "")

		return input

	},
}

func History(w http.ResponseWriter, req *http.Request) {

	content, err := ioutil.ReadFile("/home/eramirez/workspace/go/cyoa/gopher.json")

	util.CheckIfError(err)

	var data Chronicle
	err = json.Unmarshal(content, &data)

	util.CheckIfError(err)
	templateResult, err := template.New("history.html").Funcs(templateFunctions).ParseFiles("/home/eramirez/workspace/go/cyoa/web/template/history.html")
	util.CheckIfError(err)

	v := reflect.ValueOf(data)
	arcs := make([]ArcWithName, 0)
	for i := 0; i < v.NumField(); i++ {
		//fmt.Println(v.Type().Field(i).Name)

		arcs = append(arcs, ArcWithName{Name: v.Type().Field(i).Name, Arc: Arc{
			Title:   v.Field(i).FieldByName("Title").String(),
			Story:   v.Field(i).FieldByName("Story").Interface().([]string),
			Options: v.Field(i).FieldByName("Options").Interface().([]Options),
		}})
	}

	err = templateResult.Execute(w, arcs)
	util.CheckIfError(err)
}
