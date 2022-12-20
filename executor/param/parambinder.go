package executor

import (
	"errors"
	"strings"
	"taskmanager/pipeline"
)

type ScriptParamBinder interface {
	Bind(script *string) (string, error)
}

type scriptparambinder struct {
	variablemanager *pipeline.VariableManager
	tokenizer       Tokenizer
}

func NewParamBinder(variableManager *pipeline.VariableManager) ScriptParamBinder {
	scriptparambinder := &scriptparambinder{}
	scriptparambinder.Init(variableManager)
	return scriptparambinder
}

func (scriptparambinder *scriptparambinder) Init(variableManager *pipeline.VariableManager) error {
	scriptparambinder.variablemanager = variableManager
	tokenizer := NewTokenizer()
	scriptparambinder.tokenizer = tokenizer
	return nil
}

func (scriptparambinder *scriptparambinder) Bind(script *string) (string, error) {
	var variablemap map[string]string = make(map[string]string)
	// parse script parameter
	params := scriptparambinder.tokenizer.Tokenize(*script)
	for _, param := range params {
		// Get variable from variable manager
		// parameter pass to GetVariable should skip $ notation
		varVal := scriptparambinder.variablemanager.GetVariable(param[1:])
		if len(varVal) == 0 {
			return "", errors.New("Undefined variable: " + param)
		}
		variablemap[param] = varVal
	}

	executableScript := scriptparambinder.replace(script, variablemap)
	return executableScript, nil
}

func (scriptparambinder scriptparambinder) replace(script *string, variablemap map[string]string) string {
	var executableScript string = *script
	for param, paramVal := range variablemap {
		executableScript = strings.ReplaceAll(executableScript, param, paramVal)
	}

	return executableScript
}
