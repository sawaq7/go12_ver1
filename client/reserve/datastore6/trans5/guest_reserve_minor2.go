package trans5

import (

//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
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
/// æEŽããæĨäģãEäēčĻæå ąãã˛ãEãã
///

func Guest_reserve_minor2( reserve_date string ,w http.ResponseWriter, r *http.Request )  ([]type6.Guest_Reserve_Minor ) {

//     IN  reserve_date : äēį´EĨ
//     IN    w      ãã: ãŦãšããŗãšãŠã¤ãŋãŧ
//     IN    r      ãã: ãĒã¯ã¨ãšãããŠãĄãŧãŋ

//     OUT guest_reserve_minor2_slice  : æ§é äŊãâã¨ãĒãĸæE ąâãEãšãŠã¤ãš

//    fmt.Fprintf( w, "trans.guest_reserve_minor2 start \n" )  // ãEãE¯
//    fmt.Fprintf( w, "trans.guest_reserve_minor2 reserve_date \n" ,reserve_date)  // ãEãE¯

    var i_count int64

    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

//	c := appengine.NewContext(r)
    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)
    if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return	nil
	}

    query := datastore.NewQuery("Guest_Reserve_Minor").Order("Date")
//	q := datastore.NewQuery("Guest_Reserve_Minor").Order("Date")

    count, err := client.Count(ctx, query)
//	count, err := q.Count(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return	nil
	}

	guest_reserve_minor      := make([]type6.Guest_Reserve_Minor, 0, count)
	guest_reserve_minor_slice := make([]type6.Guest_Reserve_Minor, 0)

    keys, err := client.GetAll(ctx, query , &guest_reserve_minor)
//	keys, err := q.GetAll(c, &guest_reserve_minor)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
//		fmt.Fprintf( w, "d_district_area_show err \n" ,err)  // ãEãE¯
		return	nil
	}

    keys_wk := make([]int64, count)

	for ii, keysw := range keys {

       keys_wk[ii] = keysw.ID

    }

	i_count = 0

	for pos, guest_reserve_minorw := range guest_reserve_minor {

//	  fmt.Fprintf( w, "trans.guest_reserve_minor guest_reserve_minorw %v\n" ,guest_reserve_minorw)  // ãEãE¯

///  æŠčEãĢãããã§ãE¯é EŽããģãE

      if reserve_date == guest_reserve_minorw.Date {

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

