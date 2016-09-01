# How it work ?
## main.go
Create new router in main() with httprouter for example "/" route to c.Index 
```go
package main

import (
  "net/http"
  "log"
  "github.com/julienschmidt/httprouter"
  c "github.com/wattanar/go-starter/controllers"
)

func main()  {
  // Create New Router
  r := httprouter.New()
  // Router
  r.GET("/", c.Index)
  // Setup Static Path
  r.NotFound = http.StripPrefix("/static/", http.FileServer(http.Dir("./assets")))
  // Listening on port 8080
  log.Fatal(http.ListenAndServe(":8080", r))
}
```
## controllers/route.go
From main() Index parse 3 argument and for example render page-index.html with html/template package
```go
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  t, _ := template.ParseFiles("views/base.tmpl", "views/index.tmpl")
  t.Execute(w, nil)
}
```
## Views
### views/index.tmpl
```html
{{define "content"}}

  <h1>Hello World!</h1>

{{end}}
```
### views/base.tmpl
```html
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Hello Go!</title>
    <link rel="stylesheet" href="/static/css/bootstrap.min.css">
    <script src="/static/js/jquery-2.2.0.min.js" charset="utf-8"></script>
    <script src="/static/js/bootstrap.min.js" charset="utf-8"></script>
  </head>
  <body>
    <div class="container">
      {{template "content" .}}
    </div>
  </body>
</html>
```
## Run
```shell
go run main.go
```
Open browser and go to localhost:8080
