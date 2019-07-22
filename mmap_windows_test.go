package gowindows

import (
	"bytes"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestAll(t *testing.T) {
	m1, err := CreateMmap("name111111", 4096, true)
	if err != nil {
		t.Fatal(err)
	}

	m2, err := CreateMmap("name111111", 4096, false)
	if err != nil {
		t.Fatal(err)
	}

	m3, err := OpenMmap("name111111", 4096, false)
	if err != nil {
		t.Fatal(err)
	}

	b1 := m1.GetBytes()
	b2 := m2.GetBytes()
	b3 := m3.GetBytes()

	b := []byte{1, 4, 7, 8, 5, 2, 9, 6, 3}
	copy(b1, b)

	if v := b2[:len(b)]; bytes.Equal(v, b) == false {
		t.Errorf("%#v!=%#v", v, b)
	}
	if v := b3[:len(b)]; bytes.Equal(v, b) == false {
		t.Errorf("%#v!=%#v", v, b)
	}

	err = m1.Close()
	if err != nil {
		t.Error(err)
	}

	/* 测试确保 runtime.SetFinalizer 正常工作
	for i := 0; i < 1000; i++ {
		runtime.GC()
		time.Sleep(1 * time.Second)
	}*/
}

// 连续测试创建+关闭
func TestCreateMmapAndOpenAndClose(t *testing.T) {
	const name = "name111111"
	m1, err := CreateMmap(name, 4096, true)
	if err != nil {
		t.Fatal(err)
	}
	defer m1.Close()

	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			var m *Mmap
			var err error

			time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)

			if rand.Int31n(2) == 1 {
				m, err = CreateMmap(name, 4096, true)
				if err != nil {
					t.Fatal(err)
				}
			} else {
				m, err = OpenMmap(name, 4096, true)
				if err != nil {
					t.Fatal(err)
				}
			}

			time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)

			err = m.Close()
			if err != nil {
				t.Fatal(err)
			}
		}()
	}

	wg.Wait()
}

func TestCreateMmapWithSecurityDescriptor(t *testing.T) {
	m, err := CreateMmapWithSecurityDescriptor("aaa", 1024, true, "")
	if err != nil {
		t.Fatal(err)
	}

	if m == nil {
		t.Error("m==nil")
	}

	// 热议用户都有权限读写
	m2, err := CreateMmapWithSecurityDescriptor("aaa2", 1024, true, "D:P(A;OICI;GWGR;;;SY)(A;OICI;GWGR;;;BA)(A;OICI;GWGR;;;IU)(A;OICI;GWGR;;;RC)")
	if err != nil {
		t.Fatal(err)
	}

	if m2 == nil {
		t.Error("m2==nil")
	}

}
