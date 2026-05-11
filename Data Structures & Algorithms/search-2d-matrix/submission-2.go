func searchMatrix(matrix [][]int, target int) bool {
	lr := 0
	lc := 0
	
	rr := len(matrix) - 1
	rc := len(matrix[0]) - 1
	
	cn := len(matrix[0]) - 1
	rn := -1

	for lr <= rr { 
		mr := lr + (rr - lr) / 2
		// fmt.Println(lr, mr ,rr)
		if target >= matrix[mr][0] && target <= matrix[mr][cn] {
			rn = mr
			break
		} else if target < matrix[mr][0] {
			rr = mr - 1
		} else {
			lr = mr + 1
		}
	}

	if rn == -1 {
		return false
	}

	// fmt.Println("rn = ", rn)

	for lc <= rc {
		mc := lc + (rc - lc) / 2
		if matrix[rn][mc] == target {
			return true
		}

		if target < matrix[rn][mc] {
			rc = mc - 1
		} else if target > matrix[rn][mc] {
			lc = mc + 1
		}
	}

	return false
}
