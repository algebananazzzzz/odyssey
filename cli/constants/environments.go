package constants

var ENVIRONMENTS = map[int]string{
	1: "Production only (quick POCs)",
	2: "Production and Pre-production",
	// 3: "Production, Pre-production, and CI",
}

// EnvList returns a slice of short environment codes based on the selection.
func EnvList(selection int) []string {
	switch selection {
	case 1:
		return []string{"prd"}
	case 2:
		return []string{"prd", "preprod"}
	// case 3:
	// 	return []string{"prd", "preprod", "ci"}
	default:
		return []string{"prd"} // fallback
	}
}
