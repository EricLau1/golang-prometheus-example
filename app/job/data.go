package job

import (
	"encoding/json"
	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"time"
)

type Data struct {
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (d *Data) Bytes() []byte {
	b, _ := json.Marshal(d)
	return b
}

func create() *Data {
	return &Data{
		ID:        uuid.NewString(),
		Content:   faker.Paragraph(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
