package html2

   const Deliver_showall1 = `
   <!DOCTYPE html>
   <html>
     <head>
        <meta charset="UTF-8">
        <title>datastore update/delete</title>

        <link rel="stylesheet" href="css/sgh/deliver_showall1.css" type="text/css">
     </head>
     <body>
       <header>
         <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">
           <h2 align="center">Key-In Todays' Deliverly number　,please</h2>
             <form method="GET" action="/deliver_showall1" >

              <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">

              <tr>  <th>car-no</th>  <th>private-no</th> <th>deliverly-no</th>  <th>access</th> </tr>

              <td > <input type="text" name="car_no" size="5" />            </td>
              <td > <input type="text" name="private_no" size="5" />            </td>
              <td > <input type="text" name="number" size="5" />            </td>
              <td > <input type="submit" size="2"  value="登録" />           </td>
            </form>
           </table>



       </header>
       <nav>

       </nav>
       <article>
         <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">
           <h2 align="center">Current Course-No</h2>

           <tr> <th>course-no</th>
           {{range .}}

             {{ if eq .Line_No 1}}


               <tr>

               <td>
                  <input type="text" name="course_no" size="5" align="center" value="{{.Course_No|html}}" />
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
               </td>

               </tr>

             {{end}}
           {{end}}

           </table>
       </article>

       <section id="main">

           <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">
           <h2 align="center">List Of Deliverly Situation</h2>

           <tr> <th>date</th>  <th>car-no</th> <th>private-no</th><th>deliverly-no</th><th>access1</th> <th>access2</th><th>access3</th></tr>
           {{range .}}
             <tr>

             <form method="GET" action="/deliver_update" >
               <td>
                  <input type="text" name="date" size="12" value="{{.Date |html}}" />
               </td>

               <td>
                  <input type="text" name="car_no" size="2" align="center" value="{{.Car_No |html}}" />
               </td>
               <td>
                  <input type="text" name="private_no" size="5" align="center" value="{{.Private_No |html}}" />
               </td>

               <td>
                  <input type="text" name="number" size="2" align="center" value="{{.Number |html}}" />
               </td>

               <td>
                  <input type="submit"  size="2" value="変更"  />
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
               </td>
             </form>

             <form method="GET" action="/deliver_delete" >
               <td>
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
                  <input type="submit"  size="2" value="削除"  />
               </td>
             </form>
             <form method="GET" action="/deliver_copy" >
               <td>
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
                  <input type="submit"  size="2" value="コピ�E"  />
               </td>
             </form>
           </tr>
           {{end}}
           </table>

        </section>
        <aside>

        </aside>

        <footer>

        </footer>

     </body>
   </html>
   `
