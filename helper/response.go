package helper

func PesanGagalHelper(msg string) map[string]any {
	return map[string]any{
		"Status": "Failed",
		"MSG":    msg,
	}
}
func PesanSuksesHelper(msg string) map[string]any {
	return map[string]any{
		"Status": "Berhasil",
		"MSG":    msg,
	}
}
func PesanDataBerhasilHelper(msg string, data any) map[string]any {
	return map[string]any{
		"Status": "Data Berhasil",
		"MSG":    msg,
		"data":   data,
	}
}
