package golangbasics

/ 1. HexToDecimal

/*
package main

import (
"fmt"
"strconv"
)

func HexToDecimal(hex string) (int64, error) {
decimal, err := strconv.ParseInt(hex, 16, 64)
if err != nil {
return 0, err
}

return decimal, nil
}

func main(){
fmt.Println(HexToDecimal("1E"))
fmt.Println(HexToDecimal("FF"))
}

*/

// 2. BinToDecimal

/*
package main

import (
"fmt"
"strconv"

)

func binToDecimal(binStr string) (int64, error) {
value, err := strconv.ParseInt(binStr, 2, 64)
if err != nil {
return 0, err
}

return value, nil
}
func main () {
fmt.Println(binToDecimal("10"))
}
*/

// 3. Capitalize a word

/*

package main

import (
"fmt"
"strings"
)

func Capitalize(word string) string {
return strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
}

func main (){
fmt.Println(Capitalize("hELLO"))
fmt.Println(Capitalize("WORLD"))
}

*/

// 4. Last two words to upper case

/*
package main

import (
"fmt"
"strings"
)
func LastTwoToUpper(s []string, n int) []string {
for i := n; i < len(s); i++ {
s[i] = strings.ToUpper(s[i])
}
return s
}

func main () {
fmt.Println(LastTwoToUpper([]string{"This", "is", "so", "exciting"}, 2))
}

*/

// 5. joinWithPunctuation

/*

package main

import (
"fmt"
"strings"
)

func joinWithPunctuation(tokens []string) string {
result := strings.Join(tokens, "")

// Replace the comma with a comma and a space
return strings.ReplaceAll(result, ",", ", ")
}


func main(){ // Removes the space after the opening single quote

fmt.Println(joinWithPunctuation([]string{"hello", ",", "world", "!"}))
}


*/

// 6. IsPunctuation

/*
package main

import (
"fmt"
)

func isPunctuation(s string) bool {
for _, r := range s {
if r == ',' || r == '!' {
return true
}
}
return false
}

func main(){
fmt.Println(isPunctuation(","))
fmt.Println(isPunctuation("!"))
fmt.Println(isPunctuation("x"))
}