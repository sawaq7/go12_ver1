package trans

import (
	    "google.golang.org/appengine"
	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
//	    "html/template"

	    "client/sgh/type2"
//	    "time"
                                                )

///
/// 該当するコースNOのエリア情報をゲットする
///

func D_area_course_no(w http.ResponseWriter, r *http.Request ,course_no  int64)  (d_area_course_no []type2.D_Area ) {

//     IN    w      　　     : レスポンスライター
//     IN    r      　　     : リクエストパラメータ
//     IN 　course_no        : コースNo
//     OUT d_area_course_no  : 該当するコースNOのエリア情報　(構造体　”エリア情報”のスライス)

//    fmt.Fprintf( w, "trans.d_area_course_no start \n" )  // デバック

	c := appengine.NewContext(r)

	q := datastore.NewQuery("D_Area").Order("District_No")

	count, err := q.Count(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	d_area      := make([]type2.D_Area, 0, count)
	d_area_course_no      = make([]type2.D_Area, 0, count)


	if keys, err := q.GetAll(c, &d_area); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	} else {

      pos := 0

	  for _, d_areaw := range d_area {

        if course_no == d_areaw.Course_No &&
           pos       == 0                    {

          d_area_course_no = append( d_area_course_no, type2.D_Area { keys[pos].IntID()     ,
                                                          d_areaw.Course_No      ,
                                                          d_areaw.District_No    ,
                                                          d_areaw.District_Name  ,
                                                          d_areaw.Area_No        ,
                                                          d_areaw.Area_Name      ,
                                                          d_areaw.Area_Detail    ,
                                                          d_areaw.Number_Total   ,
                                                          d_areaw.Time_Total     ,
                                                          d_areaw.Productivity   ,
                                                          d_areaw.Car_No           })

          pos ++

        }
	  }
	}
//    fmt.Fprintf( w, "trans.d_area_course_no normal end \n" )  // デバック

	return d_area_course_no
}
