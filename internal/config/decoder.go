package config

type Decoder struct {
	MinLen   uint8  `envconfig:"MIN_LEN" required:"true"`
	Alphabet string `envconfig:"ALPHABET" required:"true"`
}
