package datastore_copy_list_show

import (

	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/general/process3"

	    "github.com/sawaq7/go12_ver1/general/type5"
//	    "strconv"
//	    "time"

        "cloud.google.com/go/datastore"
        "context"
        "os"
                                                  )

func Datastore_copy_list_show(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "datastore_copy_list_show start \n" )

    var g type5.Ds_Copy_List

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {
//      fmt.Fprintf( w, "storage_bucket_list :  projectID unset \n"  )

      project_name = "sample-7777"

	}

    ctx := context.Background()

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
       return
    }

	g.Basic_Name = r.FormValue("basic_name")  //  get d.s. name of basic
	g.Copy_Name  = r.FormValue("copy_file")   //  get d.s. name which is cpied
	g.New_Name   = r.FormValue("new_file")    //  get new d.s. name

                                               //  get new key in d.s.
    new_key := datastore.IncompleteKey("Ds_Copy_List", nil)

//    fmt.Fprintf(w, "datastore_copy_list_show: new_key %v\n", new_key )

    if   _, err = client.Put(ctx, new_key, &g ); err != nil {

		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

//   	fmt.Fprintf( w, "datastore_copy_list_show : g.Basic_Name %v\n", g.Basic_Name )
//    fmt.Fprintf( w, "datastore_copy_list_show : g.Copy_Name %v\n", g.Copy_Name )
//    fmt.Fprintf( w, "datastore_copy_list_show : g.New_Name %v\n", g.New_Name )

///     show csv inf. on web

    process3.Datastore_copy_list_keyin(w , r )

//	fmt.Fprintf( w, "datastore_copy_list_show : normal end \n" )




}
