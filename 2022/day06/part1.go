package day06

func FindStartOfPacketMarker(input string) (int, error) {
	return findStartOfMarker(input, packetSize)
}
