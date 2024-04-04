package value

import "testing"

func TestCreateTweetID(t *testing.T) {
	t.Run("UUIDを値としたIDが発行できること", func(t *testing.T) {
		actual := CreateTweetID()
		// UUIDのため値の検証まで行わない
		if actual.ToString() == "" {
			t.Fatalf("id faild.")
		}
	})
}
