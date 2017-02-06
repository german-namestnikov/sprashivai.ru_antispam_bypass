package main

import (
	"fmt"

	"crypto/sha256"
	"encoding/hex"

	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
)

func get_sig(username, hash string) string {
	string_to_sig := "2O889rz3A5K30hxi8cazk7UcJY8623tBNGMLW49R"
	string_to_sig += hash
	string_to_sig += username
	string_to_sig += "bDM808Lt8g647vlGhDtbCRhvxBkJyNkmjn8574zT"

	sig := sha256.Sum256([]byte(string_to_sig))
	sig_string := hex.EncodeToString(sig[:])

	return sig_string
}

func get_hash(username string) string {
	url_string := "http://sprashivai.ru/" + username

	resp, _ := http.Get(url_string)
	bytes, _ := ioutil.ReadAll(resp.Body)

	html_content := string(bytes)

	defer resp.Body.Close()

	r, _ := regexp.Compile("Responses.ask\\('(.*?)', '(.*?)'\\);")
	hash_string := r.FindStringSubmatch(html_content)[2]

	return hash_string
}

func send_question(username, question string) {
	url_string := "http://sprashivai.ru/question/ask"
	hash_string := get_hash(username)
	sig_string := get_sig(username, hash_string)

	data := url.Values{}
	data.Add("username", username)
	data.Add("anonymously", "yes")
	data.Add("question", question)
	data.Add("hash", hash_string)
	data.Add("sig", sig_string)

	resp, _ := http.PostForm(url_string, data)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body[:]))
}

func main() {
	username := "AnotherStupidUser"
	send_question(username, "Selfmessaging is not onanism!")

}
