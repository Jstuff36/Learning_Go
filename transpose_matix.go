func transpose(A [][]int) [][]int {
    transposedMatrix := make([][]int, len(A[0]))
    for i := range(A[0]) {
        transposedMatrix[i] = make([]int, len(A))
    }
    for x, row := range A {
        for y, num := range row {
            transposedMatrix[y][x] = num
        }
    }
    return transposedMatrix
}