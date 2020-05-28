package trans5

import (

	    "net/http"
//	    "fmt"
//	    "html/template"
//	    "web/htmls/sgh"
	    "github.com/sawaq7/go12_ver1/client/reserve/type6"

	    "cloud.google.com/go/datastore"
	    "context"
	    "os"

                                                )

///
///        when record equal reserve_date , select data for Guest_Reserve_Minor in d.s.
///

func Guest_reserve_minor2( reserve_date string ,w http.ResponseWriter, r *http.Request )  ([]type6.Guest_Reserve_Minor ) {

//     IN  reserve_date : reservation date
//     IN    w      　　: response-writer
//     IN    r      　　: request-parameter

//     OUT guest_reserve_minor2_slice  : slice of struct (Guest_Reserve_Minor)

//    fmt.Fprintf( w, "trans.guest_reserve_minor2 start \n" )
//    fmt.Fprintf( w, "trans.guest_reserve_minor2 reserve_date \n" ,reserve_date)

    var i_count int64

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

    query := datastore.NewQuery("Guest_Reserve_Minor").Order("Date")

    count, err := client.Count(ctx, query)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return	nil
	}

	guest_reserve_minor      := make([]type6.Guest_Reserve_Minor, 0, count)
	guest_reserve_minor_slice := make([]type6.Guest_Reserve_Minor, 0)

    keys, err := client.GetAll(ctx, query , &guest_reserve_minor)

    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
//		fmt.Fprintf( w, "d_district_area_show err \n" ,err)  // チE��チE��
		return	nil
	}

    keys_wk := make([]int64, count)

	for ii, keysw := range keys {

       keys_wk[ii] = keysw.ID

    }

///
///   	when record equal reservation date , select data for Guest_Reserve_Minor in d.s.
///

	i_count = 0

	for pos, guest_reserve_minorw := range guest_reserve_minor {

//	  fmt.Fprintf( w, "trans.guest_reserve_minor guest_reserve_minorw %v\n" ,guest_reserve_minorw)

      if reserve_date == guest_reserve_minorw.Date {    // select by reserve_date

         i_count ++

         guest_reserve_minorw.Line_No = i_count

         guest_reserve_minor_slice = append(guest_reserve_minor_slice, type6.Guest_Reserve_Minor { keys_wk[pos]   ,

                                                          guest_reserve_minorw.Line_No     ,
                                                          guest_reserve_minorw.Date        ,
                                                          guest_reserve_minorw.Guest_No    ,
                                                          guest_reserve_minorw.Guest_Name  ,
                                                          guest_reserve_minorw.Start_Time  ,
                                                          guest_reserve_minorw.End_Time           })

      }
	}

    return	guest_reserve_minor_slice
}