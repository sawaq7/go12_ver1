package process3

import (

	    "net/http"
//	    "fmt"
	    "html/template"
	    "github.com/sawaq7/go12_ver1/general/html5"
	    "storage2"
	    "github.com/sawaq7/go12_ver1/general/type5"
//	    "time"
                                                )

///
/// ζE€Ίγγγγ±γEεEEγͺγγΈγ§γ―γγγγ¦γ¨γδΈγ«θ‘¨η€ΊγγγE///

func Storage_object_show ( w http.ResponseWriter, r *http.Request ,project string ,bucket string ) {

//     IN    w      : γ¬γΉγγ³γΉγ©γ€γΏγΌ
//     IN    r      : γͺγ―γ¨γΉγγγ©γ‘γΌγΏ
//     IN  project  : γγ­γΈγ§γ―γε
//     IN  bucket   : γγ±γEεE
//    fmt.Fprintf( w, "process3.storage_object_show start \n" )  // γEγE―

//    var t_dmy   time.Time

//    var idmy int64

    storage_b_o_view := make([]type5.Storage_B_O_View, 0) // γγ±γEγͺγΉγγEθ‘¨η€Ίη¨γ¨γͺγ’γη’ΊδΏE
//    objects_minor , _ := storage2.Storage_basic( "list3" ,bucket ,idmy, w , r  )

//    objects, _ := objects_minor.([]type5.Storage_B_O_View)  // γ€γ³γΏγΌγγ§γ€γΉεγεε€ζ

    objects :=  storage2.Object_List_Detail ( w  ,r , bucket )  // ζE€Ίγγγγ±γEεEEγͺγγΈγ§γ―γζε ±γγ²γEγγ

    for pos , objectsw := range objects {

//      fmt.Fprintf( w, "storage_bucket_list : pos: %v\n", pos )

      storage_b_o_view = append(storage_b_o_view, type5.Storage_B_O_View { int64(pos+1)            ,
                                                                              project              ,
                                                                              bucket               ,
                                                                              objectsw.Object_Name ,
                                                                              objectsw.Created       })

    }

    monitor := template.Must(template.New("html").Parse(html5.Storage_object_list))

	err := monitor.Execute( w, storage_b_o_view )
	if err != nil {
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}

//	fmt.Fprintf( w, "process3.storage_object_show normal end \n" )  // γEγE―
}

