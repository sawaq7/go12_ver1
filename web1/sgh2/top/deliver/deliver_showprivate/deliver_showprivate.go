package sky

import (

	    "net/http"
	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/sgh/process"
	    "strconv"
                         )

///
/// 　　   show delivery inf. for each private
///

func init() {
	http.HandleFunc("/deliver_showprivate", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "deliver_showprivate start \n" )

	private_no := r.FormValue("private_no")
//	fmt.Fprintf( w, "deliver_showprivate : private_no %v\n", private_no )

	private_now ,err := strconv.Atoi(private_no)
	if err != nil {
//		http.Error(w,err.Error(), http.StatusInternalServerError)
       fmt.Fprintf( w, "deliver_showprivate : a private_no must be half-width characters %v\n"  )
		return
	}

	private_now2 := int64(private_now)

    //    show delivery inf. for each private on web
	process.Deliver_showprivate(w , r ,private_now2 )

//	fmt.Fprintf( w, "deliver_showprivate : normal end \n" )




}
