package _go

import "testing"

const (
	TEST_BUCKET = "mybucket"
)

func Test_MakeBucket(t *testing.T) {
	err := MakeBucket(TEST_BUCKET)
	if err != nil {
		t.Fatal("MakeBucket err:", err)
	}
	t.Log("MakeBucket Success.")
}

func Test_DeleteBucket(t *testing.T) {
	err := DeleteBucket(TEST_BUCKET)
	if err != nil {
		t.Fatal("DeleteBucket err:", err)
	}
	t.Log("DeleteBucket Success.")
}
