package operator

//Add is the addition operator, 'a+b'
func Add(a, b int) int { return a + b }

//Mul is the multiplication operator, 'a*b'
func Mul(a, b int) int { return a * b }

//Div is the division operator, 'a/b'
func Div(a, b int) int { return a / b }

//Mod is the mod operator, 'a%b'
func Mod(a, b int) int { return a % b }

//Sub is the subtraction operator, 'a-b'
func Sub(a, b int) int { return a - b }

//LT is Less, the comparison operator, 'a<b'
func LT(a, b int) bool { return a < b }

//LTE is Less than or Equal, the comparison operator, 'a<=b'
func LTE(a, b int) bool { return a <= b }

//GT is Greater than, '>'
func GT(a, b int) bool { return a > b }

//GTE is Greater than or Equal, '>='
func GTE(a, b int) bool { return a >= b }
