package private_showall2

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/sgh/process"
	    "github.com/sawaq7/go12_ver1/client/sgh/type2"
	    "strconv"

        "cloud.google.com/go/datastore"
	    "context"
	    "os"

                                                  )

func Private_showall2(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "private_showall2 start \n" )  // 繝・ヰ繝・け

	var private type2.Private

	private.Worker_Name = r.FormValue("worker_name")  /// 蛟倶ｺｺ蜷阪ｒ繧ｲ繝・ヨ
//	fmt.Fprintf( w, "private_showall2 : worker_name %v\n", private.Worker_Name )  // 繝・ヰ繝・け

	worker_no := r.FormValue("worker_no")             /// 蛟倶ｺｺNo.繧偵ご繝・ヨ
//	fmt.Fprintf( w, "private_showall2 : worker_no %v\n", worker_no )  // 繝・ヰ繝・け


	worker_now ,err := strconv.Atoi(worker_no)           // 譁・ｭ励・謨ｴ謨ｰ蛹・	if err != nil {

//       fmt.Fprintf( w, "private_showall2 : a number must be half-width characters %v\n"  )
		return
	}

	private.Worker_No = int64(worker_now)                // 謨ｴ謨ｰ縺ｮ64繝薙ャ繝亥喧

	private.Worker_Type = r.FormValue("worker_type")   /// 繝ｯ繝ｼ繧ｫ繝ｼ繧ｿ繧､繝励ｒ繧ｲ繝・ヨ

	worker_salary_str  := r.FormValue("worker_salary") /// 繝ｯ繝ｼ繧ｫ繝ｼ繧ｵ繝ｩ繝ｪ繝ｼ繧偵ご繝・ヨ

	private.Worker_Salary , _ = strconv.ParseFloat( worker_salary_str,64 )  // 繝ｯ繝ｼ繧ｫ繝ｼ繧ｵ繝ｩ繝ｪ繝ｼ繧断loat64縺ｫ螟画鋤

	private.Worker_Twh  = 50.0 * 52.14                 /// 蟷ｴ髢鍋ｷ丞感蜒肴凾髢薙ｒ險育ｮ・
	private.Worker_H_Pay  = private.Worker_Salary * 10000. / private.Worker_Twh  /// 譎らｵｦ繧定ｨ育ｮ励

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {

      project_name = "sample-7777"

	}

//	c := appengine.NewContext(r)
    ctx := context.Background()

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w,err.Error(), http.StatusInternalServerError)
		return
    }



/// 繝・・繧ｿ繧ｹ繝医い繝ｼ縺ｫ繝・・繧ｿ繧偵そ繝・ヨ ///

    new_key := datastore.IncompleteKey("Private", nil)

    if _, err = client.Put(ctx, new_key, &private ); err != nil {
//	if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Private", nil), &private); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

/// 繝｢繝九ち繝ｼ縲蜀崎｡ｨ遉ｺ ///

	process.Private_showall1(w , r )

//	fmt.Fprintf( w, "private_showall2 : normal end \n" )  // 繝・ヰ繝・け




}
