// html
package main

import (
	"fmt"
)

func main() {
	name := "Ian Simpson"
	fmt.Println(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Go Rocks!</title>
	</head>
	<body>
	<h1>` +
	name +
	`</h1>
	</body>
	</html>
	`)
}
