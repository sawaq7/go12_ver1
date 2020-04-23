package suiri

import (

//	"fmt"
	"github.com/sawaq7/go12_ver1/client/tokura/equation"
	"strconv"
	"strings"
)

///
///       make water-slope-line
///

func Kansui1(wdeta string) ([]string, []string, []string, []string, []string, []string, []string) {

	//     IN  wdeta :    water data

	//    OUT  one   : point loss
	//    OUT  two   : line loss
	//    OUT  three : velocity head
	//    OUT  four  : energy line down
	//    OUT  five  : energy line up
	//    OUT  six   : water-slope-line down
	//    OUT  seven : water-slope-line up

	var f_coeff, velocity, s_coeff, diameter, length, b_length float64
	var x_eneup, y_eneup, x_enedown, y_enedown float64
	var x_glineup, y_glineup, x_glinedown, y_glinedown float64
	var Hmax, hp, hl, b_hl, vhead float64
	var tflag, wflag, eflag, index int
	var char string

	//   fmt.Println ("func kansui1 wdeta　",wdeta )

	// change water data from string-type to string-array-type by spliting brank
	str := strings.Fields(wdeta)

	//   fmt.Println ("kansui1 nummax" ,len(str))

	///
	///   make work area with string-slice
	///

	ad_hp := make([]string, 20, 50)        // ①　hp
	ad_hl := make([]string, 20, 50)        // ②　hl
	ad_vhead := make([]string, 20, 50)     // ③　vhead
	ad_eneup := make([]string, 20, 50)     // ④　eneup
	ad_enedown := make([]string, 20, 50)   // ⑤　enedown
	ad_glineup := make([]string, 20, 50)   // ⑥　glineup
	ad_glinedown := make([]string, 20, 50) // ⑦　glinedown
	ad_wk := make([]float64, 2, 5)
	ad_wk2 := make([]string, 2, 5)

	index = 0
	eflag = 0

///
///     make various data from one-string-record
///

	for i := 0; i < len(str); i++ {

		// separate string data by spliting comma
		char = str[i]
		str2 := strings.Split(char, ",")

		//      fmt.Println("kansui1 str2" ,str2)
		//      fmt.Println ("kansui1 num2 " ,len(str2))

		tflag = 0

		for j := 0; j < len(str2); j++ {

			char2 := str2[j]
			//          fmt.Println ("kansui1 char2 " ,char2)
			str3 := strings.Split(char2, ":")

			//          fmt.Println ("kansui1 num3 " ,len(str3))
			//          fmt.Println("kansui1 str3" ,j ,str3)

			switch str3[0] {

			//    high
			case "H":

				Hmax, _ = strconv.ParseFloat(str3[1], 64)
				//             fmt.Println("kansui1 Hmax" ,Hmax)

				break

			//　coefficient
			case "n":

				s_coeff, _ = strconv.ParseFloat(str3[1], 64)
				//             fmt.Println("kansui1 s_coeff" ,s_coeff)

				break

			// point
			case "pt":

				tflag = 1

				break

				// f_coeff
			case "f":

				f_coeff, _ = strconv.ParseFloat(str3[1], 64)
				//             fmt.Println("kansui1 f_coeff" ,f_coeff)

				break

				//　velocity
			case "v":

				velocity, _ = strconv.ParseFloat(str3[1], 64)
				//             fmt.Println("kansui1 velocity" ,velocity)

				break

			//  line
			case "len":

				tflag = 2

				break

			//    diameter
			case "d":

				diameter, _ = strconv.ParseFloat(str3[1], 64)
				//             fmt.Println("kansui1 diameter" ,diameter)

				break

			//  length
			case "l":

				length, _ = strconv.ParseFloat(str3[1], 64)
				//             fmt.Println("kansui1 length" ,length)

				break

			}

			///    since adjust various data , do various calculations

			if j == 2 {

				if tflag == 1 { // calculate point loss

					vhead = equation.Suiri_Vhead(velocity) //  calculate  velocity head
					hp = f_coeff * vhead
//					fmt.Println("kansui1 hp", hp)

				} else if tflag == 2 { // make line loss

					ramuda := equation.Suiri_Manningu2(s_coeff, diameter) // calculate frictional coefficient
					vhead := equation.Suiri_Vhead( velocity )  //   calculate  velocity head
					hl = ramuda * (length / diameter) * vhead
//					fmt.Println("kansui1 hl", hl)

				}
			}
		}

		//  whether or not various data write

		if tflag == 2 {

			wflag = 1

		} else if i == len(str)-1 {

			wflag = 1
			eflag = 1
		}

	   // if wflag equal one ,write various-work-area
		if wflag == 1 {

			ad_hp[index] = strconv.FormatFloat(hp, 'f', 8, 64)
//			fmt.Println("kansui1 hp(ad)", ad_hp)

			if eflag == 1 { // when the data is last , the data is irregular process

				hl = 0.0
				vhead = 0.0
			}
			ad_hl[index] = strconv.FormatFloat(hl, 'f', 8, 64)
//			fmt.Println("kansui1 hl(ad)", ad_hl) //

			ad_vhead[index] = strconv.FormatFloat(vhead, 'f', 8, 64)
//			fmt.Println("kansui1 vhead(ad)", ad_vhead)

///　     make energy line up

			if index == 0 {
				b_length = 0.0
				x_eneup = 0.0
				y_eneup = Hmax
			} else {
				y_eneup = y_enedown - b_hl
			}
			x_eneup = x_eneup + b_length

			b_length = length
			b_hl = hl

			ad_wk[0] = x_eneup
			ad_wk[1] = y_eneup

			ad_wk2[0] = strconv.FormatFloat(ad_wk[0], 'f', 8, 64)
			ad_wk2[1] = strconv.FormatFloat(ad_wk[1], 'f', 8, 64)

			ad_eneup[index] = strings.Join(ad_wk2, ",") //　make x,y coordinate
//			fmt.Println("kansui1 eneup(ad)", ad_eneup)

///　     make energy-line-down

			x_enedown = x_eneup
			y_enedown = y_eneup - hp
			ad_wk[0] = x_enedown
			ad_wk[1] = y_enedown

			ad_wk2[0] = strconv.FormatFloat(ad_wk[0], 'f', 8, 64)
			ad_wk2[1] = strconv.FormatFloat(ad_wk[1], 'f', 8, 64)

			ad_enedown[index] = strings.Join(ad_wk2, ",") //　make x,y coordinate

//			fmt.Println("kansui1 enedown(ad)", ad_enedown)

///　     make water-slope-line-up

			x_glineup = x_eneup
			y_glineup = y_eneup - vhead
			ad_wk[0] = x_glineup
			ad_wk[1] = y_glineup

			ad_wk2[0] = strconv.FormatFloat(ad_wk[0], 'f', 8, 64)
			ad_wk2[1] = strconv.FormatFloat(ad_wk[1], 'f', 8, 64)

			ad_glineup[index] = strings.Join(ad_wk2, ",") //　make x,y coordinate

//			fmt.Println("kansui1 glinedown(ad)", ad_glineup)

///　    make water-slope-line-down

			x_glinedown = x_eneup
			y_glinedown = y_glineup - hp

			ad_wk[0] = x_glinedown
			ad_wk[1] = y_glinedown

			ad_wk2[0] = strconv.FormatFloat(ad_wk[0], 'f', 8, 64)
			ad_wk2[1] = strconv.FormatFloat(ad_wk[1], 'f', 8, 64)

			ad_glinedown[index] = strings.Join(ad_wk2, ",") //　make x,y coordinate

			//         fmt.Println("kansui1 glinedown(ad)" ,ad_glinedown)

			wflag = 0
			index++

		}
	}

	return ad_hp, ad_hl, ad_vhead, ad_eneup, ad_enedown, ad_glineup, ad_glinedown
}
