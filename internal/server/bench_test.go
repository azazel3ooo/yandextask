package server

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/azazel3ooo/yandextask/internal/models"
	"github.com/ugorji/go/codec"
)

const SetSize = 100000

func getData() []models.CustomIDSet {
	res := make([]models.CustomIDSet, SetSize)

	for i := 0; i < SetSize; i++ {
		res[i] = models.CustomIDSet{
			CorrelationID: "qwertyqwertyqwerty",
			ShortURL:      "http://localhost/qqq",
			OriginalURL:   "http://google.com",
		}
	}
	return res
}

func getByteData() []byte {
	res, _ := json.Marshal(getData())
	return res
}

func BenchmarkMarshal(b *testing.B) {
	var target []byte
	handle := new(codec.JsonHandle)
	obj := getData()

	b.ResetTimer()
	b.Run("std", func(b *testing.B) {
		for i := 0; i < 10; i++ {
			target, err := json.Marshal(obj)
			if err != nil {
				log.Println(err)
			}

			_ = target
		}
	})
	b.Run("codec", func(b *testing.B) {
		for i := 0; i < 10; i++ {
			encoder := codec.NewEncoderBytes(&target, handle)

			if err := encoder.Encode(obj); err != nil {
				log.Println(err)
			}
		}
	})
}

func BenchmarkUnmarshal(b *testing.B) {
	handle := new(codec.JsonHandle)
	var res []models.CustomIDSet

	data := getByteData()

	b.ResetTimer()
	b.Run("std", func(b *testing.B) {
		for i := 0; i < 10; i++ {
			if err := json.Unmarshal(data, &res); err != nil {
				log.Println(err)
			}
		}
	})
	b.Run("codec", func(b *testing.B) {
		for i := 0; i < 10; i++ {
			decoder := codec.NewDecoderBytes(data, handle)
			if err := decoder.Decode(&res); err != nil {
				log.Println(err)
			}
		}
	})
}
