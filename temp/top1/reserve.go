// Copyright Â©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package top1

import (
	"errors"
	"image/color"
	"math"
//	"net/http"
	 "fmt"
    "math/rand"
	"github.com/gonum/plot"

	"github.com/gonum/plot/vg"
	"github.com/gonum/plot/vg/draw"
	"github.com/gonum/plot/plotter"
 	"github.com/gonum/plot/plotutil"
                                            )

// A BarChart presents grouped data with rectangular bars
// with lengths proportional to the data values.
type BarChart struct {
	Values  []float64


	// Width is the width of the bars.
	Width vg.Length

	// Color is the fill color of the bars.
	Color color.Color

	// LineStyle is the style of the outline of the bars.
	draw.LineStyle

	// Offset is added to the X location of each bar.
	// When the Offset is zero, the bars are drawn
	// centered at their X location.
	Offset vg.Length

	// XMin is the X location of the first bar.  XMin
	// can be changed to move groups of bars
	// down the X axis in order to make grouped
	// bar charts.
	XMin float64

	// Horizontal dictates whether the bars should be in the vertical
	// (default) or horizontal direction. If Horizontal is true, all
	// X locations and distances referred to here will actually be Y
	// locations and distances.
	Horizontal bool

	// stackedOn is the bar chart upon which
	// this bar chart is stacked.
	stackedOn *BarChart
}
var (
	// DefaultLineStyle is the default style for drawing
	// lines.
	DefaultLineStyle = draw.LineStyle{
		Color:    color.Black,
		Width:    vg.Points(1),
		Dashes:   []vg.Length{},
		DashOffs: 0,
	}
                                        )

func Reserve() {

//    fmt.Println ("Reserve start ãƒ˜ãƒƒãƒ€ãƒ¼write " ,line)  // ãƒEƒãƒE‚¯
    fmt.Println ("Reserve start " )  // ãƒEƒãƒE‚¯


// func reserve( w http.ResponseWriter ,r *http.Request ,ad_eneup []type3.Point ) {

    rand.Seed(int64(0))

///
/// ã‚°ãƒ©ãƒ•ãEæ ã‚’ä½œæEã€€
///

 	p, err := plot.New()
 	if err != nil {

	   return

 	}

 	p.Title.Text = "reserve situation"
 	p.X.Label.Text = "time"
 	p.Y.Label.Text = "guest"

// 	p.X.Min = 0
//    p.X.Max = 200
//    p.Y.Min = 15
//    p.Y.Max = 20

    p.BackgroundColor = color.RGBA{R: 102,G: 204, B: 255}

    name1 := "sawamoto"

    p.NominalY( name1, "yoshida", "tanaka", "kikuti", "okada") // åE€¤ã®ãƒ©ãƒ™ãƒ«(Xè»¸)



    size :=10.

    wide := vg.Points(size) // æ£’ã‚°ãƒ©ãƒ•ãEå¹E

  	// time_maxã®åE€¤ã«ã¤ãE¦æ£’ã‚°ãƒ©ãƒ•ã‚’ç”ŸæE

//  	time_max1 := 20

//  	var guest_num int
//    guest_num = 5

//    guest_num := 100

    wide_num :=  100 / vg.Inch

//    wide_num :=  guest_num / vg.Inch

    fmt.Println ("Reserve start ,vg.Inch" ,vg.Inch )  // ãƒEƒãƒE‚¯
    fmt.Println ("Reserve start ,vg.Inch2" ,vg.Inch*2 )  // ãƒEƒãƒE‚¯
    fmt.Println ("Reserve start ,wide_num" ,wide_num )  // ãƒEƒãƒE‚¯




  	time_max := make (plotter.Values ,5)
  	time_max[0] = 10.
  	time_max[1] = 13.
  	time_max[2] = 18.
  	time_max[3] = 20.
  	time_max[4] = 21.
//  	time_max := plotter.Values{ float64(time_max1), 21 ,22 ,23 ,24 } // æç”»å¯¾è±¡

//  	bars, _ := plotter.NewBarChart(time_max, wide)

  	bars, _ := NewBarChart(time_max, wide)

  	bars.LineStyle.Width = vg.Length(0) // æ£’ã‚°ãƒ©ãƒ•ãEæ ç·šãEå¤ªãE

  	bars.Color = plotutil.Color(3)      // æ£’ã‚°ãƒ©ãƒ•ãEè‰²ã€Eã‹ã‚‰6ã¾ã§ã§plotutilã«SoftColorsã¨ã—ã¦å®šç¾©ã•ã‚Œã¦ãE‚‹ã€E

//    fmt.Println ("Reserve main bars.Values " ,bars.Values )  // ãƒEƒãƒE‚¯

//  	bars.Offset = wide * 2                    // æ£’ã‚°ãƒ©ãƒ•ã‚’è¡¨ç¤ºã™ã‚‹ä½ç½®ã®ã‚ªãƒ•ã‚»ãƒEƒˆ(Xæ–¹åEã€‚è¤E•°ã®ã‚°ãƒ«ãƒ¼ãƒ—ã‚’ä¸¦ã¹ãŸã„ã¨ããE-wãªã©ã§ä½ç½®ã‚’èª¿æ•´ã™ã‚‹ã€E

    bars.Offset = 0
//    bars.XMin = 2

// 	bars.Horizontal = false             // trueã«ã™ã‚‹ã¨æ¨ªå‘ãã®æ£’ã‚°ãƒ©ãƒ•ã«ãªã‚‹ã€E

 	bars.Horizontal = true

	p.Add(bars)

//    show_size := float64(guest_num) * size * 1.5

    err = p.Save(5*vg.Inch, 5*vg.Inch, "C:/Go_Original/sample/reservation_system/reservation.png")


	if err != nil {
		fmt.Println ("Reserve err " ,err)  // ãƒEƒãƒE‚¯
	}

    fmt.Println ("Reserve end " )  // ãƒEƒãƒE‚¯

}

