package d_schedule_keyin

import (

	    "net/http"
	    "github.com/sawaq7/go12_ver1/client/sgh/process"
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "github.com/sawaq7/go12_ver1/client/sgh/datastore2"
  	    "strconv"
//	    "fmt"

	    "github.com/sawaq7/go12_ver1/client/sgh/type2"

	    "cloud.google.com/go/datastore"
        "context"
        "os"
                                                 )

///
/// 持E��した地域�Eスケジュールを表示
///

func D_schedule_keyin(w http.ResponseWriter, r *http.Request) {

    var idmy int64

    var g type2.D_District

    var g2 type2.D_District_Temp

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {

      project_name = "sample-7777"

	}

    ctx := context.Background()
//	c := appengine.NewContext(r)

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
       return
    }


/// 地区惁E��ファイルから地区NO をGET

	idw , err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    id := int64(idw)
//    fmt.Fprintf( w, "sky d_schedule_keyin : id %v\n", id )  // チE��チE��

    key := datastore.IDKey("D_District", id, nil)

    if err := client.Get(ctx, key , &g ) ; err != nil {
//    key := datastore.NewKey(c, "D_District", "", id, nil)
//    if err := datastore.Get(c, key ,&g); err!= nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

///
/// カレント�E地区惁E��をアチE�EチE�EチE///

    g2.District_No   = g.District_No
    g2.District_Name = g.District_Name
//    fmt.Fprintf( w, "sky d_schedule_keyin : g2.District_No %v\n", g2.District_No )  // チE��チE��

// temporary-fileをイニシャライズ

    _ = datastore2.Datastore_sgh( "D_District_Temp" ,"initialize" ,idmy , w , r  )

     new_key := datastore.IncompleteKey("D_District_Temp", nil)

    if _, err = client.Put(ctx, new_key, &g2 ); err != nil {
//	if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "D_District_Temp", nil), &g2); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

/// スケジュールを表示 ///

   process.D_schedule_showall(w , r , g.District_No )

}
