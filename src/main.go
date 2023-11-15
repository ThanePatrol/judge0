package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"net/http"
	"bytes"
	"os"
	"time"
)

//{
//  "source_code": "#include <stdio.h>\n\nint main(void) {\n  char name[10];\n  scanf(\"%s\", name);\n  printf(\"hello, %s\n\", name);\n  return 0;\n}",
//  "language_id": 4,
//  "number_of_runs": 1,
//  "stdin": "Judge0",
//  "expected_output": "hello, Judge0",
//  "cpu_time_limit": 1,
//  "cpu_extra_time": 0.5,
//  "wall_time_limit": 100000,
//  "memory_limit": 128000,
//  "stack_limit": 128000,
//  "enable_per_process_and_thread_time_limit": false,
//  "enable_per_process_and_thread_memory_limit": false,
//  "max_file_size": 1024
//}
//https://localhost:2358/submissions/?base64_encoded=false&wait=false

type token struct {
	Token string `json:"token"`
}

type result struct {
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
	CompileOutput string `json:"compile_output"`
	ExitCode int `json:"exit_code"`
	Finished string `json:"finished_at"` 
	Runtime string `json:"time"`
}

type Submission struct {
    SourceCode               string `json:"source_code"`
    LanguageID               int    `json:"language_id"`
    NumberOfRuns             int    `json:"number_of_runs"`
    RedirectStderrToStdout   bool   `json:"redirect_stderr_to_stdout"`
}

type UserSubmission struct {
	SourceCode string `json:"source_code"`
	Lang string `json:"language"`
}

var mapLang = map[string]int{
	"java": 62,
	"python": 71,
	"cpp": 54,
}

func index(w http.ResponseWriter, req *http.Request) {
	file, err := os.ReadFile("resources/index.html")
	if err != nil {
		fmt.Fprintf(w, "<h1>hello, couldn't read index.html</h1>")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} 
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(file)
}

// TODO parse the request and send it to the judge0 api
func postSubmission(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var sub UserSubmission
	err := decoder.Decode(&sub)
	if err != nil {
		fmt.Println("err: ",err)
		return
	}
	subOutput, err := doSubmission(sub)

	htmlStr := "<h2>Submission Output</h2><br>"
	if subOutput.Stderr != "" {
		s := strings.Replace(subOutput.CompileOutput, "\n", "<br>", -1)
		htmlStr += "<h3>compile output:" + s + "</h3><br>"
		s = strings.Replace(subOutput.Stderr, "\n", "<br>", -1)
		htmlStr += "<h3>stderr:" + s + "</h3><br>"
	} else {
		htmlStr += "<h3>stdout:" + subOutput.Stdout + "</h3><br>"
		htmlStr += "<h3>runtime:" + subOutput.Runtime + "s</h3><br>"
	}
	fmt.Println(htmlStr)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, htmlStr)
}

func doSubmission(sub UserSubmission) (result, error) {
	
	baseUrl := "http://localhost:2358/"
	postObj := Submission{
		SourceCode: sub.SourceCode,
		LanguageID: mapLang[sub.Lang],
		NumberOfRuns: 1,
		RedirectStderrToStdout: false,
	}
	jsonValue, _ := json.Marshal(postObj)


	rsp, err := http.Post(baseUrl + "submissions/?base64_encoded=false&wait=false", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println("err: ",err)
		return result{}, err
	}
	jsonParser := json.NewDecoder(rsp.Body)


	var tokenStr token 

	if err = jsonParser.Decode(&tokenStr); err != nil {
		fmt.Println("err: ",err)
	}		
	getSubUrl := baseUrl + "submissions/" + tokenStr.Token + "?base64_encoded=false&fields=stdout,stderr,exit_code,time,compile_output,finished_at"

	rsp.Body.Close()
	
	// give it some time to compile and execute
	time.Sleep(1 * time.Second)

	rsp, err = http.Get(getSubUrl)
	jsonParser = json.NewDecoder(rsp.Body)
	subResult := result{}
	if err = jsonParser.Decode(&subResult); err != nil {
		fmt.Println("err: ",err)
	}		

	return subResult, nil
}

func main() {
	fs := http.FileServer(http.Dir("resources"))
	http.Handle("/resources/", http.StripPrefix("/resources/", fs))


	http.HandleFunc("/", index)
	http.HandleFunc("/api/submit", postSubmission)
	http.ListenAndServe(":8080", nil)
}
