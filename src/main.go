package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"bytes"
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

func main() {
	baseUrl := "http://localhost:2358/"
	postStr := map[string]string{
//		"source_code": `#include <stdio.h>\n\nint main(void) {\n  char name[10];\n  scanf(\"%s\", name);\n  printf(\"hello, %s\n\", name);\n  return 0;\n}`,
//		"language_id": "4",
//		"stdin": "world",
//	}

		"source_code": "print('Hello World!')",
		"language_id": "71", //python3
		//"number_of_runs": "1",
		//"stdin": "",
		//"expected_output": "Hello World!",
	}
	jsonValue, _ := json.Marshal(postStr)


	rsp, err := http.Post(baseUrl + "submissions/?base64_encoded=false&wait=false", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println("err: ",err)
		return
	}

	jsonParser := json.NewDecoder(rsp.Body)

	var tokenStr token 

	if err = jsonParser.Decode(&tokenStr); err != nil {
		fmt.Println("err: ",err)
	}		
	fmt.Println(tokenStr)

	rsp, err = http.Get(baseUrl + "/languages/71")

	if err != nil {
		fmt.Println("err: ",err)
		return
	}
	//fmt.Println(rsp)
}
