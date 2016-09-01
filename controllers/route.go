package controllers

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"strings"
)

func Landing(rw http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	t, _ := template.ParseFiles("views/base.tmpl", "views/index.tmpl")
	t.Execute(rw, nil)
}

func Theaters(rw http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	type theater struct {
		Title string
		Link  string
	}

	all_theater := []theater{}

	doc, err := goquery.NewDocument("http://www.thanacineplex.com/theater.php")

	if err != nil {
		json.NewEncoder(rw).Encode(all_theater)
	}

	doc.Find("#branchSelect option").Each(func(i int, s *goquery.Selection) {
		title := strings.Replace(s.Text(), "\t", "", -1)
		link, _ := s.Attr("value")
		link_replace := strings.Replace(link, "/theater.php?theater_branch=", "", -1)
		all_theater = append(all_theater, theater{title, link_replace})
	})

	json.NewEncoder(rw).Encode(all_theater)
}

func Movies(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	type movie struct {
		Title_EN string
		Title_TH string
		Time     string
		Thumb    string
	}

	all_movies := []movie{}

	doc, err := goquery.NewDocument("http://www.thanacineplex.com/movie.php?movie=&branch=" + ps.ByName("id"))

	if err != nil {
		json.NewEncoder(rw).Encode(all_movies)
	}

	doc.Find("#today ul li").Each(func(i int, s *goquery.Selection) {

		title_en := strings.Replace(s.Find(".title-en").Text(), "\t", "", -1)
		title_en_clean := strings.Replace(title_en, "\n", "", -1)
		title_th := strings.Replace(s.Find(".title-th").Text(), "\t", "", -1)
		title_th_clean := strings.Replace(title_th, "\n", "", -1)
		time := s.Find(".time ul li").Text()

		thumb, _ := s.Find(".thumb img").Attr("src")

		if title_en != "" {
			all_movies = append(all_movies, movie{title_en_clean, title_th_clean, time, thumb})
		}

	})

	json.NewEncoder(rw).Encode(all_movies)
}
