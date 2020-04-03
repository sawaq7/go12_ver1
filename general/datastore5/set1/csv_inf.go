package set1

import (

	    "net/http"
//	    "fmt"
//	    "html/template"

	    "github.com/sawaq7/go12_ver1/general/type5"

//	    "time"
                                                )

///
/// csv繝輔ぃ繧､繝ｫ縺ｫ蛻怜腰菴阪〒繝・・繧ｿ繧偵そ繝・ヨ縺吶ｋ
///

func Csv_inf ( csv_inf []type5.Csv_Inf ,csv_inf_join []string ,column_no int ,w http.ResponseWriter, r *http.Request )  ([]type5.Csv_Inf ) {

//     IN  csv_inf  縲縲: csv諠・ｱ窶昴・繧ｹ繝ｩ繧､繧ｹ
//     IN  csv_inf_join : 霑ｽ蜉縺吶ｋ蛻玲ュ蝣ｱ縲縲縲縲
//     IN  column_no  縲縲: 蟇ｾ雎｡縺ｮ陦・//     IN    w      縲縲: 繝ｬ繧ｹ繝昴Φ繧ｹ繝ｩ繧､繧ｿ繝ｼ
//     IN    r      縲縲: 繝ｪ繧ｯ繧ｨ繧ｹ繝医ヱ繝ｩ繝｡繝ｼ繧ｿ

//     OUT csv_inf_new  : 繝・・繧ｿ霑ｽ蜉蠕後・csv諠・ｱ窶昴・繧ｹ繝ｩ繧､繧ｹ

//    fmt.Fprintf( w, "set1.csv_inf start %v\n" )  // 繝・ヰ繝・け
//    fmt.Fprintf( w, "set1.csv_inf csv_inf_join %v\n" ,csv_inf_join )  // 繝・ヰ繝・け
//    fmt.Fprintf( w, "set1.csv_inf column_no %v\n" ,column_no)  // 繝・ヰ繝・け

    str_work := make([]string ,10 )  // 繝ｯ繝ｼ繧ｯ繧ｨ繝ｪ繧｢繧堤｢ｺ菫・
	csv_inf2 := make([]type5.Csv_Inf, 0)

	index := 0

	for _, csv_infw := range csv_inf {

//	  fmt.Fprintf( w, "trans.csv_inf csv_infw %v\n" ,csv_infw)  // 繝・ヰ繝・け

      str_work[0]  = csv_infw.Column1
	  str_work[1]  = csv_infw.Column2
	  str_work[2]  = csv_infw.Column3
	  str_work[3]  = csv_infw.Column4
	  str_work[4]  = csv_infw.Column5
	  str_work[5]  = csv_infw.Column6
	  str_work[6]  = csv_infw.Column7
	  str_work[7]  = csv_infw.Column8
	  str_work[8]  = csv_infw.Column9
	  str_work[9]  = csv_infw.Column10

	  str_work[column_no-1]  = csv_inf_join[index]

	  index ++

      csv_inf2 = append(csv_inf2, type5.Csv_Inf { csv_infw.Id           ,
                                                     csv_infw.Line_No      ,
                                                     csv_infw.File_Name    ,
                                                     csv_infw.Column_Num   ,
                                                     str_work[0]           ,
                                                     str_work[1]           ,
                                                     str_work[2]           ,
                                                     str_work[3]           ,
                                                     str_work[4]           ,
                                                     str_work[5]           ,
                                                     str_work[6]           ,
                                                     str_work[7]           ,
                                                     str_work[8]           ,
                                                     str_work[9]           })




	}

//    fmt.Fprintf( w, "set1.csv_inf csv_inf2 %v\n" ,csv_inf2)  // 繝・ヰ繝・け

    return	csv_inf2
}

