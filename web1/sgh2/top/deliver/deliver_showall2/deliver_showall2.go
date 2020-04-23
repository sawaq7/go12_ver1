package deliver_showall2

import (

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
/// 　　   show delivery inf.
///

func Deliver_showall2(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "sky deliver_showall2 start \n" )

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

    idw , err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    id := int64(idw)
//    fmt.Fprintf( w, "sky deliver_showall2 : id %v\n", id )


    key := datastore.IDKey("D_Area", id, nil)

    //get area inf in d.s.
    if err := client.Get(ctx, key , &g ) ; err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

/// 地区・エリア惁E��を、temporary-fileに保留　　set district and area inf. in temp.

    g2.District_No   = g.District_No
    g2.District_Name = g.District_Name
    g2.Area_No       = g.Area_No
	g2.Area_Name     = g.Area_Name
	g2.Number_Total  = g.Number_Total
	g2.Time_Total    = g.Time_Total
	g2.Productivity  = g.Productivity

//    fmt.Fprintf( w, "sky deliver_showall2 : g2.District_No %v\n", g2.District_No )
//    fmt.Fprintf( w, "sky deliver_showall2 : g2.District_Name  %v\n", g2.District_Name )
//    fmt.Fprintf( w, "sky deliver_showall2 : g2.Area_No %v\n", g2.Area_No )
//    fmt.Fprintf( w, "sky deliver_showall2 : g2.Area_Name %v\n", g2.Area_Name )

//    initialize.D_area_temp (w , r )

//  clear existing temp.-file

    _ = datastore2.Datastore_sgh( "D_Area_Temp" ,"initialize" ,idmy , w , r  )

///    make couse-no

    course_no := g2.District_No * 100 + g2.Area_No
    g2.Course_No =  course_no
//    fmt.Fprintf( w, "sky deliver_showall2 : g2.Area_Name %v\n", g2.Area_Name )

///
/// チE�EタストアにセチE��　
///

    new_key := datastore.IncompleteKey("D_Area_Temp", nil)

    if _, err = client.Put(ctx, new_key, &g2 ); err != nil {

		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

    //    show delivery inf.  on web
	process.Deliver_showall2(course_no ,w , r )

//	fmt.Fprintf( w, "sky deliver_showall2 : normal end \n" )

}
