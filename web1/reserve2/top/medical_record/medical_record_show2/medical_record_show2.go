package medical_record_show2

import (

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

//    fmt.Fprintf( w, "medical_record_show2 start \n" )


	var guest_medical_record type6.Guest_Medical_Record

//    flexible_out := datastore2.Datastore_sgh( "D_District_Temp" ,"check" ,idmy , w , r  )

    general_work := check5.Guest_temp (w , r )

//    value2, _ := flexible_out.([]type5.General_Work)

    guest_medical_record.Guest_No   = general_work[0].Int64_Work
    guest_medical_record.Guest_Name = general_work[0].String_Work

    guest_no := general_work[0].Int64_Work

//	fmt.Fprintf( w, "medical_record_show2 : guest_medical_record.Guest_No %v\n", guest_medical_record.Guest_No )
//	fmt.Fprintf( w, "medical_record_show2 : guest_medical_record.Guest_Name %v\n", guest_medical_record.Guest_Name )

//    value, _ := count.(int64)   // get value from interface data

//	fmt.Fprintf( w, "medical_record_show2 count %v   \n" , count  )
//	fmt.Fprintf( w, "medical_record_show2 district_no %v   \n" , district_no  )

    guest_medical_record.Date   = r.FormValue("date")
    guest_medical_record.Text   = r.FormValue("text2")

//	fmt.Fprintf( w, "medical_record_show2 : guest_medical_record.Text %v\n", guest_medical_record.Text )

///
///      set medical record .inf. in d.s
///

    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)
    if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    new_key := datastore.IncompleteKey("Guest_Medical_Record", nil)

    if _, err = client.Put(ctx, new_key, &guest_medical_record ); err != nil {

		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

///
///         show medical recode inf. on web
///

    process4.Medical_record_show(w , r ,guest_no)

//	fmt.Fprintf( w, "medical_record_show2 : normal end \n" )

}
