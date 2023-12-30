package cmd

import (
	"image/png"
	"os"

	"github.com/ancuongnguyen07/BlurHashGo/blurhash"
	"github.com/spf13/cobra"
)

type decodeOptions struct {
	hash   string
	dest   string
	width  int
	height int
	punch  int
}

func decodeCmd() *cobra.Command {
	options := decodeOptions{}

	cmd := &cobra.Command{
		Use:   "decode",
		Short: "Decode an image from local file or downloaded url",
		Long:  `Decode an image from local file or downloaded url. Prints the output to stdout`,
		Example: `
		blurhashgo decode --hash <BLURHASH> --width <WIDTH> --height <HEIGHT> --punch <PUNCH>
		`,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return decodeExecute(options)
		},
	}

	flag := cmd.Flags()
	flag.StringVar(&options.hash, "hash", "", "blurhash value of a particular image")
	flag.StringVar(&options.dest, "dest", "", "the filepath for the output image")
	flag.IntVar(&options.width, "width", 0, "width of an output image (pixel)")
	flag.IntVar(&options.height, "height", 0, "height of an output image (pixel)")
	flag.IntVar(&options.punch, "punch", 1, "a factor that modifies the constrast level")

	cmd.MarkFlagRequired("hash")
	cmd.MarkFlagRequired("width")
	cmd.MarkFlagRequired("height")
	cmd.MarkFlagRequired("punch")

	return cmd
}

func init() {
	rootCmd.AddCommand(decodeCmd())
}

func decodeExecute(options decodeOptions) (err error) {
	decodedImg, err := blurhash.Decode(
		options.hash,
		options.width,
		options.height,
		options.punch,
	)
	if err != nil {
		return err
	}

	var dest string
	if options.dest == "" {
		// use the current working directory as no dest is specified
		dest, err = os.Getwd()
		if err != nil {
			return err
		}
		dest = dest + "/output.png"
	} else {
		// use the specififed destination path
		dest = options.dest
	}

	// create a new empty file
	outFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// write the image to output file
	err = png.Encode(outFile, decodedImg)

	return err
}
