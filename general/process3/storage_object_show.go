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
/// 謖・､ｺ縺励◆繝舌こ繝・ヨ蜀・・繧ｪ繝悶ず繧ｧ繧ｯ繝医ｒ縲√え繧ｨ繝紋ｸ翫↓陦ｨ遉ｺ縺吶ｋ縲・///

func Storage_object_show ( w http.ResponseWriter, r *http.Request ,project string ,bucket string ) {

//     IN    w      : 繝ｬ繧ｹ繝昴Φ繧ｹ繝ｩ繧､繧ｿ繝ｼ
//     IN    r      : 繝ｪ繧ｯ繧ｨ繧ｹ繝医ヱ繝ｩ繝｡繝ｼ繧ｿ
//     IN  project  : 繝励Ο繧ｸ繧ｧ繧ｯ繝亥錐
//     IN  bucket   : 繝舌こ繝・ヨ蜷・
//    fmt.Fprintf( w, "process3.storage_object_show start \n" )  // 繝・ヰ繝・け

//    var t_dmy   time.Time

//    var idmy int64

    storage_b_o_view := make([]type5.Storage_B_O_View, 0) // 繝舌こ繝・ヨ繝ｪ繧ｹ繝医・陦ｨ遉ｺ逕ｨ繧ｨ繝ｪ繧｢繧堤｢ｺ菫・
//    objects_minor , _ := storage2.Storage_basic( "list3" ,bucket ,idmy, w , r  )

//    objects, _ := objects_minor.([]type5.Storage_B_O_View)  // 繧､繝ｳ繧ｿ繝ｼ繝輔ぉ繧､繧ｹ蝙九ｒ蝙句､画鋤

    objects :=  storage2.Object_List_Detail ( w  ,r , bucket )  // 謖・､ｺ縺励◆繝舌こ繝・ヨ蜀・・繧ｪ繝悶ず繧ｧ繧ｯ繝域ュ蝣ｱ繧偵ご繝・ヨ縺吶ｋ

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

//	fmt.Fprintf( w, "process3.storage_object_show normal end \n" )  // 繝・ヰ繝・け
}

