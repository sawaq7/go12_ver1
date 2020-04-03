package payment_register_excute

import (

//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
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

func Payment_register_excute(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "payment_register_excute start \n" )  // 繝・ヰ繝・け

	var guest_payment type6.Guest_Payment

///  temporary-file繧医ｊ縲∝慍蛹ｺNO繝ｻ蝨ｰ蛹ｺ蜷阪ｒGET

//    flexible_out := datastore2.Datastore_sgh( "D_District_Temp" ,"check" ,idmy , w , r  )

    general_work := check5.Guest_temp (w , r )

//    value2, _ := flexible_out.([]type5.General_Work)

    guest_payment.Guest_No   = general_work[0].Int64_Work
    guest_payment.Guest_Name = general_work[0].String_Work

//	fmt.Fprintf( w, "payment_register_excute : guest_payment.Guest_No %v\n", guest_payment.Guest_No )  // 繝・ヰ繝・け
//	fmt.Fprintf( w, "payment_register_excute : guest_payment.Guest_Name %v\n", guest_payment.Guest_Name )  // 繝・ヰ繝・け

// 遨ｺ繧､繝ｳ繧ｿ繝ｼ繝輔ぉ繧､繧ｹ螟画焚繧医ｊ繝舌Μ繝･繝ｼ蛟､繧偵ご繝・ヨ

//    value, _ := count.(int64)

//	fmt.Fprintf( w, "payment_register_excute count %v   \n" , count  )  // 繝・ヰ繝・け
//	fmt.Fprintf( w, "payment_register_excute district_no %v   \n" , district_no  )  // 繝・ヰ繝・け

    guest_payment.Date   = r.FormValue("date")

    guest_payment.Item   = r.FormValue("item")

	amount               := r.FormValue("amount")
	amountw ,err := strconv.Atoi(amount)  // 譁・ｭ励・謨ｴ謨ｰ蛹・	if err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)

		return
	}

	guest_payment.Amount = int64(amountw)   // 謨ｴ謨ｰ縺ｮ64繝薙ャ繝亥喧

//	fmt.Fprintf( w, "payment_register_excute : guest_payment.Item %v\n", guest_payment.Item )  // 繝・ヰ繝・け
//	fmt.Fprintf( w, "payment_register_excute : guest_payment.Amount %v\n", guest_payment.Amount )  // 繝・ヰ繝・け

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

    new_key := datastore.IncompleteKey("Guest_Payment", nil)

    if _, err = client.Put(ctx, new_key, &guest_payment ); err != nil {
//	if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Guest_Payment", nil), &guest_payment); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

/// 繝｢繝九ち繝ｼ縲蜀崎｡ｨ遉ｺ ///

	process4.Payment_register(w , r ,guest_payment.Guest_No)

//	fmt.Fprintf( w, "payment_register_excute : normal end \n" )  // 繝・ヰ繝・け

}
