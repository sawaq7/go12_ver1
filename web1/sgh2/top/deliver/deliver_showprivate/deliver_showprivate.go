package sky

import (

	    "net/http"
	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/sgh/process"
	    "strconv"
                         )

/// main car no ごとの　配達チE�Eタを表示する　///

func init() {
	http.HandleFunc("/deliver_showprivate", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "deliver_showprivate start \n" )  // チE��チE��

	private_no := r.FormValue("private_no")         // car NoをゲチE��
//	fmt.Fprintf( w, "deliver_showprivate : private_no %v\n", private_no )  // チE��チE��

	private_now ,err := strconv.Atoi(private_no)  // 個人Noの整数匁E	if err != nil {
//		http.Error(w,err.Error(), http.StatusInternalServerError)
       fmt.Fprintf( w, "deliver_showprivate : a private_no must be half-width characters %v\n"  )
		return
	}

	private_now2 := int64(private_now)   // 整数の64ビット化

/// モニター　再表示 ///

	process.Deliver_showprivate(w , r ,private_now2 )

//	http.Redirect(w, r, "/", http.StatusFound)
//	fmt.Fprintf( w, "deliver_showprivate : normal end \n" )  // チE��チE��




}
