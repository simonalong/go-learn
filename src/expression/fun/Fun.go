package fun

type CustomizeJudge struct {
}

func (customizeJudge *CustomizeJudge) Test1(name string) bool {
	if name == "zhou" {
		return true
	}
	return false
}

//func TestRun_fast_methods(t *testing.T) {
//	input := `hello() + world()`
//
//	tree, err := parser.Parse(input)
//	require.NoError(t, err)
//
//	env := map[string]interface{}{
//		"hello": func(...interface{}) interface{} { return "hello " },
//		"world": func(...interface{}) interface{} { return "world" },
//	}
//	funcConf := conf.New(env)
//	_, err = checker.Check(tree, funcConf)
//	require.NoError(t, err)
//
//	program, err := compiler.Compile(tree, funcConf)
//	require.NoError(t, err)
//
//	out, err := vm.Run(program, env)
//	require.NoError(t, err)
//
//	require.Equal(t, "hello world", out)
//}
