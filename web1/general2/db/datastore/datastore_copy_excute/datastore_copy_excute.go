package datastore_copy_excute

import (

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

//    fmt.Fprintf( w, "datastore_copy_excute start \n" )

///
///   set error message
///

var (

	  Err1 = errors.New("can't find this datastore's file (/datastore_copy_excute)")

	                                                                        )

var g type5.Ds_Copy_List   // allocate work area for d.s. copy list

    copyidw , err := strconv.Atoi(r.FormValue("id"))

	if err  != nil {
//	   fmt.Fprintf( w, "datastore_copy_excute :error copyidw %v\n", copyidw )
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}

    copyid := int64(copyidw)

///
///   get project name
///
    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {
//      fmt.Fprintf( w, "datastore_copy_excute :  projectID unset \n"  )

      project_name = "sample-7777"

	}

    ctx := context.Background()

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w,err.Error(), http.StatusInternalServerError)
		return
    }

///
///       get copy list inf.
///

    key := datastore.IDKey("Ds_Copy_List", copyid, nil)

    if err := client.Get(ctx, key , &g ) ; err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

//    fmt.Fprintf( w, "datastore_copy_excut g.Basic_Name %v\n" ,g.Basic_Name)
//	fmt.Fprintf( w, "datastore_copy_excut g.New_Name %v\n" ,g.New_Name)

///
///    copy and make new file　
///

    switch g.Basic_Name {   //   select by file　name

      case "Deliver" :

        copy3.Deliver( w , r  ,g.Basic_Name ,g.Copy_Name ,g.New_Name )

      case "D_District" :

        copy3.D_district( w , r  ,g.Basic_Name ,g.Copy_Name ,g.New_Name )

      default :                          //   other

        http.Error(w, Err1.Error(), http.StatusInternalServerError)
        return

    }

//	fmt.Fprintf( w, "datastore_copy_excute normal end \n" )

}
