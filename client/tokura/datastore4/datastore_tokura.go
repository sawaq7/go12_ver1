package datastore4

import (

        "net/http"
//	    "fmt"

        "client/sgh/datastore2/check"
        "client/sgh/datastore2/initialize"

	    "client/sgh/datastore2/trans"
	    "client/tokura/datastore4/trans2"
	    "client/tokura/datastore4/initialize2"
//	    "general/type5"

                                                )

///
///     クライアントのデータストアのアクセスのメインルーチン （　for tokura　)
///     Datastore_client


func Datastore_tokura( fname string ,function string ,flexible_in interface{} ,w http.ResponseWriter, r *http.Request )  (flexible_out interface{} ) {

//     IN    fname       : データストアのファイル名
//     IN    function    : ファンクション　
//        　　　　　　　　　＊　trans ,check ,initialize ,sort　etc
//     IN flexible_in　  : 各種インプットデータ　（別紙参照）
//     IN    w      　　 : レスポンスライター
//     IN    r      　　 : リクエストパラメータ

//     out flexible_out  : 各種アウトプットデータ　（別紙参照）

//    fmt.Fprintf( w, "datastore_tokura start \n" )  // デバック
//    fmt.Fprintf( w, "datastore_tokura function \n" ,function )  // デバック
//    fmt.Fprintf( w, "datastore_tokura fname \n" ,fname )  // デバック

///
///  ファンクション・ファイル名により各種処理を分岐して行う。
///

	switch function {

///
///  トランスの場合
///

      case "trans" :     // trans の場合

        switch fname {

//          case "Deliver" :  // 該当する配達情報をゲット

            // 空インターフェイス変数よりバリュー値をゲット

//            value, _ := flexible_in.([]type5.General_Work)
//            flexible_out = trans.Deliver ( value[0].Int64_Work  ,value[1].Int64_Work  ,w ,r )

//          break;

          case "Water_Line" :     // 水路ラインファイルの場合

            value, _ := flexible_in.(string)

            flexible_out = trans2.Water_line ( 1  ,value , w ,r )

          break;

          case "Water_Slope" :     // 導水勾配線ファイルの場合

            flexible_out = trans2.Water_slope ( w ,r )

          break;

          case "Water2" :        // 水路ファイル

            flexible_out = trans2.Water2 ( w ,r )

          break;

          case "Car" :  // 号車情報をゲット

            value, _ := flexible_in.(int64)
            flexible_out = trans.Car_district ( value ,w ,r )

          break;

          case "Sgh_Ai" :  // AI情報をゲット

            value, _ := flexible_in.(int64)
            flexible_out = trans.Sgh_ai( value  ,w , r  )

          break;

        }
      break;

      case "trans2" :       // trans2 の場合

        switch fname {

          case "Deliver" :  // 該当する配達情報をゲット



          break;

          case "D_Area" :  // 号車情報をゲット

            value, _ := flexible_in.(int64)
            flexible_out = trans.D_area_district ( w ,r ,value )


          break;

          case "Private" :  // workerの個人情報をゲット



          break;

          case "Car" :  // 号車情報をゲット



          break;

          case "D_District_Temp" :     // 地区情報の一時ファイルの場合



          break;
        }
      break;

///
///  初期化の場合
///

      case "initialize" :

        switch fname {

          case "Water2_Temp" :   // 水路ファイルの場合　

            initialize2.Water2_temp (w , r )   // 水路情報をイニシャライズ  & セット　

          break;

          case "D_District_Temp" :     // エリア情報の場合

            initialize.D_district_temp (w , r ) // temporary-fileをイニシャライズ

          break;

          case "Sgh_Ai" :     // エリア情報の場合

            value, _ := flexible_in.(int64)
            initialize.Sgh_ai( value ,w , r )  //   既存のAI情報を削除

          break;

        }

      break;


///
///  チェックの場合
///

      case "check" :

        switch fname {

          case "D_Area" :   // 配達地区のエリアのMAX値をゲット

            // 空インターフェイス変数よりバリュー値をゲット

            value, _ := flexible_in.(int64)
            flexible_out = check.D_area ( w , r  ,value )

          break;

          case "D_District_Temp" :   // 一時ファイルよりカレントの地区の情報をゲット　

            flexible_out = check.D_district_temp (w , r )

          break;

          case "Car" :   // 配達地区カーNOのMAX値をゲット

            // 空インターフェイス変数よりバリュー値をゲット

            value, _ := flexible_in.(int64)
            flexible_out = check.Car_no_max(w , r  ,value)

          break;

        }

      break;

    }

	return flexible_out

}

