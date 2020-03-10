package sky

import (

	    "net/http"
	    "fmt"
	    "client/sgh/process"
                                                  )

func init() {

	http.HandleFunc("/deliver_showall_exam", handler)

}

func handler(w http.ResponseWriter, r *http.Request) {

    fmt.Fprintf( w, "deliver_showall_exam start \n" )  // デバック


/// モニター　再表示 ///

    process.Deliver_showall_exam( w , r )

	fmt.Fprintf( w, "deliver_showall_exam : normal end \n" )  // デバック

}
