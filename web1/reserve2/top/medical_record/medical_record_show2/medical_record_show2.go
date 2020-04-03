package medical_record_show2

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

func Medical_record_show2(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "medical_record_show2 start \n" )  // 繝・ヰ繝・け


	var guest_medical_record type6.Guest_Medical_Record

//    flexible_out := datastore2.Datastore_sgh( "D_District_Temp" ,"check" ,idmy , w , r  )

    general_work := check5.Guest_temp (w , r )

//    value2, _ := flexible_out.([]type5.General_Work)

    guest_medical_record.Guest_No   = general_work[0].Int64_Work
    guest_medical_record.Guest_Name = general_work[0].String_Work

    guest_no := general_work[0].Int64_Work

//	fmt.Fprintf( w, "medical_record_show2 : guest_medical_record.Guest_No %v\n", guest_medical_record.Guest_No )  // 繝・ヰ繝・け
//	fmt.Fprintf( w, "medical_record_show2 : guest_medical_record.Guest_Name %v\n", guest_medical_record.Guest_Name )  // 繝・ヰ繝・け

// 遨ｺ繧､繝ｳ繧ｿ繝ｼ繝輔ぉ繧､繧ｹ螟画焚繧医ｊ繝舌Μ繝･繝ｼ蛟､繧偵ご繝・ヨ

//    value, _ := count.(int64)

//	fmt.Fprintf( w, "medical_record_show2 count %v   \n" , count  )  // 繝・ヰ繝・け
//	fmt.Fprintf( w, "medical_record_show2 district_no %v   \n" , district_no  )  // 繝・ヰ繝・け

    guest_medical_record.Date   = r.FormValue("date")
    guest_medical_record.Text   = r.FormValue("text2")

//	fmt.Fprintf( w, "medical_record_show2 : guest_medical_record.Text %v\n", guest_medical_record.Text )  // 繝・ヰ繝・け

///
///        繝・・繧ｿ繧ｹ繝医い繝ｼ縺ｫ繝・・繧ｿ繧偵そ繝・ヨ
///

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

    new_key := datastore.IncompleteKey("Guest_Medical_Record", nil)

    if _, err = client.Put(ctx, new_key, &guest_medical_record ); err != nil {
//	if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Guest_Medical_Record", nil), &guest_medical_record); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

///
///       繝｢繝九ち繝ｼ縲蜀崎｡ｨ遉ｺ
///

    process4.Medical_record_show(w , r ,guest_no)

//	fmt.Fprintf( w, "medical_record_show2 : normal end \n" )  // 繝・ヰ繝・け

}
