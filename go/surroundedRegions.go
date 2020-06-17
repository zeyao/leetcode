func solve(board [][]byte) {
	//https: //leetcode.com/problems/surrounded-regions/
	// & 指针只用定义一次
	//对于board slice[i] = '*'，在方法调用的上下文中，调用函数对slice引用的改表是看得见的
	//在对slice进行append操作的时候 如果不加 * &， 改变的就看不到了
	//原因就是append操作会返回这个扩展了的slice的引用， 必须让原引用重新赋值为新slice的引用
	//传递过来的这个指针原来指了内存中的A区域，A区域是原数组的真正所在。经过一次 append之后，要把这个指针改为指向B，B对应append后新的slice的引用。但是方法调用的上下文里的slice指针还是指向了老的A内存区域。

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			list := []pair{}
			if dfs(i, j, board, &list) {
				for k := 0; k < len(list); k++ {
					board[list[k].row][list[k].col] = 'X'
				}
			}
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '*' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfs(row int, col int, board [][]byte, list *[]pair) bool {
	if row < 0 || row >= len(board) || col < 0 || col >= len(board[0]) {
		return false
	}
	if board[row][col] != 'O' {
		return true
	}
	board[row][col] = '*'
	*list = append(*list, pair{row, col})
	l := dfs(row+1, col, board, list)
	r := dfs(row-1, col, board, list)
	u := dfs(row, col+1, board, list)
	d := dfs(row, col-1, board, list)
	return l && r && u && d
}

type pair struct {
	row int
	col int
}