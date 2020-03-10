package main

import (

	    "net/http"
	    "fmt"
	    "log"
	    "os"

      ///
	  ///  ---------------------  sample multi-struct appri menu   ----------------------------
      ///

      "web1/sample1/top/hello_sample"




                                                                                   )
///
///     process list
///

func main() {

   ///
   ///  ---------------------   application medical/xray process  ----------------------------
   ///

   http.HandleFunc("/hello_sample", hello_sample.Hello_sample)

///
///       ポートをセット
///

  port := os.Getenv("PORT")

  if port == "" {
        port = "8080"
        log.Printf("Defaulting to port %s", port)
  }

  log.Printf("Listening on port %s", port)
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))

}
