package html2

   const D_district_area = `
   <!DOCTYPE html>
   <html>
     <head>
        <meta charset="UTF-8">
        <title>データストアの更新・削除</title>
        <link rel="stylesheet" href="css/sgh/d_district_area.css" type="text/css">
     </head>
     <body>
       <header>



       </header>
       <nav>


	   </nav>
       <article>
         <table border="2" cellpadding="12" align="center" bgcolor="#00ced1">
         <h2 align="center">District Name</h2>

           <tr> <th>district-no</th> <th>district-name</th> </tr>
           {{range .}}
             {{ if eq .Area_No 1}}

             <tr>

             <form method="GET" action="/d_district_area_update" >


               <td>
                  <input type="text" name="district_no" size="5" align="center" value="{{.District_No|html}}" />
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>

               </td>
               <td>
                  <input type="text" name="district_name" size="10" align="center" value="{{.District_Name|html}}" />
               </td>
              </form>

            </tr>
            {{end}}
           {{end}}
           </table>


       </article>
       <section id="main">
           <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">
           <h2 align="center">Area Information</h2>

           <tr> <th>area-no</th><th>area-name  </th> <th>area-detail</th> <th>number-total</th><th>time-total</th> <th>productivity</th>

                                                <th>access1</th> <th>access2</th> <th>access3</th> <th>access4</th></tr>
           {{range .}}

             <tr>

             <form method="GET" action="/d_district_area_update" >

               <td>
                  <input type="text" name="area_no" size="5" align="center" value="{{.Area_No|html}}" />
               </td>
               <td>
                  <input type="text" name="area_name" size="15" align="center" value="{{.Area_Name|html}}" />
               </td>
               <td>
                  <input type="text" name="area_detail" size="30" align="center" value="{{.Area_Detail|html}}" />
               </td>
               <td>
                  <input type="text" name="area_detail" size="5" align="center" value="{{.Number_Total|html}}" />
               </td>
               <td>
                  <input type="text" name="area_detail" size="5" align="center" value="{{.Time_Total|html}}" />
               </td>
               <td>
                  <input type="text" name="area_detail" size="5" align="center" value="{{.Productivity|html}}" />
               </td>
               <td>
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
                  <input type="submit"  size="2" value="変更"  />
               </td>
             </form>

             <form method="GET" action="/d_district_area_delete" >

               <td>
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
                  <input type="submit"  size="2" value="削除"  />
               </td>
             </form>
             <form method="GET"  align="center" action="/deliver_showall2">
               <td >
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
                  <input type="submit" value="配達日報" />
               </td>
             </form>
             <form method="GET" action="/ai_sgh_ex_show" >
               <td>
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
                  <input type="submit"  size="2" value="expression-show"  />
               </td>
             </form>

           </tr>

           {{end}}
          </table>
          <table>
           <h2 align="center">Input New Data ,please</h2>
          <form method="GET" action="/d_district_area_show" >

              <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">

              <tr>   <th>area_name</th> <th>area-detail</th>  <th>access</th> </tr>

              <td > <input type="text" name="area_name" size="15" />            </td>
              <td > <input type="text" name="area_detail" size="30" />            </td>
              <td > <input type="submit" size="2"  value="登録" />           </td>
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
