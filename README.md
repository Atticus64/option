# Option in Golang

* Option structure inspired in Rust :crab:

Example:

```go
func main() {
	opt := Optional[string]("hello")

	if opt.IsSome() {
		value, _ := opt.GetValue()
		fmt.Println("value:", value)
	} else {
		fmt.Println("None")
	}
}
```

