package content

import (
	"testing"
	"fmt"
)

const testHTML = `
<html>
  <head>
    <title>Sample "Hello, World" Application</title>
  </head>
  <body bgcolor=white>

    <table border="0" cellpadding="10">
      <tr>
        <td>
          <img src="images/springsource.png">
        </td>
        <td>
          <h1>Sample "Hello, World" Application</h1>
        </td>
      </tr>
    </table>
    <div id="0">
      <div id="1">Just two divs peacing out</div>
    </div>
    check
    <div id="2">One more</div>
    <p>This is the home page for the HelloWorld Web application. </p>
    <p>To prove that they work, you can execute either of the following links:
    <ul>
      <li>To a <a href="hello.jsp">JSP page</a></li>
      <li>To a <a href="hello">servlet</a></li>
    </ul>
    </p>
    <div id="3">
      <div id="4">Last one</div>
    </div>

  </body>
</html>
`

func TestNewContent(t *testing.T) {
	var data []byte = []byte(testHTML)
	cnt:= NewContent(data,"sad")
	//fmt.Println(cnt.GetString())
	d,e := cnt.GetDocument()
	if e != nil {
		t.Fail()
	}
	str,_ := d.Find("table[border='0']").Attr("cellpadding")
	fmt.Println(str)

}


