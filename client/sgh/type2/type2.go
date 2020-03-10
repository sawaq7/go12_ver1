package type2

import (
          "time"
                  )
//
// the collection of struct information for sagawa express
//

//
// deliver district information　（地区情報）
//

type D_District struct {               /// データストア用

       Id               int64           //　データid
       District_No      int64           // 配達地域NO.
	   District_Name    string          // 配達地域名


   }

type D_District_Temp struct {         /// 一時ファイル用

       Id             int64           //　データid
       District_No    int64           // 配達地域NO.
	   District_Name  string          // 配達地域名


   }

type D_District_View struct {               /// 表示用

       Id               int64           //　データid
       District_No      int64           // 配達地域NO.
	   District_Name    string          // 配達地域名
       D_Area_Slice     []D_Area        //　エリア情報のスライス

   }

//
// deliver area information　（エリア情報）
//

type D_Area struct {               /// 常用ファイル用

       Id             int64           // データid
       Course_No      int64           // コースNO.
       District_No    int64           // 配達地域NO.
       District_Name  string          // 配達地域名
       Area_No        int64           // 配達エリアNO.
	   Area_Name      string          // 配達エリア名
	   Area_Detail    string          // 配達エリアの詳細
	   Number_Total   int64           // 宅配配達総数
	   Time_Total     float64         // 宅配配達総時間
	   Productivity   float64         // 宅配生産性
       Car_No         int64           // レギュラー号車
   }

//
// deliver area information　（エリア情報）
//

type D_Area_Temp struct {           /// 一時ファイル用

       Id           int64           // データid
       Course_No    int64           // コースNO.
       District_No    int64         // 配達地域NO.
       District_Name  string          // 配達地域名
       Area_No      int64           // 配達エリアNO.
	   Area_Name    string          // 配達エリア名
	   Area_Detail  string          // 配達エリアの詳細
	   Number_Total   int64           // 宅配配達総数
	   Time_Total     float64         // 宅配配達総時間
	   Productivity   float64         // 宅配生産性
       Car_No       int64           // レギュラー号車
   }

//
// deliver information　（配達）
//

type Deliver struct {

       Id           int64           // データid
       Line_No      int64           // 行NO.
       Course_No    int64           // コースNO.
       District_No    int64         // 配達地域NO.
       District_Name  string          // 配達地域名
       Area_No      int64           // 配達エリアNO.
	   Area_Name    string          // 配達エリア名
       Date         string         // 配達日
       Date_Real    time.Time      // 配達日(タイムデータ）
	   Car_No       int64           // 号車
	   Private_No   int64           // 個人番号
	   Car_Division int64           // 車両区分　０：軽自動車　１：ワゴン　２：トラック
	   Number       int64           // 宅配配達数

   }


//
// deliver's Schedule information　（配達 スケジュール）
//

type Private struct {

       Id             int64           //　データid
       Worker_No      int64           // ワーカーNO.
       Worker_Name    string          // ワーカー名
       Worker_Type    string          // ワーカータイプ　　SD : セールスドライバー　DD : 宅配ドライバー
                                      // 　　　　　　　　　OD : 外注ドライバー
       Worker_Salary  float64         // ワーカーサラリー（年収）
       Worker_Twh     float64         // ワーカー総労働時間（年間）Twh : total working hours
       Worker_H_Pay   float64         // ワーカー時給　H_Pay : hourly　pay
       Number_Total   int64           // 宅配配達総数
	   Time_Total     float64         // 宅配配達総時間
	   Productivity   float64         // 労働生産性

   }


//
//   car information　（号車情報）
//

type Car struct {

       Id                int64           //　データid
       District_No       int64           // 配達地域NO.
       District_Name     string          // 配達地域名
       Car_No            int64           // 号車NO.(シーケンシャルなNO)
	   Car_Name          string          // 号車名
	   Car_Explain       string          // 号車の走行ルート等の説明
	   Number_Total      int64           // 号車の宅配配達総数
	   Time_Total        float64         // 宅配配達総時間
	   Productivity      float64         // 労働生産性

   }



//
// deliver's Schedule information　（配達 スケジュール）
//

type D_Schedule struct {

       Id             int64           //　データid
       District_No    int64           // 配達地域NO.
       District_Name  string          // 配達地域名
       Date           string          // 配達日
       Date_Real    time.Time         // 配達日(タイムデータ）
       Expected_Num   float64         // 荷物の予想個数
       Judge          string          // 配達の判定
       Course_Num     int64           // コース数
	   Course_01      string          // コース1を乗車するドライバー
	   Car_Name_01     string          // コース1の号車名
	   Course_02      string          // コース2を乗車するドライバー
	   Car_Name_02     string          // コース1の号車名
	   Course_03      string          // コース3を乗車するドライバー
	   Car_Name_03     string          // コース1の号車名
	   Course_04      string          // コース4を乗車するドライバー
	   Car_Name_04     string          // コース1の号車名
       Course_05      string          // コース5を乗車するドライバー
       Car_Name_05     string          // コース1の号車名
	   Course_06      string          // コース6を乗車するドライバー
	   Car_Name_06     string          // コース1の号車名
	   Course_07      string          // コース7を乗車するドライバー
	   Car_Name_07     string          // コース1の号車名
	   Course_08      string          // コース8を乗車するドライバー
	   Car_Name_08     string          // コース1の号車名
	   Course_09      string          // コース9を乗車するドライバー
	   Car_Name_09     string          // コース1の号車名
	   Course_10      string          // コース10を乗車するドライバー
	   Car_Name_10     string          // コース1の号車名

   }
//
// the collection of some condition expression　（条件式）
//

type Sgh_Ai struct {

       Id              int64          // データid
       Course_No       int64          // コースNO.
       District_No     int64          // 配達地域NO.
       District_Name   string         // 配達地域名
       Area_No         int64          // 配達エリアNO.
	   Area_Name       string         // 配達エリア名
       Date_Basic      string         // 基準日(関数の始点）
       Date_Basic_Real time.Time      // 基準日(タイムデータ）
       Ex_Type         string         // 条件式のタイプ
                                         // 1. function : 関数
       Expression    string         // 条件式
	   Item_Num      int64          // 項数
	   Item1_Name    string         // 項名1
	   Item1_Factor  float64        // 項の係数1
	   Item2_Name    string         // 項名1
	   Item2_Factor  float64        // 項の係数1
	   Item3_Name    string         // 項名1
	   Item3_Factor  float64        // 項の係数1
	   Item4_Name    string         // 項名1
	   Item4_Factor  float64        // 項の係数1
	   Item5_Name    string         // 項名1
	   Item5_Factor  float64        // 項の係数1


   }



