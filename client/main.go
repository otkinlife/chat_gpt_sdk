package main

import (
	"flag"
	"fmt"
	"os"
	"syscall"

	chatgpt "github.com/golang-infrastructure/go-ChatGPT"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	oldTermState *terminal.State
)

const prompt = "ChatGPT> "

func main() {
	var token string
	flag.StringVar(&token, "t", "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6Ik1UaEVOVUpHTkVNMVFURTRNMEZCTWpkQ05UZzVNRFUxUlRVd1FVSkRNRU13UmtGRVFrRXpSZyJ9.eyJodHRwczovL2FwaS5vcGVuYWkuY29tL3Byb2ZpbGUiOnsiZW1haWwiOiJpYWJ5Y3h2MDU3MzlAMjFjbi5jb20iLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwiZ2VvaXBfY291bnRyeSI6IlVTIn0sImh0dHBzOi8vYXBpLm9wZW5haS5jb20vYXV0aCI6eyJ1c2VyX2lkIjoidXNlci16NWNXVzVoRjJESHdUSXZpR3lpYjJCeFAifSwiaXNzIjoiaHR0cHM6Ly9hdXRoMC5vcGVuYWkuY29tLyIsInN1YiI6ImF1dGgwfDYzOTQxYjBkZmU2OGQ0YjQ5MWI3NWVmMCIsImF1ZCI6WyJodHRwczovL2FwaS5vcGVuYWkuY29tL3YxIiwiaHR0cHM6Ly9vcGVuYWkuYXV0aDAuY29tL3VzZXJpbmZvIl0sImlhdCI6MTY3MDc0MTg0NCwiZXhwIjoxNjcwNzg1MDQ0LCJhenAiOiJUZEpJY2JlMTZXb1RIdE45NW55eXdoNUU0eU9vNkl0RyIsInNjb3BlIjoib3BlbmlkIGVtYWlsIHByb2ZpbGUgbW9kZWwucmVhZCBtb2RlbC5yZXF1ZXN0IG9yZ2FuaXphdGlvbi5yZWFkIG9mZmxpbmVfYWNjZXNzIn0.L00791-4N9VQID-VxoLvC6NEJkGl963Fo2qwlNk4RVv8m2HfqXGLvlEu8zkI3L5CbA7EpEmbMHLBP8r3f9GE20i-vpnQsR3ve6syT2jMeI_r_c21S8Sb-vZQtF7xTrLARjKKvCyAFkQoVjU-ROMMeylFd4I9JeuExlOiuFY9W3gCvETZa12AIOUBHuIrZt2_WqMqMMly1cijk3_J2F6Unn_-ibw8Z04T3g7f-_7sTI4QGt24pAKGarmi0_S2g6LO6Vr78uMPNgxID5EPYGUoMdHEDiw7LiE9RejS_GwTk10i983lQgzxAK0G4QPLppQzpPWlkSNGIe32yoV__XH5aA", "【chatGPT token】run code in chrome console with JSON.parse(document.getElementById('__NEXT_DATA__').text).props.pageProps.accessToken")
	flag.Parse()
	if token == "" {
		panic("token is empty")
	}

	var err error
	oldTermState, err = terminal.MakeRaw(syscall.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}

	term := terminal.NewTerminal(os.Stdin, prompt)
	term.Write([]byte("ChatGPT CLI <KcJia> - version 1.0\n"))
	if err != nil {
		fmt.Println(err)
		return
	}

	chat := chatgpt.NewChatGPT(token, term)

	for {
		term.Write([]byte("ChatGPT> Please input your question:\n"))
		question, err := term.ReadLine()
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		err = chat.Talk(question)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
	}
	terminal.Restore(syscall.Stdin, oldTermState)
}
