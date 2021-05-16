package factorymethod

type IRuleConfigParserFactory interface {
	CreateParser() IRuleConfigParserFactory
}

type jsonRuleConfigParserFactory struct{}

func (J jsonRuleConfigParserFactory) CreateParser() IRuleConfigParserFactory {
	return jsonRuleConfigParserFactory{}
}

type yamlRuleConfigParserFactory struct{}

func (Y yamlRuleConfigParserFactory) CreateParser() IRuleConfigParserFactory {
	return yamlRuleConfigParserFactory{}
}

//使用一个简单工厂来创建
func NewIRulrConfigParserFactory(t string) IRuleConfigParserFactory {
	switch t {
	case "json":
		return jsonRuleConfigParserFactory{}
	case "yaml":
		return yamlRuleConfigParserFactory{}
	}
	return nil
}
