package base

func GenerateLabel(id int) string {
	return string(byte('A') + byte(id))
}
