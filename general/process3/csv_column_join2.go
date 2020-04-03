package process3

import (

//        "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"

	    "github.com/sawaq7/go12_ver1/general/datastore5/reformat"

	    "github.com/sawaq7/go12_ver1/general/datastore5/trans3"
	    "github.com/sawaq7/go12_ver1/general/datastore5/set1"

	    "cloud.google.com/go/datastore"
        "context"
        "os"
                                                )

///
///     1åãEãEEã¿ãè¿½å ããEEolumn/move)
///

func Csv_column_join2(w http.ResponseWriter, r *http.Request , column_no1 int ,column_no2 int ) {

//     IN    w      ãã: ã¬ã¹ãã³ã¹ã©ã¤ã¿ã¼
//     IN    r      ãã: ãªã¯ã¨ã¹ããã©ã¡ã¼ã¿
//     IN  column_no1ã : åé¤äºå®ãEè¡NO
//     IN  column_no2ã : è¿½å ããè¡NO

//    fmt.Fprintf( w, "process3.csv_column_join2 start \n" )  // ãEãE¯

///
///   ãã­ã¸ã§ã¯ãåãã²ãE
///

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {
//      fmt.Fprintf( w, "storage_bucket_list :  projectID unset \n"  )  // ãEãE¯

      project_name = "sample-7777"

	}

//    c := appengine.NewContext(r)
    ctx := context.Background()

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w,err.Error(), http.StatusInternalServerError)
		return
    }

///
/// ãããcsvæE ±ã®ãã©ã¼ããããæ¡å¼µããã
///

	csv_inf := trans3.Csv_inf ( w ,r )  ///      csvæE ±ãã²ãE

    csv_inf2 := reformat.Csv_inf ( 1 ,int64(column_no2) ,csv_inf ,w ,r )
                                                           /// ãã©ã¼ãããã1åæ¡å¼µãã
///
///      csvæE ±ãè¿½å ãã
///

    csv_inf_join := trans3.Csv_inf_column ( w ,r ,column_no1 )         /// è¿½å ããcsvæE ±ãã²ãE

    csv_inf_new := set1.Csv_inf (  csv_inf2 ,csv_inf_join ,column_no2 , w ,r )
                                                        /// è¿½å ãããEEã¿1åãã»ãE

///
/// ããããEEã¿ã¹ãã¢ã«ãcsvæE ±ãåEã»ãEããã
///

    for _, csv_inf_neww := range csv_inf_new {

//   	  fmt.Fprintf( w, "process3.csv_column_join2 csv_inf_neww %v\n", csv_inf_neww )  // ãEãE¯

//      key := datastore.NewKey(c, "Csv_Inf", "", csv_inf_neww.Id, nil)  //ãã¢ã¯ã»ã¹ã­ã¼ã²ãE
      key := datastore.IDKey("Csv_Inf", csv_inf_neww.Id, nil)

      if _, err := client.Put(ctx, key, &csv_inf_neww ); err != nil {
//      if _, err := datastore.Put(c, key, &csv_inf_neww); err != nil {  // ãEEã¿ã¹ãã¢ã«åã»ãE
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	  }

    }

///
/// ãããwebä¸ã«ãcsvæE ±ãè¡¨ç¤ºããã
///

//    monitor := template.Must( template.New("html").Parse( html5.Csv_show )) // ãE³ãã¬ã¼ããEãããã¼ãGET

//     err = monitor.Execute ( w, csv_inf_new )
//	 if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	 }
//	fmt.Fprintf( w, "process3.csv_column_join2 normal end \n" )  // ãEãE¯


}

