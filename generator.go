package impligen

type (
	// The generation parameters provided
	Params struct {
		// The path to the package to be imported
		Path string
		// The package name residing in the imported path
		Package string
		// The chosen type for the generic
		Type string
		// The suffix as an optional parameter, "" for the root
		Suffix string
	}

	GenResult struct {
		// TODO TBD how the result code is published
	}

	Genner interface {
		Gen(params Params) GenResult
	}
)
