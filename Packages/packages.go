package main 
  
// Importing multiple packages 
import ( 
    "bytes"
    "fmt"
    "sort"
) 
  
func main() { 
  
    // Creating and initializing slice 
    // Using shorthand declaration 
    slice_1 := []byte{'*', 'G', 'e', 'e', 'k', 's', 'f', 
        'o', 'r', 'G', 'e', 'e', 'k', 's', '^', '^'} 
    slice_2 := []string{"Gee", "ks", "for", "Gee", "ks"} 
  
    // Displaying slices 
    fmt.Println("Original Slice:") 
    fmt.Printf("Slice 1 : %s", slice_1) 
    fmt.Println("\nSlice 2: ", slice_2) 
  
    // Trimming specified leading 
    // and trailing Unicode points 
    // from the given slice of bytes 
    // Using Trim function 
    res := bytes.Trim(slice_1, "*^") 
    fmt.Printf("\nNew Slice : %s", res) 
  
    // Sorting slice 2 
    // Using Strings function 
    sort.Strings(slice_2) 
    fmt.Println("\nSorted slice:", slice_2) 
}