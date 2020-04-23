package copy3

import (
	    "google.golang.org/appengine"
	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"

//	    "github.com/sawaq7/go12_ver1/general/datastore5/check3"
        "github.com/sawaq7/go12_ver1/client/sgh/type2"

                                                )

///
///     copy d.s
///


 func All( w http.ResponseWriter, r *http.Request ,basic_name string ,copy_file string ,new_file string )( err error) {

//     IN    w        : レスポンスライター
//     IN    r        : リクエストパラメータ
//     IN  basic_name : d.s. name of basic
//     IN  copy_file  : d.s. name which is cpied
//     IN  new_file   : new d.s. name

//    OUT  err        :   error inf.

//    fmt.Fprintf( w, "copy3.all start \n" )
//    fmt.Fprintf( w, "copy3.all basic_name %v\n" ,basic_name)

    c := appengine.NewContext(r)

	q := datastore.NewQuery(copy_file)

	count, err := q.Count(c)
	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
		return  err

	}
                                  //   allocate work area for records
    ds_data := make([]type2.Deliver, 0, count)

	if _, err := q.GetAll(c, &ds_data);  err != nil {         //  get d.s. inf.

//	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return  err

	} else{
      for _, ds_dataw := range ds_data {                       //   put copy inf. for d.s.

	    if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, new_file, nil), &ds_dataw); err != nil {

//		  http.Error(w,err.Error(), http.StatusInternalServerError)
		  return  err

	    }

	  }
	}

//	fmt.Fprintf( w, "copy3.all normal end \n" )

    return nil
}

