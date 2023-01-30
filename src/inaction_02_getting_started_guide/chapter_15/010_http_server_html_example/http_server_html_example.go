package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const form =
	`<html>
		<body>
			<form action="#" method="post" name="bar">
				<input type="text" name="numbers" />
				<input type="submit" value="submit"/>
			</form>

			<table border="2px" width="600">
				<tbody>
					<tr>
						<td width="300">数据</td>
						<td>%v</td>
					</tr>
					<tr>
						<td>sum</td>
						<td>%v</td>
					</tr>
					<tr>
						<td>average</td>
						<td>%v</td>
					</tr>
				</tbody>
			</table>
		</body>
	</html>`

func main() {

	http.HandleFunc("/execute", ExecuteHandler)

	err := http.ListenAndServe("localhost:8888", nil)

	if err != nil {
		log.Fatal(err)
	}

}

func ExecuteHandler(response http.ResponseWriter, request *http.Request)  {

	params := request.FormValue("numbers")
	if params == "" {
		page := strings.ReplaceAll(form, "<td>%v</td>", "<td></td>")
		io.WriteString(response, page)
	} else {
		numbes := strings.Split(params, ",")
		sum, average := DoExecute(numbes)
		fmt.Fprintf(response, form, numbes, sum, average)
	}

}

func DoExecute(numbers []string) (int, float64)  {

	sum, average := 0, 0.0
	count := len(numbers)

	for _, number := range numbers {
		number = strings.TrimSpace(number)
		num, _ := strconv.Atoi(number)
		sum += num
	}

	average = float64(sum/count)
	return sum, average
}
