package command

// FormulaAsset defines the asset name of a formula template.
const FormulaAsset = "assets/formula.rb"

type Formula struct {
	Package  string
	UserName string
}

func (f *Formula) Generate() (res []byte, err error) {

	return generate(FormulaAsset, f)

}
