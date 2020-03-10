package type1000

import (

//	    "google.golang.org/appengine/datastore"
//        "time"
//        "client/sgh/type2"

                                                )
//
// struct information for sagawa express
//


// Book holds metadata about a book.

type Book_Test struct {
	ID            int64
	Title         string
	Author        string
	PublishedDate string
	ImageURL      string
	Description   string
	CreatedBy     string
	CreatedByID   string
	File_Name     string
}

//
// deliver district information　（地区情報・マルチ構造体バーション）
//

type D_District struct {               /// データストア用

       Id               int64           //　データid
       District_No      int64           // 配達地域NO.
	   District_Name    string          // 配達地域名
//       D_Area_Slice   []type2.D_Area    // エリア情報
       D_Area_Slice     []D_Area    // エリア情報
     D_Area_Small_Slice []D_Area_Small  // スモールエリア情報

   }

//
// deliver area information　（エリア情報）
//

type D_Area struct {               /// 常用ファイル用

       Id              int64           // データid
       Course_No       int64           // コースNO.
       District_No     int64           // 配達地域NO.
       District_Name   string          // 配達地域名
       Area_No         int64           // 配達エリアNO.
	   Area_Name       string          // 配達エリア名
	   Area_Detail     string          // 配達エリアの詳細
	   Number_Total    int64           // 宅配配達総数
	   Time_Total      float64         // 宅配配達総時間
	   Productivity    float64         // 宅配生産性
       Car_No          int64           // レギュラー号車
//    D_Area_Small_Slice []D_Area_Small  // スモールエリア情報

   }

type D_Area_Small struct {               /// 常用ファイル用

	   Area_Name       string          // 配達エリア名
	   Area_Small_Name string          // 配達エリアの詳細

   }

//
// deliver district information　（地区情報・マルチ構造体バーション）
//

type D_District_2 struct {               /// データストア用

       Id               int64           //　データid
       District_No      int64           // 配達地域NO.
	   District_Name    string          // 配達地域名
       D_Area_Slice     []D_Area_2    // エリア情報
//     D_Area_Small_Slice []D_Area_Small  // スモールエリア情報

   }

//
// deliver area information　（エリア情報）
//

type D_Area_2 struct {               /// 常用ファイル用

       Id              int64           // データid
       Course_No       int64           // コースNO.
       District_No     int64           // 配達地域NO.
       District_Name   string          // 配達地域名
       Area_No         int64           // 配達エリアNO.
	   Area_Name       string          // 配達エリア名
	   Area_Detail     string          // 配達エリアの詳細
	   Number_Total    int64           // 宅配配達総数
	   Time_Total      float64         // 宅配配達総時間
	   Productivity    float64         // 宅配生産性
       Car_No          int64           // レギュラー号車
    D_Area_Small_Slice []D_Area_Small  // スモールエリア情報

   }