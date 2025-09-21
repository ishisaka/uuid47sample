//UUID47を使用して、uuidv7をuuidv4風に変換し、また元に戻すサンプルコード

// main.go
package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/n2p5/uuid47"
)

func main() {
	// uuid v7を生成
	u, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}
	fmt.Printf("uuid v7: %s\n", u.String())

	// 変換用キーを作成する
	key := uuid47.Key{K0: 0x0123456789abcdef, K1: 0xfedcba9876543210}
	// 乱数でキーを作成する場合には以下のようにする:
	// key, _ := uuid47.NewRandomKey()

	// uuid.UUIDをuuid47.UUIDに変換
	var v7 uuid47.UUID
	copy(v7[:], u[:])

	// uuidv7をuuidv4風に変換して外部公開用にする
	v4facade := uuid47.Encode(v7, key)
	fmt.Printf("External ID(uuidv4): %s\n", v4facade)

	// uuidv4風からuuidv7に戻す
	original := uuid47.Decode(v4facade, key)
	fmt.Printf("Internal ID(uuidv7): %s\n", original)

}
