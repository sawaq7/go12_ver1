package sky

import (

	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/sgh/process"
	    "strconv"
                         )

///
/// 　　   show delivery inf. for each car-no
///

func init() {
	http.HandleFunc("/deliver_showcar", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "deliver_showcar start \n" )

	car_no := r.FormValue("car_no")
//	fmt.Fprintf( w, "deliver_showcar : car_no %v\n", car_no )

	car_now ,err := strconv.Atoi(car_no)
	if err != nil {
//		http.Error(w,err.Error(), http.StatusInternalServerError)
//       fmt.Fprintf( w, "deliver_showcar : a car_no must be half-width characters %v\n"  )
		return
	}

	car_now2 := int64(car_now)   //make an integer64

    //    show delivery inf. for each car on web
	process.Deliver_showcar(w , r ,car_now2 )

//	fmt.Fprintf( w, "deliver_showcar : normal end \n" )




}
