package day10

import (
	"strings"
)

/*
It seems like the X register controls the horizontal position of a sprite.
Specifically, the sprite is 3 pixels wide, and the X register sets the horizontal position of the middle of that
sprite.
(In this system, there is no such thing as "vertical position": if the sprite's horizontal position puts its pixels
where the CRT is currently drawing, then those pixels will be drawn.)

You count the pixels on the CRT: 40 wide and 6 high.
This CRT screen draws the top row of pixels left-to-right, then the row below that, and so on.
The left-most pixel in each row is in position 0, and the right-most pixel in each row is in position 39.

Like the CPU, the CRT is tied closely to the clock circuit: the CRT draws a single pixel during each cycle.
Representing each pixel of the screen as a #, here are the cycles during which the first and last pixel in each row
are drawn:

Cycle   1 -> ######################################## <- Cycle  40
Cycle  41 -> ######################################## <- Cycle  80
Cycle  81 -> ######################################## <- Cycle 120
Cycle 121 -> ######################################## <- Cycle 160
Cycle 161 -> ######################################## <- Cycle 200
Cycle 201 -> ######################################## <- Cycle 240

So, by carefully timing the CPU instructions and the CRT drawing operations, you should be able to determine whether the sprite is visible the instant each pixel is drawn.
If the sprite is positioned such that one of its three pixels is the pixel currently being drawn, the screen produces a lit pixel (#); otherwise, the screen leaves the pixel dark (.).
*/

func (cc *ClockCircuit) RenderPixels() string {
	tape := cc.Tape()
	screen := strings.Builder{}

	for y := 0; y < 6; y++ {
		for x := 0; x < 40; x++ {
			spritePosition := tape[x+(y*40)]

			if spritePosition-1 <= x && x <= spritePosition+1 {
				screen.WriteString("#")
			} else {
				screen.WriteString(".")
			}

		}

		screen.WriteString("\n")
	}

	return screen.String()
}

func RenderImage(input string) (string, error) {
	cc := NewClockCircuit()

	if err := cc.ParseProgram(input); err != nil {
		return "", err
	}

	return cc.RenderPixels(), nil
}
