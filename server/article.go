

 

func GetArticle(w http.ResponseWriter, r *http.Request) {
  dateParam := chi.URLParam(r, "date")
  slugParam := chi.URLParam(r, "slug")
  article, err := database.GetArticle(date, slug)

  if err != nil {
    w.WriteHeader(422)
    w.Write([]byte(fmt.Sprintf("error fetching article %s-%s: %v", dateParam, slugParam, err)))
    return
  }
  
  if article == nil {
    w.WriteHeader(404)
    w.Write([]byte("article not found"))
    return
  }
  w.Write([]byte(article.Text()))
})
