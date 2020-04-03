package sky

import (

	    "net/http"
	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/sgh/process"
                                                  )

func init() {

	http.HandleFunc("/deliver_showall_exam", handler)

}

func handler(w http.ResponseWriter, r *http.Request) {

    fmt.Fprintf( w, "deliver_showall_exam start \n" )  // 繝・ヰ繝・け


/// 繝｢繝九ち繝ｼ縲蜀崎｡ｨ遉ｺ ///

    process.Deliver_showall_exam( w , r )

	fmt.Fprintf( w, "deliver_showall_exam : normal end \n" )  // 繝・ヰ繝・け

}
