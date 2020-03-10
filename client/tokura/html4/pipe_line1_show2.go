    ///////////////////////////////////////////////////////
   ///                                                 ///
  ///     template　for ( pipe_line1_show )            ///
 ///                                                  ///
////////////////////////////////////////////////////////

package html4

const Pipe_line1_show2 = `
   <!DOCTYPE html>
   <html>

     <head>
       <meta charset="UTF-8">
       <title>pipe_line1_show</title>
       <link rel="stylesheet" href="css/css/tokura.css" type="text/css">
     </head>

     <body>


       <section>
          <table border="2" cellpadding="8" align="center" bgcolor="#cd5c5c">

          <h2 align="center"> excute all </h2>
          <tr> <th>水路名</th> <th>実行</th> </tr>

          <tr>

          {{.Name|html}}
          {{.No|html}}
              <
          </tr>

          </table>
       </section>
     </body>
   </html>
`
