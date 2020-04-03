package process3

import (

	    "net/http"
//	    "fmt"
	    "html/template"
	    "github.com/sawaq7/go12_ver1/general/html5"
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "github.com/sawaq7/go12_ver1/general/type5"

	    "storage2"
	    "strconv"

	    "os"
	    "log"
        "cloud.google.com/go/datastore"
        "context"
                                                )

func Storage_object_rename_keyin(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "process3.storage_object_rename_keyin start \n" )  // 繝・ヰ繝・け

    var sdmy string

    var storage_b_o_temp2 type5.Storage_B_O_Temp

    line_no := r.FormValue("line_no")

//    fmt.Fprintf( w, "storage_object_show : line_no %v\n", line_no )  // 繝・ヰ繝・け

	select_id ,_ := strconv.Atoi(line_no)

//    fmt.Fprintf( w, "storage_object_show : select_id %v\n", select_id )  // 繝・ヰ繝・け

///
///   繝励Ο繧ｸ繧ｧ繧ｯ繝亥錐繧偵ご繝・ヨ
///
     projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {
//      fmt.Fprintf( w, "storage_bucket_list :  projectID unset \n"  )  // 繝・ヰ繝・け

      projectID = "sample-7777"

	}
    ctx := context.Background()
//    c := appengine.NewContext(r)

    query := datastore.NewQuery("Storage_B_O_Temp")

    client, err := datastore.NewClient(ctx, projectID)
    if err != nil {
       log.Fatalf("Failed to create client: %v", err)
    }

//	q := datastore.NewQuery("Storage_B_O_Temp")

//	count, err := q.Count(c)
	count, err := client.Count(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	storage_b_o_temp      := make([]type5.Storage_B_O_Temp, 0, count)

    if _, err = client.GetAll(ctx, query , &storage_b_o_temp) ; err != nil {
//	if keys, err := q.GetAll(c, &storage_b_o_temp);  err != nil {

	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return

	} else{

      for pos, storage_b_o_tempw := range storage_b_o_temp {

        if pos == 0 {

           storage_b_o_temp2.Line_No = 1
           storage_b_o_temp2.Project_Name = storage_b_o_tempw.Project_Name
           storage_b_o_temp2.Bucket_Name = storage_b_o_tempw.Bucket_Name

           updid := datastore.IDKey( "Storage_B_O_Temp",int64(pos+1) , nil)

           objects :=  storage2.Object_List ( w  ,r , storage_b_o_tempw.Bucket_Name )

           for pos2 , objectsw := range objects {

             if pos2+1 == select_id {

                storage_b_o_temp2.Object_Name = objectsw

             }

           }

           if _, err := client.Put(ctx, updid, &storage_b_o_temp2 ); err != nil {

             log.Fatalf("Failed to save task: %v", err)
             return
           }
//           updid = keys[pos].IntID()
//           storage_b_o_temp2.Line_No = 1
//           storage_b_o_temp2.Project_Name = storage_b_o_tempw.Project_Name
//           storage_b_o_temp2.Bucket_Name = storage_b_o_tempw.Bucket_Name

//           bucket    = storage_b_o_tempw.Bucket_Name    // 繝舌こ繝・ヨ蜷阪ｒ繧ｲ繝・ヨ

        }
	  }
	}

///
/// 繝・・繧ｿ繧ｹ繝医い縺ｮ縲・繝ｬ繧ｳ繝ｼ繝峨い繝・・繝・・繝・///

//    key := datastore.NewKey(c, "Storage_B_O_Temp", "", updid, nil)

//	if _, err := datastore.Put(c, key, &storage_b_o_temp2); err != nil {
//		http.Error(w,err.Error(), http.StatusInternalServerError)
//		return
//	}

///
///    繝・Φ繝励Ξ繝ｼ繝医・繝倥ャ繝繝ｼ繧竪ET
///

     monitor := template.Must(template.New("html").Parse(html5.Storage_object_rename_keyin))

///
///       繝｢繝九ち繝ｼ縺ｫ陦ｨ遉ｺ
///

	err = monitor.Execute(w, sdmy)
	if err != nil {
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
}