// NewBarChart returns a new bar chart with a single bar for each value.
// The bars heights correspond to the values and their x locations correspond
// to the index of their value in the Valuer.
func NewBarChart(vs plotter.Valuer, width vg.Length) (*BarChart, error) {

    fmt.Println ("NewBarChart start " )  // ãƒEƒãƒE‚¯

	if width <= 0 {
		return nil, errors.New("Width parameter was not positive")
	}
	values, err := plotter.CopyValues(vs)

	if err != nil {
		return nil, err
	}
	fmt.Println ("NewBarChart end " )  // ãƒEƒãƒE‚¯

	return &BarChart{
		Values:    values,
		Width:     width,
		Color:     color.Black,
		LineStyle: DefaultLineStyle,
	}, nil


}

// BarHeight returns the maximum y value of the
// ith bar, taking into account any bars upon
// which it is stacked.
func (b *BarChart) BarHeight(i int) float64 {

    fmt.Println (" BarHeight start i" ,i)  // ãƒEƒãƒE‚¯

	ht := 0.0
	if b == nil {

//	   ht= ht + 2.

	   fmt.Println (" BarHeight ht " ,ht )  // ãƒEƒãƒE‚¯

		return ht
	}
	if i >= 0 && i < len(b.Values) {
		ht += b.Values[i]
	}
	if b.stackedOn != nil {
		ht += b.stackedOn.BarHeight(i)
	}

	fmt.Println (" BarHeight end"  )  // ãƒEƒãƒE‚¯

	return ht
}

// StackOn stacks a bar chart on top of another,
// and sets the XMin and Offset to that of the
// chart upon which it is being stacked.
func (b *BarChart) StackOn(on *BarChart) {

    fmt.Println ("StackOn start " )  // ãƒEƒãƒE‚¯

	b.XMin = on.XMin
	b.Offset = on.Offset
	b.stackedOn = on
}
// Plot implements the plot.Plotter interface.
func (b *BarChart) Plot(c draw.Canvas, plt *plot.Plot) {

    fmt.Println ("Plot start " )  // ãƒEƒãƒE‚¯

   Values2 := make([]float64 ,20 ,50)

    Values2[0] = 1.
  	Values2[1] = 2.
  	Values2[2] = 3.
  	Values2[3] = 4.
  	Values2[4] = 5.

	trCat, trVal := plt.Transforms(&c)
	if b.Horizontal {
		trCat, trVal = trVal, trCat
	}

	for i, ht := range b.Values {
		catVal := b.XMin + float64(i)
		catMin := trCat(float64(catVal))
		if !b.Horizontal {
			if !c.ContainsX(catMin) {
				continue
			}
		} else {
			if !c.ContainsY(catMin) {
				continue
			}
		}
		catMin = catMin - b.Width/2 + b.Offset
		catMax := catMin + b.Width
		bottom := b.stackedOn.BarHeight(i)



		fmt.Println ("Plot bottom " ,bottom )  // ãƒEƒãƒE‚¯

//		valMin := trVal(bottom)

		valMin := trVal(bottom+Values2[i])

		valMax := trVal(bottom + ht)

//		fmt.Println ("Plot valMin " ,valMin )  // ãƒEƒãƒE‚¯

		var pts []vg.Point
		var poly []vg.Point
		if !b.Horizontal {                // ãƒãEãƒãƒ«ã‚¿ã‚¤ãƒ—ãEå ´åE

		   fmt.Println ("Plot type1 "  )  // ãƒEƒãƒE‚¯

			pts = []vg.Point{
				{catMin, valMin},
				{catMin, valMax},
				{catMax, valMax},
				{catMax, valMin},
			}
			poly = c.ClipPolygonY(pts)
		} else {                        // XYè»¸é€EEå ´åE

		   fmt.Println ("Plot type2 "  )  // ãƒEƒãƒE‚¯

			pts = []vg.Point{
				{valMin, catMin},
				{valMin, catMax},
//                {valMin+50, catMin},
//				{valMin+50, catMax},
				{valMax, catMax},
				{valMax, catMin},
			}

			fmt.Println ("Plot 1" ,valMin, catMin)  // ãƒEƒãƒE‚¯
            fmt.Println ("Plot 2" ,valMin, catMax)  // ãƒEƒãƒE‚¯
            fmt.Println ("Plot 3" ,valMax, catMax)  // ãƒEƒãƒE‚¯
            fmt.Println ("Plot 4" ,valMax, catMax)  // ãƒEƒãƒE‚¯

			poly = c.ClipPolygonX(pts)
		}

		fmt.Println ("Plot pts " ,pts )  // ãƒEƒãƒE‚¯

		c.FillPolygon(b.Color, poly)

		var outline [][]vg.Point
		if !b.Horizontal {              // ãƒãEãƒãƒ«ã‚¿ã‚¤ãƒ—ãEå ´åE

			pts = append(pts, vg.Point{X: catMin, Y: valMin})
			outline = c.ClipLinesY(pts)
		} else {                       // XYè»¸é€EEå ´åE

			pts = append(pts, vg.Point{X: valMin, Y: catMin})
			outline = c.ClipLinesX(pts)
		}

//		fmt.Println ("Plot outline " ,outline )  // ãƒEƒãƒE‚¯

		c.StrokeLines(b.LineStyle, outline...)
	}

	fmt.Println ("Plot end " )  // ãƒEƒãƒE‚¯

}


