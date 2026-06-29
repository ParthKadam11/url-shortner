package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type URL struct {
	ID          string `json:"id"`
	OriginalUrl string `json:"original_url"`
	ShortUrl    string `json:"short_url"`
	CreatedAt   time.Time
}

var urlDB = make(map[string]URL)

func generateShortUrl(OriginalUrl string) string {
	hasher := md5.New()
	hasher.Write([]byte(OriginalUrl))

	data := hasher.Sum(nil)
	fmt.Println("hasher = ", hasher)
	fmt.Println("data = ", data)
	hash := hex.EncodeToString(data)
	fmt.Println("hash = ", hash)
	fmt.Println("FinalHash = ", hash[:8])
	return hash[:8]
}

func createUrl(Original_url string) string {
	shortUrl:=generateShortUrl(Original_url)
	id:=shortUrl //for simplicity we are using short url as ID :/
	urlDB[id]=URL{
		ID: id,
		OriginalUrl: Original_url,
		ShortUrl: shortUrl,
		CreatedAt: time.Now(),
	}
	return shortUrl
}

func getUrl(id string)(URL,error){
	url,ok:=urlDB[id]
	if !ok{
		return URL{}, errors.New("Url Not Found")
	}
	return url,nil
}

//this function handles the "/" page
func RootpageHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("GET Method")
	fmt.Fprintf(w, "Hello / Route")
}

func ShortUrlHandler(w http.ResponseWriter, r *http.Request){
	var data struct{
		URL string `json:"url"`
	}

	err:=json.NewDecoder(r.Body).Decode(&data)
	if(err!=nil){
		http.Error(w,"Invalid Request Body",http.StatusBadRequest)
		return 
	}

	shortUrl_:=createUrl(data.URL)
	response:=struct{
		ShortUrl string `json:"short_url"`
	}{
		ShortUrl: shortUrl_,
	}

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(response)
}

func RedirectUrl(w http.ResponseWriter, r *http.Request){
	id:=r.URL.Path[len("/redirect/"):]
	
	url,err:=getUrl(id)
	if(err!=nil){
		http.Error(w,"Invalid Request", http.StatusNotFound)
	}

	http.Redirect(w,r,url.OriginalUrl,http.StatusFound)
}

func main() {
	//fmt.Println("Starting your url-shortner")
	//OriginalUrl := "https://www.youtube.com/watch?v=dVVJU-3eU1g"
	//generateShortUrl(OriginalUrl)

	//register a handle function to handle arr request to root URL ("/")
	http.HandleFunc("/",RootpageHandler)
	http.HandleFunc("/short",ShortUrlHandler)

	//start HTTP server on port 3000 
	fmt.Println("Server starting on port 3000...")
	err :=http.ListenAndServe(":3000",nil)
	if(err!=nil){
		fmt.Println("Error on starting server:",err)
	}
}
