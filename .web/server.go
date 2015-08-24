package main

import (
	"github.com/sisteamnik/articler"
	"html/template"
	"log"
	"net/http"
	"sort"
	"strings"
)

type Server struct {
	templates *template.Template
	parser    *articler.Parser
}

func NewServer(api *articler.Parser) *Server {
	return &Server{parser: api}
}

func (s *Server) Start() {
	http.HandleFunc("/", s.handler)
	http.HandleFunc("/show", s.showArticles)
	http.HandleFunc("/stats", s.showStats)
	http.ListenAndServe(":8080", nil)
}

func (s *Server) Templates() *template.Template {
	//return s.templates
	funcMap := template.FuncMap{
		// The name "title" is what the function will be called in the template text.
		"str": func(in []byte) string { return string(in) },
		"raw": func(text string) template.HTML {
			return template.HTML(text)
		},
	}
	return template.Must(template.New("titleTest").Funcs(funcMap).ParseGlob("templates/*"))
}

func (s *Server) showStats(w http.ResponseWriter, r *http.Request) {
	//todo remove to s.templates
	arts := s.parser.GetAll()
	sort.Sort(sort.Reverse(articler.Articles(arts)))
	res := map[string]interface{}{
		"count": len(arts),
	}
	err := s.Templates().ExecuteTemplate(w, "stats.html", res)
	if err != nil {
		log.Println(err)
	}
}

func (s *Server) showArticles(w http.ResponseWriter, r *http.Request) {
	//todo remove to s.templates
	arts := s.parser.GetAll()
	sort.Sort(sort.Reverse(articler.Articles(arts)))
	count := 2120
	if len(arts)-1 < count {
		count = len(arts) - 1
	}
	for i := range arts {
		sbody := string(arts[i].Body)
		arts[i].Body = []byte("<p>" + strings.Replace(sbody, "\n", "<p>", -1))
	}
	err := s.Templates().ExecuteTemplate(w, "list.html", arts[0:count])
	if err != nil {
		log.Println(err)
	}
}

func (s *Server) showArticle(w http.ResponseWriter, r *http.Request) {
	//todo remove to s.templates
	arts := s.parser.GetAll()
	sort.Sort(sort.Reverse(articler.Articles(arts)))
	count := 220

	if len(arts)-1 < count {
		count = len(arts) - 1
	}
	err := s.Templates().ExecuteTemplate(w, "list.html", arts[0:count])
	if err != nil {
		log.Println(err)
	}
}

func (s *Server) handler(w http.ResponseWriter, r *http.Request) {
	//todo remove to s.templates
	err := s.Templates().ExecuteTemplate(w, "index.html", "hello")
	if err != nil {
		log.Println(err)
	}
}
