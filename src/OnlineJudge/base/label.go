package base

func GenerateLabel(id int) string {
	return string(byte('A') + byte(id))
}

func LabelToInt(label string) int {
	return int(byte(label[0]) - byte('A'))
}
