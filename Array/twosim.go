func twoSum(nums []int, target int) []int {
    m:=make(map[int]int)

    for i , val := range nums{
         complement := target - val

         if idx , found:= m[complement]; found{
            return []int{i,idx}
         }

         m[val] = i
    }

    return []int{} 
}

// 🧱 1. Basics (Must Know)
// ✅ Variables & Types
// var a int = 10
// b := 20 // shorthand

// var name string = "Go"
// ✅ Arrays & Slices (VERY IMPORTANT)
// // Array (fixed size)
// var arr [3]int = [3]int{1, 2, 3}

// // Slice (dynamic - you'll use this mostly)
// nums := []int{1, 2, 3}
// nums = append(nums, 4)

// 👉 Think of slices like dynamic arrays in JavaScript.

// ✅ Looping
// for i := 0; i < len(nums); i++ {
//     fmt.Println(nums[i])
// }

// // Cleaner
// for i, val := range nums {
//     fmt.Println(i, val)
// }
// ✅ If-Else
// if x > 10 {
//     fmt.Println("big")
// } else {
//     fmt.Println("small")
// }
// 🧠 2. Functions (Core for DSA)
// func add(a int, b int) int {
//     return a + b
// }

// Multiple returns (very Go-like):

// func divide(a, b int) (int, int) {
//     return a / b, a % b
// }
// 🗺️ 3. Maps (CRITICAL for problems)
// m := make(map[int]int)

// m[1] = 100
// m[2] = 200

// val, exists := m[1]
// if exists {
//     fmt.Println(val)
// }

// 👉 You’ll use this in:

// Two Sum
// Frequency counting
// Hashing problems
// 📦 4. Strings
// s := "hello"

// // Access char
// fmt.Println(s[0]) // byte (not string)

// // Loop
// for i := 0; i < len(s); i++ {
//     fmt.Println(string(s[i]))
// }
// 🔁 5. Important Patterns
// 🔹 Swap
// a, b := 5, 10
// a, b = b, a
// 🔹 Reverse Slice
// for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
//     nums[i], nums[j] = nums[j], nums[i]
// }
// 🔹 Frequency Count
// freq := make(map[int]int)

// for _, num := range nums {
//     freq[num]++
// }