package main

import (
	_ "LSP/after/printarea"
	_ "LSP/after/rect"
	_ "LSP/after/square"
	"LSP/before/printarea"
	"LSP/before/rect"
	"LSP/before/square"
	_ "fmt"
)


func main () {

	//violate lsp
	printarea := printarea.CreatePrintArea();
	rct := &rect.Rect{}
	rct.SetHeight(5)
	rct.SetHeight(10)
	sqr := &square.Square{}
	sqr.SetHeight(10)
	sqr.SetWidth(10)
	printarea.Print(rct)
	printarea.Print(sqr)


	// true lsp 
	// printarea := printarea.CreatePrintArea();
	// rct := rect.SetRect(5,10);
	// sqr := square.SetSqr(10);
	// printarea.Print(rct)
	// printarea.Print(sqr)

}
