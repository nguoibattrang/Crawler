package output

type Output interface {
	// Produce sends data to the provided string channel
	Produce(mChan <-chan string)
}
