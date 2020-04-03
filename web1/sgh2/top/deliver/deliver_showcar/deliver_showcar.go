package sky

import (

	    "net/http"
	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/sgh/process"
	    "strconv"
                         )

/// main car no ごとの　配達チE�Eタを表示する　///

func init() {
	http.HandleFunc("/deliver_showcar", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "deliver_showcar start \n" )  // チE��チE��

	car_no := r.FormValue("car_no")         // car NoをゲチE��
//	fmt.Fprintf( w, "deliver_showcar : car_no %v\n", car_no )  // チE��チE��

	car_now ,err := strconv.Atoi(car_no)  // 個人Noの整数匁E	if err != nil {
//		http.Error(w,err.Error(), http.StatusInternalServerError)
       fmt.Fprintf( w, "deliver_showcar : a car_no must be half-width characters %v\n"  )
		return
	}

	car_now2 := int64(car_now)   // 整数の64ビット化

/// モニター　再表示 ///

	process.Deliver_showcar(w , r ,car_now2 )

//	http.Redirect(w, r, "/", http.StatusFound)
//	fmt.Fprintf( w, "deliver_showcar : normal end \n" )  // チE��チE��




}
