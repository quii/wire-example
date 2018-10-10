package main

func NewDB() DB {
	return &FakeDB{}
}

type FakeDB struct {

}

func (*FakeDB) Get() int {
	return 42
}
