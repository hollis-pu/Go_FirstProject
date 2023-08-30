package student

/**
* Description:
	student包
* @Author Hollis
* @Create 2023/8/29 20:50
*/

// student的s为小写，则在其他包中无法创建其实例变量，可以使用工厂模式来提供一个创造实例对象的方法
type student struct {
	name   string
	age    int
	gender string
}

// Factory 2023.08.29疑问：这里的返回值为什么要使用*student类型而不是student类型
// 自己觉得的可能原因：让返回值是一个指针而不是具体的student类型的数据，减少传输的数据量
func Factory(name string, age int, gender string) *student {
	return &student{name, age, gender}
}

// GetName 为私有的属性提供包外获取的GetXxx()函数
// 这里使用指针的原因同理，减少传输的数据量
func (s *student) GetName() string {
	return s.name
}

func (s *student) GetAge() int {
	return s.age
}

func (s *student) GetGender() string {
	return s.gender
}
