package reserve_register_excute

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

///
///     register reservation inf.
///

func Reserve_register_excute (w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "reserve_register_excute start \n" )

	var guest_reserve_minor type6.Guest_Reserve_Minor

///  get guest no and guest name in guest temp. file

//    flexible_out := datastore2.Datastore_sgh( "D_District_Temp" ,"check" ,idmy , w , r  )

    general_work := check5.Guest_temp (w , r )

//    value2, _ := flexible_out.([]type5.General_Work)

    guest_reserve_minor.Guest_No   = general_work[0].Int64_Work
    guest_reserve_minor.Guest_Name = general_work[0].String_Work

//	fmt.Fprintf( w, "reserve_register_excute : guest_reserve_minor.Guest_No %v\n", guest_reserve_minor.Guest_No )
//	fmt.Fprintf( w, "reserve_register_excute : guest_reserve_minor.Guest_Name %v\n", guest_reserve_minor.Guest_Name )

//    value, _ := count.(int64)

//	fmt.Fprintf( w, "reserve_register_excute count %v   \n" , count  )
//	fmt.Fprintf( w, "reserve_register_excute district_no %v   \n" , district_no  )

    guest_reserve_minor.Date   = r.FormValue("date")
    guest_reserve_minor.Start_Time   = r.FormValue("start_time")
	guest_reserve_minor.End_Time = r.FormValue("end_time")

//	fmt.Fprintf( w, "reserve_register_excute : guest_reserve_minor.Start_Time %v\n", guest_reserve_minor.Start_Time )
//	fmt.Fprintf( w, "reserve_register_excute : guest_reserve_minor.End_Time %v\n", guest_reserve_minor.End_Time )

///
///     register reservation inf. in d.s.
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

    new_key := datastore.IncompleteKey("Guest_Reserve_Minor", nil)

    if _, err = client.Put(ctx, new_key, &guest_reserve_minor ); err != nil {

		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

///
///     register reservation inf.
///

	process4.Reserve_register(w , r ,guest_reserve_minor.Guest_No)

//	fmt.Fprintf( w, "reserve_register_excute : normal end \n" )

}
