package command

// MakefileAsset defines the asset name of Makefile.
const MakefileAsset = "assets/Makefile"

type Makefile struct {
	// Directory to store package files
	Dest string
	// GitHub user name.
	UserName string
}

func (m *Makefile) Generate() (res []byte, err error) {

	return generateFromAsset(MakefileAsset, m)

}
