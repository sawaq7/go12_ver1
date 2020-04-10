package storage_bucket_list


import (
	     "net/http"
	     "html/template"

         "github.com/sawaq7/go12_ver1/general/type5"
         "github.com/sawaq7/go12_ver1/general/html5"
         "github.com/sawaq7/go12_ver1/storage2"

         "time"
         "os"
//         "fmt"
//         "context"
//         "cloud.google.com/go/storage"

	                                    )

func Storage_bucket_list(w http.ResponseWriter, r *http.Request) {

   var  sdmy string

   var t_dmy   time.Time

   projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

   if projectID == "" {
//      fmt.Fprintf( w, "storage_bucket_list :  projectID unset \n"  )  // チE��チE��

      projectID = "sample-7777"

   }


   storage_b_o_view := make([]type5.Storage_B_O_View, 0)

//   buckets := make([]string, 2)
//   buckets[0] = "sample-7777"
//   buckets[1] = "filesample3"

//	buckets_minor , _ := storage2.Storage_basic( "list" ,projectID ,sdmy, w , r  )

//    buckets, _ := buckets_minor.([]string)  // インターフェイス型を型変換

//    ctx := context.Background()
//    client, _ := storage.NewClient(ctx)

	buckets :=  storage2.Bucket_List ( w  ,r , projectID )
//	buckets :=  storage2.Bucket_List2 ( w  ,r , client , projectID )

//	fmt.Fprintf( w, "buckets: %v\n", buckets)

   for pos , bucketsw := range buckets {

	  storage_b_o_view = append(storage_b_o_view, type5.Storage_B_O_View { int64(pos+1) ,
                                                                           projectID    ,
                                                                           bucketsw     ,
                                                                           sdmy         ,
                                                                           t_dmy          })

//      fmt.Fprintf( w, "pos: %v\n", pos )

	}

// set template

    monitor := template.Must(template.New("html").Parse(html5.Storage_bucket_list))


//  show web


//    var str_dmy string
    monitor.Execute( w, storage_b_o_view )

//	monitor.Execute( w, str_dmy )

}
