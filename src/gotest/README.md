## golang test使用说明
* 每一个test文件必须import一个testing.
* test文件下每一个test case均必须以Test开头并符合TestXxx形式,否则go test会直接挑过测试不执行.
* go test会自动寻找该目录下的test文件, go test -v会详细的显示执行的过程
* test case的入参数为t testing.T或b testing.B
* t.Error为打印错误信息,并当前test case会被跳过
* t.SkipNow()为跳过test,并直接按PASS处理下一个test, 并必须写在test case的第一行, 否则无效
* go的test不会保证多个TestXxx是顺序执行,但是通常会按顺序执行, 为了,让起顺序执行,可以采用t.Run(name string, f func)来保证顺序执行
* TestMain(m *testing.M)作为初始化test,并使用m.Run()来调用其他tests可以完成一些需要初始化操作testing,比如数据库连接,文件打开,REST服务登陆,如果没有在testMain调用m.Run(),则除了TestMain以外其他test case都不会执行.

