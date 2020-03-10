package html5

   const Csv_show = `
   <!DOCTYPE html>
   <html>
     <head>
        <meta charset="UTF-8">
        <title>データストアの更新・削除</title>

        <link rel="stylesheet" href="css/general/csv_show.css" type="text/css">
     </head>
     <body>
       <header>


       </header>

       <nav>

       </nav>

       <article>

        <section >

         <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">
           <h2 align="left">column's command list</h2>

           {{range .}}
             {{ if eq .Line_No 1}}

             <tr> <th>command </th> <th>item</th> <th>input</th> </tr>

             <form method="GET" action="/csv_column_delete" >
              <tr>
               <th>
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
                  <button type="submit">delete<br></button>
               </th>

               <th>
                  <input type="text"  size="8" style="background-color:#dcdcdc" align="center" value="column-no"  />
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
               </th>

               <th>
                  <input type="text" name="delete_column" size="8" align="center"  />
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
               </th>
              </tr>
             </form>

             <form method="GET" action="/csv_column_join" >
              <tr>
               <th>
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>

                  <button type="submit">j o i n<br>f i l e</button>
               </th>
               <th>
                  <input type="text"  size="8" style="background-color:#dcdcdc" value="file-name"  />
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
                  <br><br>
                  <input type="text"  size="8" style="background-color:#dcdcdc" align="center" value="column-no" />
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
               </th>

               <th>
                  <input type="text" name="join_file" size="8" align="center"  />
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
                  <br><br>
                  <input type="text" name="join_column" size="8" align="center"  />
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
               </th>

             </form>

             <form method="GET" action="/csv_column_exchange" >

              <tr>
               <th>
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
                  <button type="submit">m o v e<br></button>

                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
                  <input type="hidden"  size="8" value="exchange(after)"  />

               </th>

               <th>
                  <input type="text"  size="8" style="background-color:#dcdcdc"  value="column1"  />
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
                  <br><br>
                  <input type="text"  size="8" value="column2" style="background-color:#dcdcdc" />
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
               </th>

               <th>
                  <input type="text" name="column1" size="8"  align="center"  />
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
                  <br><br>
                  <input type="text" name="column2" size="8" align="center"  />
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
               </th>
              </tr>

             </form>

              <tr>
               <th>
                <form method="GET" action="/csv_match_exp" >
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>

                  <button type="submit">matching<br>manual</button>
                </form>
               </th>

               <th>
                <form method="GET" action="/csv_match_wording" >
                 <input type="hidden" name="id"  value="{{.Id|html}}"/>

                 <button type="submit">matching<br>wording</button>
                </form>
               </th>

               <th>
                <form method="GET" action="/csv_match_file" >
                 <input type="hidden" name="id"  value="{{.Id|html}}"/>

                 <button type="submit">matching<br>file</button>
                </form>
               </th>

              </tr>

             <form method="GET" action="/csv_sort" >
              <tr>
               <th>
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
                  <button type="submit">s o r t<br></button>
               </th>

               <th>
                  <input type="text"  size="8" style="background-color:#dcdcdc" align="center" value="column-no"  />
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
               </th>

               <th>
                  <input type="text" name="sort_column" size="8" align="center"  />
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
               </th>
              </tr>
             </form>
             {{end}}
           {{end}}

          </table>

          <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">
           <h2 align="left">cullent file</h2>

           {{range .}}
             {{ if eq .Line_No 1}}

              <tr> <th>command </th> <th>item</th> <th>input</th> </tr>
              <form method="GET" action="/csv_make" >
              <tr>
               <th>
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>

                  <button type="submit">reserve or<br>make new-file</button>
               </th>
               <th>
                  <input type="text" style="background-color:#dcdcdc" size="8" align="center" value="file-name"  />
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
               </th>
               <th>
                  <input type="text" name="file_name" size="8" align="center" value="{{.File_Name|html}}" />
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
               </th>
              </tr>
             </form>

             {{end}}
           {{end}}

          </table>

        </section >

       </article>

       <section id="main">

           <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">
           <h2 align="center">Csv's List Information</h2>

           <tr> <th>line-no</th>  <th>column1</th> <th>column2</th> <th>column3</th><th>column4</th> <th>column5</th></tr>
           {{range .}}
             <tr>

             <form method="GET" action="/csv_update" >
               <td>
                  <input type="text"  size="5" align="center" value="{{.Line_No|html}}" />

                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
               </td>
               <td>
                  <input type="text" name="column1" size="10" align="center" value="{{.Column1|html}}" />

                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
               </td>
               <td>
                  <input type="text" name="column2" size="10" align="center" value="{{.Column2|html}}" />

                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
               </td>
               <td>
                  <input type="text" name="column3" size="10" align="center" value="{{.Column3|html}}" />

                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
               </td>
               <td>
                  <input type="text" name="column4" size="10" align="center" value="{{.Column4|html}}" />

                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
               </td>
               <td>
                  <input type="text" name="column5" size="10" align="center" value="{{.Column5|html}}" />

                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
               </td>

               <td>
                  <input type="submit"  size="2" value="update"  />
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
               </td>


             </form>

             <form method="GET" action="/csv_delete" >
               <td>
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
                  <input type="submit"  size="2" value="delete"  />
               </td>
             </form>

             <form method="GET" action="/csv_copy" >
               <td>
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
                  <input type="submit"  size="2" value="copy"  />
               </td>
             </form>

             <form method="GET" action="/csv_show" >
               <td>
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
                  <input type="submit"  size="2" value="storage access"  />
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
