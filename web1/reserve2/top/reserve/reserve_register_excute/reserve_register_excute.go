package reserve_register_excute

import (

//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"

	    "github.com/sawaq7/go12_ver1/client/reserve/process4"
        "github.com/sawaq7/go12_ver1/client/reserve/datastore6/check5"
	    "github.com/sawaq7/go12_ver1/client/reserve/type6"

//	    "strconv"

        "cloud.google.com/go/datastore"
	    "context"
	    "os"

                                                  )

func Reserve_register_excute (w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "reserve_register_excute start \n" )  // 繝・ヰ繝・け

	var guest_reserve_minor type6.Guest_Reserve_Minor

///  temporary-file繧医ｊ縲∝慍蛹ｺNO繝ｻ蝨ｰ蛹ｺ蜷阪ｒGET

//    flexible_out := datastore2.Datastore_sgh( "D_District_Temp" ,"check" ,idmy , w , r  )

    general_work := check5.Guest_temp (w , r )

//    value2, _ := flexible_out.([]type5.General_Work)

    guest_reserve_minor.Guest_No   = general_work[0].Int64_Work
    guest_reserve_minor.Guest_Name = general_work[0].String_Work


//    guest_reserve_minor.Guest_No = value2[0].Int64_Work
//    guest_reserve_minor.Guest_Name = value2[0].String_Work

//	fmt.Fprintf( w, "reserve_register_excute : guest_reserve_minor.Guest_No %v\n", guest_reserve_minor.Guest_No )  // 繝・ヰ繝・け
//	fmt.Fprintf( w, "reserve_register_excute : guest_reserve_minor.Guest_Name %v\n", guest_reserve_minor.Guest_Name )  // 繝・ヰ繝・け

// 遨ｺ繧､繝ｳ繧ｿ繝ｼ繝輔ぉ繧､繧ｹ螟画焚繧医ｊ繝舌Μ繝･繝ｼ蛟､繧偵ご繝・ヨ

//    value, _ := count.(int64)

//	fmt.Fprintf( w, "reserve_register_excute count %v   \n" , count  )  // 繝・ヰ繝・け
//	fmt.Fprintf( w, "reserve_register_excute district_no %v   \n" , district_no  )  // 繝・ヰ繝・け

    guest_reserve_minor.Date   = r.FormValue("date")
    guest_reserve_minor.Start_Time   = r.FormValue("start_time")
	guest_reserve_minor.End_Time = r.FormValue("end_time")

//	fmt.Fprintf( w, "reserve_register_excute : guest_reserve_minor.Start_Time %v\n", guest_reserve_minor.Start_Time )  // 繝・ヰ繝・け
//	fmt.Fprintf( w, "reserve_register_excute : guest_reserve_minor.End_Time %v\n", guest_reserve_minor.End_Time )  // 繝・ヰ繝・け

/// 繝・・繧ｿ繧ｹ繝医い繝ｼ縺ｫ繝・・繧ｿ繧偵そ繝・ヨ ///

    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

//	c := appengine.NewContext(r)
    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)
    if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    new_key := datastore.IncompleteKey("Guest_Reserve_Minor", nil)

    if _, err = client.Put(ctx, new_key, &guest_reserve_minor ); err != nil {
//	if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Guest_Reserve_Minor", nil), &guest_reserve_minor); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

/// 繝｢繝九ち繝ｼ縲蜀崎｡ｨ遉ｺ ///

	process4.Reserve_register(w , r ,guest_reserve_minor.Guest_No)

//	fmt.Fprintf( w, "reserve_register_excute : normal end \n" )  // 繝・ヰ繝・け

}
