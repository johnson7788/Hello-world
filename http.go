func main() {
	 http.HandleFunc("/Library/WebServer/Documents/", homePage)
	 if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal("failed to start server", err)
}
}
