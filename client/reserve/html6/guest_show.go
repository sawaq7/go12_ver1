package html6

   const Guest_show = `
   <!DOCTYPE html>
   <html>
     <head>
        <meta charset="UTF-8">
        <title>uodate/delete in D.S.</title>
        <link rel="stylesheet" href="css/member1_show.css" type="text/css">
     </head>
     <body>
       <article>
       </article>
       <section>
          <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">
             <h1  align="center">Register Your deliverly-number ,please</h1>

            <form method="GET" action="/guest_show2" >

              <tr> <th>guest-no</th> <th>guest-name</th>  <th>access</th> </tr>
              <td >
                    <input type="text" name="guest_no"  size="5" />
              </td>

              <td >
                    <input type="text" name="guest_name" size="10" />
              </td>

              <td >
                    <input type="submit" size="2"  value="register" />
              </td>

            </form>
          </table>

          <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">
            <h1  align="center">Register Your deliverly-number ,please</h1>

              <tr> <th>reserve-date</th> <th>access</th><th>reserve-date</th><th>access</th>
                                                          <th>reserve-date</th><th>access</th></tr>

              <form method="GET" action="/reserve_situation" >

               <td>
                <input type="text" name="reserve_date"  size="5" />
               </td>

               <td>
                <input type="submit"  size="2" value="reserve situation"  />
               </td>

              </form>

              <form method="GET" action="/reserve_situation2" >

               <td>
                <input type="text" name="reserve_date"  size="5" />
               </td>

               <td>
                <input type="submit"  size="2" value="reserve situation2"  />
               </td>

              </form>
              <form method="GET" action="/reserve_situation3" >

               <td>
                <input type="text" name="reserve_date"  size="5" />
               </td>

               <td>
                <input type="submit"  size="2" value="reserve situation3"  />
               </td>

              </form>

          </table>
       </section>

       <section id="main">


           <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">
           <h2 align="center">List Of Deliverly Situation</h2>

           <tr> <th>guest-no</th> <th>guest-name</th> <th>access1</th> <th>access2</th>
                                  <th>access3</th> <th>access4</th> <th>access5</th> <th>access6</th> </tr>
           {{range .}}
             <tr>

             <form method="GET" action="/d_district_update" >


               <td>
                  <input type="text" name="guest_no" size="5" align="center" value="{{.Guest_No|html}}" />
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>

               </td>
               <td>
                  <input type="text" name="guest_name" size="10" align="center" value="{{.Guest_Name|html}}" />
               </td>

               <td>
                  <input type="submit"  size="2" value="変更"  />
               </td>
             </form>

             <form method="GET" action="/guest_delete" >

               <td>
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
                  <input type="submit"  size="2" value="削除"  />
               </td>
             </form>
             <form method="GET" action="/reserve_register" >

               <td>
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
                  <input type="submit"  size="2" value="reserve"  />
               </td>
             </form>
             <form method="GET" action="/medical_record_show" >
               <td>
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
                  <input type="submit"  size="2" value="medical record"  />
               </td>
             </form>
             <form method="GET" action="/payment_register" >
               <td>
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
                  <input type="submit"  size="2" value="payment"  />
               </td>
             </form>
             <form method="GET" action="/medical_xray_show" >

               <td>
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
                  <input type="submit"  size="2" value="x-ray"  />
               </td>
             </form>

           </tr>
           {{end}}
           </table>

        </section>


     </body>
   </html>
   `
