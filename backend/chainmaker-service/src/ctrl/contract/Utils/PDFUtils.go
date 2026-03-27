package Utils

import (
	"ChainMaker_Backend_demo/Model"
	"bytes"
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"os"
)

// 根据输入数据生成PDF
func GeneratePdf(data Model.PDFData) *bytes.Buffer {

	currentDir, err := os.Getwd()
	if err != nil {
		// 如果发生错误，打印错误信息
		fmt.Println("Error getting current directory:", err)
	}
	fmt.Printf("Dir :  ", currentDir)

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	bgImagePath := "./ctrl/contract/Utils/碳排放溯源报告.png" // 确保路径正确
	pdf.Image(bgImagePath, 0, 0, 210, 297, false, "", 0, "")

	fontPath := "./ctrl/contract/Utils/雅黑.ttf" // 确保字体文件路径正确
	pdf.AddUTF8Font("ChineseFont", "", fontPath)

	pdf.SetFont("ChineseFont", "", 12) // 设置字体

	x, y := 31.0, 72.0
	fillPdfCell(pdf, &x, &y, 29, data.Reportid)
	fillPdfCell(pdf, &x, &y, 29, data.Txid)
	fillPdfCell(pdf, &x, &y, 30, data.CompanyName)
	fillPdfCell(pdf, &x, &y, 30, data.CompanyType)
	fillPdfCell(pdf, &x, &y, 29, data.CarbonModel)
	fillPdfCell(pdf, &x, &y, 29, data.TotalEmissions)
	fillPdfCell(pdf, &x, &y, 29, data.TimeStamp)

	// 输出PDF到字节缓冲区以便通过接口直接发送
	buf := new(bytes.Buffer)
	err = pdf.Output(buf)
	if err != nil {
		return nil
	}
	return buf
}

// 填充PDF单元格并更新Y坐标
func fillPdfCell(pdf *gofpdf.Fpdf, x *float64, y *float64, yIncrement float64, text string) {
	pdf.SetXY(*x, *y)
	pdf.CellFormat(190, 10, text, "0", 1, "L", false, 0, "")
	*y += yIncrement
}
