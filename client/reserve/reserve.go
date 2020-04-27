package reserve

import (
	"errors"
	"image/color"
	"math"
	"net/http"
	 "fmt"
    "math/rand"
	"github.com/gonum/plot"
    "time"

    "github.com/sawaq7/go12_ver1/general/type5"
    "github.com/sawaq7/go12_ver1/general/html5"
    "github.com/sawaq7/go12_ver1/client/reserve/datastore6/trans5"
    "github.com/sawaq7/go12_ver1/client/reserve/html6"
    "github.com/sawaq7/go12_ver1/storage2"
    "strconv"
    "html/template"

//	"github.com/gonum/plot/vg"
//	"github.com/gonum/plot/vg/draw"
//	"github.com/gonum/plot/plotter"
// 	"github.com/gonum/plot/plotutil"

 	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
    "gonum.org/v1/plot/plotter"
 	"gonum.org/v1/plot/plotutil"
                                            )

// A BarChart presents grouped data with rectangular bars
// with lengths proportional to the data values.
type BarChart struct {
	Values  []float64

	Values2  []float64


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

func Reserve(w http.ResponseWriter ,r *http.Request ) {

//     IN    w      : response-writer
//     IN    r      : request- paramete

//    fmt.Fprintf( w, "reserve start \n" )

    reserve_date  := r.FormValue("reserve_date")

    guest_reserve_minor_slice := trans5.Guest_reserve_minor2( reserve_date , w , r  )

    data_num := len(guest_reserve_minor_slice)

//    fmt.Fprintf( w, "reserve guest_reserve_minor_slice[0].Guest_Name %v\n",guest_reserve_minor_slice[0].Guest_Name )

    rand.Seed(int64(0))

///
///     make frame of graff
///

 	p, err := plot.New()
 	if err != nil {

	   return

 	}

 	p.Title.Text = "reserve situation"
 	p.X.Label.Text = "time"
 	p.Y.Label.Text = "guest"

    p.BackgroundColor = color.RGBA{R: 102,G: 204, B: 255}

    size :=30.

    wide := vg.Points(size) //   set wide for fream og graf

//    fmt.Println ("Reserve start ,vg.Inch" ,vg.Inch )
//    fmt.Println ("Reserve start ,vg.Inch2" ,vg.Inch*2 )


  	time_max := make (plotter.Values ,data_num)
  	time_min := make (plotter.Values ,data_num)
  	names := make ([]string ,data_num)


  	for pos, guest_reserve_minor_slicew := range guest_reserve_minor_slice {

       start_time_float64w,_ :=strconv.ParseFloat(guest_reserve_minor_slicew.Start_Time,64)
       end_time_float64w,_ :=strconv.ParseFloat(guest_reserve_minor_slicew.End_Time,64)

       time_min[pos] = start_time_float64w
       time_max[pos] = end_time_float64w
       names[pos] = guest_reserve_minor_slicew.Guest_Name

       p.NominalY2( names )

    }

  	bars, _ := NewBarChart(time_max, wide)

    bars.Values2 = make (plotter.Values ,data_num)

  	for pos := 0; pos < len(time_min); pos++ {

  	   bars.Values2[pos] = time_min[pos]

  	}

//    bars.Values2[0] = 1.
  	bars.LineStyle.Width = vg.Length(0) //    set line-size

  	bars.Color = plotutil.Color(3)      //    set color

//    fmt.Println ("Reserve main bars.Values " ,bars.Values )

    bars.Offset = 0
//    bars.XMin = 2

// 	bars.Horizontal = false             //  if this is "true" , the graf is turned sideways

 	bars.Horizontal = true

	p.Add(bars)
///

 	bucket := "sample-7777"

///
///      make file-name
///

 	date_w := time.Now()
    unique_no := fmt.Sprintf("%04d%02d%02d%02d%02d%02d",
		date_w.Year(), date_w.Month(),date_w.Day(), date_w.Hour(), date_w.Minute(), date_w.Second())

    f_name := "reserve_" + unique_no + ".png"

//    f_name := "reserve" + ".png"
//    fmt.Fprintf( w, "reserve : f_name %v\n", f_name )

    if err := p.Save_Storage(w ,r ,5*vg.Inch, 5*vg.Inch, bucket , f_name ); err != nil {
   ///  reserve in storage

       http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
 	}

///
///     show graf on web
///

 	var image_show type5.Image_Show

 	const publicURL = "https://storage.googleapis.com/%s/%s"

 	image_show.File_Name = f_name

	image_show.Url = fmt.Sprintf(publicURL, bucket, image_show.File_Name)

//	fmt.Fprintf( w, "pipe_line1_show_graf : image_show.File_Name %v\n", image_show.File_Name )
//	fmt.Fprintf( w, "pipe_line1_show_graf : image_show.Url %v\n", image_show.Url )

	monitor := template.Must(template.New("html").Parse(html5.Image_file_show))
//     monitor := template.Must(template.New("html").Parse(html4.Pipe_line1_show_graf))

	err = monitor.Execute(w, image_show)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

//	storage2.File_Delete ( w , r  ,bucket ,f_name  )



//    fmt.Fprintf( w, "reserve normal end \n" )

}

func Reserve2(w http.ResponseWriter ,r *http.Request ,reserve_date string ) {

//     IN    w      : response-writer
//     IN    r      : request- paramete
//     IN  reserve_date : reserve-date


//    fmt.Fprintf( w, "reserve2 start \n" )

    guest_reserve_minor_slice := trans5.Guest_reserve_minor3( reserve_date , w , r  )

    data_num := len(guest_reserve_minor_slice)

//    fmt.Fprintf( w, "reserve guest_reserve_minor_slice[0].Guest_Name %v\n",guest_reserve_minor_slice[0].Guest_Name )

    rand.Seed(int64(0))

///
///       make frame of graff
///

 	p, err := plot.New()
 	if err != nil {

	   return

 	}

 	p.Title.Text = "reserve situation"
 	p.X.Label.Text = "time"
 	p.Y.Label.Text = "guest"

// 	p.X.Min = 5
//    p.X.Max = 200
//    p.Y.Min = 5
//    p.Y.Max = 20

    p.BackgroundColor = color.RGBA{R: 102,G: 204, B: 255}

    size :=30.

    wide := vg.Points(size) //     //   set wide for fream og graf

//    fmt.Println ("reserve2 start ,vg.Inch" ,vg.Inch )
//    fmt.Println ("reserve2 start ,vg.Inch2" ,vg.Inch*2 )


  	time_max := make (plotter.Values ,data_num)
  	time_min := make (plotter.Values ,data_num)
  	names := make ([]string ,data_num)


  	for pos, guest_reserve_minor_slicew := range guest_reserve_minor_slice {

       start_time_float64w,_ :=strconv.ParseFloat(guest_reserve_minor_slicew.Start_Time,64)
       end_time_float64w,_ :=strconv.ParseFloat(guest_reserve_minor_slicew.End_Time,64)

       time_min[pos] = start_time_float64w
       time_max[pos] = end_time_float64w
       names[pos] = guest_reserve_minor_slicew.Guest_Name

       p.NominalY2( names )

    }

  	bars, _ := NewBarChart(time_max, wide)

    bars.Values2 = make (plotter.Values ,data_num)

  	for pos := 0; pos < len(time_min); pos++ {

  	   bars.Values2[pos] = time_min[pos]

  	}

  	bars.LineStyle.Width = vg.Length(0)

  	bars.Color = plotutil.Color(3)

    bars.Offset = 0
//    bars.XMin = 2

// 	bars.Horizontal = false

 	bars.Horizontal = true

	p.Add(bars)
///

 	bucket := "sample-7777"

///
///     make file name
///

    f_name := "reserve" + ".png"

    storage2.File_Delete ( w , r  ,bucket ,f_name  ) //    delete old file

//    fmt.Fprintf( w, "reserve2 : f_name %v\n", f_name )

    if err := p.Save_Storage(w ,r ,5*vg.Inch, 5*vg.Inch, bucket , f_name ); err != nil {  // reserve new file

       http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
 	}

///
///       show graf on web
///

 	const publicURL = "https://storage.googleapis.com/%s/%s"

	guest_reserve_minor_slice[0].File_Name = f_name
    guest_reserve_minor_slice[0].Url       = fmt.Sprintf(publicURL, bucket, f_name )

///    set template

    monitor2 := template.Must(template.New("html").Parse(html6.Reserve_situation2))

    err = monitor2.Execute(w, guest_reserve_minor_slice)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

//	fmt.Fprintf( w, "reserve2 normal end \n" )

}

// NewBarChart returns a new bar chart with a single bar for each value.
// The bars heights correspond to the values and their x locations correspond
// to the index of their value in the Valuer.
func NewBarChart(vs plotter.Valuer, width vg.Length) (*BarChart, error) {

//    fmt.Println ("NewBarChart start " )

	if width <= 0 {
		return nil, errors.New("Width parameter was not positive")
	}
	values, err := plotter.CopyValues(vs)

	if err != nil {
		return nil, err
	}
//	fmt.Println ("NewBarChart end " )  // チE��チE��

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

//    fmt.Println (" BarHeight start i" ,i)

	ht := 0.0
	if b == nil {

//	   ht= ht + 2.

//	   fmt.Println (" BarHeight ht " ,ht )

		return ht
	}
	if i >= 0 && i < len(b.Values) {
		ht += b.Values[i]
	}
	if b.stackedOn != nil {
		ht += b.stackedOn.BarHeight(i)
	}

//	fmt.Println (" BarHeight end"  )

	return ht
}

// StackOn stacks a bar chart on top of another,
// and sets the XMin and Offset to that of the
// chart upon which it is being stacked.
func (b *BarChart) StackOn(on *BarChart) {

//    fmt.Println ("StackOn start " )

	b.XMin = on.XMin
	b.Offset = on.Offset
	b.stackedOn = on
}

///
/// Plot implements the plot.Plotter interface.
///

func (b *BarChart) Plot(c draw.Canvas, plt *plot.Plot) {

//    fmt.Println ("Plot start " )

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

//		fmt.Println ("Plot bottom " ,bottom )

        valMin := trVal(bottom+b.Values2[i])

		valMax := trVal(bottom + ht)

//		fmt.Println ("Plot valMin " ,valMin )

		var pts []vg.Point
		var poly []vg.Point
		if !b.Horizontal {

//		   fmt.Println ("Plot type1 "  )

			pts = []vg.Point{
				{catMin, valMin},
				{catMin, valMax},
				{catMax, valMax},
				{catMax, valMin},
			}
			poly = c.ClipPolygonY(pts)
		} else {                        // case XY

//		   fmt.Println ("Plot type2 "  )

			pts = []vg.Point{
				{valMin, catMin},
				{valMin, catMax},
//                {valMin+50, catMin},
//				{valMin+50, catMax},
				{valMax, catMax},
				{valMax, catMin},
			}

//			fmt.Println ("Plot 1" ,valMin, catMin)
//            fmt.Println ("Plot 2" ,valMin, catMax)
//            fmt.Println ("Plot 3" ,valMax, catMax)
//            fmt.Println ("Plot 4" ,valMax, catMax)

			poly = c.ClipPolygonX(pts)
		}

//		fmt.Println ("Plot pts " ,pts )

		c.FillPolygon(b.Color, poly)

		var outline [][]vg.Point
		if !b.Horizontal {              //    case normal

			pts = append(pts, vg.Point{X: catMin, Y: valMin})
			outline = c.ClipLinesY(pts)
		} else {                       //  case vy

			pts = append(pts, vg.Point{X: valMin, Y: catMin})
			outline = c.ClipLinesX(pts)
		}

//		fmt.Println ("Plot outline " ,outline )

		c.StrokeLines(b.LineStyle, outline...)
	}

//	fmt.Println ("Plot end " )  // チE��チE��

}


