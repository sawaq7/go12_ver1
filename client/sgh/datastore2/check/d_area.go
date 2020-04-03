package check

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
//	    "html/template"
//	    "web/htmls/sgh"
	    "github.com/sawaq7/go12_ver1/client/sgh/type2"
//	    "time"

        "cloud.google.com/go/datastore"
	    "context"
	    "os"
                                                )

///                           ///
/// 蝨ｰ蛹ｺ縺ｮ繧ｨ繝ｪ繧｢謨ｰ繧偵ご繝・ヨ縺吶ｋ ///
///                          ///

func D_area(w http.ResponseWriter, r *http.Request ,district_no int64)  (int64 ) {

//     IN    w      縲縲: 繝ｬ繧ｹ繝昴Φ繧ｹ繝ｩ繧､繧ｿ繝ｼ
//     IN    r      縲縲: 繝ｪ繧ｯ繧ｨ繧ｹ繝医ヱ繝ｩ繝｡繝ｼ繧ｿ
//     IN 縲district_no : 蝨ｰ蛹ｺNo

//    fmt.Fprintf( w, "d_area start \n" )  // 繝・ヰ繝・け

    var area_number int64

    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

//	c := appengine.NewContext(r)
    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)

    query := datastore.NewQuery("D_Area").Order("District_No")
//	q := datastore.NewQuery("D_Area").Order("District_No")

//	count, err := q.Count(c)
	count, err := client.Count(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return area_number
	}

	d_area      := make([]type2.D_Area, 0, count)

    if _, err := client.GetAll(ctx, query , &d_area); err != nil {
//	if _, err := q.GetAll(c, &d_area); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return area_number
	} else {

	  area_number = 0

	  for _, d_areaw := range d_area {
        if district_no == d_areaw.District_No {

		  area_number ++

        }
	  }
	}
//    fmt.Fprintf( w, "d_area area_number \n" , area_number)  // 繝・ヰ繝・け
	return area_number
}

