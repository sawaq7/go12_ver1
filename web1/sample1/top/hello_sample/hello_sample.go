package hello_sample

import (

	     "fmt"
	     "net/http"
	     "education"
                                            )

func Hello_sample(w http.ResponseWriter, r *http.Request) {

    fmt.Fprintf( w, "hello_sample start \n" )  // デバック

    var circle_circum ,radius float64

    radius = 2.0
    circle_circum = education.Circle_Circum(radius )

    fmt.Fprintf( w, "hello_sample  %v\n" ,circle_circum)




}
