package common

func Min(a int32, b int32) int32 {
	if a > b {
		return b
	}
	return a
}

func Max(a int32, b int32) int32 {
	if a > b {
		return a
	}
	return b
}
