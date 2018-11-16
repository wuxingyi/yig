package _go

import (
	"testing"
)

const (
	TEST_KEY   = "testput"
	TEST_VALUE = "valueput"
)

func Test_Object_Prepare(t *testing.T) {
	sc := NewS3()
	err := sc.MakeBucket(TEST_BUCKET)
	if err != nil {
		t.Fatal("MakeBucket err:", err)
		panic(err)
	}
}

func Test_PutObject(t *testing.T) {
	sc := NewS3()
	err := sc.PutObject(TEST_BUCKET, TEST_KEY, TEST_VALUE)
	if err != nil {
		t.Fatal("PutObject err:", err)
	}
	t.Log("PutObject Success!")
}

func Test_HeadObject(t *testing.T) {
	sc := NewS3()
	err := sc.HeadObject(TEST_BUCKET, TEST_KEY)
	if err != nil {
		t.Fatal("HeadBucket err:", err)
	}
	t.Log("HeadObject Success!")
}

func Test_GetObject(t *testing.T) {
	sc := NewS3()
	v, err := sc.GetObject(TEST_BUCKET, TEST_KEY)
	if err != nil {
		t.Fatal("GetObject err:", err)
	}
	if v != TEST_VALUE {
		t.Fatal("GetObject err: value is:", v, ", but should be:", TEST_VALUE)
	}
	t.Log("GetObject Success value:", v)
}

func Test_DeleteObject(t *testing.T) {
	sc := NewS3()
	err := sc.DeleteObject(TEST_BUCKET, TEST_KEY)
	if err != nil {
		t.Fatal("DeleteObject err:", err)
	}
	err = sc.HeadObject(TEST_BUCKET, TEST_KEY)
	if err == nil {
		t.Fatal("HeadObject err:", err)
	}
	t.Log("DeleteObject Success!")
}

func Test_Object_End(t *testing.T) {
	sc := NewS3()
	err := sc.DeleteBucket(TEST_BUCKET)
	if err != nil {
		t.Fatal("DeleteBucket err:", err)
		panic(err)
	}
}
