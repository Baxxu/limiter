### limiter

# Goroutines limiter
Ограничивает максимальное количество горутин.
## Get
```
go get github.com/Baxxu/limiter
```
## Example
```
func example() {
	//max 10 000 goroutines
	Limiter := limiter.New(10 * 1000)

	for i := 0; i < 1*1000*1000; i++ {
		//before goroutine
		Limiter.Add()
		go func() {
			//when goroutine ends
			defer Limiter.Remove()
			time.Sleep(time.Second * 1)
		}()
	}
}
```
Или
```
func example1() {
	//max 10 000 goroutines
	Limiter := limiter.New(10 * 1000)

	for i := 0; i < 1*1000*1000; i++ {
		Limiter.Queue(func() {
			time.Sleep(time.Second * 1)
		})
	}
}
```
