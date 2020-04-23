package check

import (

	    "net/http"
//	    "fmt"

	    "github.com/sawaq7/go12_ver1/client/sgh/type2"
	    "github.com/sawaq7/go12_ver1/general/type5"

        "cloud.google.com/go/datastore"
	    "context"
	    "os"
                                                )

///                         　　　　
///             get area number and area name
///                         　　　　

func D_district_temp(w http.ResponseWriter, r *http.Request )  ([]type5.General_Work) {


//     IN    w      　　     : レスポンスライター
//     IN    r      　　     : リクエストパラメータ
//     OUT general_work_out  : area number /area name
//    fmt.Fprintf( w, "check/d_district_temp start \n" )

    var district_no int64
    var district_name string

    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)
    if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

    query := datastore.NewQuery("D_District_Temp").Order("District_No")

    count, err := client.Count(ctx, query)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

//    fmt.Fprintf( w, "check/d_district_temp count \n" ,count )

	d_district_temp     := make([]type2.D_District_Temp, 0, count)

    if _, err := client.GetAll(ctx, query , &d_district_temp) ; err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
    } else {
	  for _, d_district_tempw := range d_district_temp {

        district_no =    d_district_tempw.District_No
        district_name =    d_district_tempw.District_Name

//        fmt.Fprintf( w, "check/d_district_temp pos2 %v   \n" , pos2  )

      }
//      fmt.Fprintf( w, "check/d_district_temp district_no \n" ,district_no )
//      fmt.Fprintf( w, "check/d_district_temp district_name \n" ,district_name )
    }

    general_work_out := make([]type5.General_Work, 1)
    general_work_out[0].Int64_Work  = district_no
    general_work_out[0].String_Work = district_name

	return general_work_out
}

