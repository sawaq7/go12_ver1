package date1

import (

	    "net/http"
	    "strconv"
//	    "fmt"
	    "strings"

	    "time"
                                                )
///
///  日付データ（string型）　から　timeデータを作成する(　グーグル純正・time.Time型　）
///

func Date_realdata_get(w http.ResponseWriter ,date string) (date_real time.Time ){

//     IN    w       : レスポンスライター
//     IN  date　    : 日付　 ex : "2018/01/01"
//    OUT  date_real : タイムデータ

//    fmt.Fprintf( w, "date_realdata_get start \n" )  // デバック


///
/// タイムデータ作成
///

    strings_slice := strings.Split( date, "/" )   //　アップデートしたデータを分割

    loc, _ := time.LoadLocation("Local")            // ローカルのタイムゾーン情報を取得

    year_int64 ,_ :=strconv.ParseInt( strings_slice[0],10 ,0 )   // 年をint型に変換
    month_int64,_  :=strconv.ParseInt( strings_slice[1],10 ,0 )   // 月をint型に変換
    date_int64 ,_ :=strconv.ParseInt( strings_slice[2],10 ,0 )   // 日をint型に変換

    switch month_int64 {    /// time.Date の　月データが、変数だとbuildできないため、妥協案

// 1月の場合
          case 1 :

             date_real = time.Date( int(year_int64) ,1 ,int(date_int64) ,0 ,0 ,0 ,0 ,loc)


             break;

//　2月の場合
          case 2 :

             date_real = time.Date( int(year_int64) ,2 ,int(date_int64) ,0 ,0 ,0 ,0 ,loc )

             break;

// 3月の場合
          case 3 :

             date_real = time.Date( int(year_int64) ,3 ,int(date_int64) ,0 ,0 ,0 ,0 ,loc )

          break;

//　4月の場合
          case 4 :

            date_real = time.Date( int(year_int64) ,4 ,int(date_int64) ,0 ,0 ,0 ,0 ,loc )

          break;

//　5月の場合
          case 5 :

             date_real = time.Date( int(year_int64) ,5 ,int(date_int64) ,0 ,0 ,0 ,0 ,loc )

          break;

// 6月の場合
          case 6:

             date_real = time.Date( int(year_int64) ,6 ,int(date_int64) ,0 ,0 ,0 ,0 ,loc )

          break;

// 7月の場合
          case 7 :

             date_real = time.Date( int(year_int64) ,7 ,int(date_int64) ,0 ,0 ,0 ,0 ,loc )

          break;

//　8月の場合
          case 8 :

             date_real = time.Date( int(year_int64) ,8 ,int(date_int64) ,0 ,0 ,0 ,0 ,loc )

          break;
//　9月の場合
          case 9 :

             date_real = time.Date( int(year_int64) ,9 ,int(date_int64) ,0 ,0 ,0 ,0 ,loc )

          break;

// 10月の場合
          case 10:

            date_real = time.Date( int(year_int64) ,10 ,int(date_int64) ,0 ,0 ,0 ,0 ,loc )

          break;

// 11月の場合
          case 11 :

            date_real = time.Date( int(year_int64) ,11 ,int(date_int64) ,0 ,0 ,0 ,0 ,loc )


          break;

//　12月の場合
          case 12 :

             date_real = time.Date( int(year_int64) ,12 ,int(date_int64) ,0 ,0 ,0 ,0 ,loc )

          break;

          }
//    fmt.Fprintf( w, "date_realdata_get : date_real %v\n", date_real )  // デバック
//    fmt.Fprintf( w, "date_realdata_get normal end \n" )  // デバック

    return date_real


}
