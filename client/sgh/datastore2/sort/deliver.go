package sort

import (

	    "net/http"
//	    "fmt"

	    "github.com/sawaq7/go12_ver1/client/sgh/type2"

                                                )

///
/// 驟埼＃諠・ｱ繧偵・驥行ort縺吶ｋ
///           key1 : Date  , key2 : Car_No


func Deliver(w http.ResponseWriter ,deliver []type2.Deliver  )  (deliver2 []type2.Deliver ) {

//     IN    w      縲縲: 繝ｬ繧ｹ繝昴Φ繧ｹ繝ｩ繧､繧ｿ繝ｼ
//     IN    r      縲縲: 繝ｪ繧ｯ繧ｨ繧ｹ繝医ヱ繝ｩ繝｡繝ｼ繧ｿ
//     IN 縲deliver     : 讒矩菴薙窶晞・驕疲ュ蝣ｱ窶昴・繧ｹ繝ｩ繧､繧ｹ

//     OUT  deliver2    : 繧ｽ繝ｼ繝亥ｾ後・讒矩菴薙窶晞・驕疲ュ蝣ｱ窶昴・繧ｹ繝ｩ繧､繧ｹ

    var j_min , j_flag   int
    var carno_save ,line_counter int64
    var date_save  string

//    fmt.Fprintf( w, "sort.deliver start \n" )  // 繝・ヰ繝・け

	count := len(deliver)

	deliver2 = make([]type2.Deliver, 0)  /// 繧ｽ繝ｼ繝医ユ繝ｼ繝悶Ν繧堤｢ｺ菫・
	skip_check := make([]int ,count)        /// 繧ｹ繧ｭ繝・・縺ｮ蛻､螳壹ヵ繝ｩ繧ｰ繝・・繝悶Ν繧堤｢ｺ菫・
    line_counter = 0

	for  i := 0 ; i < count ; i++  {

	  j_flag = -1

	  for  j := 0 ; j < count ; j++ {



	    if skip_check[j] != 1  {  /// 繧ｽ繝ｼ繝域ｸ医∩縺句凄縺九・繝√ぉ繝・け

	      if j_flag  == -1  {     ///  蛻晄悄蛟､縺ｮ繧ｻ繝・ヨ

	        j_min = j
	        date_save  = deliver[j].Date
	        carno_save = deliver[j].Car_No
	        j_flag = 0

	        /// 譛蟆丞､縲∝・繧ｻ繝・ヨ

	      }  else if ( date_save >  deliver[j].Date )                                   ||
	                 ( date_save == deliver[j].Date && carno_save > deliver[j].Car_No )     {

	        j_min = j
	        date_save  = deliver[j].Date
	        carno_save = deliver[j].Car_No

	      }
	    }
      }

/// min蛟､繧偵ユ繝ｼ繝悶Ν縺ｫ繧ｻ繝・ヨ
      line_counter ++
      deliver[j_min].Line_No = line_counter

      deliver2 = append ( deliver2,  deliver[j_min] )

/// 繧ｽ繝ｼ繝域ｸ医ヵ繝ｩ繧ｰ繧偵そ繝・ヨ

      skip_check[j_min] = 1

	}

    return	deliver2
}

