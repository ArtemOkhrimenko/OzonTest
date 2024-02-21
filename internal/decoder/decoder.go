package decoder

import (
	"fmt"

	"github.com/sqids/sqids-go"
)

type Decoder struct {
	decoder *sqids.Sqids
}

func New(minLength uint8, Alphabet string) (*Decoder, error) {
	dec, err := sqids.New(sqids.Options{MinLength: minLength, Alphabet: Alphabet})
	if err != nil {
		return nil, fmt.Errorf("sqids.New: %w", err)
	}

	return &Decoder{
		decoder: dec,
	}, nil
}

func (d *Decoder) Encode(id int) (string, error) {
	encodeID, err := d.decoder.Encode([]uint64{uint64(id)})
	if err != nil {
		return "", fmt.Errorf("d.decoder.Encode: %w", err)
	}

	return encodeID, err

}

func (d *Decoder) Decode(id string) int {
	decodeID := d.decoder.Decode(id)

	return int(decodeID[0])
}
