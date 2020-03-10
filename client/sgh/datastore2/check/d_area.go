package check

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
//	    "html/template"
//	    "web/htmls/sgh"
	    "client/sgh/type2"
//	    "time"

        "cloud.google.com/go/datastore"
	    "context"
	    "os"
                                                )

///                           ///
/// 地区のエリア数をゲットする ///
///                          ///

func D_area(w http.ResponseWriter, r *http.Request ,district_no int64)  (int64 ) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN 　district_no : 地区No

//    fmt.Fprintf( w, "d_area start \n" )  // デバック

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
//    fmt.Fprintf( w, "d_area area_number \n" , area_number)  // デバック
	return area_number
}

