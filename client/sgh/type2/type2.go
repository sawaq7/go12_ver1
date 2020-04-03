package type2

import (
          "time"
                  )
//
// the collection of struct information for sagawa express
//

//
// deliver district information　�E�地区惁E���E�E//

type D_District struct {               /// チE�Eタストア用

       Id               int64           //　チE�Eタid
       District_No      int64           // 配達地域NO.
	   District_Name    string          // 配達地域名


   }

type D_District_Temp struct {         /// 一時ファイル用

       Id             int64           //　チE�Eタid
       District_No    int64           // 配達地域NO.
	   District_Name  string          // 配達地域名


   }

type D_District_View struct {               /// 表示用

       Id               int64           //　チE�Eタid
       District_No      int64           // 配達地域NO.
	   District_Name    string          // 配達地域名
       D_Area_Slice     []D_Area        //　エリア惁E��のスライス

   }

//
// deliver area information　�E�エリア惁E���E�E//

type D_Area struct {               /// 常用ファイル用

       Id             int64           // チE�Eタid
       Course_No      int64           // コースNO.
       District_No    int64           // 配達地域NO.
       District_Name  string          // 配達地域名
       Area_No        int64           // 配達エリアNO.
	   Area_Name      string          // 配達エリア吁E	   Area_Detail    string          // 配達エリアの詳細
	   Number_Total   int64           // 宁E�E配達総数
	   Time_Total     float64         // 宁E�E配達総時閁E	   Productivity   float64         // 宁E�E生産性
       Car_No         int64           // レギュラー号軁E   }

//
// deliver area information　�E�エリア惁E���E�E//

type D_Area_Temp struct {           /// 一時ファイル用

       Id           int64           // チE�Eタid
       Course_No    int64           // コースNO.
       District_No    int64         // 配達地域NO.
       District_Name  string          // 配達地域名
       Area_No      int64           // 配達エリアNO.
	   Area_Name    string          // 配達エリア吁E	   Area_Detail  string          // 配達エリアの詳細
	   Number_Total   int64           // 宁E�E配達総数
	   Time_Total     float64         // 宁E�E配達総時閁E	   Productivity   float64         // 宁E�E生産性
       Car_No       int64           // レギュラー号軁E   }

//
// deliver information　�E��E達！E//

type Deliver struct {

       Id           int64           // チE�Eタid
       Line_No      int64           // 行NO.
       Course_No    int64           // コースNO.
       District_No    int64         // 配達地域NO.
       District_Name  string          // 配達地域名
       Area_No      int64           // 配達エリアNO.
	   Area_Name    string          // 配達エリア吁E       Date         string         // 配達日
       Date_Real    time.Time      // 配達日(タイムチE�Eタ�E�E	   Car_No       int64           // 号軁E	   Private_No   int64           // 個人番号
	   Car_Division int64           // 車両区刁E���E�：軽自動車　�E�：ワゴン　�E�：トラチE��
	   Number       int64           // 宁E�E配達数

   }


//
// deliver's Schedule information　�E��E遁Eスケジュール�E�E//

type Private struct {

       Id             int64           //　チE�Eタid
       Worker_No      int64           // ワーカーNO.
       Worker_Name    string          // ワーカー吁E       Worker_Type    string          // ワーカータイプ　　SD : セールスドライバ�E　DD : 宁E�Eドライバ�E
                                      // 　　　　　　　　　OD : 外注ドライバ�E
       Worker_Salary  float64         // ワーカーサラリー�E�年収！E       Worker_Twh     float64         // ワーカー総労働時間（年間）Twh : total working hours
       Worker_H_Pay   float64         // ワーカー時給　H_Pay : hourly　pay
       Number_Total   int64           // 宁E�E配達総数
	   Time_Total     float64         // 宁E�E配達総時閁E	   Productivity   float64         // 労働生産性

   }


//
//   car information　�E�号車情報�E�E//

type Car struct {

       Id                int64           //　チE�Eタid
       District_No       int64           // 配達地域NO.
       District_Name     string          // 配達地域名
       Car_No            int64           // 号車NO.(シーケンシャルなNO)
	   Car_Name          string          // 号車名
	   Car_Explain       string          // 号車�E走行ルート等�E説昁E	   Number_Total      int64           // 号車�E宁E�E配達総数
	   Time_Total        float64         // 宁E�E配達総時閁E	   Productivity      float64         // 労働生産性

   }



//
// deliver's Schedule information　�E��E遁Eスケジュール�E�E//

type D_Schedule struct {

       Id             int64           //　チE�Eタid
       District_No    int64           // 配達地域NO.
       District_Name  string          // 配達地域名
       Date           string          // 配達日
       Date_Real    time.Time         // 配達日(タイムチE�Eタ�E�E       Expected_Num   float64         // 荷物の予想個数
       Judge          string          // 配達の判宁E       Course_Num     int64           // コース数
	   Course_01      string          // コース1を乗車するドライバ�E
	   Car_Name_01     string          // コース1の号車名
	   Course_02      string          // コース2を乗車するドライバ�E
	   Car_Name_02     string          // コース1の号車名
	   Course_03      string          // コース3を乗車するドライバ�E
	   Car_Name_03     string          // コース1の号車名
	   Course_04      string          // コース4を乗車するドライバ�E
	   Car_Name_04     string          // コース1の号車名
       Course_05      string          // コース5を乗車するドライバ�E
       Car_Name_05     string          // コース1の号車名
	   Course_06      string          // コース6を乗車するドライバ�E
	   Car_Name_06     string          // コース1の号車名
	   Course_07      string          // コース7を乗車するドライバ�E
	   Car_Name_07     string          // コース1の号車名
	   Course_08      string          // コース8を乗車するドライバ�E
	   Car_Name_08     string          // コース1の号車名
	   Course_09      string          // コース9を乗車するドライバ�E
	   Car_Name_09     string          // コース1の号車名
	   Course_10      string          // コース10を乗車するドライバ�E
	   Car_Name_10     string          // コース1の号車名

   }
//
// the collection of some condition expression　�E�条件式！E//

type Sgh_Ai struct {

       Id              int64          // チE�Eタid
       Course_No       int64          // コースNO.
       District_No     int64          // 配達地域NO.
       District_Name   string         // 配達地域名
       Area_No         int64          // 配達エリアNO.
	   Area_Name       string         // 配達エリア吁E       Date_Basic      string         // 基準日(関数の始点�E�E       Date_Basic_Real time.Time      // 基準日(タイムチE�Eタ�E�E       Ex_Type         string         // 条件式�EタイチE                                         // 1. function : 関数
       Expression    string         // 条件弁E	   Item_Num      int64          // 頁E��
	   Item1_Name    string         // 頁E��1
	   Item1_Factor  float64        // 頁E�E係数1
	   Item2_Name    string         // 頁E��1
	   Item2_Factor  float64        // 頁E�E係数1
	   Item3_Name    string         // 頁E��1
	   Item3_Factor  float64        // 頁E�E係数1
	   Item4_Name    string         // 頁E��1
	   Item4_Factor  float64        // 頁E�E係数1
	   Item5_Name    string         // 頁E��1
	   Item5_Factor  float64        // 頁E�E係数1


   }



