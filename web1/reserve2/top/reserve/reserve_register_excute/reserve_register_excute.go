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

//    fmt.Fprintf( w, "reserve_register_excute start \n" )  // γEγE―

	var guest_reserve_minor type6.Guest_Reserve_Minor

///  temporary-fileγγγε°εΊNOγ»ε°εΊεγGET

//    flexible_out := datastore2.Datastore_sgh( "D_District_Temp" ,"check" ,idmy , w , r  )

    general_work := check5.Guest_temp (w , r )

//    value2, _ := flexible_out.([]type5.General_Work)

    guest_reserve_minor.Guest_No   = general_work[0].Int64_Work
    guest_reserve_minor.Guest_Name = general_work[0].String_Work


//    guest_reserve_minor.Guest_No = value2[0].Int64_Work
//    guest_reserve_minor.Guest_Name = value2[0].String_Work

//	fmt.Fprintf( w, "reserve_register_excute : guest_reserve_minor.Guest_No %v\n", guest_reserve_minor.Guest_No )  // γEγE―
//	fmt.Fprintf( w, "reserve_register_excute : guest_reserve_minor.Guest_Name %v\n", guest_reserve_minor.Guest_Name )  // γEγE―

// η©Ίγ€γ³γΏγΌγγ§γ€γΉε€ζ°γγγγͺγ₯γΌε€γγ²γE

//    value, _ := count.(int64)

//	fmt.Fprintf( w, "reserve_register_excute count %v   \n" , count  )  // γEγE―
//	fmt.Fprintf( w, "reserve_register_excute district_no %v   \n" , district_no  )  // γEγE―

    guest_reserve_minor.Date   = r.FormValue("date")
    guest_reserve_minor.Start_Time   = r.FormValue("start_time")
	guest_reserve_minor.End_Time = r.FormValue("end_time")

//	fmt.Fprintf( w, "reserve_register_excute : guest_reserve_minor.Start_Time %v\n", guest_reserve_minor.Start_Time )  // γEγE―
//	fmt.Fprintf( w, "reserve_register_excute : guest_reserve_minor.End_Time %v\n", guest_reserve_minor.End_Time )  // γEγE―

/// γEEγΏγΉγγ’γΌγ«γEEγΏγγ»γE ///

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

/// γ’γγΏγΌγεθ‘¨η€Ί ///

	process4.Reserve_register(w , r ,guest_reserve_minor.Guest_No)

//	fmt.Fprintf( w, "reserve_register_excute : normal end \n" )  // γEγE―

}
