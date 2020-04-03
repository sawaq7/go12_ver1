package datastore4

import (

        "net/http"
//	    "fmt"

        "github.com/sawaq7/go12_ver1/client/sgh/datastore2/check"
        "github.com/sawaq7/go12_ver1/client/sgh/datastore2/initialize"

	    "github.com/sawaq7/go12_ver1/client/sgh/datastore2/trans"
	    "github.com/sawaq7/go12_ver1/client/tokura/datastore4/trans2"
	    "github.com/sawaq7/go12_ver1/client/tokura/datastore4/initialize2"
//	    "github.com/sawaq7/go12_ver1/general/type5"

                                                )

///
///     クライアント�EチE�Eタストアのアクセスのメインルーチン �E�　for tokura　)
///     Datastore_client


func Datastore_tokura( fname string ,function string ,flexible_in interface{} ,w http.ResponseWriter, r *http.Request )  (flexible_out interface{} ) {

//     IN    fname       : チE�Eタストアのファイル吁E//     IN    function    : ファンクション　
//        　　　　　　　　　�E�　trans ,check ,initialize ,sort　etc
//     IN flexible_in　  : 吁E��インプットデータ　�E�別紙参照�E�E//     IN    w      　　 : レスポンスライター
//     IN    r      　　 : リクエストパラメータ

//     out flexible_out  : 吁E��アウト�EチE��チE�Eタ　�E�別紙参照�E�E
//    fmt.Fprintf( w, "datastore_tokura start \n" )  // チE��チE��
//    fmt.Fprintf( w, "datastore_tokura function \n" ,function )  // チE��チE��
//    fmt.Fprintf( w, "datastore_tokura fname \n" ,fname )  // チE��チE��

///
///  ファンクション・ファイル名により吁E��処琁E��刁E��して行う、E///

	switch function {

///
///  トランスの場吁E///

      case "trans" :     // trans の場吁E
        switch fname {

//          case "Deliver" :  // 該当する�E達情報をゲチE��

            // 空インターフェイス変数よりバリュー値をゲチE��

//            value, _ := flexible_in.([]type5.General_Work)
//            flexible_out = trans.Deliver ( value[0].Int64_Work  ,value[1].Int64_Work  ,w ,r )

//          break;

          case "Water_Line" :     // 水路ラインファイルの場吁E
            value, _ := flexible_in.(string)

            flexible_out = trans2.Water_line ( 1  ,value , w ,r )

          break;

          case "Water_Slope" :     // 導水勾配線ファイルの場吁E
            flexible_out = trans2.Water_slope ( w ,r )

          break;

          case "Water2" :        // 水路ファイル

            flexible_out = trans2.Water2 ( w ,r )

          break;

          case "Car" :  // 号車情報をゲチE��

            value, _ := flexible_in.(int64)
            flexible_out = trans.Car_district ( value ,w ,r )

          break;

          case "Sgh_Ai" :  // AI惁E��をゲチE��

            value, _ := flexible_in.(int64)
            flexible_out = trans.Sgh_ai( value  ,w , r  )

          break;

        }
      break;

      case "trans2" :       // trans2 の場吁E
        switch fname {

          case "Deliver" :  // 該当する�E達情報をゲチE��



          break;

          case "D_Area" :  // 号車情報をゲチE��

            value, _ := flexible_in.(int64)
            flexible_out = trans.D_area_district ( w ,r ,value )


          break;

          case "Private" :  // workerの個人惁E��をゲチE��



          break;

          case "Car" :  // 号車情報をゲチE��



          break;

          case "D_District_Temp" :     // 地区惁E��の一時ファイルの場吁E


          break;
        }
      break;

///
///  初期化�E場吁E///

      case "initialize" :

        switch fname {

          case "Water2_Temp" :   // 水路ファイルの場合　

            initialize2.Water2_temp (w , r )   // 水路惁E��をイニシャライズ  & セチE��　

          break;

          case "D_District_Temp" :     // エリア惁E��の場吁E
            initialize.D_district_temp (w , r ) // temporary-fileをイニシャライズ

          break;

          case "Sgh_Ai" :     // エリア惁E��の場吁E
            value, _ := flexible_in.(int64)
            initialize.Sgh_ai( value ,w , r )  //   既存�EAI惁E��を削除

          break;

        }

      break;


///
///  チェチE��の場吁E///

      case "check" :

        switch fname {

          case "D_Area" :   // 配達地区のエリアのMAX値をゲチE��

            // 空インターフェイス変数よりバリュー値をゲチE��

            value, _ := flexible_in.(int64)
            flexible_out = check.D_area ( w , r  ,value )

          break;

          case "D_District_Temp" :   // 一時ファイルよりカレント�E地区の惁E��をゲチE��　

            flexible_out = check.D_district_temp (w , r )

          break;

          case "Car" :   // 配達地区カーNOのMAX値をゲチE��

            // 空インターフェイス変数よりバリュー値をゲチE��

            value, _ := flexible_in.(int64)
            flexible_out = check.Car_no_max(w , r  ,value)

          break;

        }

      break;

    }

	return flexible_out

}

