package initialize

import (
	    "google.golang.org/appengine"
	    "google.golang.org/appengine/datastore"
	    "net/http"
	    "fmt"

	    "github.com/sawaq7/go12_ver1/client/sgh/type2"
//	    "time"
                                                )

///                                   ///
/// delete all data of temporary file ///
///                                   ///

func D_area_temp(w http.ResponseWriter, r *http.Request )   {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ

    fmt.Fprintf( w, "init/D_area_temp start \n" )

	c := appengine.NewContext(r)

	q := datastore.NewQuery("D_Area_Temp").Order("District_No")

	count, err := q.Count(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    fmt.Fprintf( w, "init/D_area_temp count \n" ,count )

	d_area_temp     := make([]type2.D_Area_Temp, 0, count)
	keys, err := q.GetAll(c, &d_area_temp )
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
    }
	for pos2, _ := range d_area_temp {


/// 一時ファイルの削除
      key := datastore.NewKey(c, "D_Area_Temp"       , "", keys[pos2].IntID(), nil)

	  if err := datastore.Delete(c, key  ); err != nil {

		  http.Error(w, err.Error(), http.StatusInternalServerError)
		  return
	  }
      fmt.Fprintf( w, "init/d_area_temp pos2 %v   \n" , pos2  )


    }
	return
}

