package copy3

import (
	    "google.golang.org/appengine"
	    "google.golang.org/appengine/datastore"
	    "net/http"
	    "fmt"

//	    "github.com/sawaq7/go12_ver1/general/datastore5/check3"
        "github.com/sawaq7/go12_ver1/client/sgh/type2"

                                                )

///                          ///
/// ãEEã¿ã¹ãã¢ãã³ããEãã  ///
///                         ///

// func All( w http.ResponseWriter, r *http.Request ,basic_name string ,copy_file string ,new_file string ) {
 func All( w http.ResponseWriter, r *http.Request ,basic_name string ,copy_file string ,new_file string )( err error) {

//     IN    w        : ã¬ã¹ãã³ã¹ã©ã¤ã¿ã¼
//     IN    r        : ãªã¯ã¨ã¹ããã©ã¡ã¼ã¿
//     IN  basic_name : åºæ¬ã®ãEEã¿ã¹ãã¢åE//     IN  copy_file  : ã³ããEåEEãEEã¿ã¹ãã¢åE//     IN  new_file   : ãã¥ã¼ãEEã¿ã¹ãã¢åE//    OUT  err        : ã¨ã©ã¼ã¡ãE»ã¼ã¸

    fmt.Fprintf( w, "copy3.all start \n" )  // ãEãE¯
    fmt.Fprintf( w, "copy3.all basic_name %v\n" ,basic_name)  // ãEãE¯

///
///  ã¨ã©ã¼ã¡ãE»ã¼ã¸ãã»ãE
///

    c := appengine.NewContext(r)       ///ãã³ã³ãE­ã¹ããåå¾ããE
	q := datastore.NewQuery(copy_file) /// ã¯ã­ã¼ã³ã®ãªã¼ãã¼ãGET

	count, err := q.Count(c)           ///ãã¬ã³ã¼ãæ°ãGET
	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
		return  err

	}
    fmt.Fprintf( w, "copy3.all count %v\n" ,count)  // ãEãE¯


//   if err :=check3.Name( w , basic_name   ) ;  err != nil{
//       http.Error(w, err.Error(), http.StatusInternalServerError)

//	  return err

//	}

///
/// ã¯ã­ã¼ã³æE ±ãSETããã¯ã¼ã¯ã¨ãªã¢ãç¢ºä¿E///

    ds_data := make([]type2.Deliver, 0, count)



	if _, err := q.GetAll(c, &ds_data);  err != nil {         // ã¯ã­ã¼ã³æE ±ãGET

//	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return  err

	} else{
      for _, ds_dataw := range ds_data {                       //ããã¥ã¼ãã¡ã¤ã«ã«ã¯ã­ã¼ã³æE ±ãã»ãE

	    if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, new_file, nil), &ds_dataw); err != nil {

//		  http.Error(w,err.Error(), http.StatusInternalServerError)
		  return  err

	    }

	  }
	}

	fmt.Fprintf( w, "copy3.all normal end \n" )  // ãEãE¯

    return nil
}

