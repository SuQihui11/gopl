package main

func main() {
	type Weekday int

	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	print(Sunday)

	type Flags uint

	const (
		FlagUp           Flags = 1 << iota // is up
		FlagBroadcast                      // supports broadcast access capability
		FlagLoopback                       // is a loopback interface
		FlagPointToPoint                   // belongs to a point-to-point link
		FlagMulticast                      // supports multicast access capability
	)

}
