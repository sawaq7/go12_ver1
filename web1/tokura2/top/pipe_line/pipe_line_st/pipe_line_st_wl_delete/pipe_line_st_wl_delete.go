package pipe_line_st_wl_delete

import (

	"net/http"
	"strconv"
//	"fmt"
	"github.com/sawaq7/go12_ver1/client/tokura/suiri/process2"
	"github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
//	"github.com/sawaq7/go12_ver1/client/tokura/storage3/trans4"
	"github.com/sawaq7/go12_ver1/client/tokura/storage3"

                                            )

func Pipe_line_st_wl_delete(w http.ResponseWriter, r *http.Request) {

//   fmt.Fprintf( w, "pipe_line_st_wl_delete start \n" )  // 繝・ヰ繝・け

    var idmy1 ,idmy2 int64

///
///       繧ｫ繝ｬ繝ｳ繝医・豌ｴ霍ｯ諠・ｱ繧偵ご繝・ヨ
///

      water2_temp_minor , _ := storage3.Storage_tokura( "Water2_Temp" ,"trans" ,idmy1 , idmy2 , w , r  )

      water2_temp, _ := water2_temp_minor.([]type4.Water2_Temp)  // 繧､繝ｳ繧ｿ繝ｼ繝輔ぉ繧､繧ｹ蝙九ｒ蝙句､画鋤

//    water2_temp := trans4.Water2_temp( w , r  )

///
///       謖・､ｺ縺励◆id繧竪ET
///

    id := r.FormValue("id")
//    fmt.Fprintf( w, "pipe_line_st_wl_delete : id %v\n", id )  // 繝・ヰ繝・け

	delidw ,_ := strconv.Atoi(id)
	delid := int64(delidw)

//    fmt.Fprintf( w, "pipe_line_st_wl_delete : delidw %v\n", delidw )  // 繝・ヰ繝・け
//    fmt.Fprintf( w, "pipe_line_st_wl_delete : delid %v\n", delid )  // 繝・ヰ繝・け

///
///     謖・､ｺ縺励◆豌ｴ霍ｯ繝ｩ繧､繝ｳ繝・・繧ｿ繧貞炎髯､
///

      _ , _ = storage3.Storage_tokura( "Water_Line" ,"delete" ,delid , water2_temp[0].Name , w , r  )

//    delete1.Water_line( w , r ,delid , water2_temp[0].Name )

///
///        繝｢繝九ち繝ｼ縲蜀崎｡ｨ遉ｺ
///

	process2.Pipe_line_st_wl_show ( water2_temp[0].Name ,w , r )

}
