package main

const (
	URL                  = "https://github.com/Williamjacobsen/ClosedAI/tree/main"
	BRANCH               = "main"
	LOGGING              = true
	SAVE_RESULTS_TO_FILE = true
)

func main() {
	fileUrls := getAllFileUrls()

	// For testing:
	//fileUrls := []string {
	//	"https://raw.githubusercontent.com/Williamjacobsen/ClosedAI/refs/heads/main/WebApp%2Fmicroservices%2Fredis_testing.py",
	//	"https://raw.githubusercontent.com/Williamjacobsen/ClosedAI/refs/heads/main/WebApp%2Fbackend2%2Fclosedai%2Fsrc%2Fmain%2Fjava%2Fcom%2Fclosedai%2Fclosedai%2Fsse%2FSseService.java",
	//	"https://raw.githubusercontent.com/Williamjacobsen/ClosedAI/refs/heads/main/WebApp%2Fstart.bat",
	//}

	fileExtensions := getFileExtensions()

	analyseFiles(fileUrls, fileExtensions)
}