// DataRange implements the plot.DataRanger interface.
func (b *BarChart) DataRange() (xmin, xmax, ymin, ymax float64) {

    fmt.Println ("DataRange start " )  // ãƒEƒãƒE‚¯


	catMin := b.XMin
	catMax := catMin + float64(len(b.Values)-1)

	valMin := math.Inf(1)
	valMax := math.Inf(-1)
	for i, val := range b.Values {
		valBot := b.stackedOn.BarHeight(i)
		valTop := valBot + val
		valMin = math.Min(valMin, math.Min(valBot, valTop))
		valMax = math.Max(valMax, math.Max(valBot, valTop))
	}
	if !b.Horizontal {                              // ãƒãEãƒãƒ«ã‚¿ã‚¤ãƒ—ãEå ´åE
		return catMin, catMax, valMin, valMax
	}
	fmt.Println ("DataRange  valMin" ,valMin)  // ãƒEƒãƒE‚¯
    fmt.Println ("DataRange  valMax" ,valMax)  // ãƒEƒãƒE‚¯
    fmt.Println ("DataRange  catMin" ,catMin)  // ãƒEƒãƒE‚¯
    fmt.Println ("DataRange  catMax" ,catMax)  // ãƒEƒãƒE‚¯

    fmt.Println ("DataRange end " )  // ãƒEƒãƒE‚¯

	return valMin, valMax, catMin, catMax
}

// GlyphBoxes implements the GlyphBoxer interface.
func (b *BarChart) GlyphBoxes(plt *plot.Plot) []plot.GlyphBox {

    fmt.Println ("GlyphBoxes start " )  // ãƒEƒãƒE‚¯
//    fmt.Println ("GlyphBoxes plt " ,plt )  // ãƒEƒãƒE‚¯

	boxes := make([]plot.GlyphBox, len(b.Values))
	for i := range b.Values {
		cat := b.XMin + float64(i)
		if !b.Horizontal {
			boxes[i].X = plt.X.Norm(cat)
			boxes[i].Rectangle = vg.Rectangle{
				Min: vg.Point{X: b.Offset - b.Width/2},
				Max: vg.Point{X: b.Offset + b.Width/2},
			}
		} else {
			boxes[i].Y = plt.Y.Norm(cat)

			fmt.Println ( "GlyphBoxes boxes[i].Y" ,boxes[i].Y)  // ãƒEƒãƒE‚¯

			boxes[i].Rectangle = vg.Rectangle{
				Min: vg.Point{Y: b.Offset - b.Width/2},
				Max: vg.Point{Y: b.Offset + b.Width/2},
			}
		}
	}

	fmt.Println ("GlyphBoxes end "  )  // ãƒEƒãƒE‚¯

	return boxes
}

// Thumbnail fulfills the plot.Thumbnailer interface.
func (b *BarChart) Thumbnail(c *draw.Canvas) {

    fmt.Println ("Thumbnail start "  )  // ãƒEƒãƒE‚¯

	pts := []vg.Point{
		{c.Min.X, c.Min.Y},
		{c.Min.X, c.Max.Y},
		{c.Max.X, c.Max.Y},
		{c.Max.X, c.Min.Y},
	}
	poly := c.ClipPolygonY(pts)
	c.FillPolygon(b.Color, poly)

	pts = append(pts, vg.Point{X: c.Min.X, Y: c.Min.Y})
	outline := c.ClipLinesY(pts)
	c.StrokeLines(b.LineStyle, outline...)

	fmt.Println ("Thumbnail end "  )  // ãƒEƒãƒE‚¯
}
