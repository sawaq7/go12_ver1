package initialize

import (

	      "google.golang.org/appengine"
	      "google.golang.org/appengine/datastore"
	      "net/http"
	      "fmt"

	      "github.com/sawaq7/go12_ver1/client/sgh/type2"

                                                )

///
///　　　コースNoの計算式を削除
///

func Ai_sgh( course_no int64 ,w http.ResponseWriter, r *http.Request )   {

//     IN  course_no  　: コースNo
//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ

    fmt.Fprintf( w, "initialize.Ai_sgh start \n" )  // チE��チE��

	c := appengine.NewContext(r)

	q := datastore.NewQuery("Sgh_Ai").Order("Course_No")

	count, err := q.Count(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    fmt.Fprintf( w, "initialize.Ai_sgh count \n" ,count )  // チE��チE��

	sgh_ai     := make([]type2.Sgh_Ai, 0, count)
	keys, err := q.GetAll(c, &sgh_ai )
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
    }

///
/// コースNoの計算式を削除
///

	for pos2, sgh_aiw := range sgh_ai {

      if  course_no == sgh_aiw.Course_No {

        key := datastore.NewKey(c, "Sgh_Ai" , "", keys[pos2].IntID(), nil)

	    if err := datastore.Delete(c, key  ); err != nil {

		  http.Error(w, err.Error(), http.StatusInternalServerError)
		  return
	    }
	  }
//      fmt.Fprintf( w, "initialize.Ai_sgh pos2 %v   \n" , pos2  )  // チE��チE��

    }

	return
}

