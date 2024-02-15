# Think Money checkout kata

### Running the solution:
`go run .`

expected output:

```
total price:  240
unrecognized SKUs:  [O P Z: x2 T]
```

### Testing:
In root folder: `go test ./...`

### Potential changes:

- Have the product list in either a separate go file struct or 
 .json file and read from / unmarshal the data when creating a checkout with `New()` 

- Find a better solution than the `equalSlices` helper for comparing the`unrecognizedSKU` and `tc.ExpectedUnrecognized` slices - forgot DeepEqual didn't order slices when comparing
- Maybe learn how to make the program a CLI tool to pass in individual SKUs with a command to calculate the total when finished,
rather than having to declare a list of scanned items prior to running the program