package to

/*
resolve the string as follow
key1=key2
key2=key3
*/

type Env struct{}

func (e Env) Encode(v interface{}) string {

	panic("implement me")
}

func (e Env) Decode(data []byte, v interface{}) {
	panic("implement me")
}
