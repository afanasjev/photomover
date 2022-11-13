package photomover

import (
    "github.com/dsoprea/go-exif/v3"
    "github.com/dsoprea/go-exif/v3/common"
    "github.com/dsoprea/go-logging"
    "fmt"
)

func HelloWorld () {
    fmt.Println("Hello world")
}

func GetPhotoDate(filepath string) string {
    rawExif, err := exif.SearchFileAndExtractExif(filepath)
    log.PanicIf(err)

    im, err := exifcommon.NewIfdMappingWithStandard()
    log.PanicIf(err)

    ti := exif.NewTagIndex()

    _, index, err := exif.Collect(im, ti, rawExif)
    log.PanicIf(err)

    tagName := "DateTimeOriginal"
    // rootIfd := index.RootIfd
    rootIfd := index.Lookup["IFD/Exif"]

    // We know the tag we want is on IFD0 (the first/root IFD).
    results, err := rootIfd.FindTagWithName(tagName)
    log.PanicIf(err)

    // This should never happen.
    if len(results) != 1 {
        log.Panicf("there wasn't exactly one result")
    }

    ite := results[0]

    valueRaw, err := ite.Value()
    log.PanicIf(err)

    value := valueRaw.(string)
    return value
}
