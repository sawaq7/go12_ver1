package pipe_line_ds_update

import (

	    "strconv"

	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/process2"


                                                   )

func Pipe_line_ds_update(w http.ResponseWriter, r *http.Request) {

//	fmt.Fprintf( w, "sky_pipe_line_ds_update start %v\n" )  // 繝・ヰ繝・け

///
///      謖・ｮ壹＠縺溘ョ繝ｼ繧ｿid繧竪ET
///

    updidw , err := strconv.Atoi(r.FormValue("id"))
	if err  != nil {

//	   fmt.Fprintf( w, "sky_pipe_line_ds_update :error updidw %v\n", updidw )  // 繝・ヰ繝・け

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    updid := int64(updidw)

///
///       驟埼＃諠・ｱ縺ｮ螟画峩
///

	process2.Pipe_line_ds_update(w , r ,updid)

///
///        繝｢繝九ち繝ｼ縲蜀崎｡ｨ遉ｺ
///

	process2.Pipe_line_ds_show(w , r )

}
