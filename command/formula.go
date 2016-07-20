package command

// FormulaTemplateAsset defines the asset name of a formula template.
const FormulaTemplateAsset = "assets/formula.rb"

type FormulaTemplate struct {
	Package  string
	UserName string
}

type Formula struct {
	Version     string
	FileName64  string
	FileName386 string
	Hash64      string
	Hash386     string
}

func (f *FormulaTemplate) Generate() (res []byte, err error) {

	return generateFromAsset(FormulaTemplateAsset, f)

}

func (f *Formula) Generate(path string) (ref []byte, err error) {

	return generateFromFile(path, f)

}
