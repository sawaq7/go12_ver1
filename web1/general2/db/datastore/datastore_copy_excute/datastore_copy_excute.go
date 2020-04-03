package datastore_copy_excute

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
	    "strconv"
        "errors"
	    "github.com/sawaq7/go12_ver1/general/datastore5/copy3"
	    "github.com/sawaq7/go12_ver1/general/type5"

	    "os"

        "cloud.google.com/go/datastore"
        "context"

                                                  )

func Datastore_copy_excute(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "datastore_copy_excute start \n" )  // ãEãE¯

///
///  ã¨ã©ã¼ã¡ãE»ã¼ã¸ãã»ãE
///

var (

	  Err1 = errors.New("can't find this datastore's file (/datastore_copy_excute)")

	                                                                        )

var g type5.Ds_Copy_List /// ãEEã¿ã¹ãã¢ã®ã³ããEãªã¹ããEã¯ã¼ã¯ã¨ãªã¢ç¢ºä¿E
///                      ///
/// æE®ãããã¼ã¿idãGET ///
///                      ///

    copyidw , err := strconv.Atoi(r.FormValue("id"))
	if err  != nil {
//	   fmt.Fprintf( w, "datastore_copy_excute :error copyidw %v\n", copyidw )  // ãEãE¯
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    copyid := int64(copyidw)

///
///   ãã­ã¸ã§ã¯ãåãã²ãE
///
    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {
//      fmt.Fprintf( w, "datastore_copy_excute :  projectID unset \n"  )  // ãEãE¯

      project_name = "sample-7777"

	}

//	c := appengine.NewContext(r)
    ctx := context.Background()

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w,err.Error(), http.StatusInternalServerError)
		return
    }

///
///       ã³ããEãªã¹ãæå ±ãGET
///

    key := datastore.IDKey("Ds_Copy_List", copyid, nil)
//    key := datastore.NewKey(c, "Ds_Copy_List", "", copyid, nil)

    if err := client.Get(ctx, key , &g ) ; err != nil {
//	if err := datastore.Get(c, key, &g); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

//    fmt.Fprintf( w, "datastore_copy_excut g.Basic_Name %v\n" ,g.Basic_Name)  // ãEãE¯
//	fmt.Fprintf( w, "datastore_copy_excut g.New_Name %v\n" ,g.New_Name)  // ãEãE¯

///
/// ã³ããEãã¦ããã¥ã¼ãã¡ã¤ã«ãä½æEã///
///


    switch g.Basic_Name {

      case "Deliver" :                   // ééæE ±ãEEgh )

        copy3.Deliver( w , r  ,g.Basic_Name ,g.Copy_Name ,g.New_Name )

      case "D_District" :                // å°åºæE ±ãEEgh )

        copy3.D_district( w , r  ,g.Basic_Name ,g.Copy_Name ,g.New_Name )

      default :                          // è©²å½ãããã¡ã¤ã«ãªãE        http.Error(w, Err1.Error(), http.StatusInternalServerError)
        return

    }
//    copy3.All( w , r  ,g.Basic_Name ,g.Copy_Name ,g.New_Name )
//    if err := copy3.All( w , r  ,g.Basic_Name ,g.Copy_Name ,g.New_Name ); err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}

//	fmt.Fprintf( w, "datastore_copy_excute normal end \n" )  // ãEãE¯
}
