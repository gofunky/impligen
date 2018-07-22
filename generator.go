package impligen

const (
	// If the repository is not public or doesn't meet the requirements although required as such from the generator
	ErrNotFound = 404
	// If the code generation fails
	ErrInternalServerError = 500
)

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
		// Target path of an existing git repository where the files are to be generated
		Target string
	}

	// The result of the code generation
	GenResult struct {
		// If the generation fails or the parameters don't meet the requirements, otherwise nil
		Err error
	}

	// The generator-specific settings for impligen.
	Settings struct {
		// If only publically available repositories are allowed as generation argument
		OnlyPublished bool
	}

	// To be implemented by the generator.
	Genner interface {
		// Get the parameters from the generator
		Settings() Settings

		// Generate the code now
		Gen(params Params) GenResult
	}
)
