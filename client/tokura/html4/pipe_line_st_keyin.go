package html4

   const Pipe_line_st_keyin = `
   <!DOCTYPE html>
   <html>
     <head>
        <meta charset="UTF-8">
        <title>update/delete Datastore</title>
        <link rel="stylesheet" href="css/tokura.css" type="text/css">
     </head>
     <body>
       <article>
       </article>

          <section >
            <form method="GET" action="/pipe_line_st_show" >

              <table border="2" cellpadding="8" align="center" bgcolor="#cd5c5c">
              <h1  align="center">Register Your deliverly-number ,please</h1>
              <tr> <th>water-name</th> <th>water-high</th>  <th>r-facter</th> <th>access</th> </tr>

              <td > <input type="text" name="water_name"  size="10" />               </td>
              <td > <input type="text" name="water_high" size="5" />            </td>
              <td > <input type="text" name="r_facter" size="5" />            </td>

              <td > <input type="submit" size="2"  value="登録" />           </td>
            </form>
          </section>
          <section id="main">


           <table border="2" cellpadding="8" align="center" bgcolor="#cd5c5c">
           <h2 align="center">List Of Deliverly Situation</h2>

           <tr> <th>water-name</th> <th>water-high</th>  <th>r-facter</th>
                                    <th>access1</th> <th>access2</th> <th>access3</th></tr>
           {{range .}}
             <tr>

             <form method="GET" action="/pipe_line_st_update" >


               <td>
                  <input type="text" name="water_name" size="5" align="center" value="{{.Name|html}}" />
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>

               </td>
               <td>
                  <input type="text" name="water_high" size="10" align="center" value="{{.High|html}}" />
               </td>
               <td>
                  <input type="text" name="r_facter" size="10" align="center" value="{{.Roughness_Factor|html}}" />
               </td>

               <td>
                  <input type="submit"  size="2" value="変更"  />

               </td>
             </form>

             <form method="GET" action="/pipe_line_st_delete" >

               <td>
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
                  <input type="submit"  size="2" value="削除"  />
               </td>
             </form>
             <form method="GET" action="/pipe_line_st_wl_keyin" >

               <td>
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
                  <input type="submit"  size="2" value="water-line"  />
               </td>
             </form>

           </tr>
           {{end}}
           </table>

        </section>



     </body>
   </html>
   `
