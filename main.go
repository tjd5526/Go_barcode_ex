package main

import(
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/boombuler/barcode/pdf417"
	"github.com/boombuler/barcode/datamatrix"
)

func main(){
	qrcode, _ := qr.Encode("Hell world",qr.Q, qr.Auto)
	qrcode, _ = barcode.Scale(qrcode, 200,200)
	file, _ := os.Create("Qrcode.png")
	defer file.Close()

	png.Encode(file, qrcode)

	pdf417, _ :=pdf417.Encode("Hell world", 5)
	pdf417, _ = barcode.Scale(pdf417, 300,100)
	file2, _ := os.Create("pdf417.png")
	defer file2.Close()

	png.Encode(file2, pdf417)

	matrix, _ := datamatrix.Encode("Hell world")
	matrix, _ = barcode.Scale(matrix, 200,200)
	file3, _ := os.Create("matrix.png")
	defer file3.Close()

	png.Encode(file3, matrix)

}
