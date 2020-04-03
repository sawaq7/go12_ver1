package suiri

import (
	    "net/http"
//	    "fmt"
	    "html/template"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"

        "github.com/sawaq7/go12_ver1/client/tokura/html4"
                                      )


func Pipe_line1_show( w http.ResponseWriter , wnumber int , water []type4.Water  ) {

//     IN    w     : 繝ｬ繧ｹ繝昴Φ繧ｹ繝ｩ繧､繧ｿ繝ｼ
//     IN  wnumbwr : 繝昴う繝ｳ繝域錐螟ｱ縺ｮ繧ｹ繝ｩ繧､繧ｹ
//     IN  water   : 豌ｴ霍ｯ諠・ｱ鄒､縺ｮ繧ｹ繝ｩ繧､繧ｹ



//    fmt.Fprintf( w, "suiri/Pipe_line1_show start \n" )  // 繝・ヰ繝・け
//    fmt.Fprintf( w, "suiri/Pipe_line1_show wnumber  %v \n" , wnumber)  // 繝・ヰ繝・け

    water2 := make([]type4.Water, wnumber)

    for pos, waterwk := range water {

			water2[pos].No  = waterwk.No
			water2[pos].Name = waterwk.Name
			water2[pos].High = waterwk.High
            water2[pos].Roughness_factor = waterwk.Roughness_factor
//			fmt.Fprintf( w, "suiri/Pipe_line1_show pos ,NO  ,Name %v %v %v \n" , pos ,water2[pos].No ,water2[pos].Name)  // 繝・ヰ繝・け
//			fmt.Fprintf( w, "suiri/Pipe_line1_show name High ,Roughness_factor %v %v  \n" , water2[pos].High ,water2[pos].Roughness_factor)  // 繝・ヰ繝・け
	}
// 繝・Φ繝励Ξ繝ｼ繝医・繝倥ャ繝繝ｼ繧竪ET

	monitor := template.Must(template.New("html").Parse( html4.Pipe_line1_show))
//	monitor := template.Must(template.New("html").Parse( html4.Pipe_line1_show2))



// 蜷・ｨｮ蜈･蜉帙ョ繝ｼ繧ｿ繧定｡ｨ遉ｺ

	err := monitor.Execute(w, water2)

//    var str_dmy string
//	err := monitor.Execute(w, water)

    if err != nil {
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

