package tool

func CheckVideoType(str string) bool {
	videoExt := []string{"mp4", "avi", "rmvb", "mov", "mkv", "flv"}
	for _, v := range videoExt {
		if str == v {
			return true
		}
	}
	return false
}
