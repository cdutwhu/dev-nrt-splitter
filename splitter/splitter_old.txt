package splitter

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	csvtool "github.com/cdutwhu/csv-tool"
	"github.com/digisan/gotk/slice/ts"
)

func csvSplitter(csvfile, outdir string, categories ...string) error {
	// headers, _, err := csvtool.FileInfo(csvfile)
	// if err != nil {
	// 	return err
	// }
	// fmt.Println(headers)

	// CAT0 ------------------------------------------ //

	outdir = strings.TrimSuffix(outdir, "/") + "/"
	basename := filepath.Base(csvfile)

	cat0 := categories[0]
	_, rows0, err := csvtool.Subset(csvfile, true, []string{cat0}, false, nil, "")
	if err != nil {
		return err
	}
	rows0 = ts.MkSet(rows0...) // make unique
	// ts.FM(rows0, nil, func(i int, e string) string { fmt.Println(e); return "" })

	wg := &sync.WaitGroup{}
	wg.Add(len(rows0))

	for _, cat0Item := range rows0 {

		go func(cat0Item string, wg *sync.WaitGroup) {
			defer func() {
				fmt.Println(cat0Item, "done")
				wg.Done()
			}()

			outcsv := fmt.Sprintf("%s%s/%s", outdir, cat0Item, basename)
			// fmt.Println(outcsv)

			_, _, err := csvtool.Query(csvfile,
				false,
				[]string{cat0},
				'&',
				[]csvtool.Condition{{Hdr: cat0, Val: cat0Item, ValTyp: "string", Rel: "="}},
				outcsv,
				nil,
			)
			if err != nil {
				panic(err)
			}

			// CAT1 ------------------------------------------ //

			func() error {
				csvfile := outcsv
				defer func() { os.RemoveAll(csvfile) }()

				_, rows1, err := csvtool.Subset(csvfile, true, []string{categories[1]}, false, nil, "")
				if err != nil {
					return err
				}

				for _, cat1Item := range ts.MkSet(rows1...) {
					outcsv := fmt.Sprintf("%s%s/%s/%s", outdir, cat0Item, cat1Item, basename)
					// fmt.Println(outcsv)

					_, _, err := csvtool.Query(csvfile,
						false,
						[]string{categories[1]},
						'&',
						[]csvtool.Condition{{Hdr: categories[1], Val: cat1Item, ValTyp: "string", Rel: "="}},
						outcsv,
						nil,
					)
					if err != nil {
						panic(err)
					}

					// CAT2 ------------------------------------------ //

					func() error {
						csvfile := outcsv
						defer func() { os.RemoveAll(csvfile) }()

						_, rows2, err := csvtool.Subset(csvfile, true, []string{categories[2]}, false, nil, "")
						if err != nil {
							return err
						}

						for _, cat2Item := range ts.MkSet(rows2...) {
							outcsv := fmt.Sprintf("%s%s/%s/%s/%s", outdir, cat0Item, cat1Item, cat2Item, basename)
							// fmt.Println(outcsv)

							_, _, err := csvtool.Query(csvfile,
								false,
								[]string{categories[2]},
								'&',
								[]csvtool.Condition{{Hdr: categories[2], Val: cat2Item, ValTyp: "string", Rel: "="}},
								outcsv,
								nil,
							)
							if err != nil {
								panic(err)
							}
						}
						return nil
					}()
				}
				return nil
			}()

		}(cat0Item, wg)
	}

	wg.Wait()
	return nil
}
