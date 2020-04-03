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

//    fmt.Fprintf( w, "medical_record_show2 start \n" )  // γEγE―


	var guest_medical_record type6.Guest_Medical_Record

//    flexible_out := datastore2.Datastore_sgh( "D_District_Temp" ,"check" ,idmy , w , r  )

    general_work := check5.Guest_temp (w , r )

//    value2, _ := flexible_out.([]type5.General_Work)

    guest_medical_record.Guest_No   = general_work[0].Int64_Work
    guest_medical_record.Guest_Name = general_work[0].String_Work

    guest_no := general_work[0].Int64_Work

//	fmt.Fprintf( w, "medical_record_show2 : guest_medical_record.Guest_No %v\n", guest_medical_record.Guest_No )  // γEγE―
//	fmt.Fprintf( w, "medical_record_show2 : guest_medical_record.Guest_Name %v\n", guest_medical_record.Guest_Name )  // γEγE―

// η©Ίγ€γ³γΏγΌγγ§γ€γΉε€ζ°γγγγͺγ₯γΌε€γγ²γE

//    value, _ := count.(int64)

//	fmt.Fprintf( w, "medical_record_show2 count %v   \n" , count  )  // γEγE―
//	fmt.Fprintf( w, "medical_record_show2 district_no %v   \n" , district_no  )  // γEγE―

    guest_medical_record.Date   = r.FormValue("date")
    guest_medical_record.Text   = r.FormValue("text2")

//	fmt.Fprintf( w, "medical_record_show2 : guest_medical_record.Text %v\n", guest_medical_record.Text )  // γEγE―

///
///        γEEγΏγΉγγ’γΌγ«γEEγΏγγ»γE
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
///       γ’γγΏγΌγεθ‘¨η€Ί
///

    process4.Medical_record_show(w , r ,guest_no)

//	fmt.Fprintf( w, "medical_record_show2 : normal end \n" )  // γEγE―

}
