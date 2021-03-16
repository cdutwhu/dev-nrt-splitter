package splitter

import (
	"os"
	"path/filepath"
	"strings"
	"sync"

	csvtool "github.com/cdutwhu/csv-tool"
	"github.com/digisan/gotk/slice/ts"
)

func CsvSplitter(csvfile, outdir string, categories ...string) error {

	outdir = strings.TrimSuffix(outdir, "/") + "/"
	basename := filepath.Base(csvfile)
	return split(0, csvfile, outdir, basename, categories)

}

func split(rl int, csvfile, outdir, basename string, categories []string, pCatItems ...string) error {
	if rl >= len(categories) {
		return nil
	}

	defer func() {
		if rl > 1 && rl <= len(categories) {
			os.RemoveAll(csvfile)
		}
	}()

	cat := categories[rl]
	rl++

	_, rows, err := csvtool.Subset(csvfile, true, []string{cat}, false, nil, "")
	if err != nil {
		return err
	}

	unirows := ts.MkSet(rows...)

	wg := &sync.WaitGroup{}
	wg.Add(len(unirows))

	for _, catItem := range unirows {

		go func(wg *sync.WaitGroup, catItem string) {
			defer wg.Done()

			outcsv := outdir
			for _, pcItem := range pCatItems {
				outcsv += pcItem + "/"
			}
			outcsv += catItem + "/" + basename
			// fmt.Println(outcsv)

			_, _, err := csvtool.Query(csvfile,
				false,
				[]string{cat},
				'&',
				[]csvtool.Condition{{Hdr: cat, Val: catItem, ValTyp: "string", Rel: "="}},
				outcsv,
				nil,
			)
			if err != nil {
				panic(err)
			}

			split(rl, outcsv, outdir, basename, categories, append(pCatItems, catItem)...)

		}(wg, catItem)
	}

	wg.Wait()

	return nil
}
