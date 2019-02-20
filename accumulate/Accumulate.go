package accumulate

// make function that takes collection of strings and does an operation on them.


func Accumulate (collection []string , operation func(string)string ) []string {

	for i, w := range collection {
		collection [i] = operation(w)
	}


	return collection
}