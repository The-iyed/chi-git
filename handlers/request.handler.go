func RequestHandler(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	ctx := r.Context()
	key := ctx.Value("key").(string)
	w.Write([]byte(fmt.Sprintf("hi %v, %v", userID, key)))
  }