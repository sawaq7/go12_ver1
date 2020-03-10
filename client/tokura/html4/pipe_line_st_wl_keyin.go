package html4

   const Pipe_line_st_wl_keyin = `
   <!DOCTYPE html>
   <html>
     <head>
        <meta charset="UTF-8">
        <title>water line access</title>
        <link rel="stylesheet" href="css/tokura.css" type="text/css">
     </head>
     <body>
       <article>
       </article>

          <section >
            <table border="2" cellpadding="8" align="center" bgcolor="#cd5c5c">
            <h1  align="center">Register water line some data ,please</h1>

              <tr>  <th>section</th>   <th>f-facter</th>  <th>velocity</th>
                                       <th>p-diameter</th> <th>p-length</th><th>access1</th> <th>access2</th></tr>
            <form method="GET" action="/pipe_line_st_wl_show" >

              <td > <input type="text" name="section"  size="10" />        </td>

              <td > <input type="text" name="f_facter" size="5" />            </td>

              <td > <input type="text" name="velocity" size="5" />            </td>

              <td > <input type="text" name="p_diameter" size="5" />          </td>

              <td > <input type="text" name="p_length" size="5" />            </td>

              <td > <input type="submit" size="2"  value="register" />           </td>
            </form>

            <form method="GET" action="/pipe_line_ds_cal" >

               <td>

                  <input type="submit"  size="2" value="cal"  />
               </td>
             </form>
          </section>

          <section id="main">
           <table border="2" cellpadding="8" align="center" bgcolor="#cd5c5c">
           <h2 align="center">List Of Deliverly Situation</h2>

           <tr> <th>water-name</th>  <th>section</th>  <th>f-facter</th>  <th>velocity</th>
                <th>p-diameter</th>  <th>p-length</th> <th>access1</th>   <th>access2</th>  </tr>
           {{range .}}
             <tr>

             <form method="GET" action="/pipe_line_st_wl_update" >


               <td>
                  <input type="text" name="water_name" size="5" align="center" value="{{.Name|html}}" />
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>

               </td>
               <td>
                  <input type="text" name="section" size="10" align="center" value="{{.Section|html}}" />
               </td>
               <td>
                  <input type="text" name="f_facter" size="10" align="center" value="{{.Friction_Factor|html}}" />
               </td>
               <td>
                  <input type="text" name="velocity" size="10" align="center" value="{{.Velocity|html}}" />
               </td>
               <td>
                  <input type="text" name="p_diameter" size="10" align="center" value="{{.Pipe_Diameter|html}}" />
               </td>
               <td>
                  <input type="text" name="p_length" size="10" align="center" value="{{.Pipe_Length|html}}" />
               </td>
               <td>
                  <input type="submit"  size="2" value="change"  />

               </td>
             </form>

             <form method="GET" action="/pipe_line_st_wl_delete" >

               <td>
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
                  <input type="submit"  size="2" value="delete"  />
               </td>
             </form>

           </tr>
           {{end}}
           </table>

        </section>



     </body>
   </html>
   `