// DataRange implements the plot.DataRanger interface.
func (b *BarChart) DataRange() (xmin, xmax, ymin, ymax float64) {

//    fmt.Println ("DataRange start " )


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
	if !b.Horizontal {                              //  case normal
		return catMin, catMax, valMin, valMax
	}
//	fmt.Println ("DataRange  valMin" ,valMin)
//    fmt.Println ("DataRange  valMax" ,valMax)
//    fmt.Println ("DataRange  catMin" ,catMin)
//    fmt.Println ("DataRange  catMax" ,catMax)

//    fmt.Println ("DataRange end " )

	return valMin, valMax, catMin, catMax
}

// GlyphBoxes implements the GlyphBoxer interface.

func (b *BarChart) GlyphBoxes(plt *plot.Plot) []plot.GlyphBox {

//    fmt.Println ("GlyphBoxes start " )
//    fmt.Println ("GlyphBoxes plt " ,plt )

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
			boxes[i].Rectangle = vg.Rectangle{
				Min: vg.Point{Y: b.Offset - b.Width/2},
				Max: vg.Point{Y: b.Offset + b.Width/2},
			}
		}
	}

//	fmt.Println ("GlyphBoxes end "  )

	return boxes
}

// Thumbnail fulfills the plot.Thumbnailer interface.

func (b *BarChart) Thumbnail(c *draw.Canvas) {

//    fmt.Println ("Thumbnail start "  )

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

//	fmt.Println ("Thumbnail end "  )
}
