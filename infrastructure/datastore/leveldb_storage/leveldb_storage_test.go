package leveldb_storage

import (
	"fmt"
	"github.com/photoshelf/photoshelf-storage/domain/model"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("with wrong directory (file)", func(t *testing.T) {
		dbPath := path.Join(os.TempDir(), "readonly")
		file, err := os.Create(dbPath)
		assert.NoError(t, err)
		file.Close()

		instance, err := NewLeveldbStorage(dbPath)
		if assert.Error(t, err) {
			assert.Nil(t, instance)
		}
	})

	t.Run("with correct directory", func(t *testing.T) {
		instance, err := NewLeveldbStorage(path.Join(os.TempDir(), "leveldb"))
		if assert.NoError(t, err) {
			assert.NotNil(t, instance)
		}
	})
}

func TestLeveldbStorage_Save(t *testing.T) {
	t.Run("save without identifier, generate new identifier", func(t *testing.T) {
		instance := createInstance(t)
		photo := model.NewPhoto(readTestData(t))

		identifier, err := instance.Save(*photo)
		if assert.NoError(t, err) {
			assert.NotNil(t, identifier)
		}

		instance.db.Close()
	})

	t.Run("save with identifier", func(t *testing.T) {
		instance := createInstance(t)
		photo := *model.PhotoOf(*model.IdentifierOf("testdata"), readTestData(t))

		identifier, err := instance.Save(photo)
		assert.NoError(t, err)

		t.Run("returns identifier has same value", func(t *testing.T) {
			actual := photo.Id()
			assert.EqualValues(t, actual.Value(), identifier.Value())
		})

		t.Run("stored same binary", func(t *testing.T) {
			actual, err := instance.db.Get([]byte("testdata"), nil)
			if err != nil {
				assert.Fail(t, "fail load data.")
			}
			assert.EqualValues(t, readTestData(t), actual)
		})

		instance.db.Close()
	})
}

func TestLeveldbStorage_Read(t *testing.T) {
	instance := createInstance(t)
	err := instance.db.Put([]byte("testdata"), readTestData(t), nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("with no key, returns err", func(t *testing.T) {
		_, err := instance.Read(*model.IdentifierOf("noKey"))
		assert.Error(t, err)
	})

	t.Run("returns same data with source", func(t *testing.T) {
		photo, err := instance.Read(*model.IdentifierOf("testdata"))
		if assert.NoError(t, err) {
			assert.EqualValues(t, readTestData(t), photo.Image())
		}
	})

	instance.db.Close()
}

func TestLeveldbStorage_Delete(t *testing.T) {
	instance := createInstance(t)
	err := instance.db.Put([]byte("testdata"), readTestData(t), nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("when delete existing key, returns no error", func(t *testing.T) {
		err := instance.Delete(*model.IdentifierOf("testdata"))
		if assert.NoError(t, err) {
			actual, _ := instance.db.Get([]byte("testdata"), nil)
			assert.EqualValues(t, []byte{}, actual)
		}
	})
}

func BenchmarkLeveldbStorage_Save(b *testing.B) {
	data := readTestData(b)

	b.Run("override", func(b *testing.B) {
		instance := createInstance(b)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			key := fmt.Sprintf("testdata-%d", 0)
			photo := *model.PhotoOf(*model.IdentifierOf(key), data)
			instance.Save(photo)
		}
		b.StopTimer()

		instance.db.Close()
	})

	b.Run("with new key", func(b *testing.B) {
		instance := createInstance(b)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			key := fmt.Sprintf("testdata-%d", i)
			photo := *model.PhotoOf(*model.IdentifierOf(key), data)
			instance.Save(photo)
		}
		b.StopTimer()

		instance.db.Close()
	})
}

func BenchmarkLeveldbStorage_Read(b *testing.B) {
	data := readTestData(b)
	instance := createInstance(b)
	for i := 0; i < 100; i++ {
		key := []byte(fmt.Sprintf("testdata-%d", i))
		if err := instance.db.Put(key, data, nil); err != nil {
			b.Fatal(err)
		}
	}

	b.Run("same data", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			key := fmt.Sprintf("testdata-%d", 0)
			instance.Read(*model.IdentifierOf(key))
		}
		b.StopTimer()

		instance.db.Close()
	})

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			key := fmt.Sprintf("testdata-%d", i)
			instance.Read(*model.IdentifierOf(key))
		}
		b.StopTimer()

		instance.db.Close()
	})
}

func readTestData(tb testing.TB) []byte {
	tb.Helper()

	testdataPath := path.Join(os.Getenv("GOPATH"), "src/github.com/photoshelf/photoshelf-storage", "testdata")
	body, err := os.Open(path.Join(testdataPath, "e3158990bdee63f8594c260cd51a011d"))
	if err != nil {
		tb.Fatal(err)
	}
	bytea, err := ioutil.ReadAll(body)
	if err != nil {
		tb.Fatal(err)
	}
	return bytea
}

func createInstance(tb testing.TB) *LeveldbStorage {
	tb.Helper()

	dataPath := path.Join(os.TempDir(), "leveldb")
	if err := os.RemoveAll(dataPath); err != nil {
		tb.Fatal(err)
	}

	instance, err := NewLeveldbStorage(dataPath)
	if err != nil {
		tb.Fatal(err)
	}
	return instance
}