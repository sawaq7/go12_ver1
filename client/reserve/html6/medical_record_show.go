package html6

   const Medical_record_show = `
   <!DOCTYPE html>
   <html>
     <head>
        <meta charset="UTF-8">
        <title>uodate/delete in D.S.</title>
        <link rel="stylesheet" href="css/sgh/d_district_area.css" type="text/css">
     </head>
     <body>
       <header>



       </header>
       <nav>


	   </nav>
       <article>
         <table border="2" cellpadding="12" align="center" bgcolor="#00ced1">
         <h2 align="center">Current Guest</h2>

           <tr> <th>guest-no</th> <th>guest-name</th> </tr>
           {{range .}}

            {{ if eq .Line_No 1}}

             <tr>

             <form method="GET" action="/d_district_area_update" >


               <td>
                  <input type="text" name="district_no" size="5" align="center" value="{{.Guest_No|html}}" />
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>

               </td>
               <td>
                  <input type="text" name="district_name" size="10" align="center" value="{{.Guest_Name|html}}" />
               </td>
              </form>

            </tr>

            {{end}}

           {{end}}
           </table>


       </article>
       <section id="main">
           <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">
           <h2 align="center">Medical_Record Information</h2>

           <tr> <th>ope-date</th> <th>text</th> <th>access1</th> <th>access2</th> </tr>
           {{range .}}

             <tr>

             <form method="GET" action="/d_district_area_update" >

               <td>
                   <input type="text" name="date" size="12" value="{{.Date |html}}" />
              </td>

               <td>

                  <textarea name="text2" rows="6" cols="30"  >{{.Text|html}}</textarea>
               </td>

               <td>
                  <input type="submit"  size="2" value="変更"  />
               </td>
             </form>

             <form method="GET" action="/medical_record_delete" >

               <td>
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
                  <input type="submit"  size="2" value="delete"  />
               </td>
             </form>

           </tr>

           {{end}}
          </table>
          <table>
           <h2 align="center">Input New Information ,please</h2>
          <form method="GET" action="/medical_record_show2" >

              <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">

              <tr>   <th>ope-date</th> <th>text</th>  <th>access</th> </tr>

              <td>
                   <input type="text" name="date" size="12"  />
              </td>

              <td >
                    <textarea name="text2" rows="6" cols="30" >
                    </textarea>
              </td>

              <td >
                     <input type="submit" size="2"  value="register" />
              </td>

          </form>
         </table>
        </section>


        <aside>

        </aside>

        <footer>


        </footer>
     </body>
   </html>
   `
