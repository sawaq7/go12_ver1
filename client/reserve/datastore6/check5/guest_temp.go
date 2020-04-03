package check5

import (


//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
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

///                         縲縲縲縲
/// 謖・ｮ壹＠縺溘ご繧ｹ繝・O繝ｻ繧ｲ繧ｹ繝亥錐繧偵ご繝・ヨ縺吶ｋ
///                         縲縲縲縲

func Guest_temp(w http.ResponseWriter, r *http.Request )  ([]type5.General_Work) {


//     IN    w      縲縲     : 繝ｬ繧ｹ繝昴Φ繧ｹ繝ｩ繧､繧ｿ繝ｼ
//     IN    r      縲縲     : 繝ｪ繧ｯ繧ｨ繧ｹ繝医ヱ繝ｩ繝｡繝ｼ繧ｿ
//     OUT general_work_out  : 繧ｲ繧ｹ繝・O繝ｻ繧ｲ繧ｹ繝亥錐

//    fmt.Fprintf( w, "check5/guest_temp start \n" )  // 繝・ヰ繝・け

    var guest_no int64
    var guest_name string

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

    query := datastore.NewQuery("Guest_Temp").Order("Guest_Name")
//	q := datastore.NewQuery("Guest_Temp").Order("Guest_Name")

    count, err := client.Count(ctx, query)
//	count, err := q.Count(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

//    fmt.Fprintf( w, "check5/guest_temp count \n" ,count )  // 繝・ヰ繝・け

	guest_temp     := make([]type6.Guest_Temp, 0, count)

    if _, err := client.GetAll(ctx, query , &guest_temp) ; err != nil {
//	if _, err := q.GetAll(c, &guest_temp ) ; err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
    } else {
	  for _, guest_tempw := range guest_temp {

        guest_no =    guest_tempw.Guest_No
        guest_name =    guest_tempw.Guest_Name

//        fmt.Fprintf( w, "check5/guest_temp pos2 %v   \n" , pos2  )  // 繝・ヰ繝・け

      }
//      fmt.Fprintf( w, "check5/guest_temp guest_no \n" ,guest_no )  // 繝・ヰ繝・け
//      fmt.Fprintf( w, "check5/guest_temp guest_name \n" ,guest_name )  // 繝・ヰ繝・け
    }

    general_work_out := make([]type5.General_Work, 1)
    general_work_out[0].Int64_Work  = guest_no
    general_work_out[0].String_Work = guest_name

	return general_work_out
}

