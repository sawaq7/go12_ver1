package html2

   const D_schedule_showall_05 = `
   <!DOCTYPE html>
   <html>
     <head>
        <meta charset="UTF-8">
        <title>チE�Eタストアの更新・削除</title>
        <link rel="stylesheet" href="css/member1_show.css" type="text/css">
     </head>
     <body>


       <article>
       </article>
       <section>
            <form method="GET" action="/d_schedule_showall" >

              <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">
              <h1  align="center">Register Your deliverly-number ,please</h1>
              {{range .}}
                {{ if eq .Id 77}}

                  <tr> <th>{{.Car_Name_01|html}}号軁E/th> <th>{{.Car_Name_02|html}}号軁E/th>  <th>{{.Car_Name_03|html}}号軁E/th> <th>{{.Car_Name_04|html}}号軁E/th>
                       <th>{{.Car_Name_05|html}}号軁E/th> <th>access1</th> </tr>

                {{end}}
              {{end}}
              <tr>

                <td > <input type="text" name="course_no_01" size="5" />            </td>
                <td > <input type="text" name="course_no_02" size="5" />            </td>
                <td > <input type="text" name="course_no_03" size="5" />            </td>
                <td > <input type="text" name="course_no_04" size="5" />            </td>
                <td > <input type="text" name="course_no_05" size="5" />            </td>
                <td > <input type="submit" size="2"  value="登録"  />                </td>

              </tr>


            </form>
       </section>
       <section id="main">


           <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">
           <h2 align="center">List Of Deliverly Situation</h2>

           {{range .}}
             {{ if eq .Id 77}}

               <tr> <th>district-no</th> <th>date</th>  <th>{{.Car_Name_01|html}}号軁E/th> <th>{{.Car_Name_02|html}}号軁E/th>  <th>{{.Car_Name_03|html}}号軁E/th> <th>{{.Car_Name_04|html}}号軁E/th>
                                                        <th>{{.Car_Name_05|html}}号軁E/th>

                                                    <th>ex-num</th>  <th>judge</th> <th>access1</th><th>access2</th> <th>access3</th></tr>


             {{else}}

               <tr>

               <form method="GET" action="/d_schedule_update" >
                 <td>
                   <input type="text" name="district_no" size="12" value="{{.District_No|html}}" />
                 </td>

                 <td>
                   <input type="text" name="date" size="12" value="{{.Date |html}}" />
                 </td>

                 <td>

                   <input type="text" name="course_no_01" size="3" align="center" value="{{.Course_01|html}}" />
                   <input type="hidden" name="id"  value="{{.Id|html}}"/>
                 </td>
                 <td>
                   <input type="text" name="course_no_02" size="3" align="center" value="{{.Course_02|html}}" />
                 </td>

                 <td>
                   <input type="text" name="course_no_03" size="3" align="center" value="{{.Course_03|html}}" />
                 </td>

                 <td>
                   <input type="text" name="course_no_04" size="3" align="center" value="{{.Course_04|html}}" />
                 </td>

                 <td>
                   <input type="text" name="course_no_05" size="3" align="center" value="{{.Course_05|html}}" />
                 </td>

                 <td>
                   <input type="text" name="expected_num" size="3" align="center" value="{{.Expected_Num|html}}" />
                 </td>

                 <td>
                   <input type="text" name="judge" size="15" align="center" value="{{.Judge|html}}" />
                 </td>

                 <td>
                   <input type="submit"  size="2" value="変更"  />
                 </td>
               </form>

               <form method="GET" action="/d_schedule_delete" >
                 <td>
                   <input type="hidden" name="id"  value="{{.Id|html}}"/>
                   <input type="submit"  size="2" value="削除"  />
                 </td>
               </form>
               <form method="GET" action="/d_schedule_copy" >
                 <td>
                   <input type="hidden" name="id"  value="{{.Id|html}}"/>
                   <input type="submit"  size="2" value="コピ�E"  />
                 </td>
               </form>
             </tr>

             {{end}}
           {{end}}
           </table>

        </section>
     </body>
   </html>
   `
