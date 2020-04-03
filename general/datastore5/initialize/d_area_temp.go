package initialize

import (
	    "google.golang.org/appengine"
	    "google.golang.org/appengine/datastore"
	    "net/http"
	    "fmt"
//	    "html/template"
//	    "web/htmls/sgh"
	    "github.com/sawaq7/go12_ver1/client/sgh/type2"
//	    "time"
                                                )

///                                   ///
/// delete all data of temporary file ///
///                                   ///

func D_area_temp(w http.ResponseWriter, r *http.Request )   {

//     IN    w      縲縲: 繝ｬ繧ｹ繝昴Φ繧ｹ繝ｩ繧､繧ｿ繝ｼ
//     IN    r      縲縲: 繝ｪ繧ｯ繧ｨ繧ｹ繝医ヱ繝ｩ繝｡繝ｼ繧ｿ

    fmt.Fprintf( w, "init/D_area_temp start \n" )  // 繝・ヰ繝・け

	c := appengine.NewContext(r)

	q := datastore.NewQuery("D_Area_Temp").Order("District_No")

	count, err := q.Count(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

// import struct for accessing datastore get from github.com/sawaq7/go12_ver1/client/sgh/type2/sgh.go
    fmt.Fprintf( w, "init/D_area_temp count \n" ,count )  // 繝・ヰ繝・け

	d_area_temp     := make([]type2.D_Area_Temp, 0, count)
	keys, err := q.GetAll(c, &d_area_temp )
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
    }
	for pos2, _ := range d_area_temp {


/// 荳譎ゅヵ繧｡繧､繝ｫ縺ｮ蜑企勁
      key := datastore.NewKey(c, "D_Area_Temp"       , "", keys[pos2].IntID(), nil)

	  if err := datastore.Delete(c, key  ); err != nil {

		  http.Error(w, err.Error(), http.StatusInternalServerError)
		  return
	  }
      fmt.Fprintf( w, "init/d_area_temp pos2 %v   \n" , pos2  )  // 繝・ヰ繝・け


    }
	return
}

