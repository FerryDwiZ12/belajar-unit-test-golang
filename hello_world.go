package helper

// for runing test all gunain go test
// untuk yang ada dalam folder itu saja gunain go test -v
// untuk spesitik yang memiliki kesamaan menggunakan go test -v -run=TestHelloWorld maka dia akan test func yang ada kata HelloWorld
// untuk menjalankan semua unit test dari top folder atau terluar 

func HelloWorld(name string) string  {
	return "Hello" + name
}