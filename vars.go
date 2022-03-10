package spectrogram

import "flag"

var (
	OFFSET = flag.Uint64("offset", 0, "sey begin of samples")
	LENGTH = flag.Uint64("length", 0, "set number of samples [0 means all]")

	RATIO = flag.Float64("ratio", 0.8, "set ratio")

	WIDTH  = flag.Uint("width", 2048, "set width")
	HEIGHT = flag.Uint("height", 450, "set height")

	HIDEAVG    = flag.Bool("hideavg", false, "hide average")
	HIDERULERS = flag.Bool("hiderulers", false, "hide rulers")

	OUT = flag.String("out", "out.png", "set output filename")

	BINS      = flag.Uint("bins", 512, "set freq bins")
	PREEMP    = flag.Float64("preemp", 0.95, "pre-emphasis")
	RECTANGLE = flag.Bool("rectangle", false, "use rectangle window")
	DFT       = flag.Bool("dft", false, "use dft instead of fft")
	LOG10     = flag.Bool("log10", false, "pretty")
	MAG       = flag.Bool("mag", false, "mag")

	BG0 = flag.String("BG0", "000", "set background color 0")
	BG1 = flag.String("BG1", "333", "set background color 1")

	FG0 = flag.String("FG0", "0972a2", "set forground color 0")
	FG1 = flag.String("FG1", "6b5f7e", "set forground color 1")
	RUL = flag.String("RUL", "a0b0c0", "set rulers color")
)
