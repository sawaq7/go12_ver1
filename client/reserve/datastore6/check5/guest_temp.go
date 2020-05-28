package check5

import (

	    "net/http"
//	    "fmt"
//	    "html/template"
//	    "web/htmls/sgh"
	    "github.com/sawaq7/go12_ver1/client/reserve/type6"
	    "github.com/sawaq7/go12_ver1/general/type5"

	    "cloud.google.com/go/datastore"
	    "context"
	    "os"

                                              )

///                         　　　　
///     get guest no and guest name which was selected
///                         　　　　

func Guest_temp(w http.ResponseWriter, r *http.Request )  ([]type5.General_Work) {

//     IN    w      : response-writer
//     IN    r      : request-parameter

//     OUT general_work_out  : guest no ,guest name

//    fmt.Fprintf( w, "check5/guest_temp start \n" )

    var guest_no int64
    var guest_name string

    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)
    if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return	nil
	}

    query := datastore.NewQuery("Guest_Temp").Order("Guest_Name")

    count, err := client.Count(ctx, query)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

//    fmt.Fprintf( w, "check5/guest_temp count \n" ,count )

	guest_temp     := make([]type6.Guest_Temp, 0, count)

    if _, err := client.GetAll(ctx, query , &guest_temp) ; err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
    } else {
	  for _, guest_tempw := range guest_temp {

        guest_no =    guest_tempw.Guest_No
        guest_name =    guest_tempw.Guest_Name

//        fmt.Fprintf( w, "check5/guest_temp pos2 %v   \n" , pos2  )

      }
//      fmt.Fprintf( w, "check5/guest_temp guest_no \n" ,guest_no )
//      fmt.Fprintf( w, "check5/guest_temp guest_name \n" ,guest_name )
    }

    general_work_out := make([]type5.General_Work, 1)
    general_work_out[0].Int64_Work  = guest_no
    general_work_out[0].String_Work = guest_name

	return general_work_out
}

