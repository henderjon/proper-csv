package main

import(
	"os"
	"encoding/csv"
	"flag"
	"log"
	"io"
)

var (
	inputDelim, outputDelim string
	quotes, reverse, help bool
)

func init(){
	flag.StringVar(&inputDelim, "i", "\t", "the character to use as the delimiter for input")
	flag.StringVar(&outputDelim, "o", ",", "the character to use as the delimiter for output")
	flag.BoolVar(&reverse, "rev", false, "reverse the default delimiters")
	flag.BoolVar(&quotes, "q", false, "reverse the default delimiters")
	flag.BoolVar(&help, "h", false, "display this message")
	flag.Parse()
}

func main(){
	if(help){
		log.Println("proper-csv takes a data stream of particularly delimited input and encodes it using a different delimiter. For example taking TSV data and encoding it as CSV data.")
		flag.PrintDefaults()
		return
	}

	r := csv.NewReader(os.Stdin)
	w := csv.NewWriter(os.Stdout)

	if(quotes){
		r.LazyQuotes = quotes
	}

	if(reverse){
		r.Comma = []rune(outputDelim)[0]
		w.Comma = []rune(inputDelim)[0]
	}else{
		r.Comma = []rune(inputDelim)[0]
		w.Comma = []rune(outputDelim)[0]
	}

	for {
		record, err := r.Read();
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Println("error reading record from Stdin:", err)
		}

		if err := w.Write(record); err != nil {
			log.Println("error writing record to Stdout:", err)
		}
	}

	w.Flush()

}
