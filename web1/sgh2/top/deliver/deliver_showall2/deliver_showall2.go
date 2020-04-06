package deliver_showall2

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/sgh/process"
	    "github.com/sawaq7/go12_ver1/client/sgh/datastore2"
        "github.com/sawaq7/go12_ver1/client/sgh/type2"
	    "strconv"

	    "cloud.google.com/go/datastore"
        "context"
        "os"
                                                  )

///
/// main 配達チE�Eタを表示する
///

func Deliver_showall2(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "sky deliver_showall2 start \n" )  // チE��チE��

    var g type2.D_Area

    var g2 type2.D_Area_Temp

    var idmy int64

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

    idw , err := strconv.Atoi(r.FormValue("id"))     //持E��したエリアidをGET
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    id := int64(idw)
//    fmt.Fprintf( w, "sky deliver_showall2 : id %v\n", id )  // チE��チE��


    key := datastore.IDKey("D_Area", id, nil)
//    key := datastore.NewKey(c, "D_Area", "", id, nil)

/// 持E��した地区惁E��をGET   ///

    if err := client.Get(ctx, key , &g ) ; err != nil {
//    if err := datastore.Get(c, key ,&g); err!= nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

/// 地区・エリア惁E��を、temporary-fileに保留　　　///

    g2.District_No   = g.District_No
    g2.District_Name = g.District_Name
    g2.Area_No       = g.Area_No
	g2.Area_Name     = g.Area_Name
	g2.Number_Total  = g.Number_Total
	g2.Time_Total    = g.Time_Total
	g2.Productivity  = g.Productivity

//    fmt.Fprintf( w, "sky deliver_showall2 : g2.District_No %v\n", g2.District_No )  // チE��チE��
//    fmt.Fprintf( w, "sky deliver_showall2 : g2.District_Name  %v\n", g2.District_Name )  // チE��チE��
//    fmt.Fprintf( w, "sky deliver_showall2 : g2.Area_No %v\n", g2.Area_No )  // チE��チE��
//    fmt.Fprintf( w, "sky deliver_showall2 : g2.Area_Name %v\n", g2.Area_Name )  // チE��チE��

//    initialize.D_area_temp (w , r )

//  既存�E　D_Area_Temp temporary-fileをクリアー

//    _ = datastore2.D_store( "D_Area_Temp" ,"initialize" ,idmy , w , r  )
    _ = datastore2.Datastore_sgh( "D_Area_Temp" ,"initialize" ,idmy , w , r  )

/// コースNOを作�Eする

    course_no := g2.District_No * 100 + g2.Area_No
    g2.Course_No =  course_no
//    fmt.Fprintf( w, "sky deliver_showall2 : g2.Area_Name %v\n", g2.Area_Name )  // チE��チE��

///
/// チE�EタストアにセチE��　
///

    new_key := datastore.IncompleteKey("D_Area_Temp", nil)

    if _, err = client.Put(ctx, new_key, &g2 ); err != nil {
//	if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "D_Area_Temp", nil) ,
//	                                                                  &g2); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

/// モニター　表示 ///

	process.Deliver_showall2(course_no ,w , r )

//	fmt.Fprintf( w, "sky deliver_showall2 : normal end \n" )  // チE��チE��




}