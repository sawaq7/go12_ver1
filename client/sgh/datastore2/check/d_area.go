package check

import (

	    "net/http"
//	    "fmt"

	    "github.com/sawaq7/go12_ver1/client/sgh/type2"
//	    "time"

        "cloud.google.com/go/datastore"
	    "context"
	    "os"
                                                )

///
///      get area number of district
///

func D_area(w http.ResponseWriter, r *http.Request ,district_no int64)  (int64 ) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN 　district_no : 地区No

//    fmt.Fprintf( w, "d_area start \n" )

    var area_number int64

    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)

    query := datastore.NewQuery("D_Area").Order("District_No")

	count, err := client.Count(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return area_number
	}

	d_area      := make([]type2.D_Area, 0, count)

    if _, err := client.GetAll(ctx, query , &d_area); err != nil {

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
//    fmt.Fprintf( w, "d_area area_number \n" , area_number)
	return area_number
}

