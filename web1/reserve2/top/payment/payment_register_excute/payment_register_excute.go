package payment_register_excute

import (

	    "net/http"
//	    "fmt"

	    "github.com/sawaq7/go12_ver1/client/reserve/process4"
        "github.com/sawaq7/go12_ver1/client/reserve/datastore6/check5"
	    "github.com/sawaq7/go12_ver1/client/reserve/type6"

	    "strconv"

	    "cloud.google.com/go/datastore"
	    "context"
	    "os"
                                                  )

///                         　　　　
///     register payment inf. for guest which is selected in d.s.
///

func Payment_register_excute(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "payment_register_excute start \n" )

	var guest_payment type6.Guest_Payment

///  get guest no and guest name in guest temp inf.

//    flexible_out := datastore2.Datastore_sgh( "D_District_Temp" ,"check" ,idmy , w , r  )

    general_work := check5.Guest_temp (w , r )

//    value2, _ := flexible_out.([]type5.General_Work)

    guest_payment.Guest_No   = general_work[0].Int64_Work
    guest_payment.Guest_Name = general_work[0].String_Work

//	fmt.Fprintf( w, "payment_register_excute : guest_payment.Guest_No %v\n", guest_payment.Guest_No )
//	fmt.Fprintf( w, "payment_register_excute : guest_payment.Guest_Name %v\n", guest_payment.Guest_Name )

//    value, _ := count.(int64)

//	fmt.Fprintf( w, "payment_register_excute count %v   \n" , count  )
//	fmt.Fprintf( w, "payment_register_excute district_no %v   \n" , district_no  )

    guest_payment.Date   = r.FormValue("date")

    guest_payment.Item   = r.FormValue("item")

	amount               := r.FormValue("amount")
	amountw ,err := strconv.Atoi(amount)  // make an integer
	if err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)

		return
	}

	guest_payment.Amount = int64(amountw)   // make an integer64

//	fmt.Fprintf( w, "payment_register_excute : guest_payment.Item %v\n", guest_payment.Item )
//	fmt.Fprintf( w, "payment_register_excute : guest_payment.Amount %v\n", guest_payment.Amount )

/// チE�EタストアーにチE�EタをセチE�� ///

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

    new_key := datastore.IncompleteKey("Guest_Payment", nil)

    if _, err = client.Put(ctx, new_key, &guest_payment ); err != nil {

		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}


///
///      show payment inf. on web
///

	process4.Payment_register(w , r ,guest_payment.Guest_No)

//	fmt.Fprintf( w, "payment_register_excute : normal end \n" )  // チE��チE��

}
