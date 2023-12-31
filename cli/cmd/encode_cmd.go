package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/ancuongnguyen07/BlurHashGo/blurhash"
	"github.com/ancuongnguyen07/BlurHashGo/internal/utils"
	"github.com/spf13/cobra"
)

type encodeOptions struct {
	file  string
	url   string
	xcomp int
	ycomp int
}

func encodeCmd() *cobra.Command {
	options := encodeOptions{}

	cmd := &cobra.Command{
		Use:   "encode",
		Short: "Encode an image from local file or downloaded url",
		Long:  `Encode an image from local file or downloaded url. Prints the output to stdout`,
		Example: `
		blurhashgo encode --file <FILE> --xcomp <XCOMP> --ycomp <YCOMP>
		blurhashgo encode --url <URL> --xcomp <XCOMP> --ycomp <YCOMP>
		`,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return encodeExecute(options)
		},
	}

	flags := cmd.Flags()
	flags.StringVar(&options.file, "file", "", "the relative path leading to the image file")
	flags.StringVar(&options.url, "url", "", "the URL for downloading the image file")
	flags.IntVar(&options.xcomp, "xcomp", 1, "x-component indicates how many components along x-axis you want to catch during DFT")
	flags.IntVar(&options.ycomp, "ycomp", 1, "y-component indicates how many components along y-axis you want to catch during DFT")

	cmd.MarkFlagsOneRequired("file", "url")
	cmd.MarkFlagsMutuallyExclusive("file", "url")
	cmd.MarkFlagRequired("xcomp")
	cmd.MarkFlagRequired("ycomp")

	return cmd
}

func init() {
	rootCmd.AddCommand(encodeCmd())
}

func encodeExecute(options encodeOptions) (err error) {
	var file string
	if options.file != "" {
		// local image file mode
		file = options.file
	} else {
		// download image file from the Internet
		file, err = downloadFile(options.url)
		if err != nil {
			return err
		}
	}

	// encode local file
	img, err := utils.ReadImgFile(file)
	if err != nil {
		return err
	}

	imgBlurhash, err := blurhash.Encode(options.xcomp, options.ycomp, img)
	if err != nil {
		return err
	}

	fmt.Println(imgBlurhash)
	return nil
}

// downloadFile downloads file with the given URL.
// It returns the local file path to the downloaded file.
func downloadFile(url string) (string, error) {
	// Does HTTP Get request
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// extract the file name from the given URL
	urlFields := strings.FieldsFunc(url, func(r rune) bool {
		return r == '/'
	})
	file := fmt.Sprintf("/tmp/%s", urlFields[len(urlFields)-1])

	// create a new empty file
	fout, err := os.Create(file)
	if err != nil {
		return "", err
	}
	defer fout.Close()

	// write the resquest body to the newly created file
	_, err = io.Copy(fout, resp.Body)
	if err != nil {
		return "", err
	}
	return file, nil
}
