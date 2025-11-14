package gw_cache

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var keats string = `A thing of beauty is a joy for ever:\n
Its loveliness increases; it will never\n
Pass into nothingness; but still will keep\n
A bower quiet for us, and a sleep\n
Full of sweet dreams, and health, and quiet breathing.\n
Therefore, on every morrow, are we wreathing\n
A flowery band to bind us to the earth,\n
Spite of despondence, of the inhuman dearth\n
Of noble natures, of the gloomy days,\n
Of all the unhealthy and o'er-darkened ways\n
Made for our searching: yes, in spite of all,\n
Some shape of beauty moves away the pall\n
From our dark spirits. Such the sun, the moon,\n
Trees old, and young, sprouting a shady boon\n
For simple sheep; and such are daffodils\n
With the green world they live in; and clear rills\n
That for themselves a cooling covert make\n
'Gainst the hot season; the mid forest brake,\n
Rich with a sprinkling of fair musk-rose blooms:\n
And such too is the grandeur of the dooms\n
We have imagined for the mighty dead;\n
All lovely tales that we have heard or read:\n
An endless fountain of immortal drink,\n
Pouring unto us from the heaven's brink.\n`

func TestLZ4Compressor_CompressAndDecompress(t *testing.T) {
	assert := assert.New(t)
	compressor := lz4Compressor{}

	// Test case 1: Compress and decompress an empty input
	input1 := []byte{}
	expected1 := []byte{}
	compressed1, err := compressor.Compress(input1)
	assert.NoError(err, "Error compressing input")
	decompressed1, err := compressor.Decompress(compressed1)
	assert.NoError(err, "Error decompressing input")
	assert.Equal(expected1, decompressed1, "Unexpected result")

	// Test case 2: Compress and decompress non-empty input
	input2 := []byte("Hello, World!")
	expected2 := input2
	compressed2, err := compressor.Compress(input2)
	assert.NoError(err, "Error compressing input")
	decompressed2, err := compressor.Decompress(compressed2)
	assert.NoError(err, "Error decompressing input")
	assert.Equal(expected2, decompressed2, "Unexpected result")

	//Test case 3: Compress and decompress big input
	input3 := []byte(keats)
	expected3 := input3
	compressed3, err := compressor.Compress(input3)
	assert.NoError(err, "Error compressing input")
	decompressed3, err := compressor.Decompress(compressed3)
	assert.NoError(err, "Error decompressing input")
	assert.Equal(expected3, decompressed3, "Unexpected result")

	// Test case 4: Decompress invalid input
	_, err = compressor.Decompress([]byte("invalid input"))

	assert.Error(err, "Expected error decompressing invalid input")
}

type errorWriter struct{}

func (ew *errorWriter) Write(p []byte) (n int, err error) {
	return 0, errors.New("forced error")
}
