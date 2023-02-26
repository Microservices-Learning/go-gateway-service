package model

import (
	"encoding/json"
	"google.golang.org/grpc/metadata"
)

type CacheableMetadata map[string]string

func (s CacheableMetadata) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (s CacheableMetadata) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &s)
}

func (s CacheableMetadata) AsGrpcMetadata() metadata.MD {
	md := make(metadata.MD)

	for key, value := range s {
		md.Set(key, value)
	}

	return md
}

func (s CacheableMetadata) Put(key string, value string) {
	s[key] = value
}

func (s CacheableMetadata) Get(key string) string {
	return s[key]
}
