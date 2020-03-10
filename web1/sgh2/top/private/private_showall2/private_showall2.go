package private_showall2

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
	    "client/sgh/process"
	    "client/sgh/type2"
	    "strconv"

        "cloud.google.com/go/datastore"
	    "context"
	    "os"

                                                  )

func Private_showall2(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "private_showall2 start \n" )  // デバック

	var private type2.Private

	private.Worker_Name = r.FormValue("worker_name")  /// 個人名をゲット
//	fmt.Fprintf( w, "private_showall2 : worker_name %v\n", private.Worker_Name )  // デバック

	worker_no := r.FormValue("worker_no")             /// 個人No.をゲット
//	fmt.Fprintf( w, "private_showall2 : worker_no %v\n", worker_no )  // デバック


	worker_now ,err := strconv.Atoi(worker_no)           // 文字の整数化
	if err != nil {

//       fmt.Fprintf( w, "private_showall2 : a number must be half-width characters %v\n"  )
		return
	}

	private.Worker_No = int64(worker_now)                // 整数の64ビット化

	private.Worker_Type = r.FormValue("worker_type")   /// ワーカータイプをゲット

	worker_salary_str  := r.FormValue("worker_salary") /// ワーカーサラリーをゲット

	private.Worker_Salary , _ = strconv.ParseFloat( worker_salary_str,64 )  // ワーカーサラリーをfloat64に変換

	private.Worker_Twh  = 50.0 * 52.14                 /// 年間総労働時間を計算

	private.Worker_H_Pay  = private.Worker_Salary * 10000. / private.Worker_Twh  /// 時給を計算　

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



/// データストアーにデータをセット ///

    new_key := datastore.IncompleteKey("Private", nil)

    if _, err = client.Put(ctx, new_key, &private ); err != nil {
//	if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Private", nil), &private); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

/// モニター　再表示 ///

	process.Private_showall1(w , r )

//	fmt.Fprintf( w, "private_showall2 : normal end \n" )  // デバック




}
