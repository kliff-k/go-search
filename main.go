//package main
//
//import (
//	"os"
//
//	"github.com/therecipe/qt/widgets"
//)
//
//func main() {
//
//
//
//// needs to be called once before you can start using the QWidgets
//	app := widgets.NewQApplication(len(os.Args), os.Args)
//
//	// create a window
//	// with a minimum size of 250*200
//	// and sets the title to "Hello Widgets Example"
//	window := widgets.NewQMainWindow(nil, 0)
//	window.SetMinimumSize2(500, 500)
//	window.SetWindowTitle("Hot damn!!")
//
//	// create a regular widget
//	// give it a QVBoxLayout
//	// and make it the central widget of the window
//	widget := widgets.NewQWidget(nil, 0)
//	widget.SetLayout(widgets.NewQVBoxLayout())
//	window.SetCentralWidget(widget)
//
//	// create a line edit
//	// with a custom placeholder text
//	// and add it to the central widgets layout
//	input := widgets.NewQLineEdit(nil)
//	input.SetPlaceholderText("Write something ...")
//	widget.Layout().AddWidget(input)
//
//	// create a button
//	// connect the clicked signal
//	// and add it to the central widgets layout
//	button := widgets.NewQPushButton2("and click me!", nil)
//	button.ConnectClicked(func(bool) {
//		widgets.QMessageBox_Information(nil, "OK", input.Text(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
//	})
//	widget.Layout().AddWidget(button)
//
//	// make the window visible
//	window.Show()
//
//	// start the main Qt event loop
//	// and block until app.Exit() is called
//	// or the window is closed by the user
//	app.Exec()
//}
//
//package main
//
//import (
//"bytes"
//"flag"
//"fmt"
//"io/ioutil"
//"log"
//"os"
//"path/filepath"
//"strings"
//"time"
//
//"golang.org/x/net/context"
//"golang.org/x/sync/errgroup"
//)
//
//var (
//	Version   = "N/A"
//	BuildTime = "N/A"
//)
//
//func main() {
//	duration := flag.Duration("timeout", 500*time.Millisecond, "timeout in milliseconds")
//	flag.Usage = func() {
//		fmt.Printf("%s by Brian Ketelsen\n", os.Args[0])
//		fmt.Printf("Version %s, Built: %s \n", Version, BuildTime)
//		fmt.Println("Usage:")
//		fmt.Printf("	gogrep [flags] path pattern \n")
//		fmt.Println("Flags:")
//		flag.PrintDefaults()
//	}
//	flag.Parse()
//	if flag.NArg() != 2 {
//		flag.Usage()
//		os.Exit(-1)
//	}
//	path := flag.Arg(0)
//	pattern := flag.Arg(1)
//	ctx, cf := context.WithTimeout(context.Background(), *duration)
//	defer cf()
//	m, err := search(ctx, path, pattern)
//	if err != nil {
//		log.Fatal(err)
//	}
//	for _, name := range m {
//		fmt.Println(name)
//	}
//	fmt.Println(len(m), "hits")
//}
//
//func search(ctx context.Context, root string, pattern string) ([]string, error) {
//	g, ctx := errgroup.WithContext(ctx)
//	paths := make(chan string, 100)
//	// get all the paths
//
//	g.Go(func() error {
//		defer close(paths)
//
//		return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
//			if err != nil {
//				return err
//			}
//			if !info.Mode().IsRegular() {
//				return nil
//			}
//			if !info.IsDir() && !strings.HasSuffix(info.Name(), ".go") {
//				return nil
//			}
//
//			select {
//			case paths <- path:
//			case <-ctx.Done():
//				return ctx.Err()
//			}
//			return nil
//		})
//
//	})
//
//	c := make(chan string, 100)
//	for path := range paths {
//		p := path
//		g.Go(func() error {
//			data, err := ioutil.ReadFile(p)
//			if err != nil {
//				return err
//			}
//			if !bytes.Contains(data, []byte(pattern)) {
//				return nil
//			}
//			select {
//			case c <- p:
//			case <-ctx.Done():
//				return ctx.Err()
//			}
//			return nil
//		})
//	}
//	go func() {
//		g.Wait()
//		close(c)
//	}()
//
//	var m []string
//	for r := range c {
//		m = append(m, r)
//	}
//	return m, g.Wait()
//}
//

//package main
//
//import (
//	"bufio"
//	"flag"
//	"fmt"
//	"os"
//	"path/filepath"
//	"strings"
//	"sync"
//)
//
//var root, query string
//var found = 1
//var wg sync.WaitGroup
//
//func readFile(wg *sync.WaitGroup, path string) {
//	defer wg.Done()
//
//	file, err := os.Open(path)
//	defer file.Close()
//
//	if err != nil {
//		return
//	}
//	scanner := bufio.NewScanner(file)
//	for i := 1; scanner.Scan(); i++ {
//		if strings.Contains(scanner.Text(), query) {
//			found = 0
//			fmt.Printf("%s/%s:%d: %s\n", root, path, i, scanner.Text())
//		}
//	}
//}
//
//func main() {
//	flag.Parse()
//	query = flag.Arg(0)
//	root = flag.Arg(1)
//
//	filepath.Walk(root, func(path string, file os.FileInfo, err error) error {
//		if !file.IsDir() {
//			wg.Add(1)
//			go readFile(&wg, path)
//		}
//		return nil
//	})
//	wg.Wait()
//	defer os.Exit(found)
//}

package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	var pathList []string
	pathList = append(pathList, "E:\\Windows\\Games\\Touhou")
	pathList = append(pathList, "E:\\Windows\\Games\\RagnarokBattleOffline")

	for _, path := range pathList {
		files, err := ioutil.ReadDir(path)

		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			fmt.Println(file.Name())
		}
	}
}
