package differenceofsquares

func SquareOfSum(n int) int {
	result:= 0

    for i:=1; i<=n; i++ {	
        result = result + i
    }

    return result * result;
}

func SumOfSquares(n int) int {
	result := 0

    for i:=1; i<=n; i++ {	
        result = result + (i * i)
    }

    return result
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
