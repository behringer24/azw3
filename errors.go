package azw3

import "errors"

var (
	// ErrNoContent is returned by Write/Serialize when the book has no
	// chapters. A KF8 book with no text content is malformed.
	ErrNoContent = errors.New("azw3: book must have at least one chapter")

	// ErrUnsupportedImageFormat is returned by AddImage/AddImageFile/
	// SetCoverImage when image data is not one of the supported formats
	// (JPEG, PNG, GIF).
	ErrUnsupportedImageFormat = errors.New("azw3: unsupported image format")

	// ErrInvalidLanguage is returned by AddLanguage when the given tag is
	// not a valid BCP 47 / RFC 3066 language tag.
	ErrInvalidLanguage = errors.New("azw3: invalid language tag")

	// ErrInvalidContributorRole is returned by AddContributor when role is
	// not a recognized MARC relator code.
	ErrInvalidContributorRole = errors.New("azw3: invalid contributor role")

	// ErrDuplicatePath is returned by AddChapter/AddImage/AddStylesheet
	// when path is already in use by a previously added resource.
	ErrDuplicatePath = errors.New("azw3: path already in use")

	// ErrChapterNotFound is returned by Write/Serialize when a Navpoint's
	// target does not match the Id of any chapter added via AddChapter.
	ErrChapterNotFound = errors.New("azw3: navpoint target does not match any chapter")
)
